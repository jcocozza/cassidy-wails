#!/bin/bash

# This script sets up the .db file that is distributed with the application binary
# It requires that the sqlite3 command is available

CassidyDB=".cassidy.db"
MACOS_BUILD_DB="../build/bin/cassidy.app/Contents/Resources/$CassidyDB"
schema_dir="../internal/sqlcode/schema"

# Check if sqlite3 command is available
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
# Check if the database file exists and remove it if it does
if [ -f "$MACOS_BUILD_DB" ]; then
    rm "$MACOS_BUILD_DB" || { echo "Error: Failed to remove existing database file '$MACOS_BUILD_DB'."; exit 1; }
fi
# Create a new database file
touch "$MACOS_BUILD_DB" || { echo "Error: Failed to create database file '$MACOS_BUILD_DB'."; exit 1; }

# Setup the schema
for file in "$schema_dir"/*; do
    if [ -f "$file" ]; then
        create_schema "$file" "$MACOS_BUILD_DB"
    fi
done