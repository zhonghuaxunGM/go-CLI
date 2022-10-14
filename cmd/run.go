package cmd

import (
	"cli/model"
	"cli/util"
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

// RunCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   util.COMMAND_RUN,
	Short: "run file includes nexus harbor image rancher kubectl...",
	Long: `run For example:
		cli run nexus
		cli run kubectl
		cli run rancher
		`,
	Example: `cli run for example:
	cli run nexus  <flag>
	cli run kubectl  <flag>
	cli run rancher  <flag>
	`,
	Args:      cobra.RangeArgs(0, 3),
	ValidArgs: []string{"nexus", "harbor", "image", "rancher", "kubectl"},
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
				fmt.Println("run option", v)
				fmt.Println("============= The above shows the required information ========= client end ...")
			} else {
				fmt.Println("============= This command has not option list ========= client end ...")
			}
			return
		}

		for _, v := range args {
			if !util.ContainsArray(cmd.ValidArgs, v) {
				fmt.Printf("WARNING: args %s not in valid Args, please <command> -o for help\n", v)

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
				fmt.Println(fmt.Errorf("run Lstat: %s", err.Error()))
				return
			}
			data = model.Do(args, cmd.Use, ffile)
		}
		fmt.Println(data)
		fmt.Println("============= rsp callback successfully ========= client end ...")
	},
}

func init() {
	RootCmd.AddCommand(RunCmd)
	RunCmd.Flags().BoolP("skip", "s", false, "Help message for skip")
	RunCmd.Flags().BoolP("option", "o", false, "Help message for option")

	RunCmd.Flags().BoolP("result", "r", false, "Help message for result")
	RunCmd.Flags().StringP("file", "f", "./userdata.yaml", "Help message for file")
}
