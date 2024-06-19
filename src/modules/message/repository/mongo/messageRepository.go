package rMongo

import (
	"context"
	"social-network/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (storage *mongodbStorage) GetConversations(ctx context.Context, authUserId string) ([]*domain.Message, error) {
	objectID, err := primitive.ObjectIDFromHex(authUserId)
	if err != nil {
		return nil, err
	}

	collection := storage.db.Collection(domain.Message{}.TableName())

	filter := bson.A{
		bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "$or",
						Value: bson.A{
							bson.D{{Key: "receiver", Value: objectID}},
							bson.D{{Key: "sender", Value: objectID}},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "sender", Value: 1},
					{Key: "receiver", Value: 1},
					{Key: "message", Value: 1},
					{Key: "createdAt", Value: 1},
					{Key: "seen", Value: 1},
					{Key: "senderReceiver",
						Value: bson.A{
							"$sender",
							"$receiver",
						},
					},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$senderReceiver"}}}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "senderReceiver", Value: 1}}}},
		bson.D{
			{Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "$_id"},
					{Key: "senderReceiver", Value: bson.D{{Key: "$push", Value: "$senderReceiver"}}},
					{Key: "message", Value: bson.D{{Key: "$first", Value: "$message"}}},
					{Key: "createdAt", Value: bson.D{{Key: "$first", Value: "$createdAt"}}},
					{Key: "sender", Value: bson.D{{Key: "$first", Value: "$sender"}}},
					{Key: "receiver", Value: bson.D{{Key: "$first", Value: "$receiver"}}},
					{Key: "seen", Value: bson.D{{Key: "$first", Value: "$seen"}}},
				},
			},
		},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "createdAt", Value: -1}}}},
		bson.D{
			{Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "$senderReceiver"},
					{Key: "message", Value: bson.D{{Key: "$first", Value: "$message"}}},
					{Key: "createdAt", Value: bson.D{{Key: "$first", Value: "$createdAt"}}},
					{Key: "sender", Value: bson.D{{Key: "$first", Value: "$sender"}}},
					{Key: "receiver", Value: bson.D{{Key: "$first", Value: "$receiver"}}},
					{Key: "seen", Value: bson.D{{Key: "$first", Value: "$seen"}}},
				},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "sender"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "sender"},
				},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "receiver"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "receiver"},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$sender"}}}},
		bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$receiver"}}}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "createdAt", Value: -1}}}},
	}

	cursor, err := collection.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}

	var messages []*domain.Message

	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (storage *mongodbStorage) GetMessages(ctx context.Context, authUserId, userId string) ([]*domain.Message, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$and": []bson.M{
					{
						"$or": []bson.M{
							{"sender": authUserId},
							{"receiver": authUserId},
						},
					},
					{
						"$or": []bson.M{
							{"sender": userId},
							{"receiver": userId},
						},
					},
				},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "sender",
				"foreignField": "_id",
				"as":           "sender",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "receiver",
				"foreignField": "_id",
				"as":           "receiver",
			},
		},
		{
			"$project": bson.M{
				"sender.password":   0,
				"receiver.password": 0,
			},
		},
		{
			"$sort": bson.M{
				"updatedAt": 1,
			},
		},
	}

	collection := storage.db.Collection(domain.Message{}.TableName())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var messages []*domain.Message

	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}
