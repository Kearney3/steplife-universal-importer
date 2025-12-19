package main

import (
	"fmt"
	"os"
	"github.com/eiannone/keyboard"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	consts "steplife-universal-importer/internal/const"
	"steplife-universal-importer/internal/gui"
	"steplife-universal-importer/internal/model"
	"steplife-universal-importer/internal/server"
	xif "steplife-universal-importer/internal/utils/if"
	"steplife-universal-importer/internal/utils/logx"
	timeUtils "steplife-universal-importer/internal/utils/time"
	"time"
)

func main() {
	// 检查命令行参数
	if len(os.Args) > 1 && os.Args[1] == "--cli" {
		// 命令行模式
		runCommandLineMode()
	} else {
		// GUI模式（默认）
		runGUIMode()
	}
}

func runGUIMode() {
	guiApp := gui.NewGUI()
	guiApp.Run()
}

func runCommandLineMode() {
	println("\n.---..---..---..---..-.   .-..---..---.   .-..-.-.-..---..----..---. .---..---..---. ")
	println(" \\ \\ `| |'| |- | |-'| |__ | || |- | |- ###| || | | || |-'| || || |-< `| |'| |- | |-< ")
	println("`---' `-' `---'`-'  `---'   `-'`-'-'-'`-'  `----'`-'`-' `-' `---'`-'`-'\n")

	logx.Info("命令行模式执行中......")
	config, err := initConfig()
	if err != nil {
		logx.ErrorF("初始化配置失败：%v", err)
		panic(err)
	}

	err = server.Run(config)
	if err != nil {
		logx.ErrorF("Run error: %v", err)
		panic(err)
	}

	exit()
}

func initConfig() (model.Config, error) {
	var config model.Config

	// 尝试多个可能的config.ini位置
	configPaths := []string{
		"config.ini",              // 当前目录
		"./config.ini",           // 当前目录（明确）
		"../config.ini",          // 父目录
	}

	var cfg *ini.File
	var err error

	for _, path := range configPaths {
		cfg, err = ini.Load(path)
		if err == nil {
			logx.InfoF("成功加载配置文件：%s", path)
			break
		}
	}

	if err != nil {
		logx.ErrorF("在所有可能位置都找不到config.ini文件，最后尝试：%v", err)
		return config, errors.Wrap(err, "Failed to load config from any location")
	}

	err = cfg.MapTo(&config)
	if err != nil {
		logx.ErrorF("Failed to map config: %v", err)
		return config, errors.Wrap(err, "Failed to map config")
	}

	if config.PathStartTime != "" {
		config.PathStartTimestamp, err = timeUtils.ToTimestamp(config.PathStartTime)
		if err != nil {
			logx.ErrorF("时间解析失败：%s", err)
			return config, errors.Wrap(err, "时间解析失败")
		}
	} else {
		config.PathStartTimestamp = time.Now().Unix()
	}

	config.InsertPointDistance = xif.Int(
		config.InsertPointDistance < consts.MinInsertPointDistance,
		consts.DefaultInsertPointDistance,
		config.InsertPointDistance,
	)

	return config, nil
}

func exit() {
	fmt.Println("Press any key to exit...")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	_, _, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	fmt.Println("Exiting program...")
}
