SELECT
    CASE
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 0
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 1
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 2
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 3
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 4
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 5
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 6
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 7
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 8
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 9
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 10
        WHEN DATE(activity.date) BETWEEN ? AND ? THEN 11
    END AS date_bin,
    COALESCE(SUM(planned.distance), 0) AS total_planned_distance,
    COALESCE(SUM(planned.duration), 0) AS total_planned_duration,
    COALESCE(SUM(planned.vertical), 0) AS total_planned_vertical,
    COALESCE(SUM(completed.distance), 0) AS total_completed_distance,
    COALESCE(SUM(completed.duration), 0) AS total_completed_duration,
    COALESCE(SUM(completed.vertical), 0) AS total_completed_vertical
FROM activity
INNER JOIN planned ON planned.activity_uuid = activity.uuid
INNER JOIN completed ON completed.activity_uuid = activity.uuid
WHERE DATE(activity.date) BETWEEN ? AND ?
    AND activity.user_uuid = ?
GROUP BY date_bin
ORDER BY date_bin;