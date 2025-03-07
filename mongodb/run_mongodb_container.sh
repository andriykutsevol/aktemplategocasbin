#!/bin/bash


# When using the default bridge network in Docker, 
# you cannot use container names to address other containers. 
# This feature is only available when you create a custom network. 
# The default bridge network does not provide automatic DNS resolution for container names.

# $ docker network create template_go_react_network
# $ docker network ls
# $ docker network inspect template_go_react_network

# Запускается уже на запущенном контейнере
# docker exec -it template_go_react_react_cui sh

# # I careated volumes from docker desktop
# # These mounpoints from container inspection, and we just created volumes for them,
# # because if we do not do, they will be created with names like b05ee1bf...
# docker run \
#     --name template_go_react_mongodb \
#     --network template_go_react_network \
#     --volume template_go_react_mongodb_volume:/data/db \
#     --volume template_go_react_mongodb_config_volume:/data/configdb \
#     -d -p 27017:27017 mongo:latest
# # In this case volumes will remain after container deleted.


docker stop template_go_react_mongodb
docker rm template_go_react_mongodb

# In this case two volumes will be created with names b05ee1bf...
# And they will be deleted when container deleted.
docker run \
    --name template_go_react_mongodb \
    --network template_go_react_network \
    -d -p 27017:27017 mongo:latest

# docker start mongodb
