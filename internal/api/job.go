package api

import (
  "fmt"
  "log"
  "sync"
  "time"

  "github.com/gofiber/fiber/v2"
  "github.com/gogf/gf/v2/container/garray"
  "github.com/gogf/gf/v2/util/gconv"
  "jks/internal/global"
  "jks/internal/model"
  "jks/internal/model/resp"
  "jks/internal/service"
)

func Job(c *fiber.Ctx) error {
  ctx := c.Context()
  req := model.JobDetailReq{}
  if err := c.BodyParser(&req); err != nil {
    return resp.FailWithMessage(err.Error(), c)
  }
  jobDetail, err := service.Service.Jobs.JobDetail(ctx, req.Name)
  ids, err := jobDetail.GetAllBuildIds(ctx)
  if err != nil {
    return resp.FailWithMessage(err.Error(), c)
  }
  resp1 := &model.JobDetailRes{}
  r := garray.NewSortedArray(func(a, b interface{}) int {
    if a.(*model.ItemObj).Number > b.(*model.ItemObj).Number {
      return -1
    }
    if a.(*model.ItemObj).Number < b.(*model.ItemObj).Number {
      return 1
    }
    return 0
  }, true)
  start := time.Now()
  if req.Count == 0 {
    req.Count = 15
  }
  count := req.Count
  var wg sync.WaitGroup
  for _, id := range ids {
    if count <= 0 {
      break
    }
    wg.Add(1)
    id := id
    go func() {
      defer wg.Done()
      build, err := jobDetail.GetBuild(ctx, id.Number)
      if err != nil {
        log.Println(err)
      }
      item := &model.ItemObj{}
      item.Result = build.GetResult()
      item.IsRunning = build.Raw.Building
      item.Number = id.Number
      item.Url = id.URL
      item.StartTime = build.GetTimestamp().Format("2006-01-02 15:04:05")
      item.JobName = req.Name
      item.Raw = gconv.Map(build.Raw)
      item.RawJob = gconv.Map(build.Job.Raw)
      parameters := build.GetParameters()
      for _, parameter := range parameters {
        if parameter.Name == "BRANCH" {
          item.Branch = parameter.Value
          break
        }
      }
      item.BuildParams = gconv.Maps(parameters)
      causes := model.GetCauses(build)
      for _, cause := range causes {
        item.Username = gconv.String(cause["userName"])
      }
      r.Append(item)
    }()
    count--
  }
  wg.Add(1)
  go func() {
    defer wg.Done()
    params, _ := service.Service.Jobs.BuildParamsInfo(ctx, req.Name)
    resp1.Params = params
  }()
  wg.Wait()
  fmt.Println("该函数执行完成耗时：", time.Since(start))
  gconv.Structs(r, &resp1.BuildList)
  return resp.OkWithData(resp1, c)
}

func JobBuildDetail(c *fiber.Ctx) error {
  ctx := c.Context()
  req := model.JobBuildDetailReq{}
  c.BodyParser(&req)
  jobDetail, err := service.Service.Jobs.JobDetail(ctx, req.JobName)
  if err != nil {
    return resp.FailWithMessage(err.Error(), c)
  }

  build, err := jobDetail.GetBuild(ctx, gconv.Int64(req.BuildId))
  if err != nil {
    return resp.FailWithMessage(err.Error(), c)
  }
  item := &model.ItemObj{}
  item.Result = build.GetResult()
  item.JobName = req.JobName
  item.IsRunning = build.IsRunning(ctx)
  item.Number = build.GetBuildNumber()
  causes := model.GetCauses(build)
  for _, cause := range causes {
    item.Username = gconv.String(cause["userName"])
  }
  parameters := build.GetParameters()
  for _, parameter := range parameters {
    if parameter.Name == "BRANCH" {
      item.Branch = parameter.Value
      break
    }
  }
  item.BuildParams = gconv.Maps(parameters)
  item.Url = build.GetUrl()
  item.StartTime = build.GetTimestamp().Format("2006-01-02 15:04:05")
  return resp.OkWithData(gconv.Map(item), c)
}

func JobConsoleOutput(c *fiber.Ctx) error {
  ctx := c.Context()
  req := model.JobConsoleOutputReq{}
  c.BodyParser(&req)
  jobBuildInfo, err := service.Service.Jobs.GetBuild(ctx, req.JobName, req.BuildId)
  if err != nil {
    return err
  }
  consoleOutput := jobBuildInfo.GetConsoleOutput(ctx)
  return resp.OkWithData(map[string]interface{}{
    "text":      consoleOutput,
    "result":    jobBuildInfo.GetResult(),
    "timestamp": jobBuildInfo.GetTimestamp().Format("2006-01-02 15:04:05"),
    "number":    jobBuildInfo.GetBuildNumber(),
    "url":       jobBuildInfo.GetUrl(),
  }, c)
}

func BuildJob(c *fiber.Ctx) error {
  ctx := c.Context()
  req := model.JobBuildReq{}
  c.BodyParser(&req)
  build, err := service.Service.Jobs.Build(ctx, req.Name, req.Param)
  if err != nil {
    return err
  }
  global.Logger.Info(fmt.Sprintf("%v调用了构建服务:%v参数是%#v构建的queueId:%v用户ID:%v", req.Name, req.Param, build))
  return resp.OkWithData(build, c)
}

func CancelJob(c *fiber.Ctx) error {
  ctx := c.Context()
  req := model.JobCancelBuildReq{}
  c.BodyParser(&req)
  ok := service.Service.Jobs.CancelBuild(ctx, req.Name, req.BuildId, req.Param)
  if !ok {
    return resp.FailWithMessage("取消失败", c)
  }
  return resp.OkWithData(ok, c)
}
