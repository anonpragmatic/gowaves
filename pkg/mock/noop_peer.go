package mock

import "github.com/anonpragmatic/gowaves/pkg/proto"

type NoOpPeer struct {
}

func (NoOpPeer) SendMessage(proto.Message) {

}
