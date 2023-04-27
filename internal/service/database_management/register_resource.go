// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_database_management_db_management_private_endpoint", DatabaseManagementDbManagementPrivateEndpointResource())
	tfresource.RegisterResource("oci_database_management_external_asm", DatabaseManagementExternalAsmResource())
	tfresource.RegisterResource("oci_database_management_external_cluster", DatabaseManagementExternalClusterResource())
	tfresource.RegisterResource("oci_database_management_external_cluster_instance", DatabaseManagementExternalClusterInstanceResource())
	tfresource.RegisterResource("oci_database_management_external_db_node", DatabaseManagementExternalDbNodeResource())
	tfresource.RegisterResource("oci_database_management_external_db_system", DatabaseManagementExternalDbSystemResource())
	tfresource.RegisterResource("oci_database_management_external_db_system_connector", DatabaseManagementExternalDbSystemConnectorResource())
	tfresource.RegisterResource("oci_database_management_external_db_system_database_managements_management", DatabaseManagementExternalDbSystemDatabaseManagementsManagementResource())
	tfresource.RegisterResource("oci_database_management_external_db_system_discovery", DatabaseManagementExternalDbSystemDiscoveryResource())
	tfresource.RegisterResource("oci_database_management_external_exadata_infrastructure", DatabaseManagementExternalExadataInfrastructureResource())
	tfresource.RegisterResource("oci_database_management_external_exadata_infrastructure_database_managements_management", DatabaseManagementExternalExadataInfrastructureDatabaseManagementsManagementResource())
	tfresource.RegisterResource("oci_database_management_external_exadata_storage_connector", DatabaseManagementExternalExadataStorageConnectorResource())
	tfresource.RegisterResource("oci_database_management_external_listener", DatabaseManagementExternalListenerResource())
	tfresource.RegisterResource("oci_database_management_managed_database_group", DatabaseManagementManagedDatabaseGroupResource())
	tfresource.RegisterResource("oci_database_management_managed_databases_change_database_parameter", DatabaseManagementManagedDatabasesChangeDatabaseParameterResource())
	tfresource.RegisterResource("oci_database_management_managed_databases_reset_database_parameter", DatabaseManagementManagedDatabasesResetDatabaseParameterResource())
}
