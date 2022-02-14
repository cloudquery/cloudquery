config {

    provider "*" {
        # provider: the *provider.Provider
        # example:
        #  provider.Name is "aws"

        # resource: an entry in either provider.ResourceMap or provider.ResourceMap[].Relation
        # examples:
        #  resource.Key is the CQ name ("apigateway.api_keys")
        #  resource.Value.Options.PrimaryKeys is table primary key columns, with CQ relationship columns removed
        #  resource.Value.ColumnNames is table column names, with CQ relationship columns removed
        #  resource.Value.Name is the table name ("aws_apigateway_api_keys")

        resource "*" {
            identifiers       = resource.Value.Options.PrimaryKeys
            attributes        = resource.Value.ColumnNames
            ignore_attributes = ["creation_date", "creation_time", "created_at"]
            deep = false
        }
    }

}
