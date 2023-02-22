/*
 * Ory Hydra API
 *
 * Documentation for all of Ory Hydra's APIs. 
 *
 * API version: v1.11.8
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JSONWebKeySet It is important that this model object is named JSONWebKeySet for \"swagger generate spec\" to generate only on definition of a JSONWebKeySet. Since one with the same name is previously defined as client.Client.JSONWebKeys and this one is last, this one will be effectively written in the swagger spec.
type JSONWebKeySet struct {
	// The value of the \"keys\" parameter is an array of JWK values.  By default, the order of the JWK values within the array does not imply an order of preference among them, although applications of JWK Sets can choose to assign a meaning to the order for their purposes, if desired.
	Keys []JSONWebKey `json:"keys,omitempty"`
}

// NewJSONWebKeySet instantiates a new JSONWebKeySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONWebKeySet() *JSONWebKeySet {
	this := JSONWebKeySet{}
	return &this
}

// NewJSONWebKeySetWithDefaults instantiates a new JSONWebKeySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONWebKeySetWithDefaults() *JSONWebKeySet {
	this := JSONWebKeySet{}
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *JSONWebKeySet) GetKeys() []JSONWebKey {
	if o == nil || o.Keys == nil {
		var ret []JSONWebKey
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONWebKeySet) GetKeysOk() ([]JSONWebKey, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *JSONWebKeySet) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []JSONWebKey and assigns it to the Keys field.
func (o *JSONWebKeySet) SetKeys(v []JSONWebKey) {
	o.Keys = v
}

func (o JSONWebKeySet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableJSONWebKeySet struct {
	value *JSONWebKeySet
	isSet bool
}

func (v NullableJSONWebKeySet) Get() *JSONWebKeySet {
	return v.value
}

func (v *NullableJSONWebKeySet) Set(val *JSONWebKeySet) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONWebKeySet) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONWebKeySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONWebKeySet(val *JSONWebKeySet) *NullableJSONWebKeySet {
	return &NullableJSONWebKeySet{value: val, isSet: true}
}

func (v NullableJSONWebKeySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONWebKeySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


