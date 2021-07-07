package svc

import (
	"context"

	"e_tab_be/internal/config"
	"e_tab_be/internal/repo"
	"e_tab_be/internal/types"
	"e_tab_be/pkg/xdb"
)

type ServiceContext struct {
	Config config.Config

	repo repo.ITabRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

func (s *ServiceContext) SaveTabs(ctx context.Context,tabs []*types.Tab) error {

	return xdb.Exec(ctx, func(ctx context.Context) error {
		return s.repo.BatchInsert(ctx, tabs)
	})

}
