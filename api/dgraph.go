package api

import (
	"github.com/boniface/gonosqltests/domain"
	"github.com/kataras/iris"
	"time"
)

func DgraphCreate(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")

	results := domain.Results{
		Duration: 10,
		End:      time.Now(),
		Objects:  id,
		Start:    time.Now(),
	}
	ctx.JSON(results)
}

func DgraphRead(ctx iris.Context) {

	results := domain.Results{
		Duration: 10,
		End:      time.Now(),
		Objects:  10,
		Start:    time.Now(),
	}
	ctx.JSON(results)

}

func DgraphUpdate(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")

	results := domain.Results{
		Duration: 10,
		End:      time.Now(),
		Objects:  id,
		Start:    time.Now(),
	}
	ctx.JSON(results)

}

func DgraphDelete(ctx iris.Context) {

	results := domain.Results{
		Duration: 10,
		End:      time.Now(),
		Objects:  10,
		Start:    time.Now(),
	}
	ctx.JSON(results)

}
