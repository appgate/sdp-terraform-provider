/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// ClientConnections struct for ClientConnections
type ClientConnections struct {
	// SPA mode.
	SpaMode *string `json:"spaMode,omitempty"`
	// Client Profiles.
	Profiles *[]ClientConnectionsProfiles `json:"profiles,omitempty"`
}

// NewClientConnections instantiates a new ClientConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientConnections() *ClientConnections {
	this := ClientConnections{}
	var spaMode string = "TCP"
	this.SpaMode = &spaMode
	return &this
}

// NewClientConnectionsWithDefaults instantiates a new ClientConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientConnectionsWithDefaults() *ClientConnections {
	this := ClientConnections{}
	var spaMode string = "TCP"
	this.SpaMode = &spaMode
	return &this
}

// GetSpaMode returns the SpaMode field value if set, zero value otherwise.
func (o *ClientConnections) GetSpaMode() string {
	if o == nil || o.SpaMode == nil {
		var ret string
		return ret
	}
	return *o.SpaMode
}

// GetSpaModeOk returns a tuple with the SpaMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConnections) GetSpaModeOk() (*string, bool) {
	if o == nil || o.SpaMode == nil {
		return nil, false
	}
	return o.SpaMode, true
}

// HasSpaMode returns a boolean if a field has been set.
func (o *ClientConnections) HasSpaMode() bool {
	if o != nil && o.SpaMode != nil {
		return true
	}

	return false
}

// SetSpaMode gets a reference to the given string and assigns it to the SpaMode field.
func (o *ClientConnections) SetSpaMode(v string) {
	o.SpaMode = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *ClientConnections) GetProfiles() []ClientConnectionsProfiles {
	if o == nil || o.Profiles == nil {
		var ret []ClientConnectionsProfiles
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConnections) GetProfilesOk() (*[]ClientConnectionsProfiles, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ClientConnections) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []ClientConnectionsProfiles and assigns it to the Profiles field.
func (o *ClientConnections) SetProfiles(v []ClientConnectionsProfiles) {
	o.Profiles = &v
}

func (o ClientConnections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpaMode != nil {
		toSerialize["spaMode"] = o.SpaMode
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullableClientConnections struct {
	value *ClientConnections
	isSet bool
}

func (v NullableClientConnections) Get() *ClientConnections {
	return v.value
}

func (v *NullableClientConnections) Set(val *ClientConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableClientConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableClientConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientConnections(val *ClientConnections) *NullableClientConnections {
	return &NullableClientConnections{value: val, isSet: true}
}

func (v NullableClientConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
