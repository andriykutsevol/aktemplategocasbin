#!/bin/bash


# Use this when db is already executed.


# Короче, базу данных можно держать в docker.
# И мы создаем контейнеры и монтируем их если хотим data persistence:
# docker run \
#     --name template_go_react_mongodb \
#     --network template_go_react_network \
#     --volume template_go_react_mongodb_volume:/data/db \
#     --volume template_go_react_mongodb_config_volume:/data/configdb \
#     -d -p 27017:27017 mongo:latest

# И если не хотим data persistence
# docker run \
#     --name template_go_react_mongodb \
#     --network template_go_react_network \
#     -d -p 27017:27017 mongo:latest
# Тогда будут созданы эфемерные пустые volumes (и удалены при удалении контейнера)





workdir_path=$(pwd)

cd ./mongodb
./run_mongodb_container.sh

cd $workdir_path/golang-ddd-template/cmd/app
export MONGOURI=mongodb://localhost:27017
go run .




