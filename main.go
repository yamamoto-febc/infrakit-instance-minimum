package main

import (
	"github.com/docker/infrakit/pkg/cli"
	instance_plugin "github.com/docker/infrakit/pkg/rpc/instance"
)

func main() {
	cli.RunPlugin("instance-minimum", instance_plugin.PluginServer(NewMinimumInstancePlugin()))
}
