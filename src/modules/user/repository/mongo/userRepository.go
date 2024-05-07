package rMongo

import (
	"context"
	"social-network/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func (storage *mongodbStorage) GetAuth(ctx context.Context, email, password string) (*domain.User, error) {

	var user domain.UserWithPassword

	collection := storage.db.Collection(domain.User{}.TableName())
	filter := bson.M{"email": email}
	projection := bson.M{
		"_id":           1,
		"email":         1,
		"role":          1,
		"emailVerified": 1,
		"password":      1,
	}

	if err := collection.FindOne(ctx, filter, &options.FindOneOptions{Projection: projection}).Decode(&user); err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &domain.User{
		Id:            user.Id,
		Email:         user.Email,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
	}, nil
}

func (storage *mongodbStorage) GetAuthUser(ctx context.Context, id string) (*domain.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := storage.db.Collection(domain.User{}.TableName())
	filter := bson.M{"_id": objectID}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "isOnline", Value: true}}}}
	var user domain.User
	if err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
