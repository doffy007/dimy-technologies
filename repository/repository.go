package repository

import (
	"context"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/database/mysql"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	ctx    context.Context
	config *config.Config
	db     *sqlx.DB
}

func NewRepository(ctx context.Context, conf *config.Config) Repository {
	return repository{
		ctx:    ctx,
		config: conf,
		db:     mysql.Mysql(),
	}
}
