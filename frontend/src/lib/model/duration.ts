/**
 * Validate a duration string of the form HH:MM:SS. Return 0 if the string is empty
 * @param dur string of the form HH:MM:SS
 * @returns that duration in seconds
 */
export function ValidateDuration(dur: string): number | string {
    if (dur == "") {
        return 0
    } else if (dur == "00:00:00") {
        return 0
    }

    const durationPattern = /^[0-9]{1,2}:[0-5][0-9]:[0-5][0-9]$/;
    if (durationPattern.test(dur)) {
        let [hours, minutes, seconds] = dur.split(':').map(Number);

        let tot_seconds = (hours * 60 * 60) + (minutes * 60) + seconds
        return tot_seconds
    } else {
        return "Invalid duration format. Please use HH:MM:SS";
    }
};