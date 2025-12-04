package repository

//when changing the database change it here
import (
	"go.mongodb.org/mongo-driver/mongo"
)

// this stores all the methods from mongodb
type mongoclient struct {
	client *mongo.Client
}
