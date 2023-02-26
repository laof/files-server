#!/bin/bash

go build -o ../public/fs
cd ../public 
chmod 777 fs
chmod +x fs