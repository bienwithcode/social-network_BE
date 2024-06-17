package business

import (
	"context"
	"social-network/domain"
)

type MessageRepository interface {
	GetConversations(ctx context.Context, authUserId string) ([]*domain.Message, error)
}

type business struct {
	messageRepository MessageRepository
}

func NewBusiness(messageRepository MessageRepository) *business {
	return &business{
		messageRepository: messageRepository,
	}
}

func (biz *business) GetConversations(ctx context.Context, authUserId string) ([]*domain.Message, error) {

	user, err := biz.messageRepository.GetConversations(ctx, authUserId)

	if err != nil {
		return nil, err
	}
	return user, nil
}
