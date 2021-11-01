package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type MyConfig struct {
	Field1      string `fig:"field1"`
	Field2      string `fig:"field2"`
	Field3      int64  `fig:"field3"`
	Field4      uint64 `fig:"field4"`
}

func (c *config) MyConfig() MyConfig {
	c.once.Do(func() interface{} {
		var result MyConfig

		err := figure.Out(&result).
		With(figure.BaseHooks).
		From(kv.MustGetStringMap(c.getter, "transfer")).
		Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out transfer"))
		}
		c.myConfig = result
		return nil
	})
	return c.myConfig
}
