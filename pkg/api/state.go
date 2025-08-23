package api

import (
	"github.com/anonpragmatic/gowaves/pkg/proto"
	"github.com/pkg/errors"
)

// TODO Here should be internal message with rollback action
func (a *App) RollbackToHeight(apiKey string, height proto.Height) error {
	return errors.New("api method disabled")
}
