import { EmptyLength, type Distance } from "./distance";
import type { DateObject } from "./date";
import { model } from "../wailsjs/go/models";

/**
 * Create an empty total
 * @param usr the user
 * @returns an emtpy total
 */
export function EmptyTotal(usr: model.User): model.Totals {
    let a = {
        total_planned_distance:EmptyLength(false, usr),total_planned_duration:0,total_planned_vertical:EmptyLength(true, usr),
        total_completed_distance:EmptyLength(false, usr),total_completed_duration:0,total_completed_vertical:EmptyLength(true, usr),
    }
    return new model.Totals(a)
}
/**
 * Create an empty total differences
 * @param usr the user
 * @returns an empty totals differences
 */
export function EmptyTotalsDifferences(usr: model.User): model.TotalsDifferences {
    let a = {
        difference_planned_distance: EmptyLength(false, usr), difference_planned_duration:0, difference_planned_vertical:EmptyLength(true, usr),
        percent_difference_planned_distance: NaN, percent_difference_planned_duration: NaN, percent_difference_planned_vertical: NaN,
        difference_completed_distance: EmptyLength(false, usr), difference_completed_duration:0, difference_completed_vertical:EmptyLength(true, usr),
        percent_difference_completed_distance: NaN, percent_difference_completed_duration: NaN, percent_difference_completed_vertical: NaN,
    }
    return new model.TotalsDifferences(a)
}
