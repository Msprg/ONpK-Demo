package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	//including gorilla mux and handlers packages for HTTP routing and CORS support

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	//connections to mongo
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_DB = "platformsdb"
const MONGO_COLLECTION = "platforms"

// const MONGO_DEFAULT_CONN_STR = "mongodb://mongo-0.mongo:27017,mongo-1.mongo:27017,mongo-2.mongo:27017"

const MONGO_DEFAULT_CONN_STR = "mongodb://localhost:27017/?replicaSet=rs0&connect=direct"
const MONGO_DEFAULT_USERNAME = "admin"
const MONGO_DEFAULT_PASSWORD = "password"
const LISTEN_PORT = "9080"
const VERSION = "1.0.1"

type codedetail struct {
	Usecase  string `json:"usecase,omitempty" bson:"usecase"`
	Rank     int    `json:"rank,omitempty" bson:"rank"`
	Homepage string `json:"homepage,omitempty" bson:"homepage"`
	Download string `json:"download,omitempty" bson:"download"`
	Votes    int    `json:"votes" bson:"votes"`
}

type platform struct {
	Name   string     `json:"name,omitempty" bson:"name"`
	Detail codedetail `json:"codedetail,omitempty" bson:"codedetail"`
}

var c *mongo.Client
var listenport string

func createPlatform(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var detail codedetail
	_ = json.NewDecoder(req.Body).Decode(&detail)
	name := strings.ToLower(params["name"])

	fmt.Println(fmt.Sprintf("POST api call made to /platforms/%s", name))

	pltf := platform{name, detail}

	id := insertNewPlatform(c, pltf)

	if id == nil {
		_ = json.NewEncoder(w).Encode("{'result' : 'insert failed!'}")
	} else {
		err := json.NewEncoder(w).Encode(detail)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
	}

	return
}

func getPlatforms(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("GET api call made to /platforms")

	var pltfmap = make(map[string]*codedetail)
	pltfs, err := returnAllPlatforms(c, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	for _, pltf := range pltfs {
		pltfmap[pltf.Name] = &pltf.Detail
	}

	err = json.NewEncoder(w).Encode(pltfmap)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	return
}

func getPlatformByName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Println(fmt.Sprintf("GET api call made to /platforms/%s", name))

	pltf, _ := returnOnePlatform(c, bson.M{"name": name})
	if pltf == nil {
		_ = json.NewEncoder(w).Encode("{'result' : 'platform not found'}")
	} else {
		err := json.NewEncoder(w).Encode(*pltf)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
	}

	return
}

func deletePlatformByName(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Println(fmt.Sprintf("DELETE api call made to /platforms/%s", name))

	platformsRemoved := removeOnePlatform(c, bson.M{"name": name})

	_ = json.NewEncoder(w).Encode(fmt.Sprintf("{'count' : %d}", platformsRemoved))

	return
}

func voteOnPlatform(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	name := strings.ToLower(params["name"])

	fmt.Println(fmt.Sprintf("GET api call made to /platforms/%s/vote", name))

	//votesUpdated := updateVote(c, bson.M{"name": name})
	vchan := voteChannel()
	vchan <- name
	votesUpdated := <-vchan
	close(vchan)

	_ = json.NewEncoder(w).Encode(fmt.Sprintf("{'count' : %s}", votesUpdated))
}

func voteChannel() (vchan chan string) {
	vchan = make(chan string)
	go func() {
		name := <-vchan
		//fmt.Println(fmt.Sprintf("name is %s", name))
		votesUpdated := strconv.FormatInt((updateVote(c, bson.M{"name": name})), 10)
		vchan <- votesUpdated
	}()
	return vchan
}

func returnAllPlatforms(client *mongo.Client, filter bson.M) ([]*platform, error) {
	var pltfs []*platform
	collection := client.Database(MONGO_DB).Collection(MONGO_COLLECTION)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Err: ",err);
		return nil, errors.New("error querying documents from database")
	}
	for cur.Next(context.TODO()) {
		var pltf platform
		err = cur.Decode(&pltf)
		if err != nil {
			return nil, errors.New("error on decoding the document")
		}
		pltfs = append(pltfs, &pltf)
	}
	return pltfs, nil
}

func returnOnePlatform(client *mongo.Client, filter bson.M) (*platform, error) {
	var pltf platform
	collection := client.Database(MONGO_DB).Collection(MONGO_COLLECTION)
	singleResult := collection.FindOne(context.TODO(), filter)
	if singleResult.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("no documents found")
	}
	if singleResult.Err() != nil {
		log.Println("Find error: ", singleResult.Err())
		return nil, singleResult.Err()
	}
	singleResult.Decode(&pltf)
	return &pltf, nil
}

func insertNewPlatform(client *mongo.Client, pltf platform) interface{} {
	collection := client.Database(MONGO_DB).Collection(MONGO_COLLECTION)
	insertResult, err := collection.InsertOne(context.TODO(), pltf)
	if err != nil {
		log.Fatalln("Error on inserting new platform", err)
		return nil
	}
	return insertResult.InsertedID
}

func removeOnePlatform(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database(MONGO_DB).Collection(MONGO_COLLECTION)
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on deleting one Hero", err)
	}
	return deleteResult.DeletedCount
}

func updateVote(client *mongo.Client, filter bson.M) int64 {
	collection := client.Database(MONGO_DB).Collection(MONGO_COLLECTION)
	updatedData := bson.M{"$inc": bson.M{"codedetail.votes": 1}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, updatedData)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
	}
	return updatedResult.ModifiedCount
}

//getClient returns a MongoDB Client
func getClient() *mongo.Client {
	mongoconnstr := getEnv("MONGO_CONN_STR", MONGO_DEFAULT_CONN_STR)
	mongousername := getEnv("MONGO_USERNAME", MONGO_DEFAULT_USERNAME)
	mongopassword := getEnv("MONGO_PASSWORD", MONGO_DEFAULT_PASSWORD)

	fmt.Println("MongoDB connection details:")
	fmt.Println("MONGO_CONN_STR:" + mongoconnstr)
	fmt.Println("MONGO_USERNAME:" + mongousername)
	fmt.Println("MONGO_PASSWORD:")
	fmt.Println("attempting mongodb backend connection...")

	var cred options.Credential
	cred.AuthSource = "platformsdb"
	cred.Username = mongousername
	cred.Password = mongopassword

	clientOptions := options.Client().ApplyURI(mongoconnstr).SetAuth(cred)

	//test if auth is enabled or expected,
	//for demo purposes when we setup mongo as a replica set using a StatefulSet resource in K8s auth is disabled
	//if clientOptions.Auth != nil {
	//	clientOptions.Auth.Username = mongousername
	//	clientOptions.Auth.Password = mongopassword
	//}

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func init() {
	printInfo()

	c = getClient()
	err := c.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("couldn't connect to the database", err)
	} else {
		log.Println("connected!!")
	}
}

func ok(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK!")
	return
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func printInfo() {
	listenport = getEnv("LISTEN_PORT", LISTEN_PORT)
	version := getEnv("VERSION", VERSION)

	fmt.Printf("version %s", version)
	fmt.Printf("serving on port %s...", listenport)
	fmt.Println("tests:")
	fmt.Printf("curl -s localhost:%s/ok", listenport)
	fmt.Printf("curl -s localhost:%s/platforms", listenport)
	fmt.Printf("curl -s localhost:%s/platforms | jq .", listenport)
}

func main() {

	router := mux.NewRouter()

	//setup routes
	router.HandleFunc("/platforms/{name}", createPlatform).Methods("POST")
	router.HandleFunc("/platforms", getPlatforms).Methods("GET")
	router.HandleFunc("/platforms/{name}", getPlatformByName).Methods("GET")
	router.HandleFunc("/platforms/{name}", deletePlatformByName).Methods("DELETE")
	router.HandleFunc("/platforms/{name}/vote", voteOnPlatform).Methods("GET")
	router.HandleFunc("/ok", ok).Methods("GET")

	//required for CORS - ajax API requests originating from the react browser vote app
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET, POST"})

	//listen on port
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", listenport), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
