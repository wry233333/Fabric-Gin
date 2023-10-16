#!/bin/bash
# This script is used to start the application
cd fixtures && docker-compose up -d
cd ../ && ./fabric-sdk-go