package main

import (
	"log"
	"memoryDataBase/config"
	"time"
)

func main() {
	// 使用 wire 生成的代码初始化应用程序
	app, err := InitializeApp()
	studentService := app.StudentService
	studentRouter := app.StudentRouter
	if err != nil {
		log.Fatalf("初始化应用程序失败: %v", err)
	}

	cfg := config.GetConfig()

	//启动时加载缓存数据到内存
	if err = studentService.LoadCacheToMemory(cfg.MemoryDB.Capacity, cfg.CachePreheating.LoadRatio); err != nil {
		log.Printf("节点：%s 加载缓存到内存时失败：%v", cfg.Node.NodeId, err)
		if err = studentService.LoadDateBaseToMemory(cfg.MemoryDB.Capacity, cfg.CachePreheating.LoadRatio); err != nil {
			log.Printf("节点：%s 加载数据库中的数据到内存时失败：%v", cfg.Node.NodeId, err)
		}
		log.Printf("节点：%s 加载数据库到内存", cfg.Node.NodeId)
	}
	log.Printf("节点：%s 加载缓存到内存", cfg.Node.NodeId)

	//启动时等待10秒 第一个节点要等待领导者选举完成再获得地址 后面的节点要等待加入集群
	time.Sleep(10 * time.Second)
	log.Printf("节点：%s 等待领导者选举完成或加入集群", cfg.Node.NodeId)

	//定期清空缓存 定期清除内存中的过期键 让领导者节点提交命令给所有节点
	go func() {
		studentService.ReLoadCacheData(cfg.Server.ReloadInterval)
	}()

	//定期删除内存数据库过期键
	go func() {
		studentService.PeriodicDelete(cfg.Server.PeriodicDeleteInterval, cfg.Server.ExamineSize)
	}()

	//定期检测有没有损坏的节点 有的话删除
	go func() {
		studentService.PeriodicDeleteFatalPeel(cfg.Server.PeriodicDeleteFatalPeelInterval)
	}()

	// 启动服务器
	serverAddress := ":" + cfg.Node.PortAddress
	if err = studentRouter.Run(serverAddress); err != nil {
		log.Fatalf("节点：%s 初始化学生路由时出错：%v", cfg.Node.NodeId, err)
	}
}
