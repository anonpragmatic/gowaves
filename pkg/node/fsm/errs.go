package fsm

import (
	"errors"

	"github.com/anonpragmatic/gowaves/pkg/proto"
)

var TimeoutErr = proto.NewInfoMsg(errors.New("timeout"))
