# Vending Machine Service API
- using `nodemon` for real-time update
```
nodemon --exec go run main.go --signal SIGTERM
```

- after build Dockerfile using the below command to start the container
```
docker run -d -p 8080:8080 --env-file .env api
// or
docker compose up
```