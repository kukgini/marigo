// Package marigo (Aries Agent REST Server) of aries-framework-go.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	"github.com/kukgini/marigo/mylog"
	"github.com/spf13/cobra"
)

// This is an application which starts Aries agent controller API on given port.
func main() {
	fmt.Printf("begin main")

	rootCmd := &cobra.Command{
		Use: "marigo",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	fmt.Printf("rootCmd created %v", rootCmd)

	log.Initialize(mylog.NewProvider())

	logger := log.New("marigo")

	logger.Debugf("create start command")
	logger.Panicf("stop here!!!")

	// startCmd, err := startcmd.Cmd(&startcmd.HTTPServer{})
	// if err != nil {
	// 	logger.Fatalf(err.Error())
	// }

	// rootCmd.AddCommand(startCmd)

	// if err := rootCmd.Execute(); err != nil {
	// 	logger.Fatalf("Failed to run marigo: %s", err)
	// }
}
