package cmd

import (
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "1.0.0"
var cfgFile string
var component string
var directory string
var indexFile string
var useJavascript bool

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//go:embed templates/component.tsx.tmpl
var componentTsTemplate string

//go:embed templates/component.jsx.tmpl
var componentJsTemplate string

//go:embed templates/index.ts.tmpl
var indexTsTemplate string

//go:embed templates/index.js.tmpl
var indexJsTemplate string

var rootCmd = &cobra.Command{
	Use:     "rgc",
	Short:   "A component generator for React",
	Version: version,
	Long: `rgc is a component generator for React. It generates a the following:
				- a file including a component with a default export ({NewComponent}.{tsx|jsx})
				- an index.{ts|js} file that exports the component
				- an entry in the index.{ts|js} file in the root of the components directory
					that exports the component
				`,
	Run: func(cmd *cobra.Command, args []string) {
		if component == "" {
			log.Fatal("component name is required")
		}

		if directory == "" {
			log.Fatal("component directory is required")
		}

		fileExtension := "ts"
		componentFileTemplate := componentTsTemplate
		indexFileTemplate := indexTsTemplate
		if useJavascript {
			fileExtension = "js"
			componentFileTemplate = componentJsTemplate
			indexFileTemplate = indexJsTemplate
		}

		if indexFile == "" {
			indexFile = fmt.Sprintf("%s/index.%s", directory, fileExtension)
		}

		componentFilePath := fmt.Sprintf("%s/%s/%s.%s", directory, component, component, fileExtension+"x")
		componentIndexFilePath := fmt.Sprintf("%s/%s/index.%s", directory, component, fileExtension)
		indexFilePath := fmt.Sprintf("%s/index.%s", directory, fileExtension)

		// Check if the component directory exists
		if _, err := os.Stat(fmt.Sprintf("%s/%s", directory, component)); !os.IsNotExist(err) {
			log.Fatalf("Component %s already exists", component)
		}

		// Create the component directory
		err := os.MkdirAll(fmt.Sprintf("%s/%s", directory, component), fs.FileMode(0755))
		if err != nil {
			log.Fatal(err)
		}

		// Create the component file
		componentFileHandle, err := os.Create(componentFilePath)
		if err != nil {
			log.Fatal(err)
		}

		componentTemplate, err := template.New("component").Parse(componentFileTemplate)
		if err != nil {
			log.Fatal(err)
		}

		err = componentTemplate.ExecuteTemplate(componentFileHandle, "component", component)
		if err != nil {
			log.Fatal(err)
		}
		componentFileHandle.Close()

		// Create the component index file
		componentIndexFileHandle, err := os.Create(componentIndexFilePath)
		if err != nil {
			log.Fatal(err)
		}

		componentIndexTemplate, err := template.New("index").Parse(indexFileTemplate)
		if err != nil {
			log.Fatal(err)
		}

		err = componentIndexTemplate.Execute(componentIndexFileHandle, component)
		if err != nil {
			log.Fatal(err)
		}
		componentIndexFileHandle.Close()

		// Update the index file
		info, err := os.Stat(indexFilePath)
		if err != nil {
			log.Fatal(err)
		}

		mode := info.Mode()
		indexFileHandle, err := os.OpenFile(indexFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode)
		if err != nil {
			log.Fatal(err)
		}

		_, err = indexFileHandle.WriteString(fmt.Sprintf("export { default as %s } from './%s'\n", component, component))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Created %s\n", componentFilePath)
		fmt.Printf("Created %s\n", componentIndexFilePath)
		fmt.Printf("Updated %s\n", indexFile)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .rgc.yaml)")
	rootCmd.PersistentFlags().StringVarP(&component, "component", "c", "", "component name (required)")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", "", "component directory")
	rootCmd.PersistentFlags().StringVarP(&indexFile, "index", "i", "", "index file (default is {directory}/index.{ts|js})")
	rootCmd.PersistentFlags().BoolVar(&useJavascript, "use-javascrpt", false, "use javascript instead of typescript (default is false)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rgc")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	directory = viper.GetString("directory")

	if directory != "" {
		rootCmd.PersistentFlags().Set("directory", directory)
	}
}
