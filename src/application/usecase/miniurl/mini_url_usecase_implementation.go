package miniurl

import (
	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/database/repositories"
)

type miniURLUseCase struct {
	miniURLRepository repositories.MiniURLsRepository
}

func NewMiniURLUseCase(miniURLRepository repositories.MiniURLsRepository) *miniURLUseCase {
	return &miniURLUseCase{
		miniURLRepository: miniURLRepository,
	}
}
