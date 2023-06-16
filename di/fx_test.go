package di_test

import (
	"github.com/stretchr/testify/assert"
	"go-demo/di"
	"go.uber.org/config"
	"go.uber.org/fx"
	"gorm.io/gorm"

	//"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"testing"
)

type BaseConfigResult struct {
	dig.Out
	Path string `name:"path"`
}

func ProvideBaseConfig() BaseConfigResult {
	return BaseConfigResult{
		Path: "../resources/config.yaml",
	}
}

func TestNewGorm(t *testing.T) {
	fx.New(
		di.Module,
		fx.Invoke(func(db *gorm.DB) {
			assert.NotNil(t, db)
		}),
	)
}

func TestNewConfigure(t *testing.T) {
	fx.New(
		di.Module,
		fx.Invoke(func(config *config.YAML) {
			assert.NotNil(t, config)
		}),
	)
}
