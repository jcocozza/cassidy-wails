import type { DateObject } from "./date"
import type { Distance } from "./distance";
/**
 * Represents a summary of the passed N cycles
 */
export type NCycleSummary = {
    start_date_list: DateObject[];
    planned_distances: Distance[];
    planned_durations: number[];
    planned_verticals: Distance[];
    completed_distances: Distance[];
    completed_durations: number[];
    completed_verticals: Distance[];
}
