package router

import (
	"context"
	"net/http"

	"github.com/doffy007/dimy-technologies/config"
	"github.com/doffy007/dimy-technologies/handler"
	"github.com/gin-gonic/gin"
)

type router struct {
	ctx     context.Context
	config  *config.Config
	route   *gin.Engine
	handler handler.Handler
}

func Register(ctx context.Context, conf *config.Config, route *gin.Engine) Router {
	return &router{
		ctx:     ctx,
		config:  conf,
		route:   route,
		handler: handler.NewHandler(ctx, conf),
	}
}

type Router interface {
	All()
	BaseRouter()
	NotFound()
}

func (r router) All() {
	r.BaseRouter()
	r.NotFound()
}

func (r router) NotFound() {
	r.route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Page not found",
		})
	})
}
