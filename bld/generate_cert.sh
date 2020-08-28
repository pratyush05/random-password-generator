#!/bin/bash

# Check if certificates already exists
if [[ -f ${SERVICE_DIR}/server.key && -f ${SERVICE_DIR}/server.key ]]; then
    return 0
fi

# Generate and save the self-signed public and private key
openssl genrsa -out ${SERVICE_DIR}/server.key 2048
openssl req -new -x509 -sha256 -key ${SERVICE_DIR}/server.key -out ${SERVICE_DIR}/server.crt -days 3650