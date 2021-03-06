package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	componentNamePtr := flag.String("c", "", "New component name. (Required)")
	destinationPtr := flag.String("d", "./", "Destination directory.")
	dryRunPtr := flag.Bool("dry", false, "Run command as dry run (no filesystem changes).")
	flag.Parse()

	if *componentNamePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	rootPath := filepath.Join(*destinationPtr, *componentNamePtr)
	if !(*dryRunPtr) {
		err := os.Mkdir(rootPath, 0755)
		check(err)
	}
	fmt.Printf("Created %s\n", rootPath)

	componentPath := filepath.Join(rootPath, fmt.Sprintf("%s.tsx", *componentNamePtr))
	componentTemplate := []byte(fmt.Sprintf(`export default function %s() {
	return (

	)
}
`, *componentNamePtr))

	if !(*dryRunPtr) {
		err := ioutil.WriteFile(componentPath, componentTemplate, 0644)
		check(err)
	}
	fmt.Printf("Created %s, wrote %d bytes\n", componentPath, len(componentTemplate))

	indexPath := filepath.Join(rootPath, "index.ts")
	indexTemplate := []byte(fmt.Sprintf(`export { default } from './%s'
`, *componentNamePtr))

	if !(*dryRunPtr) {
		err := ioutil.WriteFile(indexPath, indexTemplate, 0644)
		check(err)
	}
	fmt.Printf("Created %s, wrote %d bytes\n", indexPath, len(indexTemplate))
}
