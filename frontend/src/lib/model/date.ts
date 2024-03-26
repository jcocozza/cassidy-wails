/**
 * A date object is a date and its day of week
 */
export type DateObject = {
    day_of_week: string;
    date: string;
}

export type NextPrevious = {
    next_start_date: string;
    next_end_date: string;
    previous_start_date: string;
    previous_end_date: string;
}
/**
 * Convert a durations in seconds to a HH:MM:SS string.
 * If the duration is negative, return -HH:MM:SS string
 * If Duration is 0, return an empty string
 * @param duration a duration in seconds
 */
export function ConvertDuration(duration: number): string {

    // Calculate hours, minutes, and remaining seconds
    const hours: number = Math.floor(duration / 3600);
    const minutes: number = Math.floor((duration % 3600) / 60);
    const remainingSeconds: number = Math.floor(duration % 60);

    if (duration < 0) {
        return "-" + ConvertDuration(Math.abs(duration))
    }

    let hour_string: string = "";
    if (hours < 1) {
        hour_string = "";
    } else {
        hour_string = String(hours) + ":";
    };

    let hh_mm_ss: string = ""
    if (hour_string == "") {
        hh_mm_ss = `${hour_string}${String(minutes).padStart(2, '0')}:${String(remainingSeconds).padStart(2, '0')} min`
    } else {
        // Format the result as HH:MM:SS
        hh_mm_ss = `${hour_string}${String(minutes).padStart(2, '0')}:${String(remainingSeconds).padStart(2, '0')} hr`
    }
    return hh_mm_ss;
}
/**
 * Same as ConvertDuration(), but without hr, min labels
 * @param duration A duration in seconds
 * @returns a string representation of the duration
 */
export function FormatDurationSimple(duration: number): string {
    if (duration === 0) {
        return ""
    }

    // Calculate hours, minutes, and remaining seconds
    const hours: number = Math.floor(duration / 3600);
    const minutes: number = Math.floor((duration % 3600) / 60);
    const remainingSeconds: number = Math.floor(duration % 60);

    if (duration < 0) {
        return "-" + ConvertDuration(Math.abs(duration))
    }

    let hour_string: string = "";
    if (hours < 1) {
        hour_string = "";
    } else {
        hour_string = String(hours).padStart(2, '0') + ":";
    };

    let hh_mm_ss: string = ""
    if (hour_string == "") {
        hh_mm_ss = `${hour_string}${String(minutes).padStart(2, '0')}:${String(remainingSeconds).padStart(2, '0')}`
    } else {
        // Format the result as HH:MM:SS
        hh_mm_ss = `${hour_string}${String(minutes).padStart(2, '0')}:${String(remainingSeconds).padStart(2, '0')}`
    }
    return hh_mm_ss;
}

/**
 * Determine if a date is today
 * @param date A date in MM-DD-YYYY form
 * @returns if the date is today or not
 */
export function IsToday(date: string): boolean {
    // Get the current date
    const currentDate = new Date();

    // Create a date object for the target date "2024-01-01"
    const targetDate = new Date(date);

    // Check if the target date is equal to today's date
    if (
    targetDate.getDate() === currentDate.getDate() &&
    targetDate.getMonth() === currentDate.getMonth() &&
    targetDate.getFullYear() === currentDate.getFullYear()
    ) {
        return true;
    } else {
        return false;
    }
}