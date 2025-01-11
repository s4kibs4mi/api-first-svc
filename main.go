package main

import (
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/s4kibs4mi/api-first-svc/cmd"
	"github.com/s4kibs4mi/api-first-svc/configs"
)

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, options *configs.Config) {})
	cli.Root().AddCommand(cmd.ApiCmd)
	cli.Run()
}
