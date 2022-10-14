package cmd

import (
	"cli/model"
	"cli/util"
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

// CheckCmd represents the check command
var CheckCmd = &cobra.Command{
	Use:   util.COMMAND_CHECK,
	Short: "check userdata or systeminfo",
	Long: `check userdata or systeminfo. For example:
		cli check userdata
		cli check system
	`,
	Example: `cli check for example:
	cli check userdata  <flag>
	cli check system  <flag>
	`,
	Args:      cobra.RangeArgs(0, 2),
	ValidArgs: []string{"userdata", "system"},

	Run: func(cmd *cobra.Command, args []string) {

		fskip, err := cmd.Flags().GetBool("skip")
		cobra.CheckErr(err)
		if fskip {
			fmt.Println("============= This stage has skipped successfully ============= client end ...")
			return
		}

		foption, err := cmd.Flags().GetBool("option")
		cobra.CheckErr(err)
		if foption {
			if v, ok := util.CommandList[cmd.Use]; ok {
				fmt.Println("check option", v)

				fmt.Println("============= The above shows the required information ========= client end ...")
			} else {
				fmt.Println("============= This command has not option list ========= client end ...")
			}
			return
		}

		for _, v := range args {
			if !util.ContainsArray(cmd.ValidArgs, v) {
				fmt.Printf("WARNING: args %s not in valid Args, please <command> -o for help\n", v)
				return
			}
		}

		var ffile string
		fresult, err := cmd.Flags().GetBool("result")
		cobra.CheckErr(err)

		var data interface{}
		if fresult {
			data = model.GetStat(args, cmd.Use)
		} else {
			ffile, err = cmd.Flags().GetString("file")
			cobra.CheckErr(err)
			_, err := os.Lstat(ffile)
			if err != nil {
				fmt.Println(fmt.Errorf("check Lstat: %s", err.Error()))
				return
			}
			data = model.Do(args, cmd.Use, ffile)
		}
		fmt.Println(data)
		fmt.Println("============= rsp callback successfully ========= client end ...")
	},
}

func init() {
	RootCmd.AddCommand(CheckCmd)
	CheckCmd.Flags().BoolP("skip", "s", false, "Help message for skip")
	CheckCmd.Flags().BoolP("option", "o", false, "Help message for option")

	CheckCmd.Flags().BoolP("result", "r", false, "Help message for result")
	CheckCmd.Flags().StringP("file", "f", "./userdata.yaml", "Help message for file")
}
