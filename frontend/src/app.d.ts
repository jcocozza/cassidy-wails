// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import { User } from "./lib/model/user";
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			session: {
				user: User;
				is_authenticated: boolean;
			}
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};