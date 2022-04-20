package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tix-rest-project/internal/controller"
	"tix-rest-project/internal/di"
	"tix-rest-project/internal/shared"
)

func main() {
	var (
		container = di.Container
	)

	err := container.Invoke(func(deps shared.Deps, ch controller.Holder) error {
		var (
			sig = make(chan os.Signal, 1)
		)

		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		ch.RegisterRoutes()

		go func() {
			if err := deps.Echo.Start(fmt.Sprintf(":%d", deps.Config.HttpServerPort)); err != nil {
				deps.Logger.Errorf("http server failed to start at port %d", deps.Config.HttpServerPort)
				sig <- syscall.SIGTERM
			}
		}()

		deps.Logger.Infof("http server started at port %d", deps.Config.HttpServerPort)

		// - gracefull shutdown & closing resources
		<-sig
		deps.Close()

		return nil
	})

	if err != nil {
		panic(err)
	}
}
