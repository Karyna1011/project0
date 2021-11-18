package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
	"time"
)

type TransferConfig struct {
	Block *big.Int      `fig:"block"`
	Time  time.Duration `fig:"time"`
}

func (c *config) TransferConfig() TransferConfig {
	c.onceTransfer.Do(func() interface{} {
		var result TransferConfig

		err := figure.Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "transfer")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out transfer"))
		}
		c.transferConfig = result
		return nil
	})
	return c.transferConfig
}
