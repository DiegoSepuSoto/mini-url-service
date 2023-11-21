package miniurl

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-service/src/domain/models"
)

type miniURLUseCaseMock struct {
	mock.Mock
}

func (m *miniURLUseCaseMock) GetMinifiedURL(ctx context.Context, miniURL string) (*models.MinifiedURLResponse, error) {
	args := m.Called(ctx, miniURL)

	if args.Get(1) == nil {
		return args.Get(0).(*models.MinifiedURLResponse), nil
	}

	return nil, args.Error(1)
}
