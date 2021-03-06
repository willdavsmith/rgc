# rgc
React-Generate-Component - A CLI tool to quickly generate a new React component.

## Installation
```
go get github.com/willdavsmith/rgc
```

## Manual Installation
1. Clone the repository.
```
git clone https://github.com/SeedlingPartnerships/seedling-ui.git
```

2. Build and install the executable.
```
go install
```

## Usage
```
$ rgc
  -c string
        New component name. (Required)
  -d string
        Destination directory. (default "./")
  -dry
        Run command as dry run (no filesystem changes).
```

## Examples
```bash
$ rgc -c NewComponent
Created C:\Users\Will\react-project\NewComponent
Created C:\Users\Will\react-project\NewComponent\NewComponent.tsx, wrote 57 bytes
Created C:\Users\Will\react-project\NewComponent\index.ts, wrote 41 bytes

$ rgc -c NewComponent -d components
Created C:\Users\Will\react-project\components\NewComponent
Created C:\Users\Will\react-project\components\NewComponent\NewComponent.tsx, wrote 57 bytes
Created C:\Users\Will\react-project\components\NewComponent\index.ts, wrote 41 bytes
```
