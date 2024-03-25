import { get, writable, type Writable } from "svelte/store";
import { page } from "$app/stores";
import type { model } from "../../wailsjs/go/models";

// The user data object can be found at:
// get(page).data.session.user.uuid

export function GetUserUuid(): string {
    let uuid: string = get(page).data.session.user.uuid
    return uuid
}

export function GetUserUnits(): string {
    let units: string = get(page).data.session.user.units
    return units
}
export type AuthState = {
    is_authenticated: boolean;
    user: model.User | null;
}

let stored_auth_state: AuthState = {
    is_authenticated: false,
    user: null
};

export const auth_store: Writable<AuthState> = writable(stored_auth_state)
