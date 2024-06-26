CREATE TABLE IF NOT EXISTS user (
    uuid TEXT PRIMARY KEY NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    units TEXT NOT NULL CHECK (units IN ('imperial', 'metric')),
    cycle_start TEXT NOT NULL,
    cycle_days INTEGER NOT NULL,
    initial_start_date TEXT
);
CREATE TABLE IF NOT EXISTS strava_token (
    user_uuid TEXT PRIMARY KEY NOT NULL,
    access_token TEXT NOT NULL,
    token_type TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    expiry TEXT NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);
CREATE TABLE IF NOT EXISTS persisted_user_login (
    -- ensure that there can only ever be 1 row
    id INTEGER PRIMARY KEY NOT NULL CHECK (id = 0),
    user_uuid TEXT,
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);
