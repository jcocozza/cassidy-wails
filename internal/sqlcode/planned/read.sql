SELECT planned.activity_uuid, planned.distance, planned.distance_unit, planned.duration, planned.vertical, planned.vertical_unit
FROM planned
WHERE activity_uuid = ?