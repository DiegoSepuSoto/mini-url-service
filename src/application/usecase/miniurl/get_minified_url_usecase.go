package miniurl

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-service/src/domain/models"
	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
	"go.opentelemetry.io/otel"
)

func (u *miniURLUseCase) GetMinifiedURL(ctx context.Context, miniURL string) (*models.MinifiedURLResponse, error) {
	ctx, span := otel.Tracer(shared.TracerName).Start(ctx, "GetMinifiedURLUsecase")
	defer span.End()

	minifiedURL, err := u.miniURLRepository.GetMinifiedURL(ctx, miniURL)
	if err != nil {
		return nil, err
	}

	return &models.MinifiedURLResponse{
		MinifiedURL: minifiedURL,
	}, nil
}
