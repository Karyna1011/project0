package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/subgroup/project/internal/config"
	"gitlab.com/tokend/subgroup/project/internal/data/postgres"
	"gitlab.com/tokend/subgroup/project/internal/service/handlers"
)

func (s *service) router(cfg config.Config) chi.Router {

	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxPerson(postgres.NewPersonQ(cfg.DB())),
		),
	)

	r.Route("/integrations/project", func(r chi.Router) {
		r.Post("/add", handlers.Add)
		r.Get("/list", handlers.List)
		r.Get("/get/{id}", handlers.GetByIndex)
		r.Get("/info", handlers.Info)
	})

	return r
}
