package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"knowledgemap_backend/microservices/common/conf"

	"knowledgemap_backend/agent/knowledgemap/server/http"

	"github.com/labstack/gommon/log"
)

func main() {
	conf.Init()
	webSrv := http.Init()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("vip-service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//			rpcSvr.Close()
			//		ws.Shutdown(context.Background())
			webSrv.Shutdown(context.Background())
			time.Sleep(time.Second * 2)
			log.Info("vip-service exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}

}
