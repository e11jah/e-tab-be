package impl

import (
	"context"
	"log"

	"e_tab_be/internal/model"
	"e_tab_be/internal/types"
	"e_tab_be/pkg/url"
	"e_tab_be/pkg/xdb"
)



type TabRepo struct {
}

func (r *TabRepo) BatchInsert(ctx context.Context, tabs []*types.Tab) error {
	tabModels := make([]*model.Tab, 0, len(tabs))

	for _, tab := range tabs {
		domain, err := url.ParseUrl(tab.Url)
		if err != nil {
			log.Printf ("parse url '%s' error, err: %s", tab.Url, err)
			continue
		}
		tabModels = append(tabModels, tab.ToModel(domain))
	}

	db := xdb.GetDB(ctx)

	return db.Create(tabModels).Error
}
