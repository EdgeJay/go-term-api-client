package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootApp struct {
	cfgFile string
	cmd     *cobra.Command
}

// Returns new instance of cli RootApp that acts as base app of
// cli interface.
func NewRootApp(appName, shortDesc, longDesc string) *RootApp {
	cmd := &cobra.Command{
		Use:   appName,
		Short: shortDesc,
		Long:  longDesc,
	}

	return &RootApp{
		cmd: cmd,
	}
}

// Must call this before running calling Execute() method.
func (app *RootApp) Configure() {
	app.cmd.PersistentFlags().StringVar(
		&app.cfgFile,
		"config",
		"",
		fmt.Sprintf("config file (default is $HOME/.%s.yaml)", app.cmd.Use),
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cobra.OnInitialize(func() {

		if app.cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(app.cfgFile)
		} else {
			// Find home directory.
			home, err := os.UserHomeDir()
			cobra.CheckErr(err)

			// Search config in home directory with name ".go-term-api-client" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigType("yaml")
			viper.SetConfigName(fmt.Sprintf(".%s", app.cmd.Use))
		}

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	})
}

// Add command to app
func (app *RootApp) AddCommand(cmdName, shortDesc, longDesc string, fn func(args []string)) {
	cmd := &cobra.Command{
		Use:   cmdName,
		Short: shortDesc,
		Long:  longDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fn(args)
		},
	}
	app.cmd.AddCommand(cmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func (app *RootApp) Execute() {
	err := app.cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
