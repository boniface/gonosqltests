package api

import "github.com/kataras/iris"

type Api interface {
	CreateHandler(ctx iris.Context)
	ReadHandler(ctx iris.Context)
	UpdateHandler(ctx iris.Context)
	DeleteHandler(ctx iris.Context)
}
