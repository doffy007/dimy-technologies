package usecase

import (
	"context"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/repository"
)

type usecase struct {
	ctx        context.Context
	config     *config.Config
	repository repository.Repository
}

func NewUsecase(ctx context.Context, conf *config.Config) Usecase {
	return usecase{
		ctx:        ctx,
		config:     conf,
		repository: repository.NewRepository(ctx, conf),
	}
}
