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

// OutboundFee struct for OutboundFee
type OutboundFee struct {
	// the asset to display the outbound fee for
	Asset string `json:"asset"`
	// the asset's outbound fee, in (1e8-format) units of the asset
	OutboundFee string `json:"outbound_fee"`
	// Total RUNE the network has withheld as fees to later cover gas costs for this asset's outbounds
	FeeWithheldRune *string `json:"fee_withheld_rune,omitempty"`
	// Total RUNE the network has spent to reimburse gas costs for this asset's outbounds
	FeeSpentRune *string `json:"fee_spent_rune,omitempty"`
	// amount of RUNE by which the fee_withheld_rune exceeds the fee_spent_rune
	SurplusRune *string `json:"surplus_rune,omitempty"`
	// dynamic multiplier basis points, based on the surplus_rune, affecting the size of the outbound_fee
	DynamicMultiplierBasisPoints *string `json:"dynamic_multiplier_basis_points,omitempty"`
}

// NewOutboundFee instantiates a new OutboundFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundFee(asset string, outboundFee string) *OutboundFee {
	this := OutboundFee{}
	this.Asset = asset
	this.OutboundFee = outboundFee
	return &this
}

// NewOutboundFeeWithDefaults instantiates a new OutboundFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundFeeWithDefaults() *OutboundFee {
	this := OutboundFee{}
	return &this
}

// GetAsset returns the Asset field value
func (o *OutboundFee) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *OutboundFee) GetAssetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *OutboundFee) SetAsset(v string) {
	o.Asset = v
}

// GetOutboundFee returns the OutboundFee field value
func (o *OutboundFee) GetOutboundFee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutboundFee
}

// GetOutboundFeeOk returns a tuple with the OutboundFee field value
// and a boolean to check if the value has been set.
func (o *OutboundFee) GetOutboundFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutboundFee, true
}

// SetOutboundFee sets field value
func (o *OutboundFee) SetOutboundFee(v string) {
	o.OutboundFee = v
}

// GetFeeWithheldRune returns the FeeWithheldRune field value if set, zero value otherwise.
func (o *OutboundFee) GetFeeWithheldRune() string {
	if o == nil || o.FeeWithheldRune == nil {
		var ret string
		return ret
	}
	return *o.FeeWithheldRune
}

// GetFeeWithheldRuneOk returns a tuple with the FeeWithheldRune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundFee) GetFeeWithheldRuneOk() (*string, bool) {
	if o == nil || o.FeeWithheldRune == nil {
		return nil, false
	}
	return o.FeeWithheldRune, true
}

// HasFeeWithheldRune returns a boolean if a field has been set.
func (o *OutboundFee) HasFeeWithheldRune() bool {
	if o != nil && o.FeeWithheldRune != nil {
		return true
	}

	return false
}

// SetFeeWithheldRune gets a reference to the given string and assigns it to the FeeWithheldRune field.
func (o *OutboundFee) SetFeeWithheldRune(v string) {
	o.FeeWithheldRune = &v
}

// GetFeeSpentRune returns the FeeSpentRune field value if set, zero value otherwise.
func (o *OutboundFee) GetFeeSpentRune() string {
	if o == nil || o.FeeSpentRune == nil {
		var ret string
		return ret
	}
	return *o.FeeSpentRune
}

// GetFeeSpentRuneOk returns a tuple with the FeeSpentRune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundFee) GetFeeSpentRuneOk() (*string, bool) {
	if o == nil || o.FeeSpentRune == nil {
		return nil, false
	}
	return o.FeeSpentRune, true
}

// HasFeeSpentRune returns a boolean if a field has been set.
func (o *OutboundFee) HasFeeSpentRune() bool {
	if o != nil && o.FeeSpentRune != nil {
		return true
	}

	return false
}

// SetFeeSpentRune gets a reference to the given string and assigns it to the FeeSpentRune field.
func (o *OutboundFee) SetFeeSpentRune(v string) {
	o.FeeSpentRune = &v
}

// GetSurplusRune returns the SurplusRune field value if set, zero value otherwise.
func (o *OutboundFee) GetSurplusRune() string {
	if o == nil || o.SurplusRune == nil {
		var ret string
		return ret
	}
	return *o.SurplusRune
}

// GetSurplusRuneOk returns a tuple with the SurplusRune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundFee) GetSurplusRuneOk() (*string, bool) {
	if o == nil || o.SurplusRune == nil {
		return nil, false
	}
	return o.SurplusRune, true
}

// HasSurplusRune returns a boolean if a field has been set.
func (o *OutboundFee) HasSurplusRune() bool {
	if o != nil && o.SurplusRune != nil {
		return true
	}

	return false
}

// SetSurplusRune gets a reference to the given string and assigns it to the SurplusRune field.
func (o *OutboundFee) SetSurplusRune(v string) {
	o.SurplusRune = &v
}

// GetDynamicMultiplierBasisPoints returns the DynamicMultiplierBasisPoints field value if set, zero value otherwise.
func (o *OutboundFee) GetDynamicMultiplierBasisPoints() string {
	if o == nil || o.DynamicMultiplierBasisPoints == nil {
		var ret string
		return ret
	}
	return *o.DynamicMultiplierBasisPoints
}

// GetDynamicMultiplierBasisPointsOk returns a tuple with the DynamicMultiplierBasisPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundFee) GetDynamicMultiplierBasisPointsOk() (*string, bool) {
	if o == nil || o.DynamicMultiplierBasisPoints == nil {
		return nil, false
	}
	return o.DynamicMultiplierBasisPoints, true
}

// HasDynamicMultiplierBasisPoints returns a boolean if a field has been set.
func (o *OutboundFee) HasDynamicMultiplierBasisPoints() bool {
	if o != nil && o.DynamicMultiplierBasisPoints != nil {
		return true
	}

	return false
}

// SetDynamicMultiplierBasisPoints gets a reference to the given string and assigns it to the DynamicMultiplierBasisPoints field.
func (o *OutboundFee) SetDynamicMultiplierBasisPoints(v string) {
	o.DynamicMultiplierBasisPoints = &v
}

func (o OutboundFee) MarshalJSON_deprecated() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset"] = o.Asset
	}
	if true {
		toSerialize["outbound_fee"] = o.OutboundFee
	}
	if o.FeeWithheldRune != nil {
		toSerialize["fee_withheld_rune"] = o.FeeWithheldRune
	}
	if o.FeeSpentRune != nil {
		toSerialize["fee_spent_rune"] = o.FeeSpentRune
	}
	if o.SurplusRune != nil {
		toSerialize["surplus_rune"] = o.SurplusRune
	}
	if o.DynamicMultiplierBasisPoints != nil {
		toSerialize["dynamic_multiplier_basis_points"] = o.DynamicMultiplierBasisPoints
	}
	return json.Marshal(toSerialize)
}

type NullableOutboundFee struct {
	value *OutboundFee
	isSet bool
}

func (v NullableOutboundFee) Get() *OutboundFee {
	return v.value
}

func (v *NullableOutboundFee) Set(val *OutboundFee) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundFee) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundFee(val *OutboundFee) *NullableOutboundFee {
	return &NullableOutboundFee{value: val, isSet: true}
}

func (v NullableOutboundFee) MarshalJSON_deprecated() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


