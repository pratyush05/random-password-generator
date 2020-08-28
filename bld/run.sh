#!/bin/bash

# Switch to the directory where this script exist
export SCRIPT_DIR=$(dirname "$0")
export BUILD_DIR=$(cd ${SCRIPT_DIR} && pwd)

# Service constants
export SERVICE_DIR=$(cd ${BUILD_DIR}/.. && pwd)
export SERVICE_NAME="rpg"
export SERVICE_PORT="9999"

# Docker container constants
export CONTAINER_BASE_IMAGE="golang:latest"
export CONTAINER_HOME="/rpg"
export CONTAINER_HOST_NAME="hostname"
export CONTAINER_NAME="rpg"
export CONTAINER_PORT="443"

# Switch to the build directory to execute the script
cd ${BUILD_DIR}

# Generate private and public keys
source ./generate_cert.sh
if [[ $? -ne 0 ]]; then
    echo 'Error while generating certificates'
    exit 1
fi

# Remove already running 
docker container rm -f ${CONTAINER_NAME}

# Run the docekr container
docker run -d \
--privileged \
--name  ${CONTAINER_NAME} \
--hostname ${CONTAINER_HOST_NAME} \
-w ${CONTAINER_HOME} \
-v ${SERVICE_DIR}:${CONTAINER_HOME} \
-p ${SERVICE_PORT}:${CONTAINER_PORT} \
-e PORT=${CONTAINER_PORT} \
-e PRIVATE_KEY=${CONTAINER_HOME}/server.key \
-e PUBLIC_KEY=${CONTAINER_HOME}/server.crt \
${CONTAINER_BASE_IMAGE} \
go run ./src/main/main.go

# Display service logs
echo "Server running on port ${SERVICE_PORT}"
docker logs -f ${CONTAINER_NAME}
