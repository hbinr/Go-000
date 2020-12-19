package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"user.service/pkg/tool/snowflake"

	"user.service/internal/service"

	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/config"
	"github.com/apache/dubbo-go/protocol/rest/server/server_impl"
	gxlog "github.com/dubbogo/gost/log"
	"github.com/emicklei/go-restful/v3"
	"golang.org/x/sync/errgroup"

	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"

	_ "github.com/apache/dubbo-go/protocol/rest"

	_ "github.com/apache/dubbo-go/registry/protocol"

	_ "github.com/apache/dubbo-go/filter/filter_impl"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"

	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/config_center/nacos"
	_ "github.com/apache/dubbo-go/registry/nacos"
)

/*
	对内服务，仅接受来自内部其他服务或者网关的请求。可以被interface\admin\job\task调用

	启动前需配置环境变量：
	export CONF_PROVIDER_FILE_PATH="../../conf/server.yml"
	export APP_LOG_CONF_FILE="../../conf/log.yml"
*/
func main() {

	if err := snowflake.Init(); err != nil {
		gxlog.CError("main: snowflake.Init() failed,err", err)
	}
	userLgc, err := initLogic()
	if err != nil {
		gxlog.CError("main: di.InitLogic() failed,err", err)
	}
	config.SetProviderService(service.NewUserServiceProvider(userLgc))
	server_impl.AddGoRestfulServerFilter(func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		chain.ProcessFilter(request, response)
	})

	config.Load()
	initSignal()
}
func initSignal() {
	g, ctx := errgroup.WithContext(context.Background())
	// 1.dubbo-go 底层已经启动了grpc server 和http server

	// 2.catch signals
	g.Go(func() error {
		exitSignals := []os.Signal{os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			select {
			case <-ctx.Done():
				logger.Info("signal ctx done.....")
				return ctx.Err()
			case <-sig:
				return nil
			}
		}
	})

	// 3.wait stop
	if err := g.Wait(); err != nil {
		logger.Infof("error: %v", err)
	}

}
