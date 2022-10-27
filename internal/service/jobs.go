package service

import (
  "context"
  "fmt"
  "strings"
  "time"

  "github.com/bndr/gojenkins"
  "github.com/gogf/gf/v2/frame/g"
  "github.com/gogf/gf/v2/util/gconv"
  "jks/internal/global"
  "jks/internal/model"
)

type Jobs struct {
  baseClient *gojenkins.Jenkins
}

func NewJobs() *Jobs {
  return &Jobs{baseClient: global.JksClient}
}

func (j *Jobs) AllViews(ctx context.Context) ([]*gojenkins.View, error) {
  return j.baseClient.GetAllViews(ctx)
}

func (j *Jobs) Build(ctx context.Context, name string, params map[string]string) (int64, error) {
  return j.baseClient.BuildJob(ctx, name, params)
}

func (j *Jobs) GetBuild(ctx context.Context, jobName string, number int64) (*gojenkins.Build, error) {
  return j.baseClient.GetBuild(ctx, jobName, number)
}

func (j *Jobs) CancelBuild(ctx context.Context, name string, buildId int, params map[string]string) bool {
  post, err := j.baseClient.Requester.Post(ctx, fmt.Sprintf("/job/%v/%v/stop", name, buildId), nil, nil, params)
  if err != nil {
    return false
  }
  if post.StatusCode >= 200 || post.StatusCode < 400 {
    return true
  }
  return false
}

func (j *Jobs) GetView(ctx context.Context, name string) (*gojenkins.View, error) {
  return j.baseClient.GetView(ctx, name)
}

func (j *Jobs) JobDetail(ctx context.Context, name string) (*gojenkins.Job, error) {
  job, err := j.baseClient.GetJob(ctx, name)
  return job, err
}

func (j *Jobs) JobDetailJson(ctx context.Context, name string) (*model.JenkinsJobJSON, error) {
  now := time.Now()
  jobUrl := fmt.Sprintf("/job/%s/api/json", name)
  var jobRes *model.JenkinsJobJSON
  _, err := global.JksClient.Requester.GetJSON(ctx, jobUrl, &jobRes, nil)
  if err != nil {
    return nil, err
  }
  fmt.Println("DetailJSON", name, time.Since(now))
  return jobRes, nil
}

func (j *Jobs) JobParamGitBranch(ctx context.Context, jobUrl string, param string) *model.ParamsBranch {
  var paramBranch *model.ParamsBranch
  if param == "" {
    param = "BRANCH"
  }
  sprintf := fmt.Sprintf("%s%s", strings.Replace(jobUrl, global.Config.Jenkins.Baseurl, "/", 1), "descriptorByName/net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition/fillValueItems")
  json, _ := j.baseClient.Requester.Post(ctx, sprintf, nil, &paramBranch, map[string]string{
    "param": param,
  })

  gconv.Struct(json.Body, &paramBranch, nil)
  return paramBranch
}

func (j *Jobs) BuildParamsInfo(ctx context.Context, name string) ([]*model.JobParams, error) {
  JobParamsArr := make([]*model.JobParams, 0)
  jobRawJson, err := j.JobDetailJson(ctx, name)
  if err != nil {
    return nil, err
  }
  for _, s := range jobRawJson.Property {
    if s.ParameterDefinitions == nil {
      continue
    }
    for _, definition := range s.ParameterDefinitions {
      JobParams := &model.JobParams{}
      JobParams.Class = definition.Class
      JobParams.Name = definition.Name
      JobParams.Type = definition.Type
      JobParams.Chooice = definition.Choices
      if definition.Type == "PT_BRANCH" {
        branchValues := j.JobParamGitBranch(ctx, jobRawJson.URL, definition.Name)
        str := g.SliceStr{}
        for _, value := range branchValues.Values {
          str = append(str, value.Value)
        }
        JobParams.Chooice = str
      }
      JobParams.Description = definition.Description
      JobParams.DefaultValue = definition.DefaultParameterValue.Value
      JobParamsArr = append(JobParamsArr, JobParams)
    }
  }
  return JobParamsArr, nil
}
