CREATE TABLE IF NOT EXISTS activity_equipment (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    activity_uuid TEXT NOT NULL,
    equipment_id INTEGER NOT NULL,
    assigned_mileage REAL NOT NULL, -- mileage for the piece of equipment for the given activity
    assigned_mileage_unit TEXT NOT NULL,
    FOREIGN KEY (activity_uuid) REFERENCES activity(uuid),
    FOREIGN KEY (equipment_id) REFERENCES equipment(id)
);

-- when equipment gets removed from an activity, it also need to decrease in mileage
CREATE TRIGGER remove_mileage
AFTER DELETE ON activity_equipment
FOR EACH ROW
BEGIN
    --DECLARE deleted_id INT;
    --DECLARE deleted_mileage REAL;
    --SET deleted_id = OLD.equipment_id;
    --SET deleted_mileage = OLD.assigned_mileage;

    UPDATE equipment
    SET mileage = mileage - OLD.assigned_mileage
    WHERE equipment.id = OLD.equipment_id;
END;
-- when equipment gets added to an activity, it also need to increase in mileage
CREATE TRIGGER add_mileage
AFTER INSERT ON activity_equipment
FOR EACH ROW
BEGIN
    --DECLARE inserted_id INT;
    --DECLARE new_mileage REAL;
    --SET inserted_id = NEW.equipment_id;
    --SET new_mileage = NEW.assigned_mileage;

    UPDATE equipment
    SET mileage = mileage + NEW.assigned_mileage
    WHERE equipment.id = NEW.equipment_id;
END;

-- when the amount of mileage of an activity is updated, it needs to be reflected in the total equipment mileage
CREATE TRIGGER update_mileage
AFTER UPDATE ON activity_equipment
FOR EACH ROW
BEGIN
    UPDATE equipment
    SET mileage = mileage + (NEW.assigned_mileage - OLD.assigned_mileage)
    WHERE equipment.id = NEW.equipment_id;
END;