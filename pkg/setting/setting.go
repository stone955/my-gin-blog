package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var AppCfg = &AppSetting{}

type AppSetting struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl     string
	ImageSavePath      string
	ImageMaxSize       int
	ImageAllowExtNames []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var ServerCfg = &ServerSetting{}

type ServerSetting struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var DatabaseCfg = &DatabaseSetting{}

type DatabaseSetting struct {
	Type        string
	Name        string
	User        string
	Password    string
	Host        string
	TablePrefix string
}

func Setup() {
	// 加载配置文件
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fatal to load 'app.ini': %v\n", err)
	}

	if err := cfg.Section("app").MapTo(AppCfg); err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppCfg.ImageMaxSize = AppCfg.ImageMaxSize * 1024 * 1024

	if err := cfg.Section("server").MapTo(ServerCfg); err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v\n", err)
	}
	ServerCfg.ReadTimeout = ServerCfg.ReadTimeout * time.Second
	ServerCfg.WriteTimeout = ServerCfg.WriteTimeout * time.Second

	if err := cfg.Section("database").MapTo(DatabaseCfg); err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v\n", err)
	}
}
