package config

import (
	"blog/config"
)

func Get() config.Config {
	return configurations
}
