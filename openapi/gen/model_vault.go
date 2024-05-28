/*
Thornode API

Thornode REST API.

Contact: devs@thorchain.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Vault struct for Vault
type Vault struct {
	BlockHeight *int64 `json:"block_height,omitempty"`
	PubKey *string `json:"pub_key,omitempty"`
	Coins []Coin `json:"coins"`
	Type *string `json:"type,omitempty"`
	Status string `json:"status"`
	StatusSince *int64 `json:"status_since,omitempty"`
	// the list of node public keys which are members of the vault
	Membership []string `json:"membership,omitempty"`
	Chains []string `json:"chains,omitempty"`
	InboundTxCount *int64 `json:"inbound_tx_count,omitempty"`
	OutboundTxCount *int64 `json:"outbound_tx_count,omitempty"`
	PendingTxBlockHeights []int64 `json:"pending_tx_block_heights,omitempty"`
	Routers []VaultRouter `json:"routers"`
	Addresses []VaultAddress `json:"addresses"`
	Frozen []string `json:"frozen,omitempty"`
}

// NewVault instantiates a new Vault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVault(coins []Coin, status string, routers []VaultRouter, addresses []VaultAddress) *Vault {
	this := Vault{}
	this.Coins = coins
	this.Status = status
	this.Routers = routers
	this.Addresses = addresses
	return &this
}

// NewVaultWithDefaults instantiates a new Vault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultWithDefaults() *Vault {
	this := Vault{}
	return &this
}

// GetBlockHeight returns the BlockHeight field value if set, zero value otherwise.
func (o *Vault) GetBlockHeight() int64 {
	if o == nil || o.BlockHeight == nil {
		var ret int64
		return ret
	}
	return *o.BlockHeight
}

// GetBlockHeightOk returns a tuple with the BlockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetBlockHeightOk() (*int64, bool) {
	if o == nil || o.BlockHeight == nil {
		return nil, false
	}
	return o.BlockHeight, true
}

// HasBlockHeight returns a boolean if a field has been set.
func (o *Vault) HasBlockHeight() bool {
	if o != nil && o.BlockHeight != nil {
		return true
	}

	return false
}

// SetBlockHeight gets a reference to the given int64 and assigns it to the BlockHeight field.
func (o *Vault) SetBlockHeight(v int64) {
	o.BlockHeight = &v
}

// GetPubKey returns the PubKey field value if set, zero value otherwise.
func (o *Vault) GetPubKey() string {
	if o == nil || o.PubKey == nil {
		var ret string
		return ret
	}
	return *o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetPubKeyOk() (*string, bool) {
	if o == nil || o.PubKey == nil {
		return nil, false
	}
	return o.PubKey, true
}

// HasPubKey returns a boolean if a field has been set.
func (o *Vault) HasPubKey() bool {
	if o != nil && o.PubKey != nil {
		return true
	}

	return false
}

// SetPubKey gets a reference to the given string and assigns it to the PubKey field.
func (o *Vault) SetPubKey(v string) {
	o.PubKey = &v
}

// GetCoins returns the Coins field value
func (o *Vault) GetCoins() []Coin {
	if o == nil {
		var ret []Coin
		return ret
	}

	return o.Coins
}

// GetCoinsOk returns a tuple with the Coins field value
// and a boolean to check if the value has been set.
func (o *Vault) GetCoinsOk() ([]Coin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coins, true
}

// SetCoins sets field value
func (o *Vault) SetCoins(v []Coin) {
	o.Coins = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Vault) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Vault) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Vault) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value
func (o *Vault) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Vault) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Vault) SetStatus(v string) {
	o.Status = v
}

// GetStatusSince returns the StatusSince field value if set, zero value otherwise.
func (o *Vault) GetStatusSince() int64 {
	if o == nil || o.StatusSince == nil {
		var ret int64
		return ret
	}
	return *o.StatusSince
}

// GetStatusSinceOk returns a tuple with the StatusSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetStatusSinceOk() (*int64, bool) {
	if o == nil || o.StatusSince == nil {
		return nil, false
	}
	return o.StatusSince, true
}

// HasStatusSince returns a boolean if a field has been set.
func (o *Vault) HasStatusSince() bool {
	if o != nil && o.StatusSince != nil {
		return true
	}

	return false
}

// SetStatusSince gets a reference to the given int64 and assigns it to the StatusSince field.
func (o *Vault) SetStatusSince(v int64) {
	o.StatusSince = &v
}

// GetMembership returns the Membership field value if set, zero value otherwise.
func (o *Vault) GetMembership() []string {
	if o == nil || o.Membership == nil {
		var ret []string
		return ret
	}
	return o.Membership
}

// GetMembershipOk returns a tuple with the Membership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetMembershipOk() ([]string, bool) {
	if o == nil || o.Membership == nil {
		return nil, false
	}
	return o.Membership, true
}

// HasMembership returns a boolean if a field has been set.
func (o *Vault) HasMembership() bool {
	if o != nil && o.Membership != nil {
		return true
	}

	return false
}

// SetMembership gets a reference to the given []string and assigns it to the Membership field.
func (o *Vault) SetMembership(v []string) {
	o.Membership = v
}

// GetChains returns the Chains field value if set, zero value otherwise.
func (o *Vault) GetChains() []string {
	if o == nil || o.Chains == nil {
		var ret []string
		return ret
	}
	return o.Chains
}

// GetChainsOk returns a tuple with the Chains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetChainsOk() ([]string, bool) {
	if o == nil || o.Chains == nil {
		return nil, false
	}
	return o.Chains, true
}

// HasChains returns a boolean if a field has been set.
func (o *Vault) HasChains() bool {
	if o != nil && o.Chains != nil {
		return true
	}

	return false
}

// SetChains gets a reference to the given []string and assigns it to the Chains field.
func (o *Vault) SetChains(v []string) {
	o.Chains = v
}

// GetInboundTxCount returns the InboundTxCount field value if set, zero value otherwise.
func (o *Vault) GetInboundTxCount() int64 {
	if o == nil || o.InboundTxCount == nil {
		var ret int64
		return ret
	}
	return *o.InboundTxCount
}

// GetInboundTxCountOk returns a tuple with the InboundTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetInboundTxCountOk() (*int64, bool) {
	if o == nil || o.InboundTxCount == nil {
		return nil, false
	}
	return o.InboundTxCount, true
}

// HasInboundTxCount returns a boolean if a field has been set.
func (o *Vault) HasInboundTxCount() bool {
	if o != nil && o.InboundTxCount != nil {
		return true
	}

	return false
}

// SetInboundTxCount gets a reference to the given int64 and assigns it to the InboundTxCount field.
func (o *Vault) SetInboundTxCount(v int64) {
	o.InboundTxCount = &v
}

// GetOutboundTxCount returns the OutboundTxCount field value if set, zero value otherwise.
func (o *Vault) GetOutboundTxCount() int64 {
	if o == nil || o.OutboundTxCount == nil {
		var ret int64
		return ret
	}
	return *o.OutboundTxCount
}

// GetOutboundTxCountOk returns a tuple with the OutboundTxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetOutboundTxCountOk() (*int64, bool) {
	if o == nil || o.OutboundTxCount == nil {
		return nil, false
	}
	return o.OutboundTxCount, true
}

// HasOutboundTxCount returns a boolean if a field has been set.
func (o *Vault) HasOutboundTxCount() bool {
	if o != nil && o.OutboundTxCount != nil {
		return true
	}

	return false
}

// SetOutboundTxCount gets a reference to the given int64 and assigns it to the OutboundTxCount field.
func (o *Vault) SetOutboundTxCount(v int64) {
	o.OutboundTxCount = &v
}

// GetPendingTxBlockHeights returns the PendingTxBlockHeights field value if set, zero value otherwise.
func (o *Vault) GetPendingTxBlockHeights() []int64 {
	if o == nil || o.PendingTxBlockHeights == nil {
		var ret []int64
		return ret
	}
	return o.PendingTxBlockHeights
}

// GetPendingTxBlockHeightsOk returns a tuple with the PendingTxBlockHeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetPendingTxBlockHeightsOk() ([]int64, bool) {
	if o == nil || o.PendingTxBlockHeights == nil {
		return nil, false
	}
	return o.PendingTxBlockHeights, true
}

// HasPendingTxBlockHeights returns a boolean if a field has been set.
func (o *Vault) HasPendingTxBlockHeights() bool {
	if o != nil && o.PendingTxBlockHeights != nil {
		return true
	}

	return false
}

// SetPendingTxBlockHeights gets a reference to the given []int64 and assigns it to the PendingTxBlockHeights field.
func (o *Vault) SetPendingTxBlockHeights(v []int64) {
	o.PendingTxBlockHeights = v
}

// GetRouters returns the Routers field value
func (o *Vault) GetRouters() []VaultRouter {
	if o == nil {
		var ret []VaultRouter
		return ret
	}

	return o.Routers
}

// GetRoutersOk returns a tuple with the Routers field value
// and a boolean to check if the value has been set.
func (o *Vault) GetRoutersOk() ([]VaultRouter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Routers, true
}

// SetRouters sets field value
func (o *Vault) SetRouters(v []VaultRouter) {
	o.Routers = v
}

// GetAddresses returns the Addresses field value
func (o *Vault) GetAddresses() []VaultAddress {
	if o == nil {
		var ret []VaultAddress
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *Vault) GetAddressesOk() ([]VaultAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *Vault) SetAddresses(v []VaultAddress) {
	o.Addresses = v
}

// GetFrozen returns the Frozen field value if set, zero value otherwise.
func (o *Vault) GetFrozen() []string {
	if o == nil || o.Frozen == nil {
		var ret []string
		return ret
	}
	return o.Frozen
}

// GetFrozenOk returns a tuple with the Frozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetFrozenOk() ([]string, bool) {
	if o == nil || o.Frozen == nil {
		return nil, false
	}
	return o.Frozen, true
}

// HasFrozen returns a boolean if a field has been set.
func (o *Vault) HasFrozen() bool {
	if o != nil && o.Frozen != nil {
		return true
	}

	return false
}

// SetFrozen gets a reference to the given []string and assigns it to the Frozen field.
func (o *Vault) SetFrozen(v []string) {
	o.Frozen = v
}

func (o Vault) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockHeight != nil {
		toSerialize["block_height"] = o.BlockHeight
	}
	if o.PubKey != nil {
		toSerialize["pub_key"] = o.PubKey
	}
	if true {
		toSerialize["coins"] = o.Coins
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StatusSince != nil {
		toSerialize["status_since"] = o.StatusSince
	}
	if o.Membership != nil {
		toSerialize["membership"] = o.Membership
	}
	if o.Chains != nil {
		toSerialize["chains"] = o.Chains
	}
	if o.InboundTxCount != nil {
		toSerialize["inbound_tx_count"] = o.InboundTxCount
	}
	if o.OutboundTxCount != nil {
		toSerialize["outbound_tx_count"] = o.OutboundTxCount
	}
	if o.PendingTxBlockHeights != nil {
		toSerialize["pending_tx_block_heights"] = o.PendingTxBlockHeights
	}
	if true {
		toSerialize["routers"] = o.Routers
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Frozen != nil {
		toSerialize["frozen"] = o.Frozen
	}
	return json.Marshal(toSerialize)
}

type NullableVault struct {
	value *Vault
	isSet bool
}

func (v NullableVault) Get() *Vault {
	return v.value
}

func (v *NullableVault) Set(val *Vault) {
	v.value = val
	v.isSet = true
}

func (v NullableVault) IsSet() bool {
	return v.isSet
}

func (v *NullableVault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVault(val *Vault) *NullableVault {
	return &NullableVault{value: val, isSet: true}
}

func (v NullableVault) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

