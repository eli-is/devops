# Basic HTTP server 
A basic http service that can be easily modified.

Include config from env or from hardcoded in the config file.


# Build
```
docker build -t  go-http-server .
```

# Run
```
docker  run -it -p 8000:8000 go-http-server
```

# Run his   

### api: 

/order
```
curl -XPOST localhost:8000/order -d '{
 "pizza-type": "margherita",
 "size": "family",
 "amount": 5
}'
```

/health
```
curl localhost:8000/health
```

# Install using helm
```
helm upgrade --install deom-app ./ --values=values.yaml --namespace=demo 
```

# Use localy on k8s
```
export POD_NAME=$(kubectl get pods --namespace demo -l "app.kubernetes.io/name=chart,app.kubernetes.io/instance=deom-app" -o jsonpath="{.items[0].metadata.name}")
export CONTAINER_PORT=$(kubectl get pod --namespace demo $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
kubectl --namespace demo port-forward $POD_NAME 8080:$CONTAINER_PORT
```