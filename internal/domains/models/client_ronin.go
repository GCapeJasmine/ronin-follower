package models

import (
	"github.com/jinzhu/copier"
)

const (
	MethodGetLatestBlockNumber = "eth_blockNumber"
	MethodGetBlockByNumber     = "eth_getBlockByNumber"
	MethodGetTransactionByHash = "eth_getTransactionByHash"
)

type GetLatestBlockNumberInput struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int64  `json:"id"`
}

type GetLatestBlockNumberOutput struct {
	Result string `json:"result"`
}

type GetBlockByNumberInput struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int64  `json:"id"`
}

type GetBlockByNumberOutput struct {
	Result struct {
		Number       string `json:"number"`
		Hash         string `json:"hash"`
		ParentHash   string `json:"parentHash"`
		GasLimit     string `json:"gasLimit"`
		GasUsed      string `json:"gasUsed"`
		Nonce        string `json:"nonce"`
		Transactions []struct {
			From     string `json:"from"`
			To       string `json:"to"`
			Gas      string `json:"gas"`
			GasPrice string `json:"gasPrice"`
			Hash     string `json:"hash"`
			Nonce    string `json:"nonce"`
		} `json:"transactions"`
	} `json:"result"`
}

type GetTransactionByHashInput struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	Id      int64  `json:"id"`
}

type GetTransactionByHashOutput struct {
	Result struct {
		From     string `json:"from"`
		To       string `json:"to"`
		Gas      string `json:"gas"`
		GasPrice string `json:"gasPrice"`
		Hash     string `json:"hash"`
		Nonce    string `json:"nonce"`
	} `json:"result"`
}

func (b *GetBlockByNumberOutput) ToBlock() *Block {
	res := &Block{}
	_ = copier.Copy(res, b.Result)
	return res
}

func (t *GetTransactionByHashOutput) ToTransaction() *Transaction {
	res := &Transaction{}
	_ = copier.Copy(res, t.Result)
	return res
}
