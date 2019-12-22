package cmd

import (
	"fmt"
	"log"
	"notify/logic/server"
	"os"

	"github.com/spf13/cobra"
)

const DefaultConfigPath = "./config.yaml"

var (
	configPath string

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "An daemon run in background to do the remind thing",
		Long: `An daemon run in background to do the remind thing`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Start Daemon")
			go func() {
				err := server.RpcHandler()
				if err != nil {
					log.Println("RpcService Start Fail：",err)
					os.Exit(-1)
				}
			}()
			server.Start()
		},
	}
)

func init() {
	serverCmd.PersistentFlags().StringVar(&configPath,"config",DefaultConfigPath,"just config file of notify")
	_, err := os.Stat(configPath)
	if err != nil && configPath == DefaultConfigPath{
		var file *os.File
		file,err = os.OpenFile(configPath,os.O_RDONLY|os.O_CREATE,0666)
		if err != nil {
			log.Println("Err：",err)
			os.Exit(-1)
		}
		_ = file.Close()
	}
	err = server.Init(configPath)
	if err != nil {
		log.Println("Err：",err)
		os.Exit(-1)
	}
	rootCmd.AddCommand(serverCmd)
}



