package mysql

import (
	"context"

	"gorm.io/gorm"
)

type IMysqlConn interface {
	Get(ctx context.Context) *gorm.DB
}


type mysqlConn struct {
	client *gorm.DB
}

func (c *mysqlConn) Get(ctx context.Context) *gorm.DB {
	return c.client.WithContext(ctx)
}
