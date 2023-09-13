// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"strings"
)

// DrsFileSortByEnum Enum with underlying type: string
type DrsFileSortByEnum string

// Set of constants representing the allowable values for DrsFileSortByEnum
const (
	DrsFileSortByBucketName   DrsFileSortByEnum = "bucketName"
	DrsFileSortByNamespace    DrsFileSortByEnum = "namespace"
	DrsFileSortByDrsFileKey   DrsFileSortByEnum = "drsFileKey"
	DrsFileSortByDrsFileName  DrsFileSortByEnum = "drsFileName"
	DrsFileSortByChecksumType DrsFileSortByEnum = "checksumType"
	DrsFileSortByIsDefault    DrsFileSortByEnum = "isDefault"
)

var mappingDrsFileSortByEnum = map[string]DrsFileSortByEnum{
	"bucketName":   DrsFileSortByBucketName,
	"namespace":    DrsFileSortByNamespace,
	"drsFileKey":   DrsFileSortByDrsFileKey,
	"drsFileName":  DrsFileSortByDrsFileName,
	"checksumType": DrsFileSortByChecksumType,
	"isDefault":    DrsFileSortByIsDefault,
}

var mappingDrsFileSortByEnumLowerCase = map[string]DrsFileSortByEnum{
	"bucketname":   DrsFileSortByBucketName,
	"namespace":    DrsFileSortByNamespace,
	"drsfilekey":   DrsFileSortByDrsFileKey,
	"drsfilename":  DrsFileSortByDrsFileName,
	"checksumtype": DrsFileSortByChecksumType,
	"isdefault":    DrsFileSortByIsDefault,
}

// GetDrsFileSortByEnumValues Enumerates the set of values for DrsFileSortByEnum
func GetDrsFileSortByEnumValues() []DrsFileSortByEnum {
	values := make([]DrsFileSortByEnum, 0)
	for _, v := range mappingDrsFileSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetDrsFileSortByEnumStringValues Enumerates the set of values in String for DrsFileSortByEnum
func GetDrsFileSortByEnumStringValues() []string {
	return []string{
		"bucketName",
		"namespace",
		"drsFileKey",
		"drsFileName",
		"checksumType",
		"isDefault",
	}
}

// GetMappingDrsFileSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrsFileSortByEnum(val string) (DrsFileSortByEnum, bool) {
	enum, ok := mappingDrsFileSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
