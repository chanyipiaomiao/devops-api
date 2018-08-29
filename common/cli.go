package common

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/chanyipiaomiao/hltool"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(AppName, AppDescription)

	inits            = app.Command("init", "Init action")
	refreshRootToken = inits.Flag("refresh-root-token", "refresh root token").Bool()
	server           = app.Command("server", "Server mode")
	logPath          = server.Flag("log", "Log Path, In Configure File, Default: logs/devops-api.log").String()
	runMode          = server.Flag("mode", "Run Mode: dev|prod|test, In Configure File, Default: dev").String()

	token       = app.Command("token", "Token Manage")
	tokenRoot   = token.Flag("root-token", "Specify Root Token").Required().String()
	tokenCreate = token.Flag("create", "Create a Token, Special a Name").String()
	tokenDelete = token.Flag("delete", "Delete a Token, Special a Name").String()

	backup   = app.Command("backup", "Backup BoltDB DB File")
	filepath = backup.Flag("filepath", "Special Backup FilePath").String()
)

// InitCli 初始化命令行参数
func InitCli() {

	app.Author(Author).Version(AppVersion)

	c, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("parse cli args error: %s\n", err)
	}

	switch c {
	case "init":
		var token *Token
		var err error
		token, err = NewToken()
		if *refreshRootToken {
			err = token.AddRootToken(true)
		} else {
			err = token.AddRootToken(false)
		}
		if err != nil {
			log.Fatalf("%s\n", err)
		}

	case "backup":
		if *filepath != "" {
			err := BackupBoltDB(*filepath)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}

	case "server":
		if EnableToken {
			token, err := NewToken()
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			r, _ := token.IsExistToken("root")
			if !r {
				log.Fatalf("root token not exist, please init \n")
			}
		}
		if *logPath != "" {
			LogPathFromCli = *logPath
		}

		// 初始化日志
		InitLog()

		// 获取服务运行模式
		if *runMode != "" {
			if ok, _ := hltool.InStringSlice([]string{"dev", "prod", "test"}, *runMode); !ok {
				log.Fatalln("run mode input error. mode: dev|prod|test")
			}

			beego.BConfig.RunMode = *runMode
		}
		beego.SetStaticPath("/api/static/download/qr", "static/download/qr")
		beego.Run()

	case "token":
		var token *Token
		var err error
		token, err = NewToken()
		if *tokenCreate != "" {
			err = token.AddToken(*tokenRoot, *tokenCreate)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}
		if *tokenDelete != "" {
			err = token.DeleteToken(*tokenRoot, *tokenDelete)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}
	}

}
