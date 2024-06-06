-- get all activities for a user between two dates
SELECT DISTINCT activity.uuid,
        activity.date, activity.`order`, activity.name, activity.description, activity.notes, activity.is_race, activity.num_strides, activity.map,
        activity_type.id, activity_type.name,
        planned.distance, planned.distance_unit, planned.duration, planned.vertical, planned.vertical_unit,
        completed.distance, completed.distance_unit, completed.moving_duration, completed.elapsed_duration, completed.vertical, completed.vertical_unit
FROM activity
INNER JOIN activity_type ON activity.activity_type_id = activity_type.id
INNER JOIN planned ON activity.uuid = planned.activity_uuid
INNER JOIN completed ON activity.uuid = completed.activity_uuid
WHERE DATE(activity.date) BETWEEN ? AND ?
    AND activity.user_uuid = ?
ORDER BY activity.date, activity.`order`
