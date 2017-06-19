package main

import (
	"fmt"
	"github.com/docker/infrakit/pkg/spi/instance"
	"github.com/docker/infrakit/pkg/types"
	"math/rand"
	"os"
	"path/filepath"
)

var (
	instanceDir = "/tmp/infrakit-dummy-instances"
)

type plugin struct{}

func NewMinimumInstancePlugin() instance.Plugin {
	return &plugin{}
}

func (p *plugin) Validate(req *types.Any) error {
	return nil
}

func (p *plugin) Provision(spec instance.Spec) (*instance.ID, error) {
	// ランダムなIDを生成
	id := instance.ID(fmt.Sprintf("instance-%d", rand.Int63()))
	path := filepath.Join(instanceDir, string(id))

	// ディレクトリ(/tmp/infrakit-dummy-instances)がなければ作成
	_, err := os.Stat(instanceDir)
	if err != nil {
		err := os.MkdirAll(instanceDir, os.FileMode(0777))
		if err != nil {
			return nil, err
		}
	}

	// ファイル作成
	if _, err := os.Stat(path); err != nil {
		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
	}
	return &id, nil
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
