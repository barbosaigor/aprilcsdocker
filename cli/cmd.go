package cli

import (
	"github.com/barbosaigor/april/destroyer"
	csd "github.com/barbosaigor/dockercs"
	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var username string
var password string
var port int

func init() {
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	rootCmd.Flags().StringVarP(&password, "password", "s", "", "Password")
	rootCmd.Flags().IntVarP(&port, "port", "p", 7071, "Server port number")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
}

var rootCmd = &cobra.Command{
	Use:   "dockercs",
	Short: "Docker chaos server terminates instances via API.",
	Long:  "Docker chaos server terminates instances via API.",
	Run: func(cmd *cobra.Command, args []string) {
		serv := destroyer.New(port, csd.DockerDestryr{})
		serv.Cred.Register(username, password)
		serv.Serve()
	},
	Version: VERSION,
}

func Execute() error {
	return rootCmd.Execute()
}
