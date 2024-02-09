// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateEmailDomainDetails The attributes to update in the email domain.
type UpdateEmailDomainDetails struct {

	// Id for Domain in Domain Management (under governance) if DOMAINID verification method used.
	DomainVerificationId *string `mandatory:"false" json:"domainVerificationId"`

	// Controls when use of a private endpoint for email routing is required.
	// SEND means all mail from senders in this email domain will be privately routed.
	// RECEIVE means all mail sent to this recipient domain will be privately routed.
	// BOTH means both rules apply.
	// This can not be set to RECEIVE or BOTH without valid domain verification.
	// This can not be set to a value other than NONE unless emailPrivateEndpointId references an ACTIVE bi-directional submission email private endpoint.
	RequirePrivatePath RequirePrivatePathTypeEnum `mandatory:"false" json:"requirePrivatePath,omitempty"`

	// Id for the bi-directional submission Email Private Endpoint to use if requiring any private path.
	EmailPrivateEndpointId *string `mandatory:"false" json:"emailPrivateEndpointId"`

	// A string that describes the details about the domain. It does not have to be unique,
	// and you can change it. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A list of custom log headers.
	// Example: `["X-Campaign-ID", "Group-ID"]`
	CustomHeaders []string `mandatory:"false" json:"customHeaders"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateEmailDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateEmailDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRequirePrivatePathTypeEnum(string(m.RequirePrivatePath)); !ok && m.RequirePrivatePath != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequirePrivatePath: %s. Supported values are: %s.", m.RequirePrivatePath, strings.Join(GetRequirePrivatePathTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
