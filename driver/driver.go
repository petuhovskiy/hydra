package driver

import (
	"github.com/petuhovskiy/hydra/driver/configuration"
)

type Driver interface {
	Configuration() configuration.Provider
	Registry() Registry
	CallRegistry() Driver
}
