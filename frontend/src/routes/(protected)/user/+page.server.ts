import { redirect, type Actions } from "@sveltejs/kit";
import { MAX_AGE } from "../../../config";
import { auth_store } from "../../../stores/auth";
import { model } from "../../../../wailsjs/go/models";
import { UpdateUser } from "../../../../wailsjs/go/controllers/UserHandler";

export const actions: Actions = {
    update: async ({ cookies, request }) => {
        const data = await request.formData()

        const uuid = data.get("uuid") as string
        const unit_class = data.get("unit_class") as string
        const cycle_start = data.get("cycle_start") as string
        const cycle_days = data.get("cycle_days") as number | null
        const initial_cycle_start = data.get("initial_cycle_start") as string


        if (cycle_days != null) {
            let usr: model.User = new model.User({
                uuid: uuid,
                username: "",
                units: unit_class,
                cycle_days: Number(cycle_days),
                cycle_start: cycle_start,
                initial_cycle_start: initial_cycle_start
            })
            const edited_user = await UpdateUser(usr)

            if (typeof edited_user == "string") {
                console.error("error updating user:" + edited_user)
                return { success: "failed" }
            } else {
                let existing_cookie = cookies.get("session")
                if (existing_cookie != undefined) {
                    const cookie = JSON.parse(existing_cookie)
                    let new_cookie = {
                        uuid: uuid,
                        username: cookie.user.username,
                        password: cookie.user.password,
                        units: unit_class,
                        cycle_start: cycle_start,
                        cycle_days: cycle_days,
                        initial_cycle_start: initial_cycle_start
                    }
                    cookies.set("session", JSON.stringify({ user: new_cookie, is_authenticated: true}), {
                        path: '/',
                        maxAge: MAX_AGE,
                        httpOnly: true,
                        sameSite: 'lax'
                    })
                    auth_store.set({ user: new_cookie, is_authenticated: true})
                    return redirect(302, "/user")
                }
            }
        }
    }
} satisfies Actions