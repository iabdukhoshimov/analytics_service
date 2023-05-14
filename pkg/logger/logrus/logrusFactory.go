package logrus

import (
	"gitlab.com/greatsoft/xif-backend/pkg/logger"
	"gitlab.com/greatsoft/xif-backend/pkg/logger/config"
)

// Factory is the receiver for logrus factory
type Factory struct{}

// Build logrus logger
func (_ *Factory) Build(cfg *config.Logging) (logger.Logger, error) {
	l, err := RegisterLogrusLog(cfg)
	if err != nil {
		return nil, err
	}

	return l, nil
}
