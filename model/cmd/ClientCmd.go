package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	opt = make(map[string]string)
)

var get = &cobra.Command{
	Use:	"get [key]",
	Short:	"Get key value from DCache",
	Long:	`This command will get value from DCache with key name`,
	Run:	getCmd,
}

var gets = &cobra.Command{
	Use:	"gets [key1 key2 key3 ...]",
	Short:	"Get keys from DCache",
	Long:	`This command will get values from DCache with keys`,
	Run:	getsCmd,
}

var getkeys = &cobra.Command{
	Use:	"getkeys",
	Short:	"Get all the keys from DCache",
	Long:	`This command will get all the keys from DCache`,
	Run:	getAllCmd,
}

var check = &cobra.Command{
	Use:	"check [key1 key2 key3 ...]",
	Short:	"Check keys from DCache",
	Long:	`This command will check the keys status from DCache`,
	Run:	checkCmd,
}

var set = &cobra.Command{
	Use:	"set [key] [value]",
	Short:	"Set the key and value into DCache",
	Long:	`This command will set the key and value into DCache`,
	Run:	setCmd,
}

var insert = &cobra.Command{
	Use:	"insert [key] [value]",
	Short:	"Insert the key and value into DCache",
	Long:	`This command will insert the key and value into DCache 
if the key does not exist otherwise return failure.`,
	Run:	insertCmd,
}

var sets = &cobra.Command{
	Use:	"sets [key1:value1 key2:value2 key3:value3 ...]",
	Short:	"Set some keys and values into DCache",
	Long:	`This command will set some keys and values into DCache. 
Example: sets name:tom age:20 ...`,
	Run:	setsCmd,
}

var del = &cobra.Command{
	Use:	"del [key]",
	Short:	"Delete key value from DCache",
	Long:	`This command will delete the value from DCache with key name`,
	Run:	delCmd,
}

var dels = &cobra.Command{
	Use:	"dels [key1 key2 key3 ...]",
	Short:	"Delete keys and values from DCache",
	Long:	`This command will delete values from DCache with keys`,
	Run:	delsCmd,
}

func cmdLineParse(ss []string) (k, v []string) {
	for _, s := range ss {
		i := strings.Split(s, ":")
		k = append(k, i[0])
		v = append(v, i[1])
	}
	return
}

func getCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	key := args[0]
	ret, value, err := r.GetKV(key, opt)
	//fmt.Println(value, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch ret {
	case 0:
		fmt.Println(value)
	case -6:
		fmt.Printf("not exist about the key '%v'\n", key)
	default:
		fmt.Println("ret: ", ret)
	}
}

func getsCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	keys := args
	_, values, err := r.GetKVBatch(keys, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", values)
}

func getAllCmd(cmd *cobra.Command, args []string) {
	// TODO 看是否有能够获取全部hash桶个数的接口
	var resp []string
	s := int32(index)
	e := int32(count)
	for {
		_, values, isEnd, err := r.GetAllKeys(s, e, opt)
		if err != nil {
			fmt.Println(err)
			return
		}
		if ! isEnd {
			s, e = e, e * 2
			resp = append(resp, values...)
		} else {
			var ss []string
			resp = append(resp, values...)
			uniqueMap := make(map[string]int)
			for _, v := range resp {
				uniqueMap[v] = 1
			}
			for k := range uniqueMap{
				ss = append(ss, k)
			}
			fmt.Println(strings.Join(ss, " "))
			return
		}
	}
}

func checkCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	keys := args
	_, value, err := r.CheckKey(keys, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}

func setCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	key := args[0]
	value := args[1]
	ret, err := r.SetKV(key, value, dirty, int32(ext), opt)
	if err != nil {
		fmt.Printf("ret: %v, err: %v\n", ret, err)
		return
	}
	switch ret {
	case 0:
		fmt.Printf("ret: %v, resp: ok\n", ret)
	default:
		fmt.Println(ret)
	}
}

func insertCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	key := args[0]
	value := args[1]
	ret, err := r.InsertKV(key, value, dirty, int32(ext), opt)
	if err != nil {
		fmt.Printf("ret: %v, err: %v\n", ret, err)
		return
	}
	switch ret {
	case 0:
		fmt.Printf("ret: %v, resp: ok\n", ret)
	case -10:
		fmt.Printf("ret: %v, resp: the key already exists\n", ret)
	default:
		fmt.Println(ret)
	}
}

func setsCmd(cmd *cobra.Command, args []string) {
	var _ds []bool
	var _es []int32

	// sets key-value example = key:value
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}

	for _, v := range args {
		if ! strings.ContainsAny(v, ":") {
			_ = cmd.Help()
			return
		}
	}

	switch dirtys {
	case true:
		_ds = func() []bool {
			d := make([]bool, len(args))
			for i := 0; i < len(args); i++ {
				d[i] = true
			}
			return d
		}()
	case false:
		_ds = make([]bool, len(args))
	}

	switch exts {
	case 0:
		_es = make([]int32, len(args))
	default:
		_es = func() []int32 {
			e := make([]int32, len(args))
			for i := 0; i < len(args); i++ {
				e[i] = int32(exts)
			}
			return e
		}()
	}
	keys, values := cmdLineParse(args)
	ret, m, err := r.SetKVBatch(keys, values, _ds, _es, opt)
	if err != nil {
		fmt.Printf("ret: %v, err: %v", ret, err)
	}
	if ret == 0 {
		fmt.Printf("ret: %v, resp: ok\n", ret)
	} else {
		fmt.Printf("ret: %v, resp: %v\n", ret, m)
	}
}

/* the method does not pass testing
func updateCmd(cmd *cobra.Command, args []string) {
}
 */

func delCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	key := args[0]
	ret, err := r.DelKV(key, opt)
	if err != nil {
		fmt.Printf("ret: %v, err: %v\n", ret, err)
		return
	}
	fmt.Printf("ret: %v, resp: ok\n", ret)
}

func delsCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Help()
		return
	}
	keys := args
	ret, m, err := r.DelKVBatch(keys, opt)
	if err != nil {
		fmt.Printf("ret: %v, err； %v\n", ret, err)
		return
	}
	fmt.Printf("ret: %v, resp: %+v\n", ret, m)
}
