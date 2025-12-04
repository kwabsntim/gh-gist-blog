package repository

//when changing the database change it here
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

// this stores all the methods from mongodb
type mongoclient struct {
	client *mongo.Client
}

// Constructor function that returns the interface
func NewMongoUserRepository(client *mongo.Client) UserRepository {
	return &mongoclient{client: client}
}

func (r *mongoclient) CreateUser(user *models.User) error {

	user.CreatedAt = time.Now()

	collection := r.client.Database("ghgistDB").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Direct insert with unique index handling
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		// Check for duplicate key error
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("user already exists")
		}
		return err
	}

	// Set ID if successful
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}
	return nil
}
func (r *mongoclient) FindUserByEmail(email string) (*models.User, error) {
	collection := r.client.Database("ghgistDB").Collection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *mongoclient) SetupIndexes() error {
	if r.client == nil {
		log.Fatal("Database client not initialized")
	}

	collection := r.client.Database("ghgistDB").Collection("profiles")
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
func (r *mongoclient) FetchAllUsers() ([]models.User, error) {
	collection := r.client.Database("ghgistDB").Collection("profiles")
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

// articles repository
func (r *mongoclient) CreateArticle(article *models.Article) error {
	article.CreatedAt = time.Now()
	collection := r.client.Database("ghgistDB").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, article)
	if err != nil {
		return err
	}
	//setting up article ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		article.ID = oid
	}
	return nil
}
func (r *mongoclient) FetchallArticles() ([]models.Article, error) {
	collection := r.client.Database("ghgistDB").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.Find().SetProjection(bson.M{})
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []models.Article
	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// editing articles
func (r *mongoclient) EditArticle(article models.Article) error {
	objID, err := primitive.ObjectIDFromHex(article.ID.Hex())
	if err != nil {
		return errors.New("Invalid article ID")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := m.client.Database("usersdb").Collection("users")
	updateFields := bson.M{}
}

//func (r *mongoclient)FindArticleBySlug(slug string)(articles *models.Article, err error){
//collection := r.client.Database("ghgistDB").Collection("articles")
//ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
//defer cancel()
//var article models.Article
//err := collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&article)
//if err != nil {
//	return nil, err
//}
//return &article, nil
//}

// get artcle by ID
func (r *mongoclient) FindArticleByID(id primitive.ObjectID) (*models.Article, error) {
	collection := r.client.Database("ghgistDB").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	var article models.Article                                         // ← Create actual instance
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&article) // ← Decode into it
	if err != nil {
		return nil, err
	}
	return &article, nil // ← Return address of the instance
}
func (r *mongoclient) FindArticlesByCategoryID(categoryID primitive.ObjectID) ([]models.Article, error) {
	collection := r.client.Database("ghgistDB").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{"category_id": categoryID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []models.Article
	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

//func (r *mongoclient) FindArticlesByAuthorID(authorID primitive.ObjectID) ([]models.Article, error) {

//}
