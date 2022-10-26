package api

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/gogf/gf/v2/util/gconv"
  "go.uber.org/zap"
  "jks/internal/global"
  "jks/internal/model/resp"
  "jks/internal/service"
)

func Views(c *fiber.Ctx) error {
  ctx := c.Context()
  views, err := service.Service.Jobs.AllViews(ctx)
  if err != nil {
    global.Logger.Error("获取Allviews出错", zap.Error(err))
  }
  var r []map[string]interface{}
  for index, view := range views {
    if view.GetName() == "all" {
      continue
    }
    jobSlice := g.SliceAny{}
    elems := g.Map{
      "index": index + 1,
      "name":  view.GetName(),
      "url":   view.GetUrl(),
      "jobs":  &jobSlice,
    }
    for i, job := range view.GetJobs() {
      m := g.Map{}
      gconv.Struct(job, &m)
      m["index"] = i + 1
      jobSlice = append(jobSlice, m)
    }
    r = append(r, elems)
  }

  return resp.OkWithData(r, c)
}

func ViewsDetail(c *fiber.Ctx) error {
  ctx := c.Context()
  query := c.Query("name")
  view, err := service.Service.Jobs.GetView(ctx, query)
  if err != nil {
    global.Logger.Error("获取ViewsDetail出错", zap.Error(err))
  }
  var r = make([]map[string]interface{}, 1)
  jobSlice := g.SliceAny{}
  elems := g.Map{
    "name": view.GetName(),
    "url":  view.GetUrl(),
    "jobs": &jobSlice,
  }
  for i, job := range view.GetJobs() {
    m := g.Map{}
    gconv.Struct(job, &m)
    m["index"] = i + 1
    jobSlice = append(jobSlice, m)
  }
  r = append(r, elems)
  return resp.OkWithData(r, c)
}
