package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/generalzy/zeropan/gateway/internal/errorx"
	"github.com/generalzy/zeropan/gateway/internal/zapx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/generalzy/zeropan/gateway/internal/config"
	"github.com/generalzy/zeropan/gateway/internal/handler"
	"github.com/generalzy/zeropan/gateway/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 设置自定义error
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, any) {
		switch e := err.(type) {
		case *errorx.ApiError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusBadRequest, nil
		}
	})
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *errorx.ApiError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusBadRequest, nil
		}
	})
	// 替换日志
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
