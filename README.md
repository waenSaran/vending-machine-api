# Vending Machine Service API
- using `nodemon` for real-time update
```
nodemon --exec go run main.go --signal SIGTERM
```
- docker stop all running container
```
docker stop $(docker ps -a -q)
```
- Remove unused images to prevent conflicts:
```
docker image prune -f
```
- local build docker
```
docker build -t api .
```
- after build Dockerfile using the below command to start the container
```
docker run -d -p 8080:8080 --env-file .env api
// or
docker compose up --build
```

- This project will push image to Google Artifact Registry and deploy on CloudRun (ref: https://docs.mikelopster.dev/c/goapi-essential/chapter-9/cloudrun)
```
# 1. Build to Artifact Registry
docker build -t asia-southeast1-docker.pkg.dev/saranya-445610/vending-machine/v1.0.0 .

# if build from Macbook M1 / M2 (ARM)
docker build -t asia-southeast1-docker.pkg.dev/saranya-445610/vending-machine/v1.0.0 . --platform linux/amd64

# 2. push image to Cloud
docker push asia-southeast1-docker.pkg.dev/saranya-445610/vending-machine/v1.0.0
```