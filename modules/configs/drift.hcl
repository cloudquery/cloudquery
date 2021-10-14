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

        resource "*" {
            ignore_identifiers = ["account_id"]
        }

        resource "ec2.instances" {
            #   identifiers       = ["id"]
            #   ignore_attributes = ["launch_time"]

            iac {
                terraform {
                    type = "aws_instance"
                    name = "aws_ec2_instances_ec2_instance"
                }
            }
        }

        resource "s3.buckets" {
            identifiers       = ["name"]
            #   ignore_attributes = ["launch_time"]

            iac {
                terraform {
                    type = "aws_s3_bucket"
                    name = "s3_bucket"
                }
            }
        }

        resource "iam.users" {
            ignore_attributes = [ "password_last_used" ]

            iac {
                terraform {
                    type = "aws_iam_user"
                    name = "users"
                }
            }
        }

        skip_resources = [
#            "ec2.instances"
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
