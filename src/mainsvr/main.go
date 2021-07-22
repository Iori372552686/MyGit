package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	`GoOne/common/misc`
	`GoOne/common/module/application`
	`GoOne/common/module/datetime`
	`GoOne/lib/logger`
	`GoOne/lib/marshal`
	`GoOne/lib/router`
	`GoOne/lib/sensitive_words`
	`GoOne/lib/sharedstruct`
	`GoOne/src/mainsvr/cmd_handler`
	config `GoOne/src/mainsvr/conf`
	`GoOne/src/mainsvr/globals`
)

var svrConfFile = flag.String("svr_conf", "./mainsvr_conf.json", "app conf file")
type MainSvrImpl struct {}

//---------------------------------- func

func onRecvSSPacket(packet *sharedstruct.SSPacket) {
	globals.TransMgr.ProcessSSPacket(packet)
	packet = nil  // packet所有权转交给transmgr，后面不能再用packet（包括data）
}

func (a *MainSvrImpl) OnInit() error {
	err := a.OnReload()
	if err != nil {
		logger.Fatalf("Failed to load config | %v", err)
		return err
	}

	if config.SvrCfg.Pprof {
		go func() {
			log.Println(http.ListenAndServe(":8088", nil))
		}()
	}

	sensitive_words.Init(config.SvrCfg.SensitiveWordsFile)
	err = router.InitAndRun(config.SvrCfg.SelfBusId,
							onRecvSSPacket,
							config.SvrCfg.RabbitMQAddr,
							misc.ServerRouteRules,
							config.SvrCfg.ZKAddr,
							)
	if err != nil {
		logger.Fatalf("Failed to initialize router | %v", err)
		return err
	}

	cmd_handler.RegisterCmd()
	globals.TransMgr.InitAndRun(misc.MaxTransNumber, true, 10)

	logger.Infof("mainsvr init success")
	return nil
}

func (a *MainSvrImpl) OnReload() error {
	err := marshal.LoadJson(*svrConfFile, &config.SvrCfg)
	if err != nil {
		logger.Fatalf("Failed to load server config | %v", err)
		return err
	}

	return nil
}

func (a *MainSvrImpl) OnProc() bool { // return: isIdle
	return true
}

func (a *MainSvrImpl) OnTick(lastMs, nowMs int64) {
	if lastMs / datetime.MS_PER_MINUTE != nowMs / datetime.MS_PER_MINUTE {   // 一分钟调用
		globals.RoleMgr.Tick()
	}
}

func (a *MainSvrImpl) OnExit() {
}

func main() {
	flag.Parse()
	defer logger.Flush()

	application.Init(&MainSvrImpl{})
	application.Run()
}