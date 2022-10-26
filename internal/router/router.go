package router

import (
  "github.com/gofiber/fiber/v2"
  "jks/internal/api"
)

func RegisterRoute(app *fiber.App) {
  {
    group := app.Group("/api")
    group.Route("/views", func(router fiber.Router) {
      router.Get("/", api.Views)
      router.Get("/detail", api.ViewsDetail)
    })
    group.Route("/job", func(router fiber.Router) {
      router.Post("/", api.Job)
      router.Post("/build/detail", api.JobBuildDetail)
      router.Post("/build", api.BuildJob)
      router.Post("/cancel", api.CancelJob)
    })
    group.Route("/online", func(router fiber.Router) {
      router.Get("/users", api.OnlineUsers)
    })
  }
}
