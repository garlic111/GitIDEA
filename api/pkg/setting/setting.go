package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize        int    //分页返回的数据个数
	RuntimeRootPath string //保存文件的跟路径
	ExpireTime      time.Duration
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

var AppSetting = &App{}

//服务配置
type Server struct {
	RunMode      string        //运行模式
	HttpPort     int           //端口号
	ReadTimeout  time.Duration //读取超时时间
	WriteTimeout time.Duration //写入超时时间
}

var ServerSetting = &Server{}

//数据库配置
type Database struct {
	Type        string //数据库类型
	User        string //数据库用户
	Password    string //数据库密码
	Host        string //数据库地址+端口号
	Name        string //数据库名称
	TablePrefix string //数据库数据表前缀
}

var DatabaseSetting = &Database{}

func SetUp() {
	cf, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("获取ini配置文件失败")
	}

	err = cf.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatal("Config 文件设置映射失败")
	}

	err = cf.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 ServerSetting 错误: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = cf.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg配置文件映射 DatabaseSetting 错误: %v", err)
	}
}
