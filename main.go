package main

import (
	_ "embed"
	"github.com/DuC-cnZj/dota2app/cmd"
)

//go:embed config_example.yaml
var configYaml []byte

func main() {
	cmd.Execute(configYaml)
}
