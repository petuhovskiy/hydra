package main

import (
	"github.com/petuhovskiy/hydra/driver"
)

type MemTestPlugin struct {
	*driver.RegistryMemory
}

func NewRegistry() driver.Registry {
	return &MemTestPlugin{RegistryMemory: driver.NewRegistryMemory()}
}

func main() {}
