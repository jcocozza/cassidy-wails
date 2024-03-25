-- get all activity type subtypes for a passed list of activity uuids
SELECT activity_type_subtype.id, activity_type_subtype.activity_uuid,
		activity_type.id, activity_type.name,
		activity_subtype.id, activity_subtype.supertype_id, activity_subtype.name
FROM activity_type_subtype
INNER JOIN activity_type ON activity_type.id = activity_type_subtype.activity_type_id
INNER JOIN activity_subtype ON activity_subtype.id = activity_type_subtype.activity_subtype_id
WHERE activity_type_subtype.activity_uuid IN (%s);