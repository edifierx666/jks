package initialize

import (
  "github.com/fsnotify/fsnotify"
  "github.com/spf13/viper"
  "go.uber.org/zap"
  "jks/internal/global"
)

func Viper() {
  v := viper.New()
  global.Viper = v
  v.SetConfigType("yaml")
  v.SetConfigFile("config.yaml")
  v.AddConfigPath(".")
  v.AddConfigPath("./config")
  if err := v.ReadInConfig(); err != nil {
    global.Logger.Error("读取配置文件出错", zap.Error(err))
  }
  v.SetDefault("server.addr", "0.0.0.0")
  v.WatchConfig()
  v.OnConfigChange(func(in fsnotify.Event) {
    if err := v.Unmarshal(&global.Config); err != nil {
      global.Logger.Error("解析配置文件出错", zap.Error(err))
    }
    global.Logger.Info("配置文件改变:", zap.Any("config", global.Config))
  })
  if err := v.Unmarshal(&global.Config); err != nil {
    global.Logger.Error("解析配置文件出错", zap.Error(err))
  }
}
