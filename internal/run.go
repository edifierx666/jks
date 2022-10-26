package internal

import (
  "fmt"

  "github.com/gofiber/fiber/v2"
  "go.uber.org/zap"
  "jks/internal/global"
  "jks/internal/initialize"
  "jks/internal/middleware"
  "jks/internal/router"
  "jks/internal/service"
)

func Run() {
  initialize.Init()
  service.Init()
  global.Logger.Info("配置文件:", zap.Any("config", global.Config))
  app := fiber.New(fiber.Config{Immutable: true, ErrorHandler: middleware.CustomErrorHandle, EnablePrintRoutes: true,
  })
  middleware.Middleware(app)
  router.RegisterRoute(app)
  listenAddr := fmt.Sprintf("%v:%v", global.Config.Server.Addr, global.Config.Server.Port)
  if err := app.Listen(listenAddr); err != nil {
    global.Logger.Error("启动服务失败", zap.Error(err))
  }
}
