package service

import (
    "net"
    "net/http"

    "gitlab.com/distributed_lab/kit/copus/types"
    "gitlab.com/distributed_lab/logan/v3"
    "gitlab.com/tokend/subgroup/project/internal/config"
)

type service struct {
    log      *logan.Entry
    copus    types.Copus
    listener net.Listener
}



func (s *service) run(cfg config.Config) error {
    s.log.Info("Running api service")

    r := s.router(cfg)

    return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
    log := cfg.Log().WithField("api-service", "")

    return &service{
        log:        log,
        copus:      cfg.Copus(),
        listener:   cfg.Listener(),
    }
}


func Run(cfg config.Config) {
    if err := newService(cfg).run(cfg); err != nil {
        panic(err)
    }
}
