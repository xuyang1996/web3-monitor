package main

import (
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/main-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// global middleware
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Infof("a new request comes, path: %s, header: %s", r.URL.Path, r.Header)
			next(w, r)
		}
	})
	logx.DisableStat()

	// Custom error
	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	switch e := err.(type) {
	//	case *util.ErrorResponse:
	//		return http.StatusOK, e
	//	default:
	//		return http.StatusInternalServerError, nil
	//	}
	//})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
