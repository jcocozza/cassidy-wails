#!/bin/bash
# This will build the wails application, set up the database, and package things up

cassidy="cassidy.app"

cd ../build/bin
rm -rf *
cd ../..

wails build --clean --platform darwin/arm64
mv "build/bin/cassidy.app" "build/bin/$cassidy"
cd scripts
# creates a .db file in the /build directory
sudo bash setup_db.sh
DATABASE_FILE_PATH="../build/.cassidy.db"

echo $(pwd)
cp "$DATABASE_FILE_PATH" "../build/bin/$cassidy/Contents/Resources/.cassidy.db"

cd ..
wails dev
