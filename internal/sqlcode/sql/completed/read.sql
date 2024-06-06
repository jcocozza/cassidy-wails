SELECT completed.distance, completed.distance_unit, completed.moving_duration, completed.elapsed_duration, completed.vertical, completed.vertical_unit
FROM completed
WHERE activity_uuid = ?
