#!/usr/bin/env bash

if [ ! -f ./bin/grule-health-alert-microservice ]; then
    echo "Binaries not found, building first..."
    ./scripts/build.sh
fi

echo "Press CTRL+C to exit..."
./bin/grule-health-alert-microservice
