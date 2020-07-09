#!/bin/bash

if [ "$(docker ps -q -f name=student-portal-mysql-db)" ]; then
  docker rm -f student-portal-mysql-db
fi

if [ "$(docker ps -q -f name=student-portal-mongo-db)" ]; then
  docker rm -f student-portal-mongo-db
fi

