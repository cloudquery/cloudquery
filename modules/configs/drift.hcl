module "drift" {

    provider "*" {
        # provider: the *provider.Provider
        # example:
        #  provider.Name is "aws"
        # special case:
        #  provider.ModuleHcl is the config provider supplies

        # resource: an entry in provider.ResourceMap
        # examples:
        #  resource.Key is the CQ name ("apigateway.api_keys")
        #  resource.Value.Name is the table name ("aws_apigateway_api_keys")

        resource "*" {
            identifiers       = resource.Value.Options.PrimaryKeys
            ignore_attributes = ["creation_date", "launch_time"]
        }

        # "source" directive evaluates the given config or statement
        source = provider.ModuleHcl
    }

    # TODO get from provider... But this could also override/decorate the * entry above, if specified
    provider "aws" {
        version = ">=0.5.10"

         resource "ec2_instances" {
        #   identifiers       = ["id"]
        #   ignore_attributes = ["launch_time"]
         }

        resource "iam_users" {
            ignore_attributes = [ "password_last_used" ]
            tf_type = "aws_iam_user"
            tf_name = "users"
        }

        skip_resources = [
            "ec2_instances"
        ]

    }


}

module "terraformer" {

    provider "aws" {
        tftemplate "*" {
#            ...
        }

        tftemplate "instance" {
#            ...
        }
    }

    provider "gcp" {
        tftemplate {
#            ...
        }
    }

}
