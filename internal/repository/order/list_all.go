package order

import (
	"context"
	domain "github.com/pda13/love-my-gf/internal/repository/model"
	"maps"
	"slices"
)

func (i *implementation) ListAll(_ context.Context) (domain.Orders, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	return slices.Collect(maps.Values(i.storage)), nil
}
