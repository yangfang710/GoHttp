package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"

	"GoHttp/conf"
	"GoHttp/env"
	"GoHttp/env/global"
	"GoHttp/handler"
	"GoHttp/router"
)

func main() {
	appConf := flag.String("app.conf", "conf/app.test.json", "app configuration file")

	flag.Parse()
	defer logrus.Exit(0)

	config, err := conf.ParseFile(*appConf)
	if err != nil {
		exit(err, true)
	}
	global.Config = *config

	if err := env.Configure(env.Option{}); err != nil {
		exit(err, false)
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)

	servlet := new(handler.Servlet)
	routeServer := router.NewRouter(servlet)

	httpSrv := &http.Server{Addr: global.Config.Addr, Handler: routeServer.GinEngine}
	go func() {
		logrus.Println("Start serving")
		httpSrv.RegisterOnShutdown(func() {
			logrus.Println("Shutdown server.")
			wg.Done()
		})
		httpSrv.ListenAndServe()
	}()

	srvs := []*http.Server{httpSrv}

	waitForTeardown(routeServer, srvs...)
	wg.Wait()
}

func waitForTeardown(router *router.Router, httpSrvs ...*http.Server) {

	sigCh := make(chan os.Signal)

	signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	<-sigCh
	if err := router.Close(); err != nil {
		logrus.WithError(err).Println("router.Close()")
	}
	ctx := context.Background()
	for _, httpSrv := range httpSrvs {
		if err := httpSrv.Shutdown(ctx); err != nil {
			logrus.WithError(err).Println("httpSrv.Shutdown(ctx)")
		}
	}
}

func exit(err error, withUsage bool) {
	logrus.Error(err)
	if withUsage {
		flag.Usage()
	}
	logrus.Exit(1)
}
