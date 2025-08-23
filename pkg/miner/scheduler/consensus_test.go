package scheduler

import (
	"testing"

	"github.com/anonpragmatic/gowaves/pkg/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMinerConsensus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock.NewMockPeerManager(ctrl)

	m.EXPECT().ConnectedCount().Return(1)
	a := NewMinerConsensus(m, 1)
	assert.True(t, a.IsMiningAllowed())

	m.EXPECT().ConnectedCount().Return(0)
	a = NewMinerConsensus(m, 1)
	assert.False(t, a.IsMiningAllowed())
}
