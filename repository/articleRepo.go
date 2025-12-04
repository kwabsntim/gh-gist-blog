package repository

import (
	"context"
	"errors"
	"ghgist-blog/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constructor function that return the interface
func NewMongoArticleRepository(client *mongo.Client) UserRepository {
	return &mongoclient{client: client}
}

// create article
func (m *mongoclient) CreateArticle(article *models.Article) error {
	article.CreatedAt = time.Now()
	collection := m.client.Database("ghgistDB").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.InsertOne(ctx, article)
	if err != nil {
		return err
	}
	//setting up article ID
	if oid, ok := cursor.InsertedID.(primitive.ObjectID); ok {
		article.ID = oid
	}
	return nil
}

// get all articles
func (m *mongoclient) FetchallArticles() ([]models.Article, error) {
	collection := m.client.Database("ghgistDB").Collection("articles")
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

// edit articles
func (m *mongoclient) EditArticle(article models.Article) error {
	objID, err := primitive.ObjectIDFromHex(article.ID.Hex())
	if err != nil {
		return errors.New("Invalid article ID")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := m.client.Database("ghgistDB").Collection("articles")
	updateFields := bson.M{}
	if article.Title != "" {
		updateFields["title"] = article.Title
	}
	if article.Content != "" {
		updateFields["content"] = article.Content
	}
	if article.ImageURL != "" {
		updateFields["image_url"] = article.ImageURL
	}
	if article.Slug != "" {
		updateFields["slug"] = article.Slug
	}
	updateFields["updated_at"] = time.Now()

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}

	cursor, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if cursor.MatchedCount == 0 {
		return errors.New("Article not found")
	}
	return nil
}

// delete article
func (m *mongoclient) DeleteArticle(article *models.Article) error {
	//delete article logic
	objID, err := primitive.ObjectIDFromHex(article.ID.Hex())
	if err != nil {
		return errors.New("invalid article ID")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := m.client.Database("ghgistDB").Collection("articles")
	filter := bson.M{"_id": objID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("article not found")
	}
	return nil
}

// get artcle by ID
func (m *mongoclient) FindArticleByID(id primitive.ObjectID) (*models.Article, error) {
	collection := m.client.Database("ghgistDB").Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	var article models.Article                                         // ← Create actual instance
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&article) // ← Decode into it
	if err != nil {
		return nil, err
	}
	return &article, nil // ← Return address of the instance
}

// get articles by category
func (m *mongoclient) FindArticlesByCategoryID(categoryID primitive.ObjectID) ([]models.Article, error) {
	collection := m.client.Database("ghgistDB").Collection("articles")
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
