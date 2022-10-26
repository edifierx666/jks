package config

type Config struct {
  Jenkins Jks    `json:"jenkins" yaml:"jenkins"`
  Server  Server `json:"server" yaml:"server"`
  Log     Log    `json:"log" yaml:"log"`
}
