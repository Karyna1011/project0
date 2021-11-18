package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type ContractConfig struct {
	Address string `fig:"address"`
}

func (c *config) ContractConfig() ContractConfig {
	c.onceTransfer.Do(func() interface{} {
		var result ContractConfig

		err := figure.Out(&result).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "contract")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out transfer"))
		}
		c.contractConfig = result
		return nil
	})
	return c.contractConfig
}
