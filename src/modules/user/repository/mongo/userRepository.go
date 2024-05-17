package rMongo

import (
	"context"
	"social-network/domain"
	"social-network/modules/auth/model"
	"social-network/utils"

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

func (storage *mongodbStorage) GetUsers(ctx context.Context, paging *utils.Pagination, filter *model.Filter) ([]*domain.User, error) {
	var authUserId string
	if auth, ok := ctx.Get("authData"); ok {
		authData, _ := auth.(map[string]interface{})
		authUserId = authData["id"].(string)
	}

	query := bson.M{"_id": bson.M{"$ne": authUserId}}

	if filter.EmailVerified != nil {
		query["emailVerified"] = *filter.EmailVerified
	}

	if filter.HideBannedUsers != nil {
		query["banned"] = bson.M{"$ne": true}
	}

	if filter.SearchQuery != nil {
		query["$or"] = []bson.M{
			{"username": bson.M{"$regex": primitive.Regex{Pattern: *filter.SearchQuery, Options: "i"}}},
			{"fullName": bson.M{"$regex": primitive.Regex{Pattern: *filter.SearchQuery, Options: "i"}}},
			{"email": bson.M{"$regex": primitive.Regex{Pattern: *filter.SearchQuery, Options: "i"}}},
		}
	}

	var users []*domain.User
	collection := storage.db.Collection(domain.User{}.TableName())

	// count number of documents
	num, err := collection.CountDocuments(ctx, query, options.Count())
	if err != nil {
		return nil, err
	}
	paging.SetTotal(num)

	// Find and paging documents
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"password": 0})
	findOptions.SetSkip(int64(paging.Page))
	findOptions.SetLimit(int64(paging.PerPage))
	findOptions.SetSort(bson.M{"createdAt": -1})

	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	// defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user domain.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
