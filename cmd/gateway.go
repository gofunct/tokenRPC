package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
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

var (
	gatewayConfig = &GatewayConfig{}
	gatewayViper = viper.New()
)

func init() {
	{
		gatewayCmd.AddCommand(gatewayServeCmd)
	}

	{
		gatewayCmd.Flags().IntVarP(&gatewayConfig.Port, "port", "p", 8080, "port to run gateway on")
		gatewayCmd.Flags().BoolVar(&gatewayConfig.LogHeaders, "log-headers", true, "log headers")
		gatewayCmd.Flags().StringVar(&gatewayConfig.SwaggerFile, "swagger-file", "swagger.json", "path to generated swagger file")
		gatewayCmd.Flags().StringVar(&gatewayConfig.CorsAllowOrigin, "allow-origin", "", "CORS origin")
		gatewayCmd.Flags().StringVar(&gatewayConfig.CorsAllowCredentials, "allow-creds", "", "CORS credentials")
		gatewayCmd.Flags().StringVar(&gatewayConfig.CorsAllowMethods, "allow-methods", "", "CORS methods")
		gatewayCmd.Flags().StringVar(&gatewayConfig.CorsAllowHeaders, "allow-headers", "", "CORS headers")
		gatewayCmd.Flags().StringVar(&gatewayConfig.ApiPrefix, "prefix", "token", "api prefix")
	}

	{
		gatewayViper.SetConfigName("hack-gateway")
		gatewayViper.AutomaticEnv()
		gatewayViper.AddConfigPath(os.Getenv("$HOME")) // name of config file (without extension)
		gatewayViper.AddConfigPath(".")
		gatewayViper.SetEnvPrefix("hack_gateway")

		gatewayViper.BindPFlag("port", gatewayCmd.Flags().Lookup("port"))
		gatewayViper.BindPFlag("log-headers", gatewayCmd.Flags().Lookup("log-headers"))
		gatewayViper.BindPFlag("swagger-file", gatewayCmd.Flags().Lookup("swagger-file"))
		gatewayViper.BindPFlag("allow-origin", gatewayCmd.Flags().Lookup("allow-origin"))
		gatewayViper.BindPFlag("allow-creds", gatewayCmd.Flags().Lookup("allow-creds"))
		gatewayViper.BindPFlag("allow-methods", gatewayCmd.Flags().Lookup("allow-methods"))
		gatewayViper.BindPFlag("allow-headers", gatewayCmd.Flags().Lookup("allow-headers"))
		gatewayViper.BindPFlag("prefix", gatewayCmd.Flags().Lookup("prefix"))

	}

	// If a config file is found, read it in.
	if err := gatewayViper.ReadInConfig(); err != nil {
		log.Println("failed to read config file, writing defaults...")
		if err := gatewayViper.WriteConfigAs("hack-gateway" + ".yaml"); err != nil {
			log.Fatal("failed to write config")
			os.Exit(1)
		}

	} else {
		log.Print("Using config file: ", gatewayViper.ConfigFileUsed())
		if err := gatewayViper.WriteConfig(); err != nil {
			log.Fatal("failed to write config file")
			os.Exit(1)
		}
	}
}

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "A REST to gRPC server gateway",
}

// gatewayCmd represents the gateway command
var gatewayServeCmd = &cobra.Command{
	Use:   "serve",
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
