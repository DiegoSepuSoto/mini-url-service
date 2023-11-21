package miniurls

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

type mongoDBRepositoryMock struct {
	mock.Mock
}

func (m *mongoDBRepositoryMock) GetMinifiedURL(ctx context.Context, miniURL string) (string, error) {
	args := m.Called(ctx, miniURL)

	if args.Get(1) != nil {
		return "", args.Get(1).(error)
	}

	return args.Get(0).(string), nil
}

type redisClientMock struct {
	mock.Mock
}

func (m *redisClientMock) Get(ctx context.Context, key string) shared.RedisStringResult {
	args := m.Called(ctx, key)

	return args.Get(0).(shared.RedisStringResult)
}

type redisStringResultMock struct {
	mock.Mock
}

func (m *redisStringResultMock) Result() (string, error) {
	args := m.Called()

	if args.Get(1) != nil {
		return "", args.Get(1).(error)
	}

	return args.Get(0).(string), nil
}
