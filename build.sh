#!/bin/bash

go build -o fs
chmod 777 fs
chmod +x fs
mv -fb fs ../public/