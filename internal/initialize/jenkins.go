package initialize

import (
  "context"

  "github.com/bndr/gojenkins"
  "github.com/edifierx666/goproject-kit/net/kclient"
  "go.uber.org/zap"
  "jks/internal/global"
  "jks/internal/model"
)

func Jks() {
  jks := global.Config.Jenkins
  baseurl := jks.Baseurl
  jksClient := NewJksClient(baseurl)
  jksClient.JenkinsAccount = &model.JenkinsAccount{
    Account: jks.Account,
    Pwd:     jks.Pwd,
  }
  jenkins, err := jksClient.Connection()
  if err != nil {
    global.Logger.Error("Jks初始化出错", zap.Error(err))
  }
  global.JksClient = jenkins
}

type JksClient struct {
  Baseurl        string `json:"baseurl,omitempty"`
  JenkinsAccount *model.JenkinsAccount
  JenkinsClient  *gojenkins.Jenkins
}

func NewJksClient(baseurl string) *JksClient {
  return &JksClient{
    Baseurl:        baseurl,
    JenkinsAccount: nil,
    JenkinsClient:  nil,
  }
}

func (c *JksClient) Connection() (*gojenkins.Jenkins, error) {
  jenkins := gojenkins.CreateJenkins(
    kclient.New().GetClient(),
    c.Baseurl,
    c.JenkinsAccount.Account,
    c.JenkinsAccount.Pwd,
  )
  var err error
  c.JenkinsClient, err = jenkins.Init(context.Background())
  return c.JenkinsClient, err
}
