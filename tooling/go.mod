module github.com/microsoft/yardl/tooling

go 1.18

require (
	github.com/alecthomas/participle/v2 v2.0.0-beta.5
	github.com/dlclark/regexp2 v1.7.0
	github.com/fsnotify/fsnotify v1.6.0
	github.com/inancgumus/screen v0.0.0-20190314163918-06e984b86ed3
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.0
	gitlab.com/akabio/iotool v0.5.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1 // indirect
)

// Replace go-yaml with a fork that contains pending PR
// https://github.com/go-yaml/yaml/pull/691
replace gopkg.in/yaml.v3 => github.com/johnstairs/go-yaml-yaml v0.0.0-20221109150101-483fca0d3ee9
