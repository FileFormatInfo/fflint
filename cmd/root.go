package cmd

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	debug        bool
	progress     bool
	showTotal    bool
	showFiles    bool
	showTests    bool
	showDetail   bool
	showPassing  bool
	outputFormat string
	fileSize     Range
	globber      Globber
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "badger",
	Short: "Badgers you if your file formats are invalid",
	Long:  `See [www.badger.sh](https://www.badger.sh/) for detailed instructions`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to execute root command: %s\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.badger.yaml)")

	//rootCmd.PersistentFlags().Int64Var(&minSize, "min", 0, "Minimum file size")
	//rootCmd.PersistentFlags().Int64Var(&maxSize, "max", 9999999999999, "Maximum file size")
	rootCmd.PersistentFlags().Var(&fileSize, "size", "Range of allowed file size")
	rootCmd.PersistentFlags().Var(&globber, "glob", "Glob algorith to use")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&showTotal, "showTotal", true, "Show total files tested/passed/failed")
	rootCmd.PersistentFlags().BoolVarP(&showFiles, "showFiles", "f", !isatty.IsTerminal(os.Stderr.Fd()), "Show each file tested")
	rootCmd.PersistentFlags().BoolVarP(&showTests, "showTests", "t", false, "Show each test performed")
	rootCmd.PersistentFlags().BoolVar(&showPassing, "showPassing", false, "Show passing files/tests")
	rootCmd.PersistentFlags().BoolVar(&showDetail, "showDetail", true, "Show detailed data about each test")
	rootCmd.PersistentFlags().BoolVar(&progress, "progress", isatty.IsTerminal(os.Stderr.Fd()), "Show progress bar")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Debugging output")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "Output format [ json | text ]")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to get home directory: %s\n", err.Error())
			os.Exit(1)
		}

		// Search config in home directory with name ".badger" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".badger")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
