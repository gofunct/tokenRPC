package cmd

import (
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	defaultConfig = viper.New()
	kitLog        = kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	ethConfig     = &EthConfig{}
)

type EthConfig struct {
	EthAddr      string
	ServerAddr   string
	PrivateKey   string
	ContractAddr string
	LogLevel     string
}

func init() {
	{
		log.SetOutput(kitlog.NewStdlibAdapter(kitLog))
		log.Println("new json logger registered")
	}

	{
		rootCmd.AddCommand(clientCmd)
		rootCmd.AddCommand(serverCmd)
		rootCmd.AddCommand(gatewayCmd)
	}

	{
		rootCmd.PersistentFlags().StringVarP(&ethConfig.EthAddr, "eth-addr", "c", "http://192.168.99.100:8545", "the ethereum client address")
		rootCmd.MarkFlagRequired("eth-addr")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.ServerAddr, "server-addr", "s", "127.0.0.1:5555", "server address")
		rootCmd.MarkFlagRequired("server-addr")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.PrivateKey, "private-key", "k", "", "deployer's private key")
		rootCmd.MarkFlagRequired("private-key")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.ContractAddr, "contract-addr", "t", "", "contract address")
		rootCmd.MarkFlagRequired("contract-addr")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.LogLevel, "log-level", "l", "", "logging verbosity level")
	}

	{
		defaultConfig.SetConfigName("hack")
		defaultConfig.AutomaticEnv()
		defaultConfig.AddConfigPath(os.Getenv("$HOME")) // name of config file (without extension)
		defaultConfig.AddConfigPath(".")
		defaultConfig.SetEnvPrefix("hack")
		defaultConfig.BindPFlag("eth-addr", rootCmd.PersistentFlags().Lookup("eth-addr"))
		defaultConfig.BindPFlag("server-addr", rootCmd.PersistentFlags().Lookup("server-addr"))
		defaultConfig.BindPFlag("private-key", rootCmd.PersistentFlags().Lookup("private-key"))
		defaultConfig.BindPFlag("contract-addr", rootCmd.PersistentFlags().Lookup("contract-addr"))
		defaultConfig.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	}

	// If a config file is found, read it in.
	if err := defaultConfig.ReadInConfig(); err != nil {
		log.Println("failed to read config file, writing defaults...")
		if err := defaultConfig.WriteConfigAs("hack" + ".yaml"); err != nil {
			log.Fatal("failed to write config")
			os.Exit(1)
		}

	} else {
		log.Print("Using config file: ", defaultConfig.ConfigFileUsed())
		if err := defaultConfig.WriteConfig(); err != nil {
			log.Fatal("failed to write config file")
			os.Exit(1)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "hack",
	Short: "A grpc based microservice for interacting with ethereum contracts",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
