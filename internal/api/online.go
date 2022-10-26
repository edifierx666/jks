package api

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gogf/gf/v2/frame/g"
  "jks/internal/model/resp"
  "jks/internal/service"
)

func OnlineUsers(c *fiber.Ctx) error {
  ctx := c.Context()
  onlineUserCache := service.Service.Cache.OnlineUserCache()
  size := onlineUserCache.MustSize(ctx)
  strings, _ := onlineUserCache.KeyStrings(ctx)
  return resp.OkWithData(g.Map{
    "count": size,
    "items": strings,
  }, c)
}
