package xdb

import (
	"context"
	"errors"
	"fmt"

	"e_tab_be/vars"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	TxCtxKey ContextKey
)

type ContextKey struct{}

func GetDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(TxCtxKey).(*gorm.DB)
	if !ok {
		return vars.MysqlConn.Get(ctx)
	}
	return db
}

func beginTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TxCtxKey, tx)
}

func checkErrIsDuplicateKey(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}

func Exec(ctx context.Context, executor func(ctx context.Context) error) (err error) {
	var dbErr error
	tx := GetDB(ctx).Begin()

	ctx = beginTx(ctx, tx)

	defer func() {
		if err != nil {
			if checkErrIsDuplicateKey(err) {
				err = errors.New("record already exist")
				return
			}
			err = errors.New("db error")
		}
		if dbErr != nil {
			err = errors.New(fmt.Sprintf("db error: %s", dbErr))
		}
	}()

	if err = executor(ctx); err != nil {
		dbErr = tx.Rollback().Error
		return
	}

	dbErr = tx.Commit().Error
	return
}
