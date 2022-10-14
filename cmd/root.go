package cmd

import (
	"cli/util"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile = "./cobra"
	RootCmd = &cobra.Command{
		Use:   "cli",
		Short: "cli short description in help output",
		Long:  "cli detail message in help <specified command>",
		Example: `cli for example:
		cli <command> <argument> <flag>
		`,
		// ValidArgs []string
		// Args PositionalArgs
		// PersistentPreRun: children of this command will inherit and execute.
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		// PersistentPreRunE: PersistentPreRun but returns an error.
		// PersistentPreRunE: func(cmd *cobra.Command, args []string) {},
		// PreRun: children of this command will not inherit.
		PreRun: func(cmd *cobra.Command, args []string) {},
		// PreRunE: PreRun but returns an error.
		// PreRunE: func(cmd *cobra.Command, args []string) {},
		Run: func(cmd *cobra.Command, args []string) {

			fversion, err := cmd.Flags().GetBool("version")
			cobra.CheckErr(err)
			if fversion {
				fmt.Println("FullVersion()")
				return
			}

			foption, err := cmd.Flags().GetBool("option")
			cobra.CheckErr(err)
			if foption {
				if v, ok := util.CommandList[cmd.Use]; ok {
					fmt.Println("util.CommandList is ok", v)
				} else {
					fmt.Println("this command has not option list")
				}
				return
			}

			fmt.Println("============= Nothing happened... =========")
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.Flags().BoolP("version", "v", false, "version info")
	RootCmd.Flags().BoolP("option", "o", false, "option list for this command")

	RootCmd.Flags().StringVarP(&cfgFile, "cfg", "f", "./cobra", "config file default is .cobra.yaml")

	// RootCmd.Flags().Bool("viper", true, "use Viper for configuration")

	// viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("viper", RootCmd.PersistentFlags().Lookup("viper"))

	// viper.SetDefault("author", "zhxu xuzhenghao@jingkun.com")
	// viper.SetDefault("license", "jingkun")
}

func initConfig() {
	if cfgFile != "" {
		// SetConfigFile explicitly defines the path
		viper.SetConfigType("yaml")
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// Search config in home directory with name ".my-calc" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	// If matching env vars are found, they are loaded into Viper.
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	cobra.CheckErr(err)
	// return config file path
	// fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
