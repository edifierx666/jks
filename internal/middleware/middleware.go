package middleware

import (
  "context"
  "fmt"
  "strings"
  "time"

  "github.com/gobuffalo/packr/v2"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/filesystem"
  "github.com/gofiber/fiber/v2/middleware/requestid"
  "go.uber.org/zap"
  "jks/internal/global"
  "jks/internal/model/resp"
  "jks/internal/service"
  recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(app *fiber.App) {
  app.Use(recover2.New(recover2.Config{
    EnableStackTrace: true,
  }), cors.New(), requestid.New(), BaseLogger, UserId)
  IndexHome(app)
}

func CustomErrorHandle(c *fiber.Ctx, err error) error {
  if e, ok := err.(*fiber.Error); ok {
    return resp.FailWithDetailed(e.Code, nil, e.Message, c)
  }
  return resp.FailWithMessage(err.Error(), c)
}

func IndexHome(app *fiber.App) {
  app.Use(filesystem.New(filesystem.Config{
    Root: packr.New("index", "/Volumes/ZY/www/www2/jenkins-ui-f/dist"),
    Next: func(c *fiber.Ctx) bool {
      if strings.Index(string(c.Request().URI().Path()), "/api") == 0 {
        return true
      }
      return false
    },
  }))
}
func UserId(c *fiber.Ctx) error {
  __id := c.Get("__id")
  if __id != "" {
    service.Service.Cache.OnlineUserCache().SetIfNotExist(context.Background(), __id, __id, time.Minute*5)
  }
  return c.Next()
}
func BaseLogger(c *fiber.Ctx) error {
  now := time.Now()
  err := c.Next()
  uri := c.Request().URI()
  rid := c.Locals("requestid")
  global.Logger.Info(
    fmt.Sprintf("[%v]", rid),
    zap.Any("URI", string(uri.Path())),
    zap.String("QUERY", string(uri.QueryString())),
    zap.Any("BODY", string(c.Body())),
    zap.Any("IP", c.IP()),
    zap.String("USER-AGENT", string(c.Request().Header.UserAgent())),
  )
  fmt.Println("耗时", rid, time.Since(now))
  return err
}
