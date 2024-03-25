CREATE TABLE IF NOT EXISTS activity_type_subtype (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    activity_uuid TEXT NOT NULL,
    activity_type_id INTEGER,
    activity_subtype_id INTEGER,
    FOREIGN KEY (activity_uuid) REFERENCES activity(uuid),
    FOREIGN KEY (activity_type_id) REFERENCES activity_type(id),
    FOREIGN KEY (activity_subtype_id) REFERENCES activity_subtype(id)
);