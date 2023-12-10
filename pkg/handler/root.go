package handler

import (
	"github.com/exit8goodgame/CobraZipSearcherGo/service"

	"go.uber.org/zap"
)

func RootHandler(input string, optTarget string, optChar int, optLength int, logger *zap.Logger) {
	logger.Info("Start RootHandler")

	s := service.NewSearchService(optTarget, optChar, optLength)
	passWord, err := s.SearchPassWord(input)
	if err != nil {
		logger.Error("Error RootHandler.", zap.String("msg", err.Error()))
		return
	}
	if len(passWord) <= 0 {
		logger.Error("Error RootHandler. No password discovered")
		return
	}

	logger.Info("End RootHandler", zap.String("PassWord", passWord))
}
