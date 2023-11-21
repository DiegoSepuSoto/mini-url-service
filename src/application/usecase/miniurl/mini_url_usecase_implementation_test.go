package miniurl

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type miniURLsRepositoryMock struct {
	mock.Mock
}

func (m *miniURLsRepositoryMock) GetMinifiedURL(ctx context.Context, originalURL string) (string, error) {
	args := m.Called(ctx, originalURL)

	if args.Get(1) == nil {
		return args.Get(0).(string), nil
	}

	return args.Get(0).(string), args.Error(1)
}
