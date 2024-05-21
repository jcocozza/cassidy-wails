CREATE TABLE IF NOT EXISTS activity (
    uuid TEXT PRIMARY KEY NOT NULL,
    user_uuid TEXT NOT NULL,
    date TEXT,
    `order` INTEGER,
    name TEXT,
    description TEXT,
    notes TEXT,
    activity_type_id INTEGER,
    is_race BOOLEAN CHECK (is_race IN (0, 1)),
    num_strides INTEGER,
    map TEXT,
    -- is_template BOOLEAN NOT NULL CHECK (is_template IN (0, 1))
    FOREIGN KEY (activity_type_id) REFERENCES activity_type(id),
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);

-- when an activity is deleted, we need to remove stuff from other tables as well
CREATE TRIGGER activity_delete
AFTER DELETE ON activity
FOR EACH ROW
BEGIN
    DELETE FROM planned
    WHERE planned.activity_uuid = OLD.uuid;

    DELETE FROM completed
    WHERE completed.activity_uuid = OLD.uuid;

    DELETE FROM activity_equipment
    WHERE activity_equipment.activity_uuid = OLD.uuid;

    DELETE FROM activity_type_subtype
    WHERE activity_type_subtype.activity_uuid = OLD.uuid;
END;

-- If we change the activity type id, then all previous subtypes need to be disassociated with that activity
CREATE TRIGGER update_activity_type
AFTER UPDATE OF activity_type_id ON activity WHEN NEW.activity_type_id <> OLD.activity_type_id
BEGIN
    DELETE FROM activity_type_subtype
    WHERE activity_type_subtype.activity_uuid = OLD.uuid;
END;