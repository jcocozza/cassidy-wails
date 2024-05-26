-- currently this doesn't work as well as it should. Two activities of the same type could easily have different units, but right now we just choose one to describe them all
SELECT activity_type.id, activity_type.name,
        SUM(planned.distance),SUM(planned.duration),SUM(planned.vertical), planned.distance_unit, planned.vertical_unit,
        SUM(completed.distance),SUM(completed.moving_duration),SUM(completed.vertical), completed.distance_unit, completed.vertical_unit
FROM activity
INNER JOIN activity_type ON activity.activity_type_id  = activity_type.id
INNER JOIN planned ON planned.activity_uuid = activity.uuid
INNER JOIN completed ON completed.activity_uuid = activity.uuid
WHERE DATE(activity.date) BETWEEN ? and ?
    AND activity.user_uuid = ?
GROUP BY activity.activity_type_id;