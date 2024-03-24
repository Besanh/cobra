package cmd

import (
	"cobra/common/env"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type Config struct {
	Port     string
	gRPCPort string
	LogLevel string
	LogFile  string
}

var config Config

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	cfg := Config{
		Port:     env.GetStringENV("PORT", "8000"),
		gRPCPort: env.GetStringENV("GRPC_PORT", "8001"),
		LogLevel: env.GetStringENV("LOG_LEVEL", "error"),
		LogFile:  env.GetStringENV("LOG_FILE", "log/console.log"),
	}
	config = cfg
}

func Execute() {
	var rootCmd = &cobra.Command{Use: "root-anh",
		Short: "Cmd root",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root success")
		},
	}
	var echoCmd = &cobra.Command{Use: "echo-cmd",
		Short: "Echo cmd",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("echo success")
		},
	}
	rootCmd.AddCommand(cmdMain, echoCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
