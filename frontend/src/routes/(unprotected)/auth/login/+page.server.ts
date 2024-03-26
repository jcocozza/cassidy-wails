import { redirect, type Actions } from '@sveltejs/kit'
import { auth_store } from '../../../../stores/auth';
import { MAX_AGE } from '../../../../config';
import { AuthenticateUser } from '$lib/wailsjs/go/controllers/UserHandler';
export const actions: Actions = {
    login: async ({ cookies, request }) => {

        const data = await request.formData()

        const username = data.get("username") as string
        const password = data.get("password") as string

        const user_auth = await AuthenticateUser({username: username, password: password})
        if (user_auth) {
            cookies.set("session", JSON.stringify(user_auth), {
                path: '/',
                maxAge: MAX_AGE,
                httpOnly: true,
                sameSite: 'lax',
            })

            auth_store.set({user: user_auth, is_authenticated: true})

            return redirect(302, "/microcycle/current")
        } else {
            return { success: "incorrect password" }
        }
    }
} satisfies Actions;
