package resp

import (
  "github.com/gofiber/fiber/v2"
)

type Response struct {
  Code int         `json:"code"`
  Data interface{} `json:"data"`
  Msg  string      `json:"msg"`
}

const (
  ERROR   = 7
  SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *fiber.Ctx) error {
  // 开始时间
  return c.JSON(Response{
    code,
    data,
    msg,
  })
}

func Ok(c *fiber.Ctx) error {
  return Result(SUCCESS, nil, "操作成功", c)
}

func OkWithMessage(message string, c *fiber.Ctx) error {
  return Result(SUCCESS, nil, message, c)
}

func OkWithData(data interface{}, c *fiber.Ctx) error {
  return Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
  return Result(SUCCESS, data, message, c)
}

func Fail(c *fiber.Ctx) error {
  return Result(ERROR, nil, "操作失败", c)
}

func FailWithMessage(message string, c *fiber.Ctx) error {
  return Result(ERROR, nil, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *fiber.Ctx) error {
  return Result(code, data, message, c)
}
