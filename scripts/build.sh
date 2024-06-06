#!/bin/bash
# This will build the wails application, set up the database, and package things up

cassidyMACOSAMD64="cassidy-amd64.app"
cassidyMACOSARM64="cassidy-arm64.app"

cd ../build/bin
rm -rf *
cd ../..

wails build --clean --platform darwin/amd64,darwin/arm64
cd scripts
# creates a .db file in the /build directory
sudo bash setup_db.sh
DATABASE_FILE_PATH="../../build/.cassidy.db"
cd ../build/bin
mkdir dist

# package macos distributions
mkdir dist/macos
# send copies of the database to each of the distros
echo $(pwd)
cp "$DATABASE_FILE_PATH" "../../build/bin/cassidy-amd64.app/Contents/Resources/.cassidy.db"
cp "$DATABASE_FILE_PATH" "../../build/bin/cassidy-arm64.app/Contents/Resources/.cassidy.db"
create-dmg "$cassidyMACOSAMD64" --dmg-title='cassidy-amd64'
mv 'cassidy 1.0.0.dmg' dist/macos/'cassidy-amd64.dmg'
create-dmg "$cassidyMACOSARM64" --dmg-title='cassidy-arm64'
mv 'cassidy 1.0.0.dmg' dist/macos/'cassidy-arm64.dmg'

# package windows distributions
cp "$DATABASE_FILE_PATH" "../../build/windows/.cassidy.db"
cd ../..
wails build -nsis --platform windows/amd64,windows/arm64
