CREATE TABLE IF NOT EXISTS activity_subtype (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    supertype_id INTEGER NOT NULL,
    name TEXT,
    FOREIGN KEY (supertype_id) REFERENCES activity_type(id)
);