package shared

import (
	"os"

	"github.com/FileFormatInfo/fflint/internal/argtype"
	"github.com/mattn/go-isatty"

	"github.com/spf13/cobra"
)

var (
	cfgFile      string
	Debug        bool
	progress     bool
	showTotal    bool
	showFiles    = argtype.NewStringSet("Verbose", "none", []string{"all", "failing", "none"})
	showTests    = argtype.NewStringSet("Verbose", "failing", []string{"all", "failing", "none"})
	showDetail   bool
	failFast     bool
	OutputFormat = argtype.NewStringSet("OutputFormat", "text", []string{"text", "json", "markdown", "filenames"})
	fileSize     argtype.Range
	globber      Globber
	//ignoreExact    []string
	ignoreFile     string
	ignoreDotFiles bool
)

func AddCommon(rootCmd *cobra.Command) {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fflint.yaml)")

	//rootCmd.PersistentFlags().Int64Var(&minSize, "min", 0, "Minimum file size")
	//rootCmd.PersistentFlags().Int64Var(&maxSize, "max", 9999999999999, "Maximum file size")
	rootCmd.PersistentFlags().Var(&fileSize, "filesize", "Range of allowed file size")
	globber.Set("doublestar")
	rootCmd.PersistentFlags().Var(&globber, "glob", "How to expand [wildcards](/files.html) in file names [ doublestar | golang | none ]")
	rootCmd.PersistentFlags().StringVar(&ignoreFile, "ignore-file", DEFAULT_IGNORE_FILE, "ignore file")
	rootCmd.PersistentFlags().BoolVar(&ignoreDotFiles, "ignore-dotfiles", true, "Ignore files/directories starting with a dot (.)")
	//rootCmd.PersistentFlags().StringSliceVar(&ignoreExact, "ignore", []string{".git"}, "Specific files to ignore")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&showTotal, "show-totals", true, "Show total files tested, passed and failed")
	rootCmd.PersistentFlags().Var(&showFiles, "show-files", "Show each file "+showFiles.HelpText())
	rootCmd.PersistentFlags().Var(&showTests, "show-tests", "Show each test "+showTests.HelpText())
	rootCmd.PersistentFlags().BoolVar(&failFast, "fail-fast", false, "Stop as soon as any test fails")
	rootCmd.PersistentFlags().BoolVar(&showDetail, "show-detail", true, "Show detailed data about each test")
	rootCmd.PersistentFlags().BoolVar(&progress, "progress", isatty.IsTerminal(os.Stderr.Fd()), "Show progress bar (default is false when stderr is piped)")
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "Debugging output")
	rootCmd.PersistentFlags().VarP(&OutputFormat, "output", "o", "Output format "+OutputFormat.HelpText())

	//LATER: executable flag: OptionalBool

}
