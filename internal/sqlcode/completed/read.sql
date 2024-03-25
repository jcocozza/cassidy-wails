SELECT completed.distance, completed.distance_unit, completed.duration, completed.vertical, completed.vertical_unit
FROM completed
WHERE activity_uuid = ?