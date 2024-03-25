import { redirect, type Actions } from '@sveltejs/kit'

export const load = async ({ cookies }) => {
	// we only use this endpoint for the api
	// and don't need to see the page
    console.log("Deleting cookie")
        cookies.delete('session',{
            path: "/",
            httpOnly: true,
            sameSite: 'lax'
        });
	throw redirect(302, '/auth/login')
}