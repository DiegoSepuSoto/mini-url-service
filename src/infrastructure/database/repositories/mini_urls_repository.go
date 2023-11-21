package repositories

import "context"

type MiniURLsRepository interface {
	GetMinifiedURL(ctx context.Context, miniURL string) (string, error)
}
