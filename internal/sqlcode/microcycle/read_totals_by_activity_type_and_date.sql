SELECT activity_type.id, activity_type.name, activity.date,
    SUM(planned.distance), SUM(planned.duration), SUM(planned.vertical), planned.distance_unit, planned.vertical_unit,
    SUM(completed.distance), SUM(completed.duration), SUM(completed.vertical), completed.distance_unit, completed.vertical_unit
FROM planned
INNER JOIN activity ON planned.activity_uuid = activity.uuid
INNER JOIN completed ON completed.activity_uuid = activity.uuid
INNER JOIN activity_type ON activity.activity_type_id = activity_type.id
WHERE DATE(activity.date) BETWEEN ? AND ?
    AND activity.user_uuid = ?
GROUP BY activity.date, activity_type.id
ORDER BY activity.date;