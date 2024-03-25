/**
 * Format a number percent that is in 100% form.
 * @param percent A percentage
 * @returns the formatted string
 */
export function FormatPercent(percent: number): string {
    if (percent == null) {
        return "n/a"
    } else {
        return String(percent) + "%"
    }
}