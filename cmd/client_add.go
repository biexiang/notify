package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"notify/logic/client"
	"os"
)

var (
	title string
	content string
	way string
	time string

	clientAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Use to add reminder",
		Long: `Use to add reminder`,
		Run: func(cmd *cobra.Command, args []string) {
			err := client.AddNotify(title,content)
			if err != nil {
				log.Println("AddNotify Failed：",err)
				os.Exit(-1)
			}
		},
	}
)

func init() {
	clientAddCmd.PersistentFlags().StringVar(&title,"title","有人说","just title of notify")
	clientAddCmd.PersistentFlags().StringVar(&content,"content","加油","just title of content")
	clientAddCmd.PersistentFlags().StringVar(&way,"way","specify","specify means the time,interval means seconds")
	clientAddCmd.PersistentFlags().StringVar(&time,"time","2019-12-31 10:00:00","just time")

	if title == "" {
		_  = clientAddCmd.Usage()
	}
	clientCmd.AddCommand(clientAddCmd)
}











