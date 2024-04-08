-- The datatbase initialization that includes the various types etc...

-- activity types
INSERT INTO activity_type (name) VALUES ("Run");      -- has id = 1
INSERT INTO activity_type (name) VALUES ("Bike");     -- has id = 2
INSERT INTO activity_type (name) VALUES ("Swim");     -- has id = 3
INSERT INTO activity_type (name) VALUES ("Hike");     -- has id = 4
INSERT INTO activity_type (name) VALUES ("Rest Day"); -- has id = 5
INSERT INTO activity_type (name) VALUES ("Strength"); -- has id = 6
INSERT INTO activity_type (name) VALUES ("Recovery"); -- has id = 7
INSERT INTO activity_type (name) VALUES ("Other");    -- has id = 8
-- activity subtypes
-- run subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Long", 1);      -- has id = 1
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Fartlek", 1);   -- has id = 2
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Tempo", 1);     -- has id = 3
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Track", 1);     -- has id = 4
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Intervals", 1); -- has id = 5
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery", 1);  -- has id = 6
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Indoor", 1);    -- has id = 7
-- bike subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Velodrome", 2); -- has id = 8
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Long Ride", 2); -- has id = 9
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery", 2);  -- has id = 10
-- swim subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Drills", 3);    -- has id = 11
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Open Water", 3);-- has id = 12
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery", 3);  -- has id = 13

-- equipment types
INSERT INTO equipment_type (name) VALUES ("Shoes"); -- has id = 1
INSERT INTO equipment_type (name) VALUES ("Bike");  -- has id = 2
INSERT INTO equipment_type (name) VALUES ("Vest");  -- has id = 3