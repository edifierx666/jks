package model

import (
  "github.com/bndr/gojenkins"
  "github.com/gogf/gf/v2/util/gconv"
)

type JobParams struct {
  Class        string   `json:"class"`
  Chooice      []string `json:"chooice"`
  Type         string   `json:"type"`
  Name         string   `json:"name"`
  Description  string   `json:"description"`
  DefaultValue string   `json:"defaultValue"`
}
type ItemObj struct {
  IsRunning   bool                     `json:"isRunning"`
  Result      string                   `json:"result"`
  Branch      string                   `json:"branch"`
  Number      int64                    `json:"number"`
  Url         string                   `json:"url"`
  Username    string                   `json:"username"`
  StartTime   string                   `json:"startTime"`
  JobName     string                   `json:"jobName"`
  BuildParams []map[string]interface{} `json:"buildParams"`
}

func GetCauses(build *gojenkins.Build) []map[string]interface{} {
  for _, a := range build.Raw.Actions {
    if a.Causes != nil {
      return gconv.Maps(a.Causes)
    }
  }
  return []map[string]interface{}{}
}
