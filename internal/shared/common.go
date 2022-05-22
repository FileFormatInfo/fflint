package shared

import (
	"fmt"
	"os"

	"github.com/fileformat/badger/internal/argtype"
	"github.com/mattn/go-isatty"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	Debug        bool
	progress     bool
	showTotal    bool
	showFiles    bool
	showTests    bool
	showDetail   bool
	showPassing  bool
	OutputFormat string
	fileSize     argtype.Range
	globber      Globber
)

func AddCommon(rootCmd *cobra.Command) {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.badger.yaml)")

	//rootCmd.PersistentFlags().Int64Var(&minSize, "min", 0, "Minimum file size")
	//rootCmd.PersistentFlags().Int64Var(&maxSize, "max", 9999999999999, "Maximum file size")
	rootCmd.PersistentFlags().Var(&fileSize, "filesize", "Range of allowed file size")
	globber.Set("doublestar")
	rootCmd.PersistentFlags().Var(&globber, "glob", "How to expand [wildcards](/files.html) in file names [ doublestar &#x7c; golang &#x7c; none ]")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&showTotal, "showTotal", true, "Show total files tested, passed and failed")
	rootCmd.PersistentFlags().BoolVarP(&showFiles, "showFiles", "f", !isatty.IsTerminal(os.Stderr.Fd()), "Show each file tested (default is false when stderr is piped)")
	rootCmd.PersistentFlags().BoolVarP(&showTests, "showTests", "t", false, "Show each test performed")
	rootCmd.PersistentFlags().BoolVar(&showPassing, "showPassing", false, "Show passing files and tests")
	rootCmd.PersistentFlags().BoolVar(&showDetail, "showDetail", true, "Show detailed data about each test")
	rootCmd.PersistentFlags().BoolVar(&progress, "progress", isatty.IsTerminal(os.Stderr.Fd()), "Show progress bar (default is false when stderr is piped)")
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "Debugging output")
	rootCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "text", "Output format [ json &#x7c; text &#x7c; filenames ]")
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
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
