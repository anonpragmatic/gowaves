package api

import (
	"testing"

	"github.com/anonpragmatic/gowaves/pkg/services"
	"github.com/stretchr/testify/require"
)

func TestAppAuth(t *testing.T) {
	app, _ := NewApp("apiKey", nil, services.Services{})
	require.Error(t, app.checkAuth("bla"))
	require.NoError(t, app.checkAuth("apiKey"))
}
