-- create users
-- main test user
INSERT INTO user (uuid, username, password, units ,cycle_start, cycle_days, initial_start_date) VALUES ("d4e0fbfe-c945-11ee-b8e2-325096b39f47", "test1@test.com", "test", "imperial", "Monday", 7, "");
-- other test users
INSERT INTO user (uuid, username, password, units ,cycle_start, cycle_days, initial_start_date) VALUES ("e790c0b8-c945-11ee-a42a-325096b39f47", "test2@test.com", "test", "metric", "Monday", 7, "");
INSERT INTO user (uuid, username, password, units ,cycle_start, cycle_days, initial_start_date) VALUES ("ee41cf88-c945-11ee-8e03-325096b39f47", "test3@test.com", "test", "imperial", "Wednesday", 10, "2024-01-01");

-- activity types
INSERT INTO activity_type (name) VALUES ("Run");      -- has id = 1
INSERT INTO activity_type (name) VALUES ("Bike");     -- has id = 2
INSERT INTO activity_type (name) VALUES ("Swim");     -- has id = 3
INSERT INTO activity_type (name) VALUES ("Hike");     -- has id = 4
INSERT INTO activity_type (name) VALUES ("Rest Day"); -- has id = 5
INSERT INTO activity_type (name) VALUES ("Strength"); -- has id = 6
INSERT INTO activity_type (name) VALUES ("Recovery"); -- has id = 7
-- activity subtypes
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Long", 1);      -- has id = 1
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Fartlek", 1);   -- has id = 2
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Tempo", 1);     -- has id = 3
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Track", 1);     -- has id = 4
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Intervals", 1); -- has id = 5
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Recovery", 1);  -- has id = 6
INSERT INTO activity_subtype (name, supertype_id) VALUES ("Velodrome", 2); -- has id = 7

-- equipment types
INSERT INTO equipment_type (name) VALUES ("Shoes"); -- has id = 1
INSERT INTO equipment_type (name) VALUES ("Bike");  -- has id = 2
INSERT INTO equipment_type (name) VALUES ("Vest");  -- has id = 3

-- user equipment
INSERT INTO equipment (equipment_type_id, user_uuid, name, brand, model, cost, size, purchase_date, notes, mileage, mileage_unit, is_retired) VALUES (1, "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "Kickers 1","Hoka","Clifton 5",120,12.5, "2024-01-01", "Notes", 337962.24, "mi", 0);
INSERT INTO equipment (equipment_type_id, user_uuid, name, brand, model, cost, size, purchase_date, notes, mileage, mileage_unit, is_retired) VALUES (1, "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "Kickers 2","Saucony","Kinvara 20",75,13, "2023-05-04", "Notes", 0, "mi", 0);
INSERT INTO equipment (equipment_type_id, user_uuid, name, brand, model, cost, size, purchase_date, notes, mileage, mileage_unit, is_retired) VALUES (1, "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "Kickers 3","Nike","Peg 12",130,11, "2020-07-19", "Notes", 10000, "km", 0);
INSERT INTO equipment (equipment_type_id, user_uuid, name, brand, model, cost, size, purchase_date, notes, mileage, mileage_unit, is_retired) VALUES (2, "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "Spinner"  ,"Trek","Domane",1500, "", "", "Nice Bike", 563270.4, "mi", 0);

-- create some test activities
-- empty activity imperial
INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides) VALUES("c6d91576-b65a-11ee-8546-325096b39f22", "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "2024-01-01", 1, "Workout Name", "Workout Description", "Notes", 1, 0, 0);
INSERT INTO planned   (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91576-b65a-11ee-8546-325096b39f22", 0, "mi", 0, 0, "ft");
INSERT INTO completed (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91576-b65a-11ee-8546-325096b39f22", 0, "mi", 0, 0, "ft");
-- empty activity metric
INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides) VALUES ("c6d91bf2-b65a-11ee-a67e-325096b39f50", "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "2024-01-01", 2, "Workout Name", "Workout Description", "Notes", 1, 0, 0);
INSERT INTO planned   (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91bf2-b65a-11ee-a67e-325096b39f50", 0, "km", 0, 0, "m");
INSERT INTO completed (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91bf2-b65a-11ee-a67e-325096b39f50", 0, "km", 0, 0, "m");
-- different day empty
-- empty activity imperial
INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides) VALUES ("c6d91576-b65a-11ee-8546-325096b39f23", "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "2024-01-05", 1, "Workout Name", "Workout Description", "Notes", 1, 0, 0);
INSERT INTO planned   (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91576-b65a-11ee-8546-325096b39f23", 0, "mi", 0, 0, "ft");
INSERT INTO completed (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91576-b65a-11ee-8546-325096b39f23", 0, "mi", 0, 0, "ft");
-- empty activity metric
INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides) VALUES ("c6d91bf2-b65a-11ee-a67e-325096b39f51", "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "2024-01-04", 1, "Workout Name", "Workout Description", "Notes", 1, 0, 0);
INSERT INTO planned   (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91bf2-b65a-11ee-a67e-325096b39f51", 0, "km", 0, 0, "m");
INSERT INTO completed (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91bf2-b65a-11ee-a67e-325096b39f51", 0, "km", 0, 0, "m");

-- full imperial activity
INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides) VALUES ("c6d91576-b65a-11ee-8546-325096b39f47", "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "2024-01-01", 3, "Workout Name", "Workout Description", "Notes", 1, 0, 0);
INSERT INTO planned   (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91576-b65a-11ee-8546-325096b39f47", 16093.44, "mi", 5000, 91.44, "ft");
INSERT INTO completed (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91576-b65a-11ee-8546-325096b39f47", 16093.44, "mi", 5000, 91.44, "ft");
INSERT INTO activity_equipment (activity_uuid, equipment_id, assigned_mileage, assigned_mileage_unit) VALUES ("c6d91576-b65a-11ee-8546-325096b39f47", 1, 8046.72, "mi");
INSERT INTO activity_equipment (activity_uuid, equipment_id, assigned_mileage, assigned_mileage_unit) VALUES ("c6d91576-b65a-11ee-8546-325096b39f47", 2, 8046.72, "mi");
INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id) VALUES ("c6d91576-b65a-11ee-8546-325096b39f47", 1, 1);
INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id) VALUES ("c6d91576-b65a-11ee-8546-325096b39f47", 1, 2);
INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id) VALUES ("c6d91576-b65a-11ee-8546-325096b39f47", 1, 3);

-- full metric activity
INSERT INTO activity (uuid, user_uuid, date, `order`, name, description, notes, activity_type_id, is_race, num_strides) VALUES ("c6d91850-b65a-11ee-b827-325096b39f47", "d4e0fbfe-c945-11ee-b8e2-325096b39f47", "2024-01-01", 4, "Workout Name", "Workout Description", "Notes", 1, 0, 0);
INSERT INTO planned   (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91850-b65a-11ee-b827-325096b39f47", 10000, "km", 5000, 300, "m");
INSERT INTO completed (activity_uuid, distance, distance_unit, duration, vertical, vertical_unit) VALUES("c6d91850-b65a-11ee-b827-325096b39f47", 10000, "km", 5000, 300, "m");
INSERT INTO activity_equipment (activity_uuid, equipment_id, assigned_mileage, assigned_mileage_unit) VALUES ("c6d91850-b65a-11ee-b827-325096b39f47", 1, 5, "km");
INSERT INTO activity_equipment (activity_uuid, equipment_id, assigned_mileage, assigned_mileage_unit) VALUES ("c6d91850-b65a-11ee-b827-325096b39f47", 2, 5, "km");
INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id) VALUES ("c6d91850-b65a-11ee-b827-325096b39f47", 1, 1);
INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id) VALUES ("c6d91850-b65a-11ee-b827-325096b39f47", 1, 2);
INSERT INTO activity_type_subtype (activity_uuid, activity_type_id, activity_subtype_id) VALUES ("c6d91850-b65a-11ee-b827-325096b39f47", 1, 3);