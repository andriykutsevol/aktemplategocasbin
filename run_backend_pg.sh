#!/bin/bash

# Use this when db is already executed.




# If I run service in docker-compose on the network:
# networks:
#   template_go_react_network:
#     driver: bridge

# How can I reach it from the host?

    # To access a service running in a Docker container from your host machine, 
    # you'll need to expose the service's port to the host. 
    # This involves mapping the container's port to a port on the host machine in your Docker Compose configuration.


workdir_path=$(pwd)

cd $workdir_path/golang-ddd-template/cmd/app
export DBTYPE=pg
export PGCASBINURI=postgres://postgres:okokokokd@localhost:5432/casbin
export PGWEATHERURI=postgres://postgres:okokokokd@localhost:5432/weather
go run .