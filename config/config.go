package config

import (
	"time"
)

// MySQLConfig 定义 MySQL 配置结构体
type MySQLConfig struct {
	DSN string
}

// RedisConfig 定义 Redis 配置结构体
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// MemoryDBConfig 定义内存数据库配置结构体
type MemoryDBConfig struct {
	Capacity   int     //内存容量
	EvictRatio float64 //触发内存淘汰时淘汰的键的比例
}

// CachePreheatingConfig 定义缓存预热配置结构体
type CachePreheatingConfig struct {
	LoadRatio float64 //加载到内存中的数据量占内存容量的比例
}

// ServerConfig 定义服务器配置结构体
type ServerConfig struct {
	ReloadInterval         time.Duration //重载缓存的时间间隔
	PeriodicDeleteInterval time.Duration //过期删除的时间间隔
	ExamineSize            int           //过期删除时检测的键的数量
}

// Node 定义当前节点信息结构体
type Node struct {
	NodeId      string
	Address     string
	PortAddress string
}

// Peer 表示集群中其他单个节点的信息 在加入集群时要包含集群中的所有其他节点
type Peer struct {
	NodeId      string
	Address     string
	PortAddress string
}

// Config 定义配置结构体
type Config struct {
	MySQL           MySQLConfig
	Redis           RedisConfig
	MemoryDB        MemoryDBConfig
	CachePreheating CachePreheatingConfig
	Server          ServerConfig
	Node            Node
	Peers           []*Peer
}

// GetConfig 获取配置实例
func GetConfig() Config {
	return Config{
		// 配置 Mysql
		MySQL: MySQLConfig{
			DSN: "root:1234@tcp(127.0.0.1:3306)/mdb?charset=utf8mb4&parseTime=True&loc=Local",
		},
		// 配置 redis
		Redis: RedisConfig{
			Addr:     "192.168.88.128:6379",
			Password: "1234",
			DB:       0,
		},
		// 配置内存数据库
		MemoryDB: MemoryDBConfig{
			Capacity:   10,
			EvictRatio: 0.2,
		},
		CachePreheating: CachePreheatingConfig{
			LoadRatio: 0.5,
		},
		Server: ServerConfig{
			ReloadInterval:         time.Hour,
			PeriodicDeleteInterval: time.Hour,
			ExamineSize:            10,
		},
		Node: Node{
			NodeId:      "节点1",
			Address:     "127.0.0.1:8080",
			PortAddress: "8080",
		},
		Peers: []*Peer{},
	}
}
