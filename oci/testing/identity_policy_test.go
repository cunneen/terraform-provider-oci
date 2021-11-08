// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package testing

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v49/common"
	oci_identity "github.com/oracle/oci-go-sdk/v49/identity"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PolicyRequiredOnlyResource = PolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Required, Create, policyRepresentation)

	policyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"name":           Representation{RepType: Optional, Create: `LaunchInstances`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{Required, policyDataSourceFilterRepresentation}}
	policyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_identity_policy.test_policy.id}`}},
	}

	policyRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.tenancy_ocid}`},
		"description":    Representation{RepType: Required, Create: `Policy for users who need to launch instances, attach volumes, manage images`, Update: `description2`},
		"name":           Representation{RepType: Required, Create: `LaunchInstances`},
		"statements":     Representation{RepType: Required, Create: []string{`Allow Group Administrators to read instances in tenancy`}},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"version_date":   Representation{RepType: Optional, Create: ``, Update: `2018-01-01`},
	}

	PolicyResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: identity/default
func TestIdentityPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestIdentityPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	tenancyId := utils.GetEnvSettingWithBlankDefault("tenancy_ocid")

	resourceName := "oci_identity_policy.test_policy"
	datasourceName := "data.oci_identity_policies.test_policies"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Create, policyRepresentation), "identity", "policy", t)

	acctest.ResourceTest(t, testAccCheckIdentityPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + PolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Required, Create, policyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "description", "Policy for users who need to launch instances, attach volumes, manage images"),
				resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + PolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + PolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Create, policyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "Policy for users who need to launch instances, attach volumes, manage images"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckNoResourceAttr(resourceName, "version_date"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithBlankDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + PolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Update, policyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "statements.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "version_date", "2018-01-01"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_identity_policies", "test_policies", Optional, Update, policyDataSourceRepresentation) +
				compartmentIdVariableStr + PolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", Optional, Update, policyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "name", "LaunchInstances"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.compartment_id", tenancyId),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "policies.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.name", "LaunchInstances"),
				resource.TestCheckResourceAttrSet(datasourceName, "policies.0.state"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.statements.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "policies.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "policies.0.version_date", "2018-01-01"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				// ETag, lastUpdateETag, and policyHash are non-API fields that
				// get computed during resource Create/Update but omitted from Get calls.
				// These are internally used for diff suppression and not needed for imports.
				// Omit them in the import verification.
				"ETag",
				"lastUpdateETag",
				"policyHash",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckIdentityPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := TestAccProvider.Meta().(*OracleClients).identityClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_identity_policy" {
			noResourceFound = false
			request := oci_identity.GetPolicyRequest{}

			tmp := rs.Primary.ID
			request.PolicyId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "identity")

			response, err := client.GetPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_identity.PolicyLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}