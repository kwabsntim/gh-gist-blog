package repository

import (
	"context"
	"errors"
	"ghgist-blog/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Constructor function that returns the interface
func NewMongoUserRepository(client *mongo.Client) UserRepository {
	return &mongoclient{client: client}
}

func (m *mongoclient) CreateUser(user *models.User) error {

	user.CreatedAt = time.Now()

	collection := m.client.Database("ghgistDB").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 	Direct insert with unique index handling
	cursor, err := collection.InsertOne(ctx, user)
	if err != nil {
		// Check for duplicate key error
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("user already exists")
		}
		return err
	}

	// Set ID if successful
	if oid, ok := cursor.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}
	return nil
}
func (m *mongoclient) FindUserByEmail(email string) (*models.User, error) {
	collection := m.client.Database("ghgistDB").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (m *mongoclient) SetupIndexes() error {
	if m.client == nil {
		log.Fatal("Database client not initialized")
	}

	collection := m.client.Database("ghgistDB").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create unique index on email
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// Create unique index on username
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}}, // Note: matches your struct tag
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{emailIndex, usernameIndex})
	if err != nil {
		log.Printf("Warning: Failed to create indexes: %v", err)
		return err
	}

	log.Println("Database indexes created successfully!")
	return nil
}
func (m *mongoclient) FetchAllUsers() ([]models.User, error) {
	collection := m.client.Database("ghgistDB").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.Find().SetProjection(bson.M{"password": 0})
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
