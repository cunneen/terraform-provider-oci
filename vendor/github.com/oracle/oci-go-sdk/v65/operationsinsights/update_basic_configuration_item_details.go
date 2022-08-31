// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package operationsinsights

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateBasicConfigurationItemDetails Configuration item details for OPSI configuration update.
type UpdateBasicConfigurationItemDetails struct {

	// Name of configuration item.
	Name *string `mandatory:"false" json:"name"`

	// Value of configuration item.
	Value *string `mandatory:"false" json:"value"`
}

func (m UpdateBasicConfigurationItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBasicConfigurationItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateBasicConfigurationItemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateBasicConfigurationItemDetails UpdateBasicConfigurationItemDetails
	s := struct {
		DiscriminatorParam string `json:"configItemType"`
		MarshalTypeUpdateBasicConfigurationItemDetails
	}{
		"BASIC",
		(MarshalTypeUpdateBasicConfigurationItemDetails)(m),
	}

	return json.Marshal(&s)
}
