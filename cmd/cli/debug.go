package main

import (
	"context"
	"github.com/anytypeio/any-sync/app"
	"github.com/anytypeio/any-sync/util/crypto"
	"github.com/anytypeio/go-anytype-middleware/core/anytype"
	"github.com/anytypeio/go-anytype-middleware/core/debug"
	"github.com/anytypeio/go-anytype-middleware/core/event"
	"github.com/anytypeio/go-anytype-middleware/pb"
	"github.com/anytypeio/go-anytype-middleware/util/console"
	"github.com/spf13/cobra"
	"os"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug commands",
}

var (
	debugRepoPath   string
	debugMnemonic   string
	debugTree       string
	debugOutputFile string
)

var dumpTree = &cobra.Command{
	Use:   "dump-tree",
	Short: "Dumps tree of changes for specific thread",
	Run: func(c *cobra.Command, args []string) {
		if debugMnemonic == "" {
			console.Fatal("please specify account")
		}
		if debugTree == "" {
			console.Fatal("please specify tree")
		}
		acc, err := crypto.Mnemonic(debugMnemonic).DeriveEd25519Key(0)
		if err != nil {
			panic(err)
		}
		comps := []app.Component{
			anytype.BootstrapConfig(false, false, true),
			anytype.BootstrapWallet(debugRepoPath, acc),
			event.NewCallbackSender(func(event *pb.Event) {}),
		}

		app, err := anytype.StartNewApp(context.Background(), comps...)
		if err != nil {
			console.Fatal("failed to start anytype: %s", err.Error())
		}

		dbg := app.MustComponent(debug.CName).(debug.Debug)

		isAnonymize := false
		dumpWithSvg := false
		filename, err := dbg.DumpTree(debugTree, debugOutputFile, isAnonymize, dumpWithSvg)
		if err != nil {
			console.Fatal("failed to dump tree: %s", err.Error())
		}
		console.Success("file saved: %s", filename)
	},
}
var dumpLocalstore = &cobra.Command{
	Use:   "dump-localstore",
	Short: "Dumps localstore for all objects",
	Run: func(c *cobra.Command, args []string) {
		if debugMnemonic == "" {
			console.Fatal("please specify account")
		}
		acc, err := crypto.Mnemonic(debugMnemonic).DeriveEd25519Key(0)
		if err != nil {
			panic(err)
		}
		comps := []app.Component{
			anytype.BootstrapConfig(false, false, true),
			anytype.BootstrapWallet(debugRepoPath, acc),
			event.NewCallbackSender(func(event *pb.Event) {}),
		}

		app, err := anytype.StartNewApp(context.Background(), comps...)
		if err != nil {
			console.Fatal("failed to start anytype: %s", err.Error())
		}

		dbg := app.MustComponent(debug.CName).(debug.Debug)

		filename, err := dbg.DumpLocalstore(nil, debugOutputFile)
		if err != nil {
			console.Fatal("failed to dump localstore: %s", err.Error())
		}
		console.Success("file saved: %s", filename)
	},
}

func init() {
	// subcommands
	homeDir, _ := os.UserHomeDir()

	debugCmd.AddCommand(dumpTree)
	debugCmd.AddCommand(dumpLocalstore)

	debugCmd.PersistentFlags().StringVarP(&debugRepoPath, "repo", "r", homeDir+"/.config/anytype2/data", "path to dir with accounts folder")
	debugCmd.PersistentFlags().StringVarP(&debugMnemonic, "mnemonic", "m", "", "account mnemonic")
	debugCmd.PersistentFlags().StringVarP(&debugTree, "tree", "t", "", "id of tree to debug")
	debugCmd.PersistentFlags().StringVarP(&debugOutputFile, "out", "o", "./", "folder to save file")
}
