# rgc
React-Generate-Component - A CLI tool to quickly generate a new React component.

## Installation
```
go install github.com/willdavsmith/rgc
```

## Usage
```
Specify your project components directory in .rgc.yaml.

Usage:
  rgc [flags]

Flags:
  -c, --component string   component name (required)
      --config string      config file (default is .rgc.yaml)
  -d, --directory string   component directory
  -h, --help               help for rgc
  -i, --index string       index file (default is {directory}/index.{ts|js})
      --use-javascrpt      use javascript instead of typescript (default is false)
  -v, --version            version for rgc
  -h, --help               show this page
```
