#!/bin/bash

MINIKUBE_IP=`minikube ip`
PSQL_PORT=`kubectl get service postgres --output='jsonpath={.spec.ports[0].nodePort}'`
psql -d postgres -h $MINIKUBE_IP -p $PSQL_PORT -U postgres -f ./configure-maestro-db.sql

# TODO: ugly :(
cd $GOPATH/src/github.com/topfreegames/maestro
MAESTRO_EXTENSIONS_PG_HOST=$MINIKUBE_IP MAESTRO_EXTENSIONS_PG_PORT=$PSQL_PORT MAESTRO_EXTENSIONS_PG_USER=maestro MAESTRO_EXTENSIONS_PG_DATABASE=maestro PGPASSWORD=""  make migrate
cd -
