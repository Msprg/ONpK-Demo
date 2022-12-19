# Backend

## Background
Provides a CRUD based API written in Go. The API is designed to read and write into a MongoDB backend database. The API is utilised by the frontend web app (api-vote-reack). The frontend is developed using React and makes AJAX requests to this API.

## API endpoints
The API provides the following endpoints:
```
GET /platforms
GET /platforms/{name}
GET /platforms/{name}/vote
GET /ok
POST /platforms/{name}
DELETE /platforms/{name}
```

## API GETs with curl and jq
The API can be used to perform GETs with **curl** and **jq** like so:
```
curl -s http://localhost:8080/platforms | jq .
curl -s http://localhost:8080/platforms/{name} | jq .
curl -s http://localhost:8080/platforms/{name}/vote | jq .
```

## API POSTs with curl
The API can be used to perform POSTs with **curl** like so:
```
curl http://localhost:8080/platforms/{name} \
--header "Content-Type: application/json" \
--request POST \
--data-binary @- <<BODY
{
    "Usecase": "Container platform"",
    "Rank": 12,
    "Homepage": "https://openshift.com,
    "Download": "https://developers.redhat.com/products/codeready-containers/overview",
    "Votes": 0
}
BODY
```

## API DELETEs with curl and jq
The API can be used to perform DELETEs with **curl** like so:
```
curl -s -X DELETE http://localhost:8080/platforms/{name}
```

## Linux Compiling
The API can be compiled using the following commands:
```
#ensure to be in the same dir as the main.go file
go get -v -t -d ./...
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api
```

## API MongoDB Database
The API is designed to read/write into a MongoDB 4.2.x database. The MongoDB database should be setup and populated using the following commands performed within the mongo shell:
```
use platformsdb;

db.createUser({user: "admin",
pwd: "password",
roles:[{role: "userAdmin" , db:"platformsdb"}]
});

db.platforms.insert({"name" : "OpenShift", "codedetail" : { "usecase" : "Container platform", "rank" : 12, "homepage" : "https://openshift.com", "download" : "https://developers.redhat.com/products/codeready-containers/overview", "votes" : 0}});
db.platforms.insert({"name" : "Kubernetes", "codedetail" : { "usecase" : "Container orchestration platform ", "rank" : 38, "homepage" : "https://kubernetes.com", "download" : "https://kubernetes.io/docs/tasks/tools/install-minikube/", "votes" : 0}});
db.platforms.insert({"name" : "Rancher", "codedetail" : { "usecase" : "Container platform management ", "rank" : 50, "homepage" : "https://rancher.com/", "download" : "https://github.com/rancher/rancher", "votes" : 0}});


show collections;

db.platforms.find();
```
**Note**: The mongodb replicaset deployment used in the K8s cluster does not have authentication enabled and therefore doesnt require the admin user to be created.

## API Environment Vars
The API looks for the following defined environment variables:
```
MONGO_CONN_STR=mongodb://localhost:27017/platformsdb
MONGO_USERNAME=admin
MONGO_PASSWORD=password
```
**Note**: The mongodb replicaset deployment used in the K8s cluster does not have authentication enabled and therefore doesnt require the environment variables MONGO_USERNAME and MONGO_USERNAME to be configured.

## API Startup
The API can be started directly using the **main.go** file like so
```
MONGO_CONN_STR=mongodb://localhost:27017/platformsdb MONGO_USERNAME=admin MONGO_PASSWORD=password go run main.go
```
or by using the binary:
```
MONGO_CONN_STR=mongodb://localhost:27017/platformsdb MONGO_USERNAME=admin MONGO_PASSWORD=password ./api
```
