// Package database 实现了数据库接口
package database

import (
	"context"

	"github.com/kvii/di/config"
)

// 键类型
type key struct{}

// Key 变量定义了从上下文中获取组件的键。
var Key = new(key)

type DB struct {
	Driver string
	DSN    string
}

// FromContext 函数根据上下文创建组件。
func FromContext(ctx context.Context) *DB {
	var c Config
	v := config.Read(ctx)
	err := v.UnmarshalKey("database", &c)
	if err != nil {
		panic(err)
	}
	return open(c.Driver, c.DSN)
}

// WithContext 函数将组件混入到上下文中。
func WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key, FromContext(ctx))
}

// Read 函数从上下文中获取组件。
func Read(ctx context.Context) *DB {
	return ctx.Value(Key).(*DB)
}

type Config struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

func open(driver, dsn string) *DB {
	return &DB{
		Driver: driver,
		DSN:    dsn,
	}
}
