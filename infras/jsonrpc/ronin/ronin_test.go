package ronin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
)

func TestRonin_GetLatestBlockNumber(t *testing.T) {
	roninConfig := &config.RoninConfig{
		BaseEndPoint: "https://api.roninchain.com/rpc",
	}
	roninClient := NewRoninClient(roninConfig)
	resp, err := roninClient.GetLatestBlockNumber(context.Background(), &models.GetLatestBlockNumberInput{})
	require.NotNil(t, resp)
	require.NoError(t, err)
}

func TestRonin_GetBlockByNumber(t *testing.T) {
	roninConfig := &config.RoninConfig{
		BaseEndPoint: "https://api.roninchain.com/rpc",
	}
	roninClient := NewRoninClient(roninConfig)
	resp, err := roninClient.GetBlockByNumber(context.Background(), &models.GetBlockByNumberInput{
		Params: []any{
			"0x212b708",
			true,
		},
	})
	require.NotNil(t, resp)
	require.NoError(t, err)
}
