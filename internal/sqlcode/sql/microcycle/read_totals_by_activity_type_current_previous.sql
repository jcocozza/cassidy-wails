-- currently this doesn't work as well as it should. Two activities of the same type could easily have different units, but right now we just choose one to describe them all
/*
SELECT
    CASE
		WHEN DATE(activity.date) BETWEEN ? AND ? THEN "previous"
		WHEN DATE(activity.date) BETWEEN ? AND ? THEN "current"
		ELSE "ignore"
	END AS cycleLocation,
        activity_type.id, activity_type.name,
        SUM(planned.distance),SUM(planned.duration),SUM(planned.vertical), planned.distance_unit, planned.vertical_unit,
        SUM(completed.distance),SUM(completed.duration),SUM(completed.vertical), completed.distance_unit, completed.vertical_unit
FROM activity
INNER JOIN activity_type ON activity.activity_type_id  = activity_type.id
INNER JOIN planned ON planned.activity_uuid = activity.uuid
INNER JOIN completed ON completed.activity_uuid = activity.uuid
WHERE DATE(activity.date) BETWEEN ? and ?
    AND activity.user_uuid = ?
GROUP BY cycleLocation, activity.activity_type_id
ORDER BY cycleLocation DESC; -- ensure that previous comes first;
*/

-- first 12 ?'s are for previous, next 12 are for current, next 2 are for whole range, last is user uuid
SELECT
    activity_type.id,
    activity_type.name,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN planned.distance ELSE 0 END), 0) AS previous_planned_distance,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN planned.duration ELSE 0 END), 0) AS previous_planned_duration,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN planned.vertical ELSE 0 END), 0) AS previous_planned_vertical,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN completed.distance ELSE 0 END), 0) AS previous_completed_distance,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN completed.moving_duration ELSE 0 END), 0) AS previous_completed_duration,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN completed.vertical ELSE 0 END), 0) AS previous_completed_vertical,

    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN planned.distance ELSE 0 END), 0) AS current_planned_distance,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN planned.duration ELSE 0 END), 0) AS current_planned_duration,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN planned.vertical ELSE 0 END), 0) AS current_planned_vertical,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN completed.distance ELSE 0 END), 0) AS current_completed_distance,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN completed.moving_duration ELSE 0 END), 0) AS current_completed_duration,
    COALESCE(SUM(CASE WHEN DATE(activity.date) BETWEEN ? AND ? THEN completed.vertical ELSE 0 END), 0) AS current_completed_vertical
    --planned.distance_unit AS distance_unit,
    --planned.vertical_unit AS vertical_unit
FROM activity
INNER JOIN activity_type ON activity.activity_type_id = activity_type.id
INNER JOIN planned ON planned.activity_uuid = activity.uuid
INNER JOIN completed ON completed.activity_uuid = activity.uuid
WHERE DATE(activity.date) BETWEEN ? AND ?
    AND activity.user_uuid = ?
GROUP BY activity_type.id, activity_type.name
ORDER BY activity_type.id;