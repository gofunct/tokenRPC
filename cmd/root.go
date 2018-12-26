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
	defaultViper = viper.New()
	kitLog        = kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	ethConfig     = &EthConfig{}
)

type EthConfig struct {
	EthAddr      string
	ServerAddr   string
	PrivateKey   string
	ContractAddr string
	Password 	string
	LogLevel     string
	TokenName 	string
	TokenSymbol string
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
		rootCmd.PersistentFlags().StringVar(&ethConfig.EthAddr, "eth-addr",  "http://192.168.99.100:8545", "the ethereum client address")
		rootCmd.MarkFlagRequired("eth-addr")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.ServerAddr, "server-addr", "s", "127.0.0.1:5555", "server address")
		rootCmd.MarkFlagRequired("server-addr")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.PrivateKey, "private-key", "k", "", "deployer's private key")
		rootCmd.MarkFlagRequired("private-key")
		rootCmd.PersistentFlags().StringVarP(&ethConfig.ContractAddr, "contract-addr", "c", "", "contract address")
		rootCmd.MarkFlagRequired("contract-addr")
		rootCmd.PersistentFlags().StringVar(&ethConfig.Password, "password", "admin", "private key passphrase")
		rootCmd.MarkFlagRequired("password")
		rootCmd.PersistentFlags().StringVar(&ethConfig.LogLevel, "log-level",  "", "logging verbosity level")
		rootCmd.PersistentFlags().StringVar(&ethConfig.TokenName, "token-name", "", "name of eth token")
		rootCmd.MarkFlagRequired("token-name")
		rootCmd.PersistentFlags().StringVar(&ethConfig.TokenSymbol, "token-symbol", "", "token symbol")
		rootCmd.MarkFlagRequired("token-symbol")
	}

	{
		defaultViper.SetConfigName("hack")
		defaultViper.AutomaticEnv()
		defaultViper.AddConfigPath(os.Getenv("$HOME")) // name of config file (without extension)
		defaultViper.AddConfigPath(".")
		defaultViper.SetEnvPrefix("hack")
		defaultViper.BindPFlag("eth-addr", rootCmd.PersistentFlags().Lookup("eth-addr"))
		defaultViper.BindPFlag("server-addr", rootCmd.PersistentFlags().Lookup("server-addr"))
		defaultViper.BindPFlag("private-key", rootCmd.PersistentFlags().Lookup("private-key"))
		defaultViper.BindPFlag("contract-addr", rootCmd.PersistentFlags().Lookup("contract-addr"))
		defaultViper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
		defaultViper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
		defaultViper.BindPFlag("token-name", rootCmd.PersistentFlags().Lookup("token-name"))
		defaultViper.BindPFlag("token-symbol", rootCmd.PersistentFlags().Lookup("token-symbol"))

	}

	// If a config file is found, read it in.
	if err := defaultViper.ReadInConfig(); err != nil {
		log.Println("failed to read config file, writing defaults...")
		if err := defaultViper.WriteConfigAs("hack" + ".yaml"); err != nil {
			log.Fatal("failed to write config")
			os.Exit(1)
		}

	} else {
		log.Print("Using config file: ", defaultViper.ConfigFileUsed())
		if err := defaultViper.WriteConfig(); err != nil {
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
