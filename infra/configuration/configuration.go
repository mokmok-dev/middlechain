package configuration

import (
	"fmt"

	domain "github.com/Pranc1ngPegasus/golang-template/domain/configuration"
	"github.com/Pranc1ngPegasus/golang-template/domain/logger"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var _ domain.Configuration = (*Configuration)(nil)

var NewConfigurationSet = wire.NewSet(
	wire.Bind(new(domain.Configuration), new(*Configuration)),
	NewConfiguration,
)

type Configuration struct {
	config *domain.Config
}

func NewConfiguration(
	logger logger.Logger,
) (*Configuration, error) {
	viper.SetConfigFile("sample.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("failed to load environment variable", err)

		return nil, fmt.Errorf("failed to load environment variable: %w", err)
	}

	var config domain.Config

	if err := viper.Unmarshal(&config); err != nil {
		logger.Error("failed to unmarshal environment variable", err)

		return nil, fmt.Errorf("failed to unmarshal environment variable: %w", err)
	}

	return &Configuration{
		config: &config,
	}, nil
}

func (c *Configuration) Config() *domain.Config {
	return c.config
}
