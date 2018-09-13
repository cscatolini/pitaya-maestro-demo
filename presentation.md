# Pitaya Demo

Go to demo folder
```sh
cd $GOPATH/src/github.com/cscatolini/pitaya-maestro-demo

```

Start deps (nats, etcd and pitaya admin server)
```sh
make deps
```

Start connector
```sh
make run-connector
```

Start room
```sh
make run-room
```

Start pitaya-cli
```sh
pitaya-cli
```

List servers
```sh
curl http://localhost:8000/servers | jq
```

## pitaya-cli

Connect
```
connect localhost:3250
```

Call frontend handler
```
request connector.connectorhandler.getsessiondata
```

You can omit the server name if calling the same server type you are connected to
```
request connectorhandler.setsessiondata {"data":{"camila":false}}
```

Retrieve session data from the backend server
```
request room.roomhandler.getsessiondata
```

Set session data using the backend server and retrieve it using the frontend server
```
request room.roomhandler.setsessiondata {"data":{"camila":true}}
request connector.connectorhandler.getsessiondata
```

## pitaya-admin

Bind user id
```
request room.roomhandler.entry camilaid
```

Send push to user 
```
curl -H 'content-type: application/json' -X POST http://localhost:8000/user/push -d '{"uids": ["camilaid"], "frontendType": "connector", "message": "oi camila", "route": "msg.route"}'| jq
```

Kick user 
```
curl -H 'content-type: application/json' -X POST http://localhost:8000/user/kick -d '{"uids": ["camilaid"], "frontendType": "connector"}'| jq
```

Add some other backendserver
```sh 
make run-room
```

## jaeger

Open Jaeger UI in the broswer and have fun: http://localhost:16686/search

# Maestro Demo

Start minikube
```sh
minikube start --vm-driver hyperkit
```

Ensure kubectl current context is correct
```sh
kubectl config current-context
```

Go to maestro folder
```sh
cd $GOPATH/src/github.com/cscatolini/pitaya-maestro-demo

```

Start deps (psql and redis)
```sh
make deps drop migrate
```

Start api
```sh
make run-dev
```

Start worker
```sh
make work-dev
```

Get current IP address
```sh
ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | head -n 1
```

Edit scheduler MAESTRO_HOST_PORT env var to point to local ip address
```sh
edit manifests/scheduler-config-7.yaml
```

Create scheduler using maestro cli
```sh
maestro -c local create manifests/scheduler-config-7.yaml
```

Watch pods and scheduler status
```sh
watch kubectl get pods -n scheduler-name
watch maestro -c local status scheduler-name
```

Occupy some rooms
```sh
curl -XPUT `minikube ip`:`kubectl get svc <POD_NAME> -n scheduler-name --output='jsonpath={.spec.ports[0].nodePort}'`/status -d '{"status": "occupied"}'
```

Free some rooms
```sh
curl -XPUT `minikube ip`:`kubectl get svc <POD_NAME> -n scheduler-name --output='jsonpath={.spec.ports[0].nodePort}'`/status -d '{"status": "ready"}'
```
