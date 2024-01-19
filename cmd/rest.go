package cmd

import (
	"github.com/sifer169966/device-interactions/protocol"

	"github.com/spf13/cobra"
)

/*
	|--------------------------------------------------------------------------
	| Command to serve REST protocol
	|--------------------------------------------------------------------------
	|
	| Here is the command that makes your application serve the REST protocol
	| for client.
	|
*/

var serveRESTCmd = &cobra.Command{
	Use:   "serve-rest",
	Short: "start a http server",
	RunE:  serveREST,
}

func serveREST(cmd *cobra.Command, args []string) error {
	return protocol.ServeREST()
}

func init() {
	rootCmd.AddCommand(serveRESTCmd)
}
