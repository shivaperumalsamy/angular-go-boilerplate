#!/bin/sh

ng serve &
# gin --port 4201 --path . --build ./src/server/ --i --all &
make build_run &

wait