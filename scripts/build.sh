#!/bin/bash
# This will build the wails application and then run the setup db script
# run this in the scripts directory

cassidyAMD64="cassidy-amd64.app"
cassidyARM64="cassidy-arm64.app"


cd ..
wails build --clean --platform darwin/amd64,darwin/arm64,windows/amd64,windows/arm64,linux/amd64,linux/arm64
cd scripts

# Set databases for MAC distros
sudo bash setup_db.sh "$cassidyAMD64"
sudo bash setup_db.sh "$cassidyARM64"
cd ../build/bin


mkdir dist

# create the dmg using an npm package https://github.com/sindresorhus/create-dmg; its pretty nice
create-dmg "$cassidyAMD64" --dmg-title='cassidy-amd64'
mv 'cassidy 1.0.0.dmg' dist/'cassidy-amd64.dmg'
create-dmg "$cassidyARM64" --dmg-title='cassidy-arm64'
mv 'cassidy 1.0.0.dmg' dist/'cassidy-arm64.dmg'
