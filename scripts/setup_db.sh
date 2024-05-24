#!/bin/bash

# This script sets up the .db file that is distributed with the application binary
# It requires that the sqlite3 command is available

CassidyDB=".cassidy.db"
schema_dir="../internal/sqlcode/schema"

DATABASE_FILE_PATH="../build/$CassidyDB"

# Check if sqlite3 is installed
if ! command -v sqlite3 &> /dev/null; then
    echo "Error: sqlite3 command not found. Please make sure sqlite3 is installed." >&2
    exit 1
fi

# function to run create schema
create_schema () {
    schema_filepath="$1"
    database_filepath="$2"

    echo "running schema creation for: $schema_filepath"
    sqlite3 "$database_filepath" <<EOF
.read "$schema_filepath"
.exit
EOF
}

# Check if schema directory exists
if [ ! -d "$schema_dir" ]; then
    echo "Error: Schema directory '$schema_dir' not found." >&2
    exit 1
fi
# Check if database file exists, remove it if it does
if [ -f "$DATABASE_FILE_PATH" ]; then
    rm "$DATABASE_FILE_PATH" || { echo "Error: Failed to remove existing database file '$DATABASE_FILE_PATH'."; exit 1; }
fi
# Create a new database file
touch "$DATABASE_FILE_PATH" || { echo "Error: Failed to create database file '$DATABASE_FILE_PATH'."; exit 1; }
# Need to let the user write to the file
chmod a+w "$DATABASE_FILE_PATH" || { echo "Error: Failed to set write privileges for '$DATABASE_FILE_PATH'."; exit 1; }

# Set up schema
for file in "$schema_dir"/*; do
    if [ -f "$file" ]; then
        create_schema "$file" "$DATABASE_FILE_PATH"
    fi
done

# Insert necessary data into the database
create_schema "$schema_dir/init/init.sql" "$DATABASE_FILE_PATH"
