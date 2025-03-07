#!/bin/bash


docker stop pgadmin-container
docker rm -v pgadmin-container

# localhost:5050
# andriyk@domain.com
# okokokok

docker run --name pgadmin-container -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=andriyk@domain.com -e PGADMIN_DEFAULT_PASSWORD=okokokok -d dpage/pgadmin4

# В веб интерфейс заходим
    # andriyk@domain.com
    # okokokok


# Connect admin to postgres
# Получаем IP контейнера бд
# $ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres_container
# 172.17.0.2
# Порт по умолчанию мы не меняли

# В веб интерфейсе нажимаем Add New server

# Name:postgres_container
# Connection
#     Host name: 172.17.0.2
#     Username: postgres  (or another name if you changed i)
#     Password: okokokok  (но это не тот пароль что для самой админки, это пароль который мы указали при создании контейнера бд)

# !!! Кстати если запускаем из docker-compose на одтельной сети, то Host Name - это имя сервайса в .yml файле.

