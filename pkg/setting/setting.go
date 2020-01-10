package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	// 加载配置文件
	var err error
	if Cfg, err = ini.Load("config/app.ini"); err != nil {
		log.Fatalf("Fatal to load 'app.ini': %v\n", err)
	}

	loadBase()

	loadServer()

	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("DEBUG")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fatal to get section 'server': %v\n", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	if ReadTimeout, err = time.ParseDuration(sec.Key("READ_TIMEOUT").MustString("60s")); err != nil {
		log.Fatalf("Fatal to parse duration 'READ_TIMEOUT': %v\n", err)
	}
	if WriteTimeout, err = time.ParseDuration(sec.Key("WRITE_TIMEOUT").MustString("60s")); err != nil {
		log.Fatalf("Fatal to parse duration 'WRITE_TIMEOUT': %v\n", err)
	}
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fatal to get section 'app': %v\n", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
