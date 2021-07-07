package repo

import (
	"context"

	"e_tab_be/internal/types"
)

type ITabRepo interface {
	BatchInsert(ctx context.Context, tabs []*types.Tab) error
}
