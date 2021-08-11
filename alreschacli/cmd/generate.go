/*
MIT License

Copyright (c) 2021 unsafe-risk

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	goalrescha "github.com/unsafe-risk/go-alrescha"
)

var InputFileName, OutputFileName, OutputLang, OutputPackage *string
var InputStdin, UnSafeModes *bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates code based on Alrescha definition files",
	Long:  `Generates code based on Alrescha definition files`,
	Run: func(cmd *cobra.Command, args []string) {
		if *InputFileName == "" && !*InputStdin {
			cmd.PrintErrln("Error: No input file specified")
			cmd.PrintErrln("Usage: alreschacli generate -i <inputfile> -o <outputfile>")
			return
		}
		if *OutputFileName == "" {
			cmd.PrintErrln("Error: No output file specified")
			cmd.PrintErrln("Usage: alreschacli generate -i <inputfile> -o <outputfile>")
			return
		}
		if *OutputLang == "" {
			langPrompt := promptui.Select{
				Label: "Select Language",
				Items: []string{"Go", "Python", "Rust", "Java", "C#", "JavaScript", "TypeScript", "Dart"},
			}
			_, lang, err := langPrompt.Run()
			if err != nil {
				cmd.PrintErrln("Error: ", err)
				return
			}
			*OutputLang = lang
		}
		if *OutputPackage == "" {
			re := regexp.MustCompile(`[a-zA-Z0-9_]+`)
			pkgPrompt := promptui.Prompt{
				Label: "Package Name",
				Validate: func(s string) error {
					if len(s) == 0 {
						return fmt.Errorf("package name cannot be empty")
					}
					if re.ReplaceAllString(s, "") != "" {
						return fmt.Errorf("package name must contain only letters, numbers and underscores")
					}
					return nil
				},
			}
			packageName, err := pkgPrompt.Run()
			if err != nil {
				cmd.PrintErrln("Error: ", err)
				return
			}
			*OutputPackage = packageName
		}
		var inputFile *os.File
		if *InputStdin {
			inputFile = os.Stdin
		} else {
			var err error
			inputFile, err = os.Open(*InputFileName)
			if err != nil {
				cmd.PrintErrln("Error: Could not open input file")
				return
			}
			defer inputFile.Close()
		}
		inputFileData, err := io.ReadAll(inputFile)
		if err != nil && err != io.EOF {
			cmd.PrintErrln("Error: Could not read input file")
			cmd.PrintErrln(err)
			return
		}
		var output []byte
		switch *OutputLang {
		case "Go":
			output, err = goalrescha.MakeGoCode(inputFileData, *OutputPackage)
			if err != nil {
				cmd.PrintErrln("Error: ", err)
				return
			}
		}
		outputFile, err := os.Create(*OutputFileName)
		if err != nil {
			cmd.PrintErrln("Error: Could not create output file")
			return
		}
		defer outputFile.Close()
		_, err = outputFile.Write(output)
		if err != nil {
			cmd.PrintErrln("Error: Could not write output file")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	InputFileName = generateCmd.Flags().StringP("input", "i", "", "Alrescha definition file")
	OutputFileName = generateCmd.Flags().StringP("output", "o", "", "Output file")
	generateCmd.MarkFlagRequired("output")
	InputStdin = generateCmd.Flags().Bool("input-stdin", false, "Read Alrescha definition file from stdin")
	UnSafeModes = generateCmd.Flags().Bool("unsafe", false, "Enable unsafe mode")
	OutputLang = generateCmd.Flags().StringP("lang", "l", "", "Language to generate code for")
	OutputPackage = generateCmd.Flags().StringP("package-name", "p", "", "Package name for generated code")
}
