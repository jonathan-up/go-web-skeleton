package config

import "go.uber.org/zap/zapcore"

type Config struct {
	Server   Server   `json:"server" yaml:"server"`
	Logger   Logger   `json:"logger" yaml:"logger"`
	Database Database `json:"database" yaml:"database"`
}

type Server struct {
	Addr string `json:"addr" yaml:"addr"`
}

type Logger struct {
	Level zapcore.Level
}

type Database struct {
	Type     string `json:"type" yaml:"type"`
	Host     string `json:"host" yaml:"host"`
	Port     uint16 `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DbName   string `json:"dbname" yaml:"dbname"`
	Prefix   string `json:"prefix" yaml:"prefix"`
	Encode   string `json:"encode" yaml:"encode"`
}

var YAML = Config{}
