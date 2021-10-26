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
        #  resource.Value.ColumnNames is table column names
        #  resource.Value.Name is the table name ("aws_apigateway_api_keys")

        resource "*" {
            identifiers       = resource.Value.Options.PrimaryKeys
            attributes        = resource.Value.ColumnNames
            ignore_attributes = ["cq_id", "meta", "creation_date"]
            deep = false
        }

        # "source" directive evaluates the given config or statement
        source = provider.ModuleHcl
    }

    # TODO get from provider... But this could also override/decorate the * entry above, if specified
    provider "aws" {
        version = ">=0.5.10"

        resource "*" {
            ignore_identifiers = [ "account_id", "region" ]
            ignore_attributes = [ "unknown_fields" ]

            iac {
                terraform {
                    # map of attributes from cloud provider to iac provider
                    attribute_map = [
                        "tags=tags_all"
                    ]
                }
            }
        }

        resource "accessanalyzer.analyzers" {
            identifiers = [ "name" ]

            iac {
                terraform {
                    type = "aws_accessanalyzer_analyzer"
                }
            }
        }

        resource "apigateway.api_keys" {
            iac {
                terraform {
                    type = "aws_api_gateway_api_key"
                }
            }
        }

        resource "apigateway.client_certificates" {
            iac {
                terraform {
                    type = "aws_api_gateway_client_certificate"
                }
            }
        }

        # Unmatched: apigateway.domain_names

        resource "apigateway.rest_apis" {
            iac {
                terraform {
                    type = "aws_api_gateway_rest_api"
                }
            }
        }

        resource "apigateway.usage_plans" {
            iac {
                terraform {
                    type = "aws_api_gateway_usage_plan"
                }
            }
        }

        resource "apigateway.vpc_links" {
            iac {
                terraform {
                    type = "aws_api_gateway_vpc_link"
                }
            }
        }

        resource "apigatewayv2.apis" {
            iac {
                terraform {
                    type = "aws_apigatewayv2_api"
                }
            }
        }

        # Unmatched: apigatewayv2.domain_names

        resource "apigatewayv2.vpc_links" {
            iac {
                terraform {
                    type = "aws_apigatewayv2_vpc_link"
                }
            }
        }

        resource "autoscaling.launch_configurations" {
            identifiers = [ "launch_configuration_name" ]

            iac {
                terraform {
                    type = "aws_launch_configuration"
                }
            }
        }

        resource "cloudfront.cache_policies" {
            iac {
                terraform {
                    type = "aws_cloudfront_cache_policy"
                }
            }
        }

        resource "cloudfront.distributions" {
            iac {
                terraform {
                    type = "aws_cloudfront_distribution"
                }
            }
        }

        resource "cloudtrail.trails" {
            identifiers = [ "name" ]

            iac {
                terraform {
                    type = "aws_cloudtrail"
                }
            }
        }

        resource "cloudwatch.alarms" {
            iac {
                terraform {
                    type = "aws_cloudwatch_metric_alarm"
                }
            }
        }

        resource "cloudwatchlogs.filters" {
            identifiers = [ "name" ]

            iac {
                terraform {
                    type = "aws_cloudwatch_log_metric_filter"
                }
            }
        }

        resource "cognito.identity_pools" {
            iac {
                terraform {
                    type = "aws_cognito_identity_pool"
                }
            }
        }

        resource "cognito.user_pools" {
            iac {
                terraform {
                    type = "aws_cognito_user_pool"
                }
            }
        }

        resource "config.configuration_recorders" {
            identifiers = [ "name" ]

            iac {
                terraform {
                    type = "aws_config_configuration_recorder"
                }
            }
        }

        resource "config.conformance_packs" {
            identifiers = [ "conformance_pack_name" ]

            iac {
                terraform {
                    type = "aws_config_conformance_pack"
                }
            }
        }

        resource "directconnect.connections" {
            iac {
                terraform {
                    type = "aws_dx_connection"
                }
            }
        }

        resource "directconnect.gateways" {
            iac {
                terraform {
                    type = "aws_dx_gateway"
                }
            }
        }

        resource "directconnect.lags" {
            iac {
                terraform {
                    type = "aws_dx_lag"
                }
            }
        }

        # Unmatched: directconnect.virtual_gateways

        resource "directconnect.virtual_interfaces" {
            iac {
                terraform {
                    type = "aws_dx_public_virtual_interface"
                }
            }
        }

        # Unmatched: ec2.byoip_cidrs

        # Unmatched: ec2.customer_gateways

        resource "ec2.ebs_volumes" {
            iac {
                terraform {
                    type = "aws_ebs_volume"
                }
            }
        }

        resource "ec2.flow_logs" {
            iac {
                terraform {
                    type = "aws_flow_log"
                }
            }
        }

        # Unmatched: ec2.images

        resource "ec2.instances" {
            ignore_attributes = ["launch_time"]

            iac {
                terraform {
                    type = "aws_instance"
                }
            }
        }

        resource "ec2.internet_gateways" {
            iac {
                terraform {
                    type = "aws_internet_gateway"
                }
            }
        }

        resource "ec2.nat_gateways" {
            iac {
                terraform {
                    type = "aws_nat_gateway"
                }
            }
        }

        resource "ec2.network_acls" {
            iac {
                terraform {
                    type = "aws_network_acl"
                }
            }
        }

        # Unmatched: ec2.regional_config

        resource "ec2.route_tables" {
            iac {
                terraform {
                    type = "aws_route_table"
                }
            }
        }

        # Unmatched: ec2.security_groups

        resource "ec2.subnets" {
            iac {
                terraform {
                    type = "aws_subnet"
                }
            }
        }

        resource "ec2.transit_gateways" {
            iac {
                terraform {
                    type = "aws_ec2_transit_gateway"
                }
            }
        }

        resource "ec2.vpc_endpoints" {
            iac {
                terraform {
                    type = "aws_vpc_endpoint"
                }
            }
        }

        resource "ec2.vpc_peering_connections" {
            iac {
                terraform {
                    type = "aws_vpc_peering_connection"
                }
            }
        }

        resource "ec2.vpcs" {
            iac {
                terraform {
                    type = "aws_vpc"
                }
            }
        }

        resource "ec2.vpn_gateways" {
            iac {
                terraform {
                    type = "aws_vpn_gateway"
                }
            }
        }

        resource "ecr.repositories" {
            identifiers = [ "name" ]

            iac {
                terraform {
                    type = "aws_ecr_repository"
                }
            }
        }

        resource "ecs.clusters" {
            identifiers = [ "arn" ]

            iac {
                terraform {
                    type = "aws_ecs_cluster"
                }
            }
        }

        resource "efs.filesystems" {
            iac {
                terraform {
                    type = "aws_efs_file_system"
                }
            }
        }

        resource "eks.clusters" {
            iac {
                terraform {
                    type = "aws_eks_cluster"
                }
            }
        }

        resource "elasticbeanstalk.environments" {
            iac {
                terraform {
                    type = "aws_elastic_beanstalk_environment"
                }
            }
        }

        resource "elasticsearch.domains" {
            identifiers = [ "arn" ]

            iac {
                terraform {
                    type = "aws_elasticsearch_domain"
                }
            }
        }

        # Unmatched: elbv1.load_balancers

        resource "elbv2.load_balancers" {
            iac {
                terraform {
                    type = "aws_lb"
                }
            }
        }

        resource "elbv2.target_groups" {
            iac {
                terraform {
                    type = "aws_lb_target_group"
                }
            }
        }

        resource "emr.clusters" {
            iac {
                terraform {
                    type = "aws_emr_cluster"
                }
            }
        }

        resource "fsx.backups" {
            iac {
                terraform {
                    type = "aws_fsx_backup"
                }
            }
        }

        # Unmatched: iam.accounts

        resource "iam.groups" {
            identifiers = [ "name" ]
            iac {
                terraform {
                    type = "aws_iam_group"
                }
            }
        }

        resource "iam.openid_connect_identity_providers" {
            iac {
                terraform {
                    type = "aws_iam_openid_connect_provider"
                }
            }
        }

        # Unmatched: iam.password_policies

        resource "iam.policies" {
            # TODO
            iac {
                terraform {
                    type = "aws_iam_policy"
                }
            }
        }

        resource "iam.roles" {
            # TODO
            identifiers = [ "name" ]
            iac {
                terraform {
                    type = "aws_iam_role"
                }
            }
        }

        resource "iam.saml_identity_providers" {
            iac {
                terraform {
                    type = "aws_iam_saml_provider"
                }
            }
        }

        resource "iam.server_certificates" {
            iac {
                terraform {
                    type = "aws_iam_server_certificate"
                }
            }
        }

        resource "iam.users" {
            identifiers       = ["user_name"]
#            ignore_attributes = ["id", "user_id", "password_last_used"]
            attributes = [ "arn", "path", "permissions_boundary_arn", "permissions_boundary_type", "tags" ]

            iac {
                terraform {
                    type = "aws_iam_user"
                }
            }
        }

        resource "aws_iam_user_groups" {
            identifiers = [ "group_name" ]

            iac {
                terraform {
                    type = "aws_iam_group"
                }
            }
        }

        # Unmatched: iam.virtual_mfa_devices

        resource "kms.keys" {
            # TODO
            identifiers = [ "key_id" ]
            iac {
                terraform {
                    type = "aws_kms_key"
                }
            }
        }

        resource "lambda.functions" {
            identifiers = [ "name" ]
            iac {
                terraform {
                    type = "aws_lambda_function"
                }
            }
        }

        # Unmatched: lambda.layers

        resource "mq.brokers" {
            iac {
                terraform {
                    type = "aws_mq_broker"
                }
            }
        }

        # Unmatched: organizations.accounts

        # Unmatched: rds.certificates

        resource "rds.clusters" {
            # TODO
            identifiers = [ "id" ]

            iac {
                terraform {
                    type = "aws_rds_cluster"
                }
            }
        }

        # Unmatched: rds.db_subnet_groups

        resource "rds.instances" {
            # TODO
            iac {
                terraform {
                    type = "aws_rds_cluster_instance"
                }
            }
        }

        resource "redshift.clusters" {
            iac {
                terraform {
                    type = "aws_redshift_cluster"
                }
            }
        }

        resource "redshift.subnet_groups" {
            iac {
                terraform {
                    type = "aws_redshift_subnet_group"
                }
            }
        }

        resource "route53.health_checks" {
            iac {
                terraform {
                    type = "aws_route53_health_check"
                }
            }
        }

        resource "route53.hosted_zones" {
            iac {
                terraform {
                    type = "aws_route53_zone"
                }
            }
        }

        # Unmatched: route53.reusable_delegation_sets

        # Unmatched: route53.traffic_policies

        resource "s3.buckets" {
            ignore_attributes = ["account_id", "name"]

            iac {
                terraform {
                    type = "aws_s3_bucket"
                }
            }
        }

        resource "sns.subscriptions" {
            iac {
                terraform {
                    type = "aws_sns_topic_subscription"
                }
            }
        }

        resource "sns.topics" {
            iac {
                terraform {
                    type = "aws_sns_topic"
                }
            }
        }

        resource "sqs.queues" {
            identifiers = [ "url" ]
            ignore_attributes = [ "policy", "redrive_policy" ] # string type in TF, json type in CQ
            iac {
                terraform {
                    type = "aws_sqs_queue"
                }
            }
        }

        resource "waf.rule_groups" {
            identifiers = [ "id" ]
            iac {
                terraform {
                    type = "aws_waf_rule_group"
                }
            }
        }

        resource "waf.rules" {
            identifiers = [ "id" ]
            iac {
                terraform {
                    type = "aws_waf_rule"
                }
            }
        }

        # Unmatched: waf.subscribed_rule_groups

        resource "waf.web_acls" {
            iac {
                terraform {
                    type = "aws_waf_web_acl"
                }
            }
        }

        # Unmatched: wafv2.managed_rule_groups

        resource "wafv2.rule_groups" {
            iac {
                terraform {
                    type = "aws_wafv2_rule_group"
                }
            }
        }

        resource "wafv2.web_acls" {
            iac {
                terraform {
                    type = "aws_wafv2_web_acl"
                }
            }
        }

        skip_resources = [
#            "ec2.instances",
#            "iam.users",
#            "s3.buckets"
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
