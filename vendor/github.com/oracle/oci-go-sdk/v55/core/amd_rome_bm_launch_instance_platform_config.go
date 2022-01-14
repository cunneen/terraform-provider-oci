// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// AmdRomeBmLaunchInstancePlatformConfig The platform configuration used when launching a bare metal instance with the AMD Rome platform.
type AmdRomeBmLaunchInstancePlatformConfig struct {

	// Whether Secure Boot is enabled on the instance.
	IsSecureBootEnabled *bool `mandatory:"false" json:"isSecureBootEnabled"`

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	IsTrustedPlatformModuleEnabled *bool `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`

	// Whether the Measured Boot feature is enabled on the instance.
	IsMeasuredBootEnabled *bool `mandatory:"false" json:"isMeasuredBootEnabled"`

	// Whether the instance is a confidential instance. If this value is `true`, the instance is a confidential instance. The default value is `false`.
	IsMemoryEncryptionEnabled *bool `mandatory:"false" json:"isMemoryEncryptionEnabled"`

	// The manufacturer specific technology used for memory encryption.
	MemoryEncryptionTechnology AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum `mandatory:"false" json:"memoryEncryptionTechnology,omitempty"`
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m AmdRomeBmLaunchInstancePlatformConfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m AmdRomeBmLaunchInstancePlatformConfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m AmdRomeBmLaunchInstancePlatformConfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

//GetIsMemoryEncryptionEnabled returns IsMemoryEncryptionEnabled
func (m AmdRomeBmLaunchInstancePlatformConfig) GetIsMemoryEncryptionEnabled() *bool {
	return m.IsMemoryEncryptionEnabled
}

func (m AmdRomeBmLaunchInstancePlatformConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmdRomeBmLaunchInstancePlatformConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum[string(m.MemoryEncryptionTechnology)]; !ok && m.MemoryEncryptionTechnology != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MemoryEncryptionTechnology: %s. Supported values are: %s.", m.MemoryEncryptionTechnology, strings.Join(GetAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AmdRomeBmLaunchInstancePlatformConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmdRomeBmLaunchInstancePlatformConfig AmdRomeBmLaunchInstancePlatformConfig
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAmdRomeBmLaunchInstancePlatformConfig
	}{
		"AMD_ROME_BM",
		(MarshalTypeAmdRomeBmLaunchInstancePlatformConfig)(m),
	}

	return json.Marshal(&s)
}

// AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum Enum with underlying type: string
type AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum string

// Set of constants representing the allowable values for AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum
const (
	AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyTsme AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum = "TSME"
	AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologySmee AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum = "SMEE"
)

var mappingAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum = map[string]AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum{
	"TSME": AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyTsme,
	"SMEE": AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologySmee,
}

// GetAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumValues Enumerates the set of values for AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum
func GetAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumValues() []AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum {
	values := make([]AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum, 0)
	for _, v := range mappingAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum {
		values = append(values, v)
	}
	return values
}

// GetAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumStringValues Enumerates the set of values in String for AmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnum
func GetAmdRomeBmLaunchInstancePlatformConfigMemoryEncryptionTechnologyEnumStringValues() []string {
	return []string{
		"TSME",
		"SMEE",
	}
}