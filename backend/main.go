package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/svc"
	"backend/internal/util"
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
			logx.Infof("a new request comes, method: %s, path: %s, header: %s", r.Method, r.URL.Path, r.Header)
			next(w, r)
		}
	})
	logx.DisableStat()

	// Custom error
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *util.ErrorResponse:
			if e.Code == util.FailToGetIdOrAccountFromPath || e.Code == util.InvalidIdArrayForThisAccount {
				return http.StatusBadRequest, e.Data()
			}
			return http.StatusInternalServerError, e.Data()
		default:
			return http.StatusInternalServerError, errors.New(util.MessageMap[util.Unknown])
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
