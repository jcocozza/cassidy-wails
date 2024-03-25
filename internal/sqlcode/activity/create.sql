INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- The following queries need to be run to fully create an activity

--INSERT INTO planned (activity_uuid, distance, duration, vertical)
--VALUES (?, ?, ?, ?);

-- Optional
--INSERT INTO activity_equipment (activity_uuid, equipment_id, assigned_mileage)
--VALUES (?,?,?);

-- Optional
--INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id)
--VALUES (?,?,?)