package types

import (
	"encoding/json"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*callRequestMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (c CallRequest) MarshalJSON() ([]byte, error) {
	type CallRequest struct {
		From                 *common.Address              `json:"from,omitempty"`
		To                   *common.Address              `json:"to,omitempty"`
		Gas                  *hexutil.Uint64              `json:"gas,omitempty"`
		GasPrice             *hexutil.Big                 `json:"gasPrice,omitempty"`
		MaxFeePerGas         *hexutil.Big                 `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas *hexutil.Big                 `json:"maxPriorityFeePerGas,omitempty"`
		Value                *hexutil.Big                 `json:"value,omitempty"`
		Nonce                *hexutil.Uint64              `json:"nonce,omitempty"`
		Data                 hexutil.Bytes                `json:"data,omitempty"`
		Input                hexutil.Bytes                `json:"input,omitempty"`
		AccessList           *types.AccessList            `json:"accessList,omitempty"`
		AuthorizationList    []types.SetCodeAuthorization `json:"authorizationList,omitempty"`
		ChainID              *hexutil.Big                 `json:"chainId,omitempty"`
		Type                 *hexutil.Uint64              `json:"type,omitempty"`
	}
	var enc CallRequest
	enc.From = c.From
	enc.To = c.To
	enc.Gas = (*hexutil.Uint64)(c.Gas)
	enc.GasPrice = (*hexutil.Big)(c.GasPrice)
	enc.MaxFeePerGas = (*hexutil.Big)(c.MaxFeePerGas)
	enc.MaxPriorityFeePerGas = (*hexutil.Big)(c.MaxPriorityFeePerGas)
	enc.Value = (*hexutil.Big)(c.Value)
	enc.Nonce = (*hexutil.Uint64)(c.Nonce)
	enc.Data = c.Data
	enc.Input = c.Input
	enc.AccessList = c.AccessList
	enc.AuthorizationList = c.AuthorizationList
	enc.ChainID = (*hexutil.Big)(c.ChainID)
	enc.Type = (*hexutil.Uint64)(c.Type)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (c *CallRequest) UnmarshalJSON(input []byte) error {
	type CallRequest struct {
		From                 *common.Address              `json:"from,omitempty"`
		To                   *common.Address              `json:"to,omitempty"`
		Gas                  *hexutil.Uint64              `json:"gas,omitempty"`
		GasPrice             *hexutil.Big                 `json:"gasPrice,omitempty"`
		MaxFeePerGas         *hexutil.Big                 `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas *hexutil.Big                 `json:"maxPriorityFeePerGas,omitempty"`
		Value                *hexutil.Big                 `json:"value,omitempty"`
		Nonce                *hexutil.Uint64              `json:"nonce,omitempty"`
		Data                 *hexutil.Bytes               `json:"data,omitempty"`
		Input                *hexutil.Bytes               `json:"input,omitempty"`
		AccessList           *types.AccessList            `json:"accessList,omitempty"`
		AuthorizationList    []types.SetCodeAuthorization `json:"authorizationList,omitempty"`
		ChainID              *hexutil.Big                 `json:"chainId,omitempty"`
		Type                 *string                      `json:"type,omitempty"`
	}
	var dec CallRequest
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.From != nil {
		c.From = dec.From
	}
	if dec.To != nil {
		c.To = dec.To
	}
	if dec.Gas != nil {
		c.Gas = (*uint64)(dec.Gas)
	}
	if dec.GasPrice != nil {
		c.GasPrice = (*big.Int)(dec.GasPrice)
	}
	if dec.MaxFeePerGas != nil {
		c.MaxFeePerGas = (*big.Int)(dec.MaxFeePerGas)
	}
	if dec.MaxPriorityFeePerGas != nil {
		c.MaxPriorityFeePerGas = (*big.Int)(dec.MaxPriorityFeePerGas)
	}
	if dec.Value != nil {
		c.Value = (*big.Int)(dec.Value)
	}
	if dec.Nonce != nil {
		c.Nonce = (*uint64)(dec.Nonce)
	}
	if dec.Data != nil {
		c.Data = *dec.Data
	}
	if dec.Input != nil {
		c.Input = *dec.Input
	}
	if dec.AccessList != nil {
		c.AccessList = dec.AccessList
	}
	if dec.AuthorizationList != nil {
		c.AuthorizationList = dec.AuthorizationList
	}
	if dec.ChainID != nil {
		c.ChainID = (*big.Int)(dec.ChainID)
	}
	if dec.Type != nil {
		valInt64, err := strconv.ParseInt(*dec.Type, 0, 64)
		if err != nil {
			return err
		}
		valUint64 := uint64(valInt64)
		c.Type = &valUint64
	}
	return nil
}
