import { redirect, type Actions } from "@sveltejs/kit";
import { auth_store } from "../../../../stores/auth";
import { MAX_AGE } from "../../../../config";
import { CreateUser } from "../../../../../wailsjs/go/controllers/UserHandler";
import { model } from "../../../../../wailsjs/go/models";

export const actions: Actions = {
    _register: async ({ cookies, request }) => {
        const data = await request.formData();

        const username = data.get("username") as string;
        const password = data.get("password") as string;
        const units = data.get("units") as string;
        const cycle_start = data.get("cycle_start") as string;
        const cycle_days = data.get("cycle_days") as unknown as number;
        const initial_cycle_start = data.get("initial_cycle_start") as string;

        let new_user: model.User = new model.User({
            username: username,
            password: password,
            units: units,
            cycle_start: cycle_start,
            cycle_days: Number(cycle_days),
            initial_cycle_start: initial_cycle_start
        });

        let create_response = await CreateUser(new_user);
        if (create_response) {
            let session = { user: create_response, is_authenticated: true }
            cookies.set("session", JSON.stringify(session), {
                path: '/',
                maxAge: MAX_AGE,
                httpOnly: true,
                sameSite: 'lax'
            });

            return redirect(302, "/home");
        } else {
            return { success: "already exists" };
        }

    },
    get register() {
        return this._register;
    },
    set register(value) {
        this._register = value;
    },
} satisfies Actions;