package factory

import (
	"fmt"

	"gitlab.com/greatsoft/xif-backend/pkg/logger"
	"gitlab.com/greatsoft/xif-backend/pkg/logger/config"
	"gitlab.com/greatsoft/xif-backend/pkg/logger/logrus"
	"gitlab.com/greatsoft/xif-backend/pkg/logger/zap"
)

// logger map to map logger code to logger builder
var logFactoryBuilderMap = map[string]loggerBuilder{
	config.LOGRUS: &logrus.Factory{},
	config.ZAP:    &zap.Factory{},
}

// interface for logger factory
type loggerBuilder interface {
	Build(cfg *config.Logging) (logger.Logger, error)
}

// accessors for factoryBuilderMap
func getLogFactoryBuilder(key string) (loggerBuilder, error) {
	logFactoryBuilder, ok := logFactoryBuilderMap[key]
	if !ok {
		return nil, fmt.Errorf("not supported logger: %s", key)
	}

	return logFactoryBuilder, nil
}

// Build logger using appropriate log factory
func Build(cfg *config.Logging) (logger.Logger, error) {
	logFactoryBuilder, err := getLogFactoryBuilder(cfg.Code)
	if err != nil {
		return nil, err
	}

	log, err := logFactoryBuilder.Build(cfg)
	if err != nil {
		return nil, err
	}

	return log, nil
}
