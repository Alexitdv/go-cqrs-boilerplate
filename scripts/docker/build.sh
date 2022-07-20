#!/bin/bash
docker build -t gitlab.dfarmer.ru:5050/dfarmer/idp:"${1:-latest}"   .
