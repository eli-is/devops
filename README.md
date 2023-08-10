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

# Run 
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