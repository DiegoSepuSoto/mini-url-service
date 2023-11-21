package miniurl

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-service/src/domain/models"
)

func (u *miniURLUseCase) GetMinifiedURL(ctx context.Context, miniURL string) (*models.MinifiedURLResponse, error) {
	minifiedURL, err := u.miniURLRepository.GetMinifiedURL(ctx, miniURL)
	if err != nil {
		return nil, err
	}

	return &models.MinifiedURLResponse{
		MinifiedURL: minifiedURL,
	}, nil
}
