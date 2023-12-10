package handler

import (
	"github.com/exit8goodgame/CobraZipSearcherGo/service"

	"go.uber.org/zap"
)

const DemoFileName = "demo.zip"

func DemoHandler(logger *zap.Logger) {
	logger.Info("Start DemoHandler")
	err := service.CrateDemoFile(DemoFileName)
	if err != nil {
		logger.Error("Error DemoHandler.", zap.String("msg", err.Error()))
		return
	}

	s := service.NewSearchService("", 1, 0)
	passWord, err := s.SearchPassWord(DemoFileName)
	if err != nil {
		logger.Error("Error DemoHandler.", zap.String("msg", err.Error()))
		return
	}
	if len(passWord) <= 0 {
		logger.Error("Error DemoHandler. No password discovered")
		return
	}

	logger.Info("End DemoHandler", zap.String("PassWord", passWord))
}
