package remote

import (
	"testing"

	"github.com/fluxcd/flux/api"
)

// Just test that the mock does its job.
func TestMock(t *testing.T) {
	ServerTestBattery(t, func(mock api.Server) api.Server { return mock })
}
