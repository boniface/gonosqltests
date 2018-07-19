package api

import "github.com/kataras/iris"

//Index: landing page
func Index(ctx iris.Context) {
	// Bind: {{.message}} with "Hello world!"
	ctx.ViewData("message", "Hello world! Test Testing ")
	// Render template file: ./views/hello.html
	ctx.View("index.html")
}
