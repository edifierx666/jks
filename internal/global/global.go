package global

import (
  "github.com/bndr/gojenkins"
  "github.com/edifierx666/goproject-kit/os/klog"
  "github.com/spf13/viper"
  "jks/internal/config"
)

// 日志
var (
  LoggerCfg  *klog.LoggerCfg
  Logger     *klog.Logger
  LoggerFile *klog.Logger
)

// Jks

var JksClient *gojenkins.Jenkins

// 配置

var (
  Config *config.Config
  Viper  *viper.Viper
)
