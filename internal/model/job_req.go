package model

type PlusAiReq struct {
}

type ViewDetailReq struct {
  Name string `json:"name"`
}

type JobDetailReq struct {
  Name  string `json:"name" validate:"required"`
  Count int    `json:"count"`
}
type JobDetailRes struct {
  Params    []*JobParams `json:"params"`
  BuildList []*ItemObj   `json:"buildList"`
}
type JobBuildDetailReq struct {
  JobName string `json:"jobName"`
  BuildId string `json:"buildId"`
}

type JobBuildReq struct {
  Name  string            `json:"name"`
  Param map[string]string `json:"param"`
}

type JobCancelBuildReq struct {
  Name    string            `json:"name"`
  BuildId int               `json:"buildId"`
  Param   map[string]string `json:"param"`
}
