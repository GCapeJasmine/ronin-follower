package ronin

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/GCapeJasmine/ronin-follower/config"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/client"
	"github.com/GCapeJasmine/ronin-follower/internal/domains/models"
	httpPkg "github.com/GCapeJasmine/ronin-follower/pkg/http"
)

const (
	JsonRPCVersion = "2.0"
	IdDefault      = 1
)

type Ronin struct {
	roninConfig *config.RoninConfig
	client      *http.Client
}

func NewRoninClient(roninConfig *config.RoninConfig) client.IClientRonin {
	client := &http.Client{Transport: http.DefaultTransport, Timeout: time.Second * 60}
	return &Ronin{
		roninConfig: roninConfig,
		client:      client,
	}
}

func (r *Ronin) GetLatestBlockNumber(ctx context.Context, input *models.GetLatestBlockNumberInput) (*models.GetLatestBlockNumberOutput, error) {
	log.Printf("GetLatestBlockNumber...")
	input.JsonRPC = JsonRPCVersion
	input.Id = IdDefault
	input.Method = models.MethodGetLatestBlockNumber

	bytesBody, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, "POST", r.roninConfig.BaseEndPoint, bytes.NewBuffer(bytesBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	bodyResp, err := httpPkg.MakeRequest(ctx, r.client, req, 5, 5)
	if err != nil {
		log.Printf("GetLatestBlockNumber got error: %v", err)
		return nil, err
	}
	output := &models.GetLatestBlockNumberOutput{}
	err = json.Unmarshal(bodyResp, &output)
	if err != nil {
		log.Printf("GetLatestBlockNumber got error: %v", err)
		return nil, err
	}
	return output, nil
}

func (r *Ronin) GetBlockByNumber(ctx context.Context, input *models.GetBlockByNumberInput) (*models.GetBlockByNumberOutput, error) {
	log.Printf("GetBlockByNumber...")
	input.JsonRPC = JsonRPCVersion
	input.Id = IdDefault
	input.Method = models.MethodGetBlockByNumber

	bytesBody, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, "POST", r.roninConfig.BaseEndPoint, bytes.NewBuffer(bytesBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	bodyResp, err := httpPkg.MakeRequest(ctx, r.client, req, 5, 5)
	if err != nil {
		log.Printf("GetBlockByNumber got error: %v", err)
		return nil, err
	}
	output := &models.GetBlockByNumberOutput{}
	err = json.Unmarshal(bodyResp, &output)
	if err != nil {
		log.Printf("GetBlockByNumber got error: %v", err)
		return nil, err
	}
	return output, nil
}

func (r *Ronin) GetTransactionByHash(ctx context.Context, input *models.GetTransactionByHashInput) (*models.GetTransactionByHashOutput, error) {
	log.Printf("GetTransactionByHash ...")
	input.JsonRPC = JsonRPCVersion
	input.Id = IdDefault
	input.Method = models.MethodGetTransactionByHash

	bytesBody, _ := json.Marshal(input)
	req, err := http.NewRequestWithContext(ctx, "POST", r.roninConfig.BaseEndPoint, bytes.NewBuffer(bytesBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	bodyResp, err := httpPkg.MakeRequest(ctx, r.client, req, 5, 5)
	if err != nil {
		log.Printf("GetTransactionByHash got error: %v", err)
		return nil, err
	}
	output := &models.GetTransactionByHashOutput{}
	err = json.Unmarshal(bodyResp, &output)
	if err != nil {
		log.Printf("GetTransactionByHash got error: %v", err)
		return nil, err
	}
	return output, nil
}
