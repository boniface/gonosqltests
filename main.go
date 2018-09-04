package main

import (
	"github.com/boniface/gonosqltests/api"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Load all templates from the "./views" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./views", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", api.Index)

	// Cassandra
	app.Get("/cassandra/create/{id}", api.CassandraCreate)
	app.Get("/cassandra/read", api.Index)
	app.Get("/cassandra/update", api.Index)
	app.Get("/cassandra/delete", api.Index)

	// Mongo DB
	app.Get("/mongodb/create/{id}", api.Index)
	app.Get("/mongodb/read", api.Index)
	app.Get("/mongodb/update", api.Index)
	app.Get("/mongodb/delete", api.Index)

	// Redis DB
	app.Get("/redis/create/{id}", api.Index)
	app.Get("/redis/read", api.Index)
	app.Get("/redis/update", api.Index)
	app.Get("/redis/delete", api.Index)

	// Dgraph
	app.Get("/dgraph/create/{id}", api.Index)
	app.Get("/dgraph/read", api.Index)
	app.Get("/dgraph/update", api.Index)
	app.Get("/dgraph/delete", api.Index)

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
