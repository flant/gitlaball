package common

import (
	"github.com/perhamm/glaball/pkg/client"
	"github.com/perhamm/glaball/pkg/config"
	"github.com/perhamm/glaball/pkg/limiter"

	"github.com/spf13/viper"
)

var (
	Config  *config.Config
	Client  *client.Client
	Limiter *limiter.Limiter
)

func Init() (err error) {
	var cfg config.Config
	if err = viper.Unmarshal(&cfg); err != nil {
		return err
	}

	Config = &cfg

	if Client, err = client.NewClient(Config); err != nil {
		return err
	}

	Limiter = limiter.NewLimiter(Config.Threads)

	return nil
}
