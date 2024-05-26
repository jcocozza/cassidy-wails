CREATE TABLE IF NOT EXISTS completed (
    activity_uuid TEXT PRIMARY KEY NOT NULL,
    distance REAL,
    distance_unit TEXT,
    moving_duration REAL,
    elapsed_duration REAL,
    vertical REAL,
    vertical_unit TEXT,
    FOREIGN KEY (activity_uuid) REFERENCES activity(uuid)
)
