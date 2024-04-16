-- The datatbase initialization that includes the various types etc...

-- activity types
INSERT INTO activity_type (name) VALUES ("Run");
INSERT INTO activity_type (name) VALUES ("Bike");
INSERT INTO activity_type (name) VALUES ("Swim");
INSERT INTO activity_type (name) VALUES ("Hike");
INSERT INTO activity_type (name) VALUES ("Rest Day");
INSERT INTO activity_type (name) VALUES ("Strength");
INSERT INTO activity_type (name) VALUES ("Other");
-- activity subtypes
-- run subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Long", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Fartlek", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Tempo", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Track", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Intervals", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery Run", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Indoor", 1);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Trails", 1);
-- bike subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Velodrome", 2);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Long Ride", 2);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery Ride", 2);
-- swim subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Drills", 3);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Open Water", 3);
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery Swim", 3);

-- equipment types
INSERT INTO equipment_type (name) VALUES ("Shoes");
INSERT INTO equipment_type (name) VALUES ("Bike");
INSERT INTO equipment_type (name) VALUES ("Vest");