#!/bin/bash


#=============================================================================
#=============================================================================

function shutdown_all(){

    container_id=template_go_react_golang
    # Get the list of volumes used by the container
    volumes=$(docker inspect $container_id --format '{{ range .Mounts }}{{ if eq .Type "volume" }}{{ .Name }} {{ end }}{{ end }}')
    # Stop and remove the container
    docker stop $container_id
    docker rm $container_id
    docker rmi $container_id

    # Remove the volumes
    for volume in $volumes; do
    docker volume rm $volume
    done



    container_id=template_go_react_pg
    # Get the list of volumes used by the container
    volumes=$(docker inspect $container_id --format '{{ range .Mounts }}{{ if eq .Type "volume" }}{{ .Name }} {{ end }}{{ end }}')
    # Stop and remove the container
    docker stop $container_id
    docker rm $container_id
    docker rmi $container_id

    # Remove the volumes
    for volume in $volumes; do
    docker volume rm $volume
    done


    container_id=template_go_react_pgadmin
    # Get the list of volumes used by the container
    volumes=$(docker inspect $container_id --format '{{ range .Mounts }}{{ if eq .Type "volume" }}{{ .Name }} {{ end }}{{ end }}')
    # Stop and remove the container
    docker stop $container_id
    docker rm $container_id
    docker rmi $container_id

    # Remove the volumes
    for volume in $volumes; do
    docker volume rm $volume
    done

    sleep 1

}



function all(){
    docker-compose -f docker-compose_pg.yml --profile all --project-name glang-ddd-react up -d
}


function db(){
    docker-compose -f docker-compose_pg.yml --profile db --project-name glang-ddd-react up -d
}




# Check if at least two arguments are provided
if [ $# -lt 1 ]; then
    echo "Error: Insufficient arguments provided."
    echo "Specify profile: all, db, backend"
    help
    exit 1
fi


# Read the command-line arguments
profile=$1




# Case statement to handle different commands
case $profile in
    all)
        all
        ;;
    db)
        db
        ;;        
    shutdown_all)
        shutdown_all
        ;;
    *)
        echo "Error: Invalid command."
        help
        exit 1
        ;;
esac




