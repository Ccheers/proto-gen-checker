package main

import (
	"flag"
	"fmt"

	"github.com/ccheers/proto-gen-checker/internal/config"
	"github.com/ccheers/proto-gen-checker/internal/generator"
	"github.com/ccheers/proto-gen-checker/internal/meta"
	interPlugin "github.com/ccheers/proto-gen-checker/internal/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	flags   flag.FlagSet
	c       config.Config
	version bool
)

func init() {
	flag.BoolVar(&version, "v", false, "print version information")
	flag.BoolVar(&version, "version", false, "print version information")
}

func main() {
	flag.Parse()
	if version {
		fmt.Printf("%+v\n", meta.AppInfo)
		return
	}
	plugin, err := interPlugin.New(
		&c,
		generator.NewU32(),
		generator.NewU64(),
		generator.NewI32(),
		generator.NewI64(),
		generator.NewF32(),
		generator.NewF64(),
	)
	if err != nil {
		panic(err)
	}
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(plugin.Generate)
}
