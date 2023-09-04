// Package config 提供了配置组件。
package config

import (
	"context"

	"github.com/spf13/viper"
)

// 键类型。
type key struct{}

// Key 变量定义了从上下文中获取组件的键。
var Key = new(key)

// FromContext 函数根据上下文创建组件。
func FromContext(ctx context.Context) *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return v
}

// WithContext 函数将组件混入到上下文中。
func WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key, FromContext(ctx))
}

// Read 函数从上下文中获取组件。
func Read(ctx context.Context) *viper.Viper {
	return ctx.Value(Key).(*viper.Viper)
}
