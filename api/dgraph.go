package api

import "github.com/kataras/iris"

func CreateHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("number")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

func ReadHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("number")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

func UpdateHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("number")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}

func DeleteHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("number")
	ctx.Writef("from "+ctx.GetCurrentRoute().Path()+" with ID: %d", id)
}
