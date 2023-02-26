#!/bin/bash

go build fs
chmod 777 fs
chmod +x fs
mv -fb fs ../public/