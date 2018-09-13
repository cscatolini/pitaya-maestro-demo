Start minikube

```sh
minikube start --vm-driver hyperkit
```

Ensure kubectl current context is correct
```sh
kubectl config current-context
```

Create deps
```sh
kubectl apply -f manifests/deps
```

Configure maestro db
```sh
cd scripts
sh ./configure-maestro-db.sh
cd -
```

Create maestro
```sh
kubectl apply -f manifests/maestro
```

Create pitaya servers
```sh
kubectl apply -f manifests/pitaya
```

Get running servers
```sh
curl `minikube ip`:`kubectl get service pitaya-admin --output='jsonpath={.spec.ports[0].nodePort}'`/servers | jq
```

Get pitaya connector URL
```sh
echo Pitaya Connector URL: `minikube ip`:`kubectl get service pitaya-connector --output='jsonpath={.spec.ports[0].nodePort}'`
```

Allow maestro to create namespaces
```sh
kubectl create clusterrolebinding add-on-cluster-admin \
  --clusterrole=cluster-admin \
  --serviceaccount=default:default
```

Create scheduler
```sh
curl -i -X POST http://`minikube ip`:`kubectl get service maestro-api --output='jsonpath={.spec.ports[0].nodePort}'`/scheduler \
  -H "Content-Type: application/json" \
  --data-binary "@manifests/schedulers/room-scheduler.json"
```

Delete scheduler
```sh
curl -X DELETE http://`minikube ip`:`kubectl get service maestro-api --output='jsonpath={.spec.ports[0].nodePort}'`/scheduler/pitaya-room-scheduler
```




curl -XPUT $(minikube ip):`kubectl get svc scheduler-name-0cbf40d6 -n scheduler-name --output='jsonpath={.spec.ports[0].nodePort}'`/status -d '{"status": "ready"}'

kubectl get svc -n scheduler-name