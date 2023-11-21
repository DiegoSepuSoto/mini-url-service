package usecase

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-service/src/domain/models"
)

type MiniURLUseCase interface {
	GetMinifiedURL(ctx context.Context, miniURL string) (*models.MinifiedURLResponse, error)
}
