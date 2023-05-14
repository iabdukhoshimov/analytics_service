package zap

import (
	"gitlab.com/greatsoft/xif-backend/pkg/logger"
	"gitlab.com/greatsoft/xif-backend/pkg/logger/config"
)

// Factory is the receiver for zap factory
type Factory struct{}

// Build zap logger
func (_ *Factory) Build(cfg *config.Logging) (logger.Logger, error) {
	l, err := RegisterLog(cfg)
	if err != nil {
		return nil, err
	}

	return l, nil
}
