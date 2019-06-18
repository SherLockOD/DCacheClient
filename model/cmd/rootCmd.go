package cmd

import (
	"DCacheClient/model/client"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	index		int
	count		int
	dirty		bool
	dirtys		bool
	ext			int
	exts		int
	basePath	string
	r			*client.Req
)

var RootCmd = &cobra.Command{
	Use:	"DCacheClient",
	Short:	"A DCache Client CommandLine",
	Long:	`This is a CommandLine Tools for DCache Client`,
	Run:	func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
	},
}


func init() {
	getkeys.Flags().IntVarP(&index, "index" ,"i",0,"set start number")
	getkeys.Flags().IntVarP(&count, "count" ,"c",100000,"set end number")
	set.Flags().BoolVarP(&dirty, "dirty", "d",true, "set dirty data switch")
	set.Flags().IntVarP(&ext, "expireTime", "e",0, "set expire time")
	sets.Flags().BoolVarP(&dirtys, "all-the-kv-dirty", "d", true, "set all the kv dirty")
	sets.Flags().IntVarP(&exts, "all-the-kv-expire", "e", 0, "set all the kv expire times")
	RootCmd.AddCommand(get)
	RootCmd.AddCommand(gets)
	RootCmd.AddCommand(getkeys)
	RootCmd.AddCommand(check)
	RootCmd.AddCommand(set)
	RootCmd.AddCommand(sets)
	RootCmd.AddCommand(insert)
	RootCmd.AddCommand(del)
	RootCmd.AddCommand(dels)
}


func Execute(req *client.Req) {
	r = req
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
