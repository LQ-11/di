// Package service 实现了服务
package service

import (
	"context"
	"log/slog"

	"github.com/kvii/di/database"
	"github.com/kvii/di/logger"
)

// 键类型
type key struct{}

// Key 变量定义了从上下文中获取组件的键。
var Key = new(key)

// FromContext 函数根据上下文创建组件。
func FromContext(ctx context.Context) *Service {
	return &Service{
		DB:     database.Read(ctx),
		Logger: logger.Read(ctx).With(slog.String("module", "service")),
	}
}

type Service struct {
	DB     *database.DB
	Logger *slog.Logger
}

func (s *Service) Hello(ctx context.Context, name string) (string, error) {
	s.Logger.Info("Hello", slog.String("name", name))
	return "Hello " + name, nil
}
