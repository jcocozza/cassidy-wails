SELECT activity_type.id, activity_type.name, activity_subtype.id, activity_subtype.name
FROM activity_type
LEFT JOIN activity_subtype ON activity_type.id = activity_subtype.supertype_id;