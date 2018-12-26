package cmd

import (
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"os"

	"github.com/gofunct/hack/contracts/token_service"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"time"
)

type GatewayConfig struct {
	Port                 int
	LogHeaders           bool
	SwaggerFile          string
	CorsAllowOrigin      string
	CorsAllowCredentials string
	CorsAllowMethods     string
	CorsAllowHeaders     string
	ApiPrefix            string
}

var gatewayConfig = &GatewayConfig{}

func init() {
	gatewayCmd.Flags().IntVarP(&gatewayConfig.Port, "port", "p", 8080, "port to run gateway on")
	gatewayCmd.PersistentFlags().BoolVarP(&gatewayConfig.LogHeaders, "log-headers", "u", true, "server address to dial")
	gatewayCmd.PersistentFlags().StringVarP(&gatewayConfig.SwaggerFile, "swagger-file", "g", "swagger.json", "path to generated swagger file")
	gatewayCmd.PersistentFlags().StringVarP(&gatewayConfig.CorsAllowOrigin, "allow-origin", "o", "", "CORS origin")
	gatewayCmd.PersistentFlags().StringVarP(&gatewayConfig.CorsAllowCredentials, "allow-creds", "d", "", "CORS credentials")
	gatewayCmd.PersistentFlags().StringVarP(&gatewayConfig.CorsAllowMethods, "allow-methods", "m", "", "CORS methods")
	gatewayCmd.PersistentFlags().StringVarP(&gatewayConfig.CorsAllowHeaders, "allow-headers", "r", "", "CORS headers")
	gatewayCmd.PersistentFlags().StringVarP(&gatewayConfig.ApiPrefix, "prefix", "x", "token", "api prefix")
}

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "A REST to gRPC server gateway",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := SetupMux(ctx)

		addr := fmt.Sprintf(":%v", gatewayConfig.Port)
		server := &http.Server{Addr: addr, Handler: mux}

		SignalRunner(
			func() {
				logrus.Infof("launching http server on %v", server.Addr)
				if err := server.ListenAndServe(); err != nil {
					logrus.Fatalf("Could not start http server: %v", err)
				}
			},
			func() {
				shutdown, _ := context.WithTimeout(ctx, 10*time.Second)
				server.Shutdown(shutdown)
			})
	},
}

func SetupMux(ctx context.Context) *http.ServeMux {

	formatter := logFormatter()

	logrus.Info("Creating grpc-gateway proxy")
	mux := http.NewServeMux()

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, gatewayConfig.SwaggerFile)
	})

	gwmux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(incomingHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(outgoingHeaderMatcher),
	)
	logrus.Infof("Proxying requests to gRPC service at '%s'", ethConfig.ServerAddr)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	// If you get a compilation error that gw.RegisterTokenHandlerFromEndpoint
	// does not exist, it's because you haven't added any google.api.http annotations
	// to your proto. Add some!
	err := token_service.RegisterTokenHandlerFromEndpoint(ctx, gwmux, ethConfig.ServerAddr, opts)
	if err != nil {
		logrus.Fatalf("Could not register gateway: %v", err)
	}

	prefix := sanitizeApiPrefix(gatewayConfig.ApiPrefix)
	logrus.Infof("API prefix is: %s", prefix)
	mux.Handle(prefix, handlers.CustomLoggingHandler(os.Stdout, http.StripPrefix(prefix[:len(prefix)-1], allowCors(gwmux)), formatter))
	return mux
}
