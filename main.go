package main

import (
	"flag"
	"fmt"
	"github.com/iris-contrib/gateway"
	"github.com/kataras/iris/v12"
)

var (
	port = flag.Int("port", -1, "specify a port")
)

func main() {

	flag.Parse()

	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)

	app.Get("/ping", status)
	app.Post("/who", whoIsGopher)

	admin := app.Party("/admin")
	admin.Get("/hello", helloGopher)
	admin.Post("/who", whoIsGopher)

	if *port != -1 {
		portNow := fmt.Sprintf(":%d", *port)
		app.Listen(portNow)
		return
	}

	runner, configurator := gateway.New(gateway.Options{
		URLPathParameter: "path",
	})
	app.Run(runner, configurator)
}

type IWhoRequest struct {
	Name string `json:"name"`
}

func whoIsGopher(context iris.Context) {
	var req IWhoRequest
	err := context.ReadJSON(&req)
	if err != nil {
		context.JSON(iris.Map{"Message": "error"})
		return
	}
	context.JSON(iris.Map{"Message": fmt.Sprintf("oh yes, he is %s",
		req.Name)})
}

func notFound(ctx iris.Context) {
	code := ctx.GetStatusCode()
	msg := iris.StatusText(code)
	if err := ctx.GetErr(); err != nil {
		msg = err.Error()
	}

	ctx.JSON(iris.Map{
		"Message": msg,
		"Code":    code,
	})
}

func index(ctx iris.Context) {
	var req map[string]interface{}
	ctx.ReadJSON(req)
	ctx.JSON(req)
}

func status(ctx iris.Context) {
	ctx.JSON(iris.Map{"Message": "OK"})
}

func helloGopher(ctx iris.Context) {
	ctx.JSON(iris.Map{"Message": "hello gopher"})
}
