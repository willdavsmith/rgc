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
		err := os.MkdirAll(rootPath, 0755)
		check(err)
	}

	absoluteRootPath, err := filepath.Abs(rootPath)
	check(err)

	fmt.Printf("Created %s\n", absoluteRootPath)

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

	absoluteComponentPath, err := filepath.Abs(componentPath)
	check(err)

	fmt.Printf("Created %s, wrote %d bytes\n", absoluteComponentPath, len(componentTemplate))

	indexPath := filepath.Join(rootPath, "index.ts")
	indexTemplate := []byte(fmt.Sprintf(`export { default } from './%s'
`, *componentNamePtr))

	if !(*dryRunPtr) {
		err := ioutil.WriteFile(indexPath, indexTemplate, 0644)
		check(err)
	}

	absoluteIndexPath, err := filepath.Abs(indexPath)
	check(err)

	fmt.Printf("Created %s, wrote %d bytes\n", absoluteIndexPath, len(indexTemplate))
}
