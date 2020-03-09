package main

import (
	"hugo-server/handler"
	"hugo-server/mid"

	"github.com/yusank/klyn"
)

func main() {
	core := klyn.Default()
	core.UseMiddleware(mid.LogMid)
	group := core.Group("")

	handler.NewRouter(group)

	if err := core.Service(":17771");err != nil {
		panic(err)
	}
}
