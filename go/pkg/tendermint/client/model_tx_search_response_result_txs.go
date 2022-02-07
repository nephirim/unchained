/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// TxSearchResponseResultTxs struct for TxSearchResponseResultTxs
type TxSearchResponseResultTxs struct {
	Hash *string `json:"hash,omitempty"`
	Height *string `json:"height,omitempty"`
	Index *int32 `json:"index,omitempty"`
	TxResult *TxSearchResponseResultTxResult `json:"tx_result,omitempty"`
	Tx *string `json:"tx,omitempty"`
	Proof *TxSearchResponseResultProof `json:"proof,omitempty"`
}

// NewTxSearchResponseResultTxs instantiates a new TxSearchResponseResultTxs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxSearchResponseResultTxs() *TxSearchResponseResultTxs {
	this := TxSearchResponseResultTxs{}
	return &this
}

// NewTxSearchResponseResultTxsWithDefaults instantiates a new TxSearchResponseResultTxs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxSearchResponseResultTxsWithDefaults() *TxSearchResponseResultTxs {
	this := TxSearchResponseResultTxs{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *TxSearchResponseResultTxs) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultTxs) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *TxSearchResponseResultTxs) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *TxSearchResponseResultTxs) SetHash(v string) {
	o.Hash = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TxSearchResponseResultTxs) GetHeight() string {
	if o == nil || o.Height == nil {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultTxs) GetHeightOk() (*string, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TxSearchResponseResultTxs) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *TxSearchResponseResultTxs) SetHeight(v string) {
	o.Height = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TxSearchResponseResultTxs) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultTxs) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TxSearchResponseResultTxs) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TxSearchResponseResultTxs) SetIndex(v int32) {
	o.Index = &v
}

// GetTxResult returns the TxResult field value if set, zero value otherwise.
func (o *TxSearchResponseResultTxs) GetTxResult() TxSearchResponseResultTxResult {
	if o == nil || o.TxResult == nil {
		var ret TxSearchResponseResultTxResult
		return ret
	}
	return *o.TxResult
}

// GetTxResultOk returns a tuple with the TxResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultTxs) GetTxResultOk() (*TxSearchResponseResultTxResult, bool) {
	if o == nil || o.TxResult == nil {
		return nil, false
	}
	return o.TxResult, true
}

// HasTxResult returns a boolean if a field has been set.
func (o *TxSearchResponseResultTxs) HasTxResult() bool {
	if o != nil && o.TxResult != nil {
		return true
	}

	return false
}

// SetTxResult gets a reference to the given TxSearchResponseResultTxResult and assigns it to the TxResult field.
func (o *TxSearchResponseResultTxs) SetTxResult(v TxSearchResponseResultTxResult) {
	o.TxResult = &v
}

// GetTx returns the Tx field value if set, zero value otherwise.
func (o *TxSearchResponseResultTxs) GetTx() string {
	if o == nil || o.Tx == nil {
		var ret string
		return ret
	}
	return *o.Tx
}

// GetTxOk returns a tuple with the Tx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultTxs) GetTxOk() (*string, bool) {
	if o == nil || o.Tx == nil {
		return nil, false
	}
	return o.Tx, true
}

// HasTx returns a boolean if a field has been set.
func (o *TxSearchResponseResultTxs) HasTx() bool {
	if o != nil && o.Tx != nil {
		return true
	}

	return false
}

// SetTx gets a reference to the given string and assigns it to the Tx field.
func (o *TxSearchResponseResultTxs) SetTx(v string) {
	o.Tx = &v
}

// GetProof returns the Proof field value if set, zero value otherwise.
func (o *TxSearchResponseResultTxs) GetProof() TxSearchResponseResultProof {
	if o == nil || o.Proof == nil {
		var ret TxSearchResponseResultProof
		return ret
	}
	return *o.Proof
}

// GetProofOk returns a tuple with the Proof field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxSearchResponseResultTxs) GetProofOk() (*TxSearchResponseResultProof, bool) {
	if o == nil || o.Proof == nil {
		return nil, false
	}
	return o.Proof, true
}

// HasProof returns a boolean if a field has been set.
func (o *TxSearchResponseResultTxs) HasProof() bool {
	if o != nil && o.Proof != nil {
		return true
	}

	return false
}

// SetProof gets a reference to the given TxSearchResponseResultProof and assigns it to the Proof field.
func (o *TxSearchResponseResultTxs) SetProof(v TxSearchResponseResultProof) {
	o.Proof = &v
}

func (o TxSearchResponseResultTxs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.TxResult != nil {
		toSerialize["tx_result"] = o.TxResult
	}
	if o.Tx != nil {
		toSerialize["tx"] = o.Tx
	}
	if o.Proof != nil {
		toSerialize["proof"] = o.Proof
	}
	return json.Marshal(toSerialize)
}

type NullableTxSearchResponseResultTxs struct {
	value *TxSearchResponseResultTxs
	isSet bool
}

func (v NullableTxSearchResponseResultTxs) Get() *TxSearchResponseResultTxs {
	return v.value
}

func (v *NullableTxSearchResponseResultTxs) Set(val *TxSearchResponseResultTxs) {
	v.value = val
	v.isSet = true
}

func (v NullableTxSearchResponseResultTxs) IsSet() bool {
	return v.isSet
}

func (v *NullableTxSearchResponseResultTxs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxSearchResponseResultTxs(val *TxSearchResponseResultTxs) *NullableTxSearchResponseResultTxs {
	return &NullableTxSearchResponseResultTxs{value: val, isSet: true}
}

func (v NullableTxSearchResponseResultTxs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxSearchResponseResultTxs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

