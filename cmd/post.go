package cmd

import (
	"cli/model"
	"cli/util"
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

// PostCmd represents the post command
var PostCmd = &cobra.Command{
	Use:   util.COMMAND_POST,
	Short: "post file includes nexus harbor image rancher certs kubectl...",
	Long: `post file For example:
		cli post nexus
		cli post certs
		cli post kubectl
		cli post rancher
		`,
	Example: `cli post for example:
	cli post nexus  <flag>
	cli post kubectl  <flag>
	cli post rancher  <flag>
	`,
	Args:      cobra.RangeArgs(0, 4),
	ValidArgs: []string{"nexus", "harbor", "image", "rancher", "certs", "kubectl"},
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
				fmt.Println("post option", v)
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
				fmt.Println(fmt.Errorf("post Lstat: %s", err.Error()))
				return
			}
			data = model.Do(args, cmd.Use, ffile)
		}
		fmt.Println(data)
		fmt.Println("============= rsp callback successfully ========= client end ...")
	},
}

func init() {
	RootCmd.AddCommand(PostCmd)
	PostCmd.Flags().BoolP("skip", "s", false, "Help message for skip")
	PostCmd.Flags().BoolP("option", "o", false, "Help message for option")

	PostCmd.Flags().BoolP("result", "r", false, "Help message for result")
	PostCmd.Flags().StringP("file", "f", "./userdata.yaml", "Help message for file")
}
