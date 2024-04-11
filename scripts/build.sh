#!/bin/bash
# This will build the wails application and then run the setup db script
# run this in the scripts directory
cd ..
wails build --clean #-platform darwin/amd64
cd scripts
sudo bash setup_db.sh
cd ../build/bin
tar -cvf cassidy.app.tar cassidy.app
mv cassidy.app.tar ~/Downloads