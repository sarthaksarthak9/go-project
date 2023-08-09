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
![1](https://github.com/sarthaksarthak9/go-project/assets/122533767/b0d9c721-d1f6-4fe7-a051-d8d52cd48d3a)



// To check everything working fine

3. Run docker-compose file

```
docker-compose build
docker-compose up -d // to run the container
docker-compose ps // to check the status of the containers
docker-compose down // stop and remove the containers 
```

4. Check your minikube IP
```
minikube ip
```

![6](https://github.com/sarthaksarthak9/go-project/assets/122533767/75e30877-6a59-4d29-808b-bddf05cab6d3)



6. Helm
```
helm template go-app/ // to check the template
helm install web go-app/ // to install the chart in go-app directory
kubectl get po, svc // it will list down all pods and services running
```
![2](https://github.com/sarthaksarthak9/go-project/assets/122533767/ee7549ec-24ad-4671-a841-e81436098ce9)

so, ***31579*** is the port in which serrvices are listening 

![3](https://github.com/sarthaksarthak9/go-project/assets/122533767/6737800b-ec63-4bea-ba79-f2a5b64ad1d4)

<p>
To have the details type ***localhost/31579/details*** on your browser

![4](https://github.com/sarthaksarthak9/go-project/assets/122533767/49f0ae4c-e4b5-43ad-a5f3-926366bd8ca1)

</p>


To have the health information type ***localhost/31579/health*** on your browser

![5](https://github.com/sarthaksarthak9/go-project/assets/122533767/8c6aab47-e83d-477e-aec2-a53f47e10859)








