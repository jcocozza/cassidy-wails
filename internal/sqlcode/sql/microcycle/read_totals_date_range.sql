SELECT
    COALESCE(SUM(planned.distance) / ?, 0) AS total_planned_distance,
    COALESCE(SUM(planned.duration) / ?, 0) AS total_planned_duration,
    COALESCE(SUM(planned.vertical) / ?, 0) AS total_planned_vertical,
    COALESCE(SUM(completed.distance) / ?, 0) AS total_completed_distance,
    COALESCE(SUM(completed.moving_duration) / ?, 0) AS total_completed_duration,
    COALESCE(SUM(completed.vertical) / ?, 0) AS total_completed_vertical
FROM planned
RIGHT JOIN activity ON activity.uuid = planned.activity_uuid AND activity.date BETWEEN ? AND ? AND activity.user_uuid = ?
LEFT JOIN completed ON completed.activity_uuid = planned.activity_uuid;