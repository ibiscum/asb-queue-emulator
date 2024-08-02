#!/bin/bash

# Script to run the integration tests

# Define the container name
container_name="rabbitmq"
project_root="$(pwd)"
test_go_file_path="${project_root}/tests/integration/main/main.go"

# Check if the container is already running
if [ "$(docker ps -q -f name=$container_name)" ]; then
    echo "Container $container_name is already running."
else
    # Run the RabbitMQ container in detached mode
    docker run -d --rm --name $container_name -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management

    # Check if the container was started successfully
    if [ $? -eq 0 ]; then
        echo "Container $container_name started successfully."
    else
        echo "Failed to start container $container_name."
        exit 1
    fi
    echo "Waiting for the rabbitmq container to start the engine"
    sleep 15
    echo "done!"
    sleep 1
fi

echo "Running integration tests"
go run $test_go_file_path

# To stop and remove the container when you're done with it
echo "Removing the ${container_name} container"
docker stop $container_name