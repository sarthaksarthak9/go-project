# Building Containerized Microservices in Golang | Dockerize and Deploy to Kubernetes using Helm

Here, in this project we defines a simple web server that serves HTTP requests using the net/http package and the Gorilla Mux router. The server has three main endpoints: /health, /, and /details, each of which is associated with a corresponding handler function.

#### Three handler functions are defined for different endpoints:
1. healthHandler: Responds to the /health endpoint with a JSON response indicating the application's health status and a timestamp.
2. rootHandler: Serves the root (/) endpoint with a simple text response confirming that the application is up and running.
3. detailsHandler: Handles the /details endpoint by fetching and displaying the hostname and IP address of the server.

## Getting Started

### Prerequisite

1. Go installed in your system
2. docker 
3. helm
4. Running kubernetes cluster(MiniKube)


### Installation

1. Clone this repository to your local machine.
```
git clone https://github.com/sarthaksarthak9/go-project.git
cd go-project
```
2. Install dependencies
```
go build main.go
./main.exe
```

3. Run docker-compose file

```
docker-compose build
docker-compose up
```




