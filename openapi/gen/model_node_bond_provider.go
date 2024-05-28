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

// NodeBondProvider struct for NodeBondProvider
type NodeBondProvider struct {
	BondAddress *string `json:"bond_address,omitempty"`
	Bond *string `json:"bond,omitempty"`
}

// NewNodeBondProvider instantiates a new NodeBondProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeBondProvider() *NodeBondProvider {
	this := NodeBondProvider{}
	return &this
}

// NewNodeBondProviderWithDefaults instantiates a new NodeBondProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeBondProviderWithDefaults() *NodeBondProvider {
	this := NodeBondProvider{}
	return &this
}

// GetBondAddress returns the BondAddress field value if set, zero value otherwise.
func (o *NodeBondProvider) GetBondAddress() string {
	if o == nil || o.BondAddress == nil {
		var ret string
		return ret
	}
	return *o.BondAddress
}

// GetBondAddressOk returns a tuple with the BondAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBondProvider) GetBondAddressOk() (*string, bool) {
	if o == nil || o.BondAddress == nil {
		return nil, false
	}
	return o.BondAddress, true
}

// HasBondAddress returns a boolean if a field has been set.
func (o *NodeBondProvider) HasBondAddress() bool {
	if o != nil && o.BondAddress != nil {
		return true
	}

	return false
}

// SetBondAddress gets a reference to the given string and assigns it to the BondAddress field.
func (o *NodeBondProvider) SetBondAddress(v string) {
	o.BondAddress = &v
}

// GetBond returns the Bond field value if set, zero value otherwise.
func (o *NodeBondProvider) GetBond() string {
	if o == nil || o.Bond == nil {
		var ret string
		return ret
	}
	return *o.Bond
}

// GetBondOk returns a tuple with the Bond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBondProvider) GetBondOk() (*string, bool) {
	if o == nil || o.Bond == nil {
		return nil, false
	}
	return o.Bond, true
}

// HasBond returns a boolean if a field has been set.
func (o *NodeBondProvider) HasBond() bool {
	if o != nil && o.Bond != nil {
		return true
	}

	return false
}

// SetBond gets a reference to the given string and assigns it to the Bond field.
func (o *NodeBondProvider) SetBond(v string) {
	o.Bond = &v
}

func (o NodeBondProvider) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BondAddress != nil {
		toSerialize["bond_address"] = o.BondAddress
	}
	if o.Bond != nil {
		toSerialize["bond"] = o.Bond
	}
	return json.Marshal(toSerialize)
}

type NullableNodeBondProvider struct {
	value *NodeBondProvider
	isSet bool
}

func (v NullableNodeBondProvider) Get() *NodeBondProvider {
	return v.value
}

func (v *NullableNodeBondProvider) Set(val *NodeBondProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeBondProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeBondProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeBondProvider(val *NodeBondProvider) *NullableNodeBondProvider {
	return &NullableNodeBondProvider{value: val, isSet: true}
}

func (v NullableNodeBondProvider) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeBondProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

