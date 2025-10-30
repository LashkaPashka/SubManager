#!/bin/bash

export MIGRATION_DIR=$migration_dir
export MIGRATION_DSN=$migration_dsn
export CONFIG=$config_path

sleep 2 && ./goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v

./subManager