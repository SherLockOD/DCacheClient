package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"DCacheClient/demo/cmd/subCmd"
)

var name string
var age int
//var key string
//var value string
//var ext int

var RootCmd = &cobra.Command{
	Use: 	"demo",
	Short: 	"A test demo",
	Long:	`Demo is a test appcation for print things`,
	Run:	func(cmd *cobra.Command, args []string) {
		//fmt.Println("Cmd: ", cmd, " Args: ", args)
		if len(name) == 0 {
			cmd.Help()
			return
		}
		subCmd.Show(name, age)
	},
}
var Set = &cobra.Command{
	Use:	"Set [key] [value] [expiretime]",
	Short:	"DCache command",
	Long:	`This command can get value from DCache with key name`,
	Run:	func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			cmd.Help()
			return
		}
		fmt.Printf("set %v: %v : %v\n", args[0], args[1], args[2])
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
func init() {
	RootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	RootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
	Set.Flags().StringVarP(&key, "key", "", "","key name")
	Set.Flags().StringVarP(&value, "value", "", "", "value name")
	Set.Flags().IntVarP(&ext, "ext", "", 0, "expire time")
	RootCmd.AddCommand(Set)
}
*/

func init() {
	RootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	RootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
	RootCmd.AddCommand(Set)
}