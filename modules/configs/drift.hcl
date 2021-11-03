module "drift" {

    provider "*" {
        # provider: the *provider.Provider
        # example:
        #  provider.Name is "aws"
        # special case:
        #  provider.ModuleHcl is the config provider supplies

        # resource: an entry in either provider.ResourceMap or provider.ResourceMap[].Relation
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
    }

    # TODO get from provider... But this could also override/decorate the * entry above, if specified
    provider "aws" {
        version = ">=0.6.0"

        resource "*" {
            ignore_identifiers = [ "account_id", "region", "user_cq_id", "api_cq_id", "api_integration_cq_id", "api_route_cq_id", "distribution_cq_id", "trail_cq_id", "alarm_cq_id", "filter_cq_id", "connection_cq_id", "directconnect_gateway_cq_id", "lag_cq_id" ]
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

        # Unmatched: apigateway.domain_names (no data in tests)

        resource "apigateway.rest_apis" {
            iac {
                terraform {
                    type = "aws_api_gateway_rest_api"
                }
            }
        }

        resource "aws_apigateway_rest_api_authorizers" {
            identifiers = [ "id" ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_authorizer"
                }
            }
        }

        resource "aws_apigateway_rest_api_deployments" {
            identifiers = [ "id" ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_deployment"
                }
            }
        }

        resource "aws_apigateway_rest_api_documentation_parts" {
            identifiers = [ sql("CONCAT(c.rest_api_id, '/', c.id)") ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_documentation_part"
                }
            }
        }

        resource "aws_apigateway_rest_api_documentation_versions" {
            identifiers = [ sql("CONCAT(c.rest_api_id, '/', c.version)") ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_documentation_version"
                }
            }
        }

        # TODO aws_apigateway_rest_api_gateway_responses (no PKs)

        resource "aws_apigateway_rest_api_models" {
            identifiers = [ "id" ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_model"
                }
            }
        }

        resource "aws_apigateway_rest_api_request_validators" {
            identifiers = [ "id" ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_request_validator"
                }
            }
        }

        # TODO aws_apigateway_rest_api_resources

        resource "aws_apigateway_rest_api_stages" {
            identifiers = [ sql("CONCAT('ags-',parent.id,'-',c.stage_name)") ]
            parent_match = "rest_api_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_stage"
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

        # TODO aws_apigateway_usage_plan_api_stages

        resource "aws_apigateway_usage_plan_keys" {
            identifiers = [ "id" ]
            parent_match = "usage_plan_cq_id"

            iac {
                terraform {
                    type = "aws_api_gateway_usage_plan_key"
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

        resource "aws_apigatewayv2_api_authorizers" {
            parent_match = "api_cq_id"

            iac {
                terraform {
                    type = "aws_apigatewayv2_authorizer"
                }
            }
        }

        resource "aws_apigatewayv2_api_deployments" {
            parent_match = "api_cq_id"

            iac {
                terraform {
                    type = "aws_apigatewayv2_deployment"
                }
            }
        }

        resource "aws_apigatewayv2_api_integrations" {
            parent_match = "api_cq_id"

            iac {
                terraform {
                    type = "aws_apigatewayv2_integration"
                }
            }
        }

        resource "aws_apigatewayv2_api_integration_responses" {
            parent_match = "api_integration_cq_id"

            iac {
                terraform {
                    type = "aws_apigatewayv2_integration_response"
                }
            }
        }

        resource "aws_apigatewayv2_api_models" {
          parent_match = "api_cq_id"

          iac {
            terraform {
              type = "aws_apigatewayv2_model"
            }
          }
        }

        resource "aws_apigatewayv2_api_routes" {
          parent_match = "api_cq_id"

          iac {
            terraform {
              type = "aws_apigatewayv2_route"
            }
          }
        }

        resource "aws_apigatewayv2_api_route_responses" {
          parent_match = "api_route_cq_id"

          iac {
            terraform {
              type = "aws_apigatewayv2_route_response"
            }
          }
        }

        resource "aws_apigatewayv2_api_stages" {
          parent_match = "api_cq_id"

          iac {
            terraform {
              type = "aws_apigatewayv2_stage"
            }
          }
        }

        # TODO: apigatewayv2.domain_names (no data in tests)

        # TODO: aws_apigatewayv2_domain_name_configurations (no data in tests)

        # TODO: aws_apigatewayv2_domain_name_rest_api_mappings (no data in tests)

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

       # Unmatched: aws_autoscaling_launch_configuration_block_device_mappings

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

        # TODO: aws_cloudfront_distribution_cache_behaviours (no data in tests)

        # TODO: aws_cache_behaviour_lambda_function_associations (no data in tests)

        # TODO: aws_cloudfront_distribution_custom_error_responses (no data in tests)

        resource "aws_cloudfront_distribution_origins" {
            identifiers = [ sql("SPLIT_PART(c.s3_origin_config_origin_access_identity,'/', 3)") ]
            parent_match = "distribution_cq_id"

            iac {
                terraform {
                    type = "aws_cloudfront_origin_access_identity"
                }
            }
        }

        # TODO: aws_cloudfront_distribution_alias_icp_recordals (no data in tests)

        # TODO: aws_cloudfront_distribution_origin_groups (no data in tests)

        resource "cloudtrail.trails" {
            identifiers = [ "name" ]

            iac {
                terraform {
                    type = "aws_cloudtrail"
                }
            }
        }

        # TODO: aws_cloudtrail_trail_event_selectors

        resource "cloudwatch.alarms" {
            iac {
                terraform {
                    type = "aws_cloudwatch_metric_alarm"
                }
            }
        }

        resource "aws_cloudwatch_alarm_metrics" {
            identifiers = [ "alarm_name" ]
            parent_match = "alarm_cq_id"

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

        # Unmatched: aws_cloudwatchlogs_filter_metric_transformations

        resource "cognito.identity_pools" {
            iac {
                terraform {
                    type = "aws_cognito_identity_pool"
                }
            }
        }

        # TODO: aws_cognito_identity_pool_cognito_identity_providers (no data in tests)

        resource "cognito.user_pools" {
            iac {
                terraform {
                    type = "aws_cognito_user_pool"
                }
            }
        }

        # Unmatched: aws_cognito_user_pool_schema_attributess

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

        # TODO aws_directconnect_connection_mac_sec_keys (no data in tests)

        resource "directconnect.gateways" {
            iac {
                terraform {
                    type = "aws_dx_gateway"
                }
            }
        }

        resource "aws_directconnect_gateway_associations" {
            identifiers = [ sql("CONCAT('ga-', c.directconnect_gateway_id, c.associated_gateway_id)") ]
            parent_match = "directconnect_gateway_cq_id"
            iac {
                terraform {
                    type = "aws_dx_gateway_association"
                }
            }
        }

        # TODO: aws_directconnect_gateway_attachments (no data in tests)

        resource "directconnect.lags" {
            iac {
                terraform {
                    type = "aws_dx_lag"
                }
            }
        }

        # TODO: aws_directconnect_lag_mac_sec_keys (no data in tests)

        # Unmatched: directconnect.virtual_gateways (aws_dx_gateway but IDs don't match)

        resource "directconnect.virtual_interfaces" {
            iac {
                terraform {
                    type = "aws_dx_public_virtual_interface"
                }
            }
        }

        # TODO: aws_directconnect_virtual_interface_bgp_peers (no data in tests)

        # Unmatched: ec2.byoip_cidrs (no data in tests)

        resource "ec2.customer_gateways" {
            iac {
                terraform {
                    type = "aws_customer_gateway"
                }
            }
        }

        resource "ec2.ebs_volumes" {
            iac {
                terraform {
                    type = "aws_ebs_volume"
                }
            }
        }

        # TODO aws_ec2_ebs_volume_attachments

        resource "ec2.flow_logs" {
            iac {
                terraform {
                    type = "aws_flow_log"
                }
            }
        }

        resource "ec2.images" {
            identifiers = [ sql("tags->>'Ec2ImageBuilderArn'") ]
            iac {
                terraform {
                    type = "aws_imagebuilder_image"
                }
            }
        }

        resource "ec2.instances" {
            ignore_attributes = ["launch_time"]

            iac {
                terraform {
                    type = "aws_instance"
                }
            }
        }

        resource "ec2.internet_gateways" {
            filters = [
                "NOT EXISTS (SELECT 1 FROM aws_ec2_internet_gateway_attachments a JOIN aws_ec2_vpcs v ON v.id=a.vpc_id WHERE a.internet_gateway_cq_id=c.cq_id AND v.is_default)",
                "NOT EXISTS (SELECT 1 FROM aws_ec2_route_table_routes WHERE gateway_id=c.id AND destination_cidr_block='0.0.0.0/0')"
            ]

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
            filters = [ "c.is_default!=true" ]

            iac {
                terraform {
                    type = "aws_network_acl"
                }
            }
        }

#        resource "aws_ec2_network_acl_entries" {
#            # TODO no CRC32 function, no data in tests to verify
#            identifiers = [ sql("CONCAT('nacl-',(CONCAT(parent.id,'-',c.rule_number,'-',CASE WHEN c.egress THEN 'true' ELSE 'false' END,'-',c.protocol,'-')))") ]
#            filters = [ "((c.cidr_block='0.0.0.0/0' AND c.rule_number=32767) OR (c.ipv6_cidr_block=':/0' AND c.rule_number=32768)) AND c.rule_action='deny' AND c.protocol='-1'" ]
#            parent_match = "network_acl_cq_id"
#
#            iac {
#                terraform {
#                    type = "aws_network_acl_rule"
#                }
#            }
#        }

        # Unmatched: ec2.regional_config

        resource "ec2.route_tables" {
            filters = [
                "NOT EXISTS (SELECT 1 FROM aws_ec2_route_table_associations WHERE route_table_cq_id=c.cq_id AND main)",
                "NOT EXISTS (SELECT 1 FROM aws_ec2_route_table_routes WHERE route_table_cq_id=c.cq_id AND origin='CreateRouteTable')"
            ]

            iac {
                terraform {
                    type = "aws_route_table"
                }
            }
        }

        resource "ec2.security_groups" {
            filters = [ "c.group_name!='default'" ]

            iac {
                terraform {
                    type = "aws_security_group"
                }
            }
        }

        resource "ec2.subnets" {
            filters = [ "c.default_for_az!=true" ]

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
            filters = [ "c.is_default!=true" ]

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

        resource "elbv1.load_balancers" {
            iac {
                terraform {
                    type = "aws_elb"
                }
            }
        }

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

        # Unmatched: iam.password_policies (no data in tests)

        resource "iam.policies" {
            identifiers = [ "arn" ]

            iac {
                terraform {
                    type = "aws_iam_policy"
                }
            }
        }

        resource "iam.roles" {
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
            attributes = [ "arn", "path", "permissions_boundary_arn", "permissions_boundary_type", "tags" ]

            iac {
                terraform {
                    type = "aws_iam_user"
                }
            }
        }

        resource "aws_iam_user_groups" {
            identifiers = [ "group_name" ]
            parent_match = "user_cq_id" # This is required until we can read resolver names from cqproto

            iac {
                terraform {
                    type = "aws_iam_group"
                }
            }
        }

        resource "aws_iam_user_access_keys" {
#            ignore_identifiers = [ "user_cq_id" ] # Ignored in provider level
            parent_match = "user_cq_id"

            iac {
                terraform {
                    type = "aws_iam_access_key"
                }
            }
        }

        resource "aws_iam_user_attached_policies" {
            identifiers = [ sql("CONCAT(parent.user_name, ':user_', c.policy_name)") ]
            parent_match = "user_cq_id"

            iac {
                terraform {
                    type = "aws_iam_user_policy"
                }
            }
        }

        resource "aws_iam_user_policies" {
            identifiers = [ sql("CONCAT(parent.user_name, ':', c.policy_name)") ]
            parent_match = "user_cq_id"

            iac {
                terraform {
                    type = "aws_iam_user_policy"
                }
            }
        }

        # Unmatched: iam.virtual_mfa_devices (no data in tests)

        resource "kms.keys" {
            identifiers = [ "id" ]
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

        resource "aws_lambda_layer_versions" {
            identifiers = [ sql("CONCAT(parent.arn, ':', c.version)") ]
            parent_match = "layer_cq_id"

            iac {
                terraform {
                    type = "aws_lambda_layer_version"
                }
            }
        }

        resource "mq.brokers" {
            iac {
                terraform {
                    type = "aws_mq_broker"
                }
            }
        }

        # Unmatched: organizations.accounts

        # Unmatched: rds.certificates (mode: data)

        resource "rds.clusters" {
            identifiers = [ "db_cluster_identifier" ]
            iac {
                terraform {
                    type = "aws_rds_cluster"
                }
            }
        }

        resource "rds.db_subnet_groups" {
            identifiers = [ "name" ]
            iac {
                terraform {
                    type = "aws_db_subnet_group"
                }
            }
        }

        resource "rds.instances" {
            identifiers = [ "db_name" ]

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

        resource "route53.reusable_delegation_sets" {
            identifiers = [ sql("SPLIT_PART(c.id, '/', 3)") ]
            iac {
                terraform {
                    type = "aws_route53_delegation_set"
                }
            }
        }

        # Unmatched: route53.traffic_policies (no data in tests)

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

        # Unmatched: waf.subscribed_rule_groups (no data in tests)

        resource "waf.web_acls" {
            iac {
                terraform {
                    type = "aws_waf_web_acl"
                }
            }
        }

        # Unmatched: wafv2.managed_rule_groups (aws_wafv2_web_acl but IDs don't match)

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
