import type { model } from "../../../../frontend/src/lib/wailsjs/go/models";
/**
 * A distance is a unit with a length
 */
export type Distance = {
    unit: string;
    length: number;
}

/**
 * return the desired units based on user preference
 * @param type - types can be: vertical, default
 * @returns the units assocaited with the type given user preference
 */
export function HandleUserUnits(usr: model.User, type: string): string {
    if (type == "vertical") {
        if (usr.units == "imperial") {
            return "ft"
        } else if (usr.units == "metric") {
            return "m"
        }
    } else if (type == "default") {
        if (usr.units == "imperial") {
            return "mi"
        } else if (usr.units == "metric") {
            return "km"
        }
    }
    return ""
}

/**
 * Returns the empty length based on user unit preferences
 * @param is_vertical - if the length is describing vertical or not
 * @param usr - the current user (stored in cookies)
 * @returns A distance object
 */
export function EmptyLength(is_vertical: boolean, usr: model.User): Distance {
    if (is_vertical) {
        return {unit: HandleUserUnits(usr, "vertical"), length: 0}
    } else {
        return {unit: HandleUserUnits(usr, "default"), length: 0}
    }
}