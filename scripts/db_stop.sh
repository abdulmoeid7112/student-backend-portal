#!/bin/bash

if [ "$(docker ps -q -f name=student-portal-mysql)" ]; then
  docker rm -f student-portal-mysql
fi

