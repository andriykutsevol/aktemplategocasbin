#!/bin/bash

docker build -t template_go_react_golang .

# docker run --name template_go_react_golang --network template_go_react_network -p 8080:8080 template_go_react_golang
# # options.Client().ApplyURI("mongodb://template_go_react_mongodb:27017")



docker run \
    --name template_go_react_golang \
    --network template_go_react_network \
    --env MONGOURI=mongodb://template_go_react_mongodb:27017 \
    -d -p 8080:8080 template_go_react_golang
# uri, ok := os.LookupEnv("MONGOURI")
# clientOptions := options.Client().ApplyURI(uri)



#docker run --name golang_ddd_template_container golang_ddd_template

# docker run --rm --name my-app --link mongodb:mongodb --link redis:redis -p 8080:8080 my-golang-app
#docker run --name golang_ddd_template_container --link mongodbM:mongodb -p 8080:8080 golang_ddd_template


# mongodb: 
    #This is the name of the container you want to link to. In this case, 
    #it refers to the MongoDB container running on your Docker host.

# :mongodb:
    # This specifies an alias that Docker assigns to the linked container. 
    # The alias allows the container you are running (my-app in this case) to refer to the linked container (mongodb) 
    # by a simpler name (mongodb) within its own network namespace

# -p 8080:8080 maps port 8080 on your host to port 8080 on the container 
# (assuming your Go application listens on port 8080).