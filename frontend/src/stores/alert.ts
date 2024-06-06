import { writable } from 'svelte/store'

// The idea is that whenever there is an error in the app, we push it to the store.
//
// There is a component watching the store that will display the errors for users to see.

export enum AlertType {
    Info = 'info',
    Success = 'success',
    Danger = 'danger',
    Warning = 'warning'
}

export interface Alert {
    Id: number;
    Message: string;
    Type: AlertType;
}

export const alerts = writable<Alert[]>([]);

function AlertId(): number {
    return Math.floor(Math.random() * 10000)
}

export function AddAlert(message: string, type: AlertType) {
    let id: number = AlertId()
    let new_alert: Alert = {Id: id, Message: message, Type: type}

    addAlert(new_alert)
}

const addAlert = (new_alert: Alert) => {
    alerts.update((all) => [new_alert, ...all]);
}

export const RemoveAlert = (id: number) => {
    alerts.update((all) => all.filter((a) => a.Id !== id))
}
