-- TODO: make this not so disgusting and horrific to look at
SELECT
    cycleLocation,
    COALESCE(SUM(total_planned_distance), 0) AS total_planned_distance,
    COALESCE(SUM(total_planned_duration), 0) AS total_planned_duration,
    COALESCE(SUM(total_planned_vertical), 0) AS total_planned_vertical,
    COALESCE(SUM(total_completed_distance), 0) AS total_completed_distance,
    COALESCE(SUM(total_completed_duration), 0) AS total_completed_duration,
    COALESCE(SUM(total_completed_vertical), 0) AS total_completed_vertical
FROM (
    SELECT
        CASE
			WHEN DATE(activity.date) BETWEEN ? AND ? THEN "previous"
			WHEN DATE(activity.date) BETWEEN ? AND ? THEN "current"
            ELSE 'ignore'
        END AS cycleLocation,
        COALESCE(SUM(planned.distance), 0) AS total_planned_distance,
        COALESCE(SUM(planned.duration), 0) AS total_planned_duration,
        COALESCE(SUM(planned.vertical), 0) AS total_planned_vertical,
        COALESCE(SUM(completed.distance), 0) AS total_completed_distance,
        COALESCE(SUM(completed.moving_duration), 0) AS total_completed_duration,
        COALESCE(SUM(completed.vertical), 0) AS total_completed_vertical
    FROM planned
    INNER JOIN activity ON planned.activity_uuid = activity.uuid
    INNER JOIN completed ON completed.activity_uuid = activity.uuid
    WHERE (DATE(activity.date) BETWEEN ? AND ?)
        AND activity.user_uuid = ?
    GROUP BY cycleLocation
    -- ensures that we will always get two rows in the event of null
    UNION ALL
    SELECT 'previous' AS cycleLocation, 0, 0, 0, 0, 0, 0
    UNION ALL
    SELECT 'current' AS cycleLocation, 0, 0, 0, 0, 0, 0) AS subquery
GROUP BY cycleLocation
ORDER BY cycleLocation DESC;
