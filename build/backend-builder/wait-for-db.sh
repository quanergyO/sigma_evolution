#!/bin/bash

HOST="$1"            
PORT="$2"            
DB_NAME="postgres"    
USERNAME="postgres"   
PASSWORD="$DB_PASSWORD" 

shift 2
COMMAND=("$@")

wait_for_db() {
    echo "Waiting for PostgreSQL to be ready..."
    until PGPASSWORD="$PASSWORD" psql -h "$HOST" -p "$PORT" -U "$USERNAME" -d "$DB_NAME" -c '\l' > /dev/null 2>&1; do
        echo "PostgreSQL is unavailable - sleeping..."
        sleep 2
    done
    echo "PostgreSQL is up and running!"
}

wait_for_db

echo "Executing command: ${COMMAND[@]}"
exec "${COMMAND[@]}"