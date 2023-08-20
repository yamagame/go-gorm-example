#!/bin/bash
export CONTAINER_NAME=go-gorm-example-db
export DB_PASS=pass
export DB_NAME=go-gorm-example
export DB_PORT=3306
docker run --name ${CONTAINER_NAME} -e MYSQL_ROOT_PASSWORD=${DB_PASS} -e MYSQL_DATABASE=${DB_NAME} -p ${DB_PORT}:3306 -d mysql:8.0.32
