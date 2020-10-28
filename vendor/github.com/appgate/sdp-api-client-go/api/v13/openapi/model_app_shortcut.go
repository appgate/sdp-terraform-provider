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

// AppShortcut Publishes the configured URL as an app on the client using the display name as the app name.
type AppShortcut struct {
	// Name for the App Shortcut which will be visible on the Client UI.
	Name string `json:"name"`
	// Description for the App Shortcut which will be visible on the Client UI.
	Description *string `json:"description,omitempty"`
	// The URL that will be triggered on the OS to be handled. For example, an HTTPS URL will start the browser for the given URL.
	Url string `json:"url"`
	// The code of the published app on the client. - 1: Light Green - 2: Green - 3: Indigo - 4: Deep Purple - 5: Yellow - 6: Lime - 7: Light Blue - 8: Blue - 9: Amber - 10: Orange - 11: Cyan - 12: Teal - 13: Deep Orange - 14: Red - 15: Gray - 16: Brown - 17: Pink - 18: Purple - 19: Blue Gray - 20: Near Black
	ColorCode *int32 `json:"colorCode,omitempty"`
}

// NewAppShortcut instantiates a new AppShortcut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppShortcut(name string, url string) *AppShortcut {
	this := AppShortcut{}
	this.Name = name
	this.Url = url
	var colorCode int32 = 1
	this.ColorCode = &colorCode
	return &this
}

// NewAppShortcutWithDefaults instantiates a new AppShortcut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppShortcutWithDefaults() *AppShortcut {
	this := AppShortcut{}
	var colorCode int32 = 1
	this.ColorCode = &colorCode
	return &this
}

// GetName returns the Name field value
func (o *AppShortcut) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppShortcut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppShortcut) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppShortcut) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppShortcut) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppShortcut) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppShortcut) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value
func (o *AppShortcut) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AppShortcut) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AppShortcut) SetUrl(v string) {
	o.Url = v
}

// GetColorCode returns the ColorCode field value if set, zero value otherwise.
func (o *AppShortcut) GetColorCode() int32 {
	if o == nil || o.ColorCode == nil {
		var ret int32
		return ret
	}
	return *o.ColorCode
}

// GetColorCodeOk returns a tuple with the ColorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppShortcut) GetColorCodeOk() (*int32, bool) {
	if o == nil || o.ColorCode == nil {
		return nil, false
	}
	return o.ColorCode, true
}

// HasColorCode returns a boolean if a field has been set.
func (o *AppShortcut) HasColorCode() bool {
	if o != nil && o.ColorCode != nil {
		return true
	}

	return false
}

// SetColorCode gets a reference to the given int32 and assigns it to the ColorCode field.
func (o *AppShortcut) SetColorCode(v int32) {
	o.ColorCode = &v
}

func (o AppShortcut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.ColorCode != nil {
		toSerialize["colorCode"] = o.ColorCode
	}
	return json.Marshal(toSerialize)
}

// AsInlineResponse20013ResultOneOf wraps this instance of AppShortcut in InlineResponse20013ResultOneOf
func (s *AppShortcut) AsInlineResponse20013ResultOneOf() InlineResponse20013ResultOneOf {
	return InlineResponse20013ResultOneOf{InlineResponse20013ResultOneOfInterface: s}
}

type NullableAppShortcut struct {
	value *AppShortcut
	isSet bool
}

func (v NullableAppShortcut) Get() *AppShortcut {
	return v.value
}

func (v *NullableAppShortcut) Set(val *AppShortcut) {
	v.value = val
	v.isSet = true
}

func (v NullableAppShortcut) IsSet() bool {
	return v.isSet
}

func (v *NullableAppShortcut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppShortcut(val *AppShortcut) *NullableAppShortcut {
	return &NullableAppShortcut{value: val, isSet: true}
}

func (v NullableAppShortcut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppShortcut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
