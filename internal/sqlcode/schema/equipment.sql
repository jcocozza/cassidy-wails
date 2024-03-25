CREATE TABLE IF NOT EXISTS equipment (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_uuid TEXT NOT NULL,
    equipment_type_id INTEGER NOT NULL,
    name TEXT,
    brand TEXT,
    model TEXT,
    cost REAL,
    size TEXT,
    purchase_date TEXT,
    notes TEXT,
    mileage REAL,
    mileage_unit TEXT,
    is_retired BOOLEAN CHECK (is_retired IN (0, 1)),
    FOREIGN KEY (equipment_type_id) REFERENCES equipment_type(id),
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);