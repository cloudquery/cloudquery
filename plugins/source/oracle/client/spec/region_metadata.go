package spec

// Region metadata.
// For more information see
// [Adding Regions](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdk_adding_new_region_endpoints.htm)
// in Oracle documentation.
type RegionMetadata struct {
	RealmKey             string `json:"realmKey" jsonschema:"required,minLength=1"`
	RealmDomainComponent string `json:"realmDomainComponent" jsonschema:"required,minLength=1"`
	RegionKey            string `json:"regionKey" jsonschema:"required,minLength=1"`
	RegionIdentifier     string `json:"regionIdentifier" jsonschema:"required,minLength=1"`
}
