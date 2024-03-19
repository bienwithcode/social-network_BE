package rMongo

import (
	"context"
	"social-network/domain"

	"go.mongodb.org/mongo-driver/bson"
)

// type UserRepository interface {
// 	Fetch(ctx context.Context, cursor string, num int64) (res []domain.Article, nextCursor string, err error)
// 	GetByID(ctx context.Context, id int64) (domain.Article, error)
// 	GetByTitle(ctx context.Context, title string) (domain.Article, error)
// 	Update(ctx context.Context, ar *domain.Article) error
// 	Store(ctx context.Context, a *domain.Article) error
// 	Delete(ctx context.Context, id int64) error
// }

func (storage *mongodbStorage) GetAuth(ctx context.Context, email, password string) (*domain.User, error) {

	var user domain.User
	collection := storage.db.Collection(domain.User{}.TableName())
	filter := bson.M{"email": email, "password": password}

	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
