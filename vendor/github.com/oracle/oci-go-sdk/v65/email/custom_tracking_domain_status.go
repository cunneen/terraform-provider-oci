// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"strings"
)

// CustomTrackingDomainStatusEnum Enum with underlying type: string
type CustomTrackingDomainStatusEnum string

// Set of constants representing the allowable values for CustomTrackingDomainStatusEnum
const (
	CustomTrackingDomainStatusInactive       CustomTrackingDomainStatusEnum = "INACTIVE"
	CustomTrackingDomainStatusActive         CustomTrackingDomainStatusEnum = "ACTIVE"
	CustomTrackingDomainStatusNeedsAttention CustomTrackingDomainStatusEnum = "NEEDS_ATTENTION"
)

var mappingCustomTrackingDomainStatusEnum = map[string]CustomTrackingDomainStatusEnum{
	"INACTIVE":        CustomTrackingDomainStatusInactive,
	"ACTIVE":          CustomTrackingDomainStatusActive,
	"NEEDS_ATTENTION": CustomTrackingDomainStatusNeedsAttention,
}

var mappingCustomTrackingDomainStatusEnumLowerCase = map[string]CustomTrackingDomainStatusEnum{
	"inactive":        CustomTrackingDomainStatusInactive,
	"active":          CustomTrackingDomainStatusActive,
	"needs_attention": CustomTrackingDomainStatusNeedsAttention,
}

// GetCustomTrackingDomainStatusEnumValues Enumerates the set of values for CustomTrackingDomainStatusEnum
func GetCustomTrackingDomainStatusEnumValues() []CustomTrackingDomainStatusEnum {
	values := make([]CustomTrackingDomainStatusEnum, 0)
	for _, v := range mappingCustomTrackingDomainStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomTrackingDomainStatusEnumStringValues Enumerates the set of values in String for CustomTrackingDomainStatusEnum
func GetCustomTrackingDomainStatusEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"ACTIVE",
		"NEEDS_ATTENTION",
	}
}

// GetMappingCustomTrackingDomainStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomTrackingDomainStatusEnum(val string) (CustomTrackingDomainStatusEnum, bool) {
	enum, ok := mappingCustomTrackingDomainStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
