package node

import (
	"log/slog"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/anonpragmatic/gowaves/pkg/mock"
	"github.com/anonpragmatic/gowaves/pkg/node/peers/storage"
	"github.com/anonpragmatic/gowaves/pkg/p2p/peer"
	"github.com/anonpragmatic/gowaves/pkg/proto"
	"github.com/anonpragmatic/gowaves/pkg/services"
)

func TestPeersAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock.NewMockPeerManager(ctrl)
	m.EXPECT().KnownPeers().Return([]storage.KnownPeer{})
	addr := proto.NewTCPAddr(net.ParseIP("127.0.0.1"), 6868).ToIpPort()
	m.EXPECT().UpdateKnownPeers([]storage.KnownPeer{storage.KnownPeer(addr)})

	_, err := PeersAction(services.Services{
		Peers: m,
	}, peer.ProtoMessage{
		Message: &proto.PeersMessage{
			Peers: []proto.PeerInfo{
				{
					Addr: net.ParseIP("127.0.0.1"),
					Port: 6868,
				},
			},
		},
	}, nil, slog.New(slog.DiscardHandler))
	require.NoError(t, err)
}
