package main

import (
	"github.com/docker/infrakit/pkg/spi/instance"
	"github.com/docker/infrakit/pkg/types"
)

type plugin struct{}

func NewMinimumInstancePlugin() instance.Plugin {
	return &plugin{}
}

func (p *plugin) Validate(req *types.Any) error {
	return nil
}

func (p *plugin) Provision(spec instance.Spec) (*instance.ID, error) {
	return nil, nil
}

func (p *plugin) Label(instance instance.ID, labels map[string]string) error {
	return nil
}

func (p *plugin) Destroy(instance instance.ID, context instance.Context) error {
	return nil
}

func (p *plugin) DescribeInstances(labels map[string]string, properties bool) ([]instance.Description, error) {
	return nil, nil
}
