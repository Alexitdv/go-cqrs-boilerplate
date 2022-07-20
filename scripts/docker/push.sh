#!/bin/bash
docker login gitlab.dfarmer.ru:5050
docker push gitlab.dfarmer.ru:5050/dfarmer/idp:"${1:-latest}"