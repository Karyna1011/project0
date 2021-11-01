package service

import (
  "github.com/go-chi/chi"
  "gitlab.com/distributed_lab/ape"
  "gitlab.com/tokend/subgroup/project/internal/config"
  "gitlab.com/tokend/subgroup/project/internal/data/functions"
  "gitlab.com/tokend/subgroup/project/internal/service/handlers"
)



func (s *service) router(cfg config.Config) chi.Router {


  r := chi.NewRouter()
  //db:=OpenConnection()

  r.Use(
    ape.RecoverMiddleware(s.log),
    ape.LoganMiddleware(s.log),
    ape.CtxMiddleware(
      handlers.CtxLog(s.log),
      handlers.CtxPerson(functions.NewPersonQ(cfg.DB())),
      ),
    )

  r.Route("/integrations/project", func(r chi.Router) {
    r.Post("/add", handlers.POSTHandler)
    r.Get("/get", handlers.GETHandler)
    r.Get("/get/{id}", handlers.GetByIndex)
    r.Get("/", handlers.GetMessage)
  })

  return r
}
