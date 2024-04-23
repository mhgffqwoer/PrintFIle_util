package PrintFile

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "PrintFile",
	Long:  "My PrintFile utility for Linux and MacOS",
	Run:   func(cmd *cobra.Command, args []string) {
		file, _ := cmd.PersistentFlags().GetString("file")
		tail, _ := cmd.Flags().GetBool("tail")
		delimiter, _ := cmd.Flags().GetString("delimiter")
		lines, _ := cmd.Flags().GetInt64("lines")
		PrintFile(file, tail, delimiter, lines)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", "", "Path to file. [required parameter]")
	rootCmd.Flags().BoolP("tail", "t", false, "Tail mode.")
	rootCmd.Flags().StringP("delimiter", "d", "\n", "Delimiter.")
	rootCmd.Flags().Int64P("lines", "l", 0, "Number of lines to print.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}