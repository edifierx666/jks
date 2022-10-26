package initialize

import (
  "github.com/edifierx666/goproject-kit/os/klog"
  "jks/internal/global"
)

func Logger() {
  global.LoggerCfg = klog.NewLoggerCfg()
  if global.Config.Log.File {
    global.LoggerCfg.SetLogInFile(true)
  }
  logger := klog.New()
  logger.SetLoggerCfg(global.LoggerCfg)
  global.Logger = logger.Build()

  cfg := global.LoggerCfg.New()
  cfg.SetLogInFile(true)
  global.LoggerFile = klog.New().SetLoggerCfg(cfg)
}
