package env

import mongo "go.mongodb.org/mongo-driver/mongo"

var MongoDBConnection *mongo.Database

const RESTPort = "REST_PORT"

var REST_Port string
