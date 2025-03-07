#!/bin/bash


# Короче, базу данных можно держать в docker.
# И мы создаем контейнеры и монтируем их если хотим data persistence:
# Создаем волюм в Docker Desktop под назаванием postgresql

# $ docker run -d \
# 	--name postgres_container \
# 	-e POSTGRES_PASSWORD=okokokok \
# 	-e PGDATA=/var/lib/postgresql/data/pgdata \
# 	-volume postgresql:/var/lib/postgresql/data \
# 	postgres


# И если не хотим data persistence
# $ docker run -d \
# 	--name postgres_container \
# 	-e POSTGRES_PASSWORD=okokokok \
# 	postgres


# It is for the first time.
# Later we just run in manually from docker desktop.

docker stop postgres_container

# The -v flag with docker rm instructs Docker to remove the container and any unnamed volumes associated with it. 
# Named volumes will remain untouched.
docker rm -v postgres_container

sleep 1

# In this case two volumes will be created with names b05ee1bf...
# And they will be deleted when container deleted.
docker run -d \
	--name postgres_container \
	-e POSTGRES_PASSWORD=okokokok \
    -p 5432:5432 \
	postgres

