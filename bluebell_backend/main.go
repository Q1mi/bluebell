package main

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/logger"
	"bluebell_backend/pkg/gen_id"
	"bluebell_backend/routers"
	"bluebell_backend/settings"
	"flag"
	"fmt"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "./conf/bluebell.ini", "配置文件")
	flag.Parse()
	// 加载配置
	if err := settings.LoadFromFile(confFile); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	if err := gen_id.Init(1); err != nil {
		fmt.Printf("init gen_id failed, err:%v\n", err)
		return
	}
	// 注册路由
	r := routers.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.ServerConfig.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
