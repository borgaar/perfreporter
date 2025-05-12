#!/usr/bin/bash

touch data.json

DOCKER_OUTPUT=$(docker ps)

while [[ "$DOCKER_OUTPUT" == *"perfreporter"* ]]; do
  echo "$(date) : Polling sensor data..."
  sensors -j > data.json
  sleep 5
  DOCKER_OUTPUT=$(docker ps)
done

