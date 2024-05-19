package main

import (
	"fmt"
	"log"
	"os"
	"time"

	localServer "com.chinatelecom.oneops.client/pkg/http"
	"com.chinatelecom.oneops.client/pkg/localdb"
	"github.com/kardianos/service"
)

var logger service.Logger

func main() {
	srvConfig := &service.Config{
		Name:        "MyUserExperienceService",
		DisplayName: "MyUserExperienceService数据提取",
		Description: "MyUserExperience数据提取服务",
	}
	prg := &program{}
	s, err := service.New(prg, srvConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		serviceAction := os.Args[1]
		switch serviceAction {
		case "install":
			err := s.Install()
			if err != nil {
				fmt.Println("安装服务失败: ", err.Error())
			} else {
				fmt.Println("安装服务成功")
			}
			return
		case "uninstall":
			err := s.Uninstall()
			if err != nil {
				fmt.Println("卸载服务失败: ", err.Error())
			} else {
				fmt.Println("卸载服务成功")
			}
			return
		case "start":
			err := s.Start()
			if err != nil {
				fmt.Println("运行服务失败: ", err.Error())
			} else {
				fmt.Println("运行服务成功")
			}
			return
		case "stop":
			err := s.Stop()
			if err != nil {
				fmt.Println("停止服务失败: ", err.Error())
			} else {
				fmt.Println("停止服务成功")
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	fmt.Println("服务运行...")
	localdb.InitDb(logger)
	localServer.StartServer()
	go p.run()
	return nil
}
func (p *program) run() error {
	logger.Infof("I'm running %v.", service.Platform())
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case tm := <-ticker.C:
			logger.Infof("Still running at %v...", tm)
			localdb.PingNetwork()

		}
	}
}

func (p *program) Stop(s service.Service) error {
	localdb.CloseDb()
	localServer.StopServer()
	return nil
}
