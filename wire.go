//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"memoryDataBase/cache"
	"memoryDataBase/config"
	"memoryDataBase/controller"
	"memoryDataBase/dao"
	"memoryDataBase/database"
	"memoryDataBase/routers"
	"memoryDataBase/service"
)

// ProviderSet 定义依赖注入的提供者集合
var ProviderSet = wire.NewSet(
	config.GetConfig,
	database.InitDB,
	cache.InitRedis,
	dao.NewMemoryDBDao,
	dao.NewStudentCacheDao,
	dao.NewStudentMysqlDao,
	service.NewStudentCacheService,
	service.NewStudentMysqlService,
	service.NewStudentMdbService,
	service.NewStudentService,
	controller.NewStudentController,
	routers.SetUpStudentRouter,
)

// App 定义一个结构体来包装应用程序的依赖项
type App struct {
	StudentRouter  *gin.Engine
	StudentService *service.StudentService
}

// InitializeApp 初始化应用程序，注入生成所需的对象
func InitializeApp() (App, error) {
	wire.Build(
		ProviderSet,
		wire.FieldsOf(new(config.Config), "MySQL", "Redis", "Node", "MemoryDB"),
		wire.Struct(new(App), "*"),
	)
	return App{}, nil
}
