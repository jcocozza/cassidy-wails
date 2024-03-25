CREATE TABLE IF NOT EXISTS completed (
    activity_uuid TEXT PRIMARY KEY NOT NULL,
    distance REAL,
    distance_unit TEXT,
    duration REAL,
    vertical REAL,
    vertical_unit TEXT,
    FOREIGN KEY (activity_uuid) REFERENCES activity(uuid)
)