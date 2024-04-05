#!/bin/bash
# This will build the wails application and then run the setup db script
# run this in the scripts directory
cd ..
wails build --clean
cd scripts
sudo bash setup_db.sh