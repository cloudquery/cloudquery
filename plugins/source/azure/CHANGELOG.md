# Changelog

All notable changes to this provider will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [4.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v4.2.0...plugins-source-azure-v4.3.0) (2023-02-28)


### Features

* **azure-resources:** Add AKS Upgrade Profile ([#8444](https://github.com/cloudquery/cloudquery/issues/8444)) ([a665f4b](https://github.com/cloudquery/cloudquery/commit/a665f4b626110838bc26dcf85522492ec41c7987))
* **azure-resources:** Add Management Groups ([#8226](https://github.com/cloudquery/cloudquery/issues/8226)) ([20a4c0b](https://github.com/cloudquery/cloudquery/commit/20a4c0b5f09d38a0eac2a629504cc01f33cb3af2))
* **azure-resources:** Add SQL Server Security Alert Policies ([#7939](https://github.com/cloudquery/cloudquery/issues/7939)) ([210cf81](https://github.com/cloudquery/cloudquery/commit/210cf811b94f474aa15bd6d4034ac7c0b428683a))
* **azure-spec:** Add cloud name configuration ([#8471](https://github.com/cloudquery/cloudquery/issues/8471)) ([d2ec0cc](https://github.com/cloudquery/cloudquery/commit/d2ec0cc531ed36ac21cadfda981e7e721d047bb9))


### Bug Fixes

* **azure-resources:** Make Monitor Diagnostic Settings a relation of generic resources ([#7943](https://github.com/cloudquery/cloudquery/issues/7943)) ([8316b87](https://github.com/cloudquery/cloudquery/commit/8316b87664d546c0da3cb3066dd39b1a573aa7a9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.0 ([#8344](https://github.com/cloudquery/cloudquery/issues/8344)) ([9c57544](https://github.com/cloudquery/cloudquery/commit/9c57544d06f9a774adcc659bcabd2518a905bdaa))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.39.1 ([#8371](https://github.com/cloudquery/cloudquery/issues/8371)) ([e3274c1](https://github.com/cloudquery/cloudquery/commit/e3274c109739bc107387627d340a713470c3a3c1))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.40.0 ([#8401](https://github.com/cloudquery/cloudquery/issues/8401)) ([4cf36d6](https://github.com/cloudquery/cloudquery/commit/4cf36d68684f37c0407332930766c1ba60807a93))

## [4.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v4.1.1...plugins-source-azure-v4.2.0) (2023-02-21)


### Features

* **azure:** Add network ExpressRoute circuit authorizations and peerings ([#8128](https://github.com/cloudquery/cloudquery/issues/8128)) ([2d4cba5](https://github.com/cloudquery/cloudquery/commit/2d4cba5dd34b4157d59924c7b4ff6f959c56305d)), closes [#7927](https://github.com/cloudquery/cloudquery/issues/7927)
* **azure:** Add network: interface_ip_configurations and virtual_network_subnets ([#8126](https://github.com/cloudquery/cloudquery/issues/8126)) ([df5e48b](https://github.com/cloudquery/cloudquery/commit/df5e48b5fde4db098c923019ef77cd6e370e224b)), closes [#7929](https://github.com/cloudquery/cloudquery/issues/7929)
* **azure:** Add postgresql databases resource ([#8125](https://github.com/cloudquery/cloudquery/issues/8125)) ([91cab61](https://github.com/cloudquery/cloudquery/commit/91cab6176826de858fda3807aae4b3a2172e3a47)), closes [#7928](https://github.com/cloudquery/cloudquery/issues/7928)


### Bug Fixes

* **azure:** Ensure spec subscriptions are unique ([#8099](https://github.com/cloudquery/cloudquery/issues/8099)) ([20dc235](https://github.com/cloudquery/cloudquery/commit/20dc235b998f0c31214f4ca1b3a6d366552f5683))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.38.2 ([#8156](https://github.com/cloudquery/cloudquery/issues/8156)) ([ac2d2d7](https://github.com/cloudquery/cloudquery/commit/ac2d2d70d5c4bc45fb8734bd4deb8a1e36074f6d))
* **deps:** Update module golang.org/x/net to v0.7.0 [SECURITY] ([#8176](https://github.com/cloudquery/cloudquery/issues/8176)) ([fc4cef8](https://github.com/cloudquery/cloudquery/commit/fc4cef86dce4ca76ca8397e897ab744e48975834))

## [4.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v4.1.0...plugins-source-azure-v4.1.1) (2023-02-14)


### Bug Fixes

* **azure:** Fix detecting logic for CIS Azure Benchmark v1.3.0 - 2.12 ([#7807](https://github.com/cloudquery/cloudquery/issues/7807)) ([56b7ee2](https://github.com/cloudquery/cloudquery/commit/56b7ee2736e1fb267098eba4dd13bd27301cf7c3))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.0 ([#7809](https://github.com/cloudquery/cloudquery/issues/7809)) ([c85a9cb](https://github.com/cloudquery/cloudquery/commit/c85a9cb697477520e94a1fd260c56b89da62fc87))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.36.1 ([#7930](https://github.com/cloudquery/cloudquery/issues/7930)) ([39dccc1](https://github.com/cloudquery/cloudquery/commit/39dccc1bf81f4eb02d181ba0c47b37038a4c5455))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.0 ([#7933](https://github.com/cloudquery/cloudquery/issues/7933)) ([dc9cffb](https://github.com/cloudquery/cloudquery/commit/dc9cffbf37bbc6fae73a20bf47e6bbf17e74d1f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.37.1 ([#8008](https://github.com/cloudquery/cloudquery/issues/8008)) ([c47aac0](https://github.com/cloudquery/cloudquery/commit/c47aac0b5e3190a04299713651b97e360043911f))

## [4.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v4.0.0...plugins-source-azure-v4.1.0) (2023-02-07)


### Features

* **azure-resources:** Add Virtual Network Gateways and Connections ([#7636](https://github.com/cloudquery/cloudquery/issues/7636)) ([97a9254](https://github.com/cloudquery/cloudquery/commit/97a9254ba557d2920c5780c9d5c1bd527ab00571))


### Bug Fixes

* **azure-resources:** Use default API version of `armhealthbot` bots client ([#7177](https://github.com/cloudquery/cloudquery/issues/7177)) ([e24af94](https://github.com/cloudquery/cloudquery/commit/e24af94dd8b10056a1ef7a48111236d903dcb025))
* **azure:** Fix Azure CIS Policies in Section 2 ([#7718](https://github.com/cloudquery/cloudquery/issues/7718)) ([3c77b2c](https://github.com/cloudquery/cloudquery/commit/3c77b2c68ba0b8eb2d13e5a3d421d3f9874cbb1e))
* **deps:** Update golang.org/x/exp digest to f062dba ([#7531](https://github.com/cloudquery/cloudquery/issues/7531)) ([59d5575](https://github.com/cloudquery/cloudquery/commit/59d55758b0951553b8d246d1e78b4e3917ff1976))
* **deps:** Update google.golang.org/genproto digest to 1c01626 ([#7533](https://github.com/cloudquery/cloudquery/issues/7533)) ([c549c27](https://github.com/cloudquery/cloudquery/commit/c549c275077f1cdfb9df0b3f3c129cbf0b150552))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azidentity to v1.2.1 ([#7540](https://github.com/cloudquery/cloudquery/issues/7540)) ([3b5c838](https://github.com/cloudquery/cloudquery/commit/3b5c83832064d729ad1097728f7d12aedbbb9400))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.0 ([#7595](https://github.com/cloudquery/cloudquery/issues/7595)) ([c5adc75](https://github.com/cloudquery/cloudquery/commit/c5adc750d4b0242563997c04c582f8da27913095))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.33.1 ([#7614](https://github.com/cloudquery/cloudquery/issues/7614)) ([2fe665c](https://github.com/cloudquery/cloudquery/commit/2fe665cdd80d88c5699bb203bd7accd604dfba99))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.34.0 ([#7719](https://github.com/cloudquery/cloudquery/issues/7719)) ([6a33085](https://github.com/cloudquery/cloudquery/commit/6a33085c75adcf2387f7bbb5aa4f7a84ce7e2957))
* **deps:** Update module github.com/golang-jwt/jwt/v4 to v4.4.3 ([#7543](https://github.com/cloudquery/cloudquery/issues/7543)) ([0607454](https://github.com/cloudquery/cloudquery/commit/060745428eda5839be801c153c2f7261fcc54abd))

## [4.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.5.0...plugins-source-azure-v4.0.0) (2023-01-31)


### ⚠ BREAKING CHANGES

* **azure:** Drop `subscription_id` from `azure_reservations_reservation` ([#7345](https://github.com/cloudquery/cloudquery/issues/7345))
* **azure:** Drop `subscription_id` from `azure_support_services` ([#7351](https://github.com/cloudquery/cloudquery/issues/7351))
* **azure:** Drop `subscription_id` from `azure_policy_data_policy_manifests` ([#7344](https://github.com/cloudquery/cloudquery/issues/7344))
* **azure:** Drop `subscription_id` from `azure_eventgrid_topic_types` ([#7343](https://github.com/cloudquery/cloudquery/issues/7343))
* **azure:** Drop `subscription_id` from `azure_cdn_edge_nodes` ([#7342](https://github.com/cloudquery/cloudquery/issues/7342))
* **azure:** Drop `subscription_id` from `azure_authorization_provider_operations_metadata` ([#7341](https://github.com/cloudquery/cloudquery/issues/7341))
* **azure:** Drop `subscription_id` column from `azure_advisor_recommendation_metadata` ([#7337](https://github.com/cloudquery/cloudquery/issues/7337))
* **azure:** Change `azure_compute_skus` PK to `(subscription_id, name, _sku_hash)` ([#7305](https://github.com/cloudquery/cloudquery/issues/7305))
* **azure:** Add `subscription_id` to `azure_security_secure_score_control_definitions` PK ([#7285](https://github.com/cloudquery/cloudquery/issues/7285))
* **azure:** Add `subscription_id` to `azure_security_topology` PK ([#7280](https://github.com/cloudquery/cloudquery/issues/7280))
* **azure:** Add `subscription_id` to `azure_policy_set_definitions` PK ([#7283](https://github.com/cloudquery/cloudquery/issues/7283))
* **azure:** Add `subscription_id` to `azure_security_assessments_metadata` PK ([#7281](https://github.com/cloudquery/cloudquery/issues/7281))
* **azure:** Change `azure_compute_skus` PK from `id` to `(subscription_id, family, kind, name)` ([#7267](https://github.com/cloudquery/cloudquery/issues/7267))
* **azure:** Add `subscription_id` to `azure_network_express_route_service_providers` PK ([#7279](https://github.com/cloudquery/cloudquery/issues/7279))
* **azure:** Add `subscription_id` to `azure_network_bgp_service_communities` PK ([#7277](https://github.com/cloudquery/cloudquery/issues/7277))
* **azure:** Add `subscription_id` to `azure_policy_definitions` PK ([#7264](https://github.com/cloudquery/cloudquery/issues/7264))
* **azure:** Add `subscription_id` to `azure_network_azure_firewall_fqdn_tags` PK ([#7278](https://github.com/cloudquery/cloudquery/issues/7278))
* **azure:** Add `subscription_id` to `azure_frontdoor_managed_rule_sets` PK ([#7276](https://github.com/cloudquery/cloudquery/issues/7276))
* **azure:** Add `subscription_id` to `azure_authorization_role_assignments` ([#7270](https://github.com/cloudquery/cloudquery/issues/7270))
* **azure:** Add `subscription_id` to `azure_cdn_managed_rule_sets` PK ([#7272](https://github.com/cloudquery/cloudquery/issues/7272))

### Features

* Azure policy implementaion for v2  ([#6557](https://github.com/cloudquery/cloudquery/issues/6557)) ([0768fe6](https://github.com/cloudquery/cloudquery/commit/0768fe60441ad1e570cc1cf8a6373405030f84b9))
* Azure resources for policy implementation v2 ([#6677](https://github.com/cloudquery/cloudquery/issues/6677)) ([581eb7d](https://github.com/cloudquery/cloudquery/commit/581eb7d60acf1450b6ccd8bae602a95093cb319b))


### Bug Fixes

* **azure:** Add `subscription_id` to `azure_authorization_role_assignments` ([#7270](https://github.com/cloudquery/cloudquery/issues/7270)) ([2d7492d](https://github.com/cloudquery/cloudquery/commit/2d7492d1a6e9e53f289656a7ccb9a095f6d69caa)), closes [#7242](https://github.com/cloudquery/cloudquery/issues/7242)
* **azure:** Add `subscription_id` to `azure_cdn_managed_rule_sets` PK ([#7272](https://github.com/cloudquery/cloudquery/issues/7272)) ([cc9b5af](https://github.com/cloudquery/cloudquery/commit/cc9b5afb6debf3dade83ca760df6ac29719bd390)), closes [#7244](https://github.com/cloudquery/cloudquery/issues/7244)
* **azure:** Add `subscription_id` to `azure_frontdoor_managed_rule_sets` PK ([#7276](https://github.com/cloudquery/cloudquery/issues/7276)) ([830d669](https://github.com/cloudquery/cloudquery/commit/830d669128099e2eb2235bf09d79c5acd66b3d62)), closes [#7247](https://github.com/cloudquery/cloudquery/issues/7247)
* **azure:** Add `subscription_id` to `azure_network_azure_firewall_fqdn_tags` PK ([#7278](https://github.com/cloudquery/cloudquery/issues/7278)) ([2b31fa4](https://github.com/cloudquery/cloudquery/commit/2b31fa471039b3896e6be62f3eb93136a0f42982)), closes [#7248](https://github.com/cloudquery/cloudquery/issues/7248)
* **azure:** Add `subscription_id` to `azure_network_bgp_service_communities` PK ([#7277](https://github.com/cloudquery/cloudquery/issues/7277)) ([50088a0](https://github.com/cloudquery/cloudquery/commit/50088a0ca818b376a07bc68b5f179fa2bb0e58e8)), closes [#7249](https://github.com/cloudquery/cloudquery/issues/7249)
* **azure:** Add `subscription_id` to `azure_network_express_route_service_providers` PK ([#7279](https://github.com/cloudquery/cloudquery/issues/7279)) ([ea5b468](https://github.com/cloudquery/cloudquery/commit/ea5b468506357a4309c030d7c433851c760b093b)), closes [#7250](https://github.com/cloudquery/cloudquery/issues/7250)
* **azure:** Add `subscription_id` to `azure_policy_definitions` PK ([#7264](https://github.com/cloudquery/cloudquery/issues/7264)) ([809ebbd](https://github.com/cloudquery/cloudquery/commit/809ebbd28d612f2d411f64db2a54975c04255dbb)), closes [#7252](https://github.com/cloudquery/cloudquery/issues/7252)
* **azure:** Add `subscription_id` to `azure_policy_set_definitions` PK ([#7283](https://github.com/cloudquery/cloudquery/issues/7283)) ([96ba863](https://github.com/cloudquery/cloudquery/commit/96ba863da220694abe3e05ccded39ad37d9f223c)), closes [#7253](https://github.com/cloudquery/cloudquery/issues/7253)
* **azure:** Add `subscription_id` to `azure_security_assessments_metadata` PK ([#7281](https://github.com/cloudquery/cloudquery/issues/7281)) ([8b0f3ed](https://github.com/cloudquery/cloudquery/commit/8b0f3ed3327a73cca3762804cff22deab395faaf)), closes [#7255](https://github.com/cloudquery/cloudquery/issues/7255)
* **azure:** Add `subscription_id` to `azure_security_secure_score_control_definitions` PK ([#7285](https://github.com/cloudquery/cloudquery/issues/7285)) ([1344bd3](https://github.com/cloudquery/cloudquery/commit/1344bd3b917639aaf28eb170a28ed106f54af7d2)), closes [#7256](https://github.com/cloudquery/cloudquery/issues/7256)
* **azure:** Add `subscription_id` to `azure_security_topology` PK ([#7280](https://github.com/cloudquery/cloudquery/issues/7280)) ([3e06103](https://github.com/cloudquery/cloudquery/commit/3e06103521cd8cb2057c637fb0102d93ef4f434f)), closes [#7257](https://github.com/cloudquery/cloudquery/issues/7257)
* **azure:** Change `azure_compute_skus` PK from `id` to `(subscription_id, family, kind, name)` ([#7267](https://github.com/cloudquery/cloudquery/issues/7267)) ([f0fd5fb](https://github.com/cloudquery/cloudquery/commit/f0fd5fb686952f5dbada3e7a7845d090b24a8122)), closes [#7245](https://github.com/cloudquery/cloudquery/issues/7245)
* **azure:** Change `azure_compute_skus` PK to `(subscription_id, name, _sku_hash)` ([#7305](https://github.com/cloudquery/cloudquery/issues/7305)) ([e8e049c](https://github.com/cloudquery/cloudquery/commit/e8e049c9e2aa2f5eee2c5755532bbdeb6b880db9)), closes [#7245](https://github.com/cloudquery/cloudquery/issues/7245)
* **azure:** Drop `subscription_id` column from `azure_advisor_recommendation_metadata` ([#7337](https://github.com/cloudquery/cloudquery/issues/7337)) ([e6d8f14](https://github.com/cloudquery/cloudquery/commit/e6d8f149120fc0f49696cc8f4d4bd5b0ffdd9641))
* **azure:** Drop `subscription_id` from `azure_authorization_provider_operations_metadata` ([#7341](https://github.com/cloudquery/cloudquery/issues/7341)) ([4d84e8e](https://github.com/cloudquery/cloudquery/commit/4d84e8e916d084a8af7ec1e48c90efd717bf481f))
* **azure:** Drop `subscription_id` from `azure_cdn_edge_nodes` ([#7342](https://github.com/cloudquery/cloudquery/issues/7342)) ([0616f9b](https://github.com/cloudquery/cloudquery/commit/0616f9b0afad52d2d6785df4f0abc1453883ddbf))
* **azure:** Drop `subscription_id` from `azure_eventgrid_topic_types` ([#7343](https://github.com/cloudquery/cloudquery/issues/7343)) ([c603801](https://github.com/cloudquery/cloudquery/commit/c603801ffc59e8e05e0a25cae4d65750b105697a))
* **azure:** Drop `subscription_id` from `azure_policy_data_policy_manifests` ([#7344](https://github.com/cloudquery/cloudquery/issues/7344)) ([31a6040](https://github.com/cloudquery/cloudquery/commit/31a604062fff5637373337310fc872714f9856ea))
* **azure:** Drop `subscription_id` from `azure_reservations_reservation` ([#7345](https://github.com/cloudquery/cloudquery/issues/7345)) ([daf7249](https://github.com/cloudquery/cloudquery/commit/daf7249ff0335df9dd973c6e62a65bf68cdce653))
* **azure:** Drop `subscription_id` from `azure_support_services` ([#7351](https://github.com/cloudquery/cloudquery/issues/7351)) ([2a91922](https://github.com/cloudquery/cloudquery/commit/2a919222e0af91d631a0c7ed5c87748cba9f2c53))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.29.0 ([#7121](https://github.com/cloudquery/cloudquery/issues/7121)) ([b7441c9](https://github.com/cloudquery/cloudquery/commit/b7441c93c274ae3a6009474a2b28f44a172dd6dc))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.30.0 ([#7222](https://github.com/cloudquery/cloudquery/issues/7222)) ([73ca21c](https://github.com/cloudquery/cloudquery/commit/73ca21c4259545f7e949c9d780d8184db475d2ac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.31.0 ([#7228](https://github.com/cloudquery/cloudquery/issues/7228)) ([36e8549](https://github.com/cloudquery/cloudquery/commit/36e8549f722658d909865723630fad1b2821db62))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.32.0 ([#7334](https://github.com/cloudquery/cloudquery/issues/7334)) ([b684122](https://github.com/cloudquery/cloudquery/commit/b68412222219f9ca160c0753290709d52de7fcd6))

## [3.5.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.4.0...plugins-source-azure-v3.5.0) (2023-01-24)


### Features

* **docs:** Add descriptions to tables `a-b` ([cf8de61](https://github.com/cloudquery/cloudquery/commit/cf8de61f40733ba1bcafde0100941c9efcf14b81))
* **docs:** Add descriptions to tables `c` ([#6932](https://github.com/cloudquery/cloudquery/issues/6932)) ([3152fcb](https://github.com/cloudquery/cloudquery/commit/3152fcb4097b9ac2dbcc8070030a194f793dbf74))
* **docs:** Add descriptions to tables `d-m` ([#6934](https://github.com/cloudquery/cloudquery/issues/6934)) ([93b41c2](https://github.com/cloudquery/cloudquery/commit/93b41c2bc1d991c42882e0e4d9fc21ba0e813e8e))
* **docs:** Add descriptions to tables `n-w` ([#6936](https://github.com/cloudquery/cloudquery/issues/6936)) ([d473aae](https://github.com/cloudquery/cloudquery/commit/d473aae7bcf5ad55d74ae22476185941b56413ae))


### Bug Fixes

* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.3.0 ([#6955](https://github.com/cloudquery/cloudquery/issues/6955)) ([66bfdf3](https://github.com/cloudquery/cloudquery/commit/66bfdf30343f6d8fc1b8cac9675631095ebfc01e))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice to v2 ([#6937](https://github.com/cloudquery/cloudquery/issues/6937)) ([b5245cf](https://github.com/cloudquery/cloudquery/commit/b5245cf4c28f6442c67b9c8d275d4d5c1afd806b))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization to v2 ([#6938](https://github.com/cloudquery/cloudquery/issues/6938)) ([8e96e67](https://github.com/cloudquery/cloudquery/commit/8e96e6750046c1d9a764e38f513f85b70989f128))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice to v1 ([#6939](https://github.com/cloudquery/cloudquery/issues/6939)) ([7c2ab93](https://github.com/cloudquery/cloudquery/commit/7c2ab936875d66f4093a08b228b59e006f6d2ac3))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute to v4 ([#6554](https://github.com/cloudquery/cloudquery/issues/6554)) ([b6c2936](https://github.com/cloudquery/cloudquery/commit/b6c29362681d7ddd3f33aac79f40ae7ef0fb1192))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance to v2 ([#6940](https://github.com/cloudquery/cloudquery/issues/6940)) ([726d86e](https://github.com/cloudquery/cloudquery/commit/726d86ec2fa53bf8620430b08c9efd4af0488594))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice to v2 ([#6941](https://github.com/cloudquery/cloudquery/issues/6941)) ([3310640](https://github.com/cloudquery/cloudquery/commit/331064075142d23ea63cf9857417b99d2e743e85))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos to v2 ([#6942](https://github.com/cloudquery/cloudquery/issues/6942)) ([99bf26b](https://github.com/cloudquery/cloudquery/commit/99bf26bda157da2c874fe07be170ac85927c530e))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory to v3 ([#6943](https://github.com/cloudquery/cloudquery/issues/6943)) ([70bba7c](https://github.com/cloudquery/cloudquery/commit/70bba7c49df0739ae90d690fd138f18bd7294ca6))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid to v2 ([#6944](https://github.com/cloudquery/cloudquery/issues/6944)) ([01ba430](https://github.com/cloudquery/cloudquery/commit/01ba430f5773de967bc591f27e349326f5312dbe))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork to v2 ([#6945](https://github.com/cloudquery/cloudquery/issues/6945)) ([aabf093](https://github.com/cloudquery/cloudquery/commit/aabf09369fd56023420c422cce101fdaf636a870))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction to v2 ([#6946](https://github.com/cloudquery/cloudquery/issues/6946)) ([2ba9008](https://github.com/cloudquery/cloudquery/commit/2ba90088c0519d3fdc193f61bbf81f0798fcaee3))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers to v2 ([#6947](https://github.com/cloudquery/cloudquery/issues/6947)) ([4a35095](https://github.com/cloudquery/cloudquery/commit/4a35095d9c7063f5c36ce7ff00289058d3da54b9))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis to v2 ([#6948](https://github.com/cloudquery/cloudquery/issues/6948)) ([d9f0d02](https://github.com/cloudquery/cloudquery/commit/d9f0d020a8dc9ea984b1b51e0f9b00cff845b997))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache to v2 ([#6949](https://github.com/cloudquery/cloudquery/issues/6949)) ([869021e](https://github.com/cloudquery/cloudquery/commit/869021e8283d7abde7776a0e85c2cd1c7a0b5706))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.28.0 ([#7009](https://github.com/cloudquery/cloudquery/issues/7009)) ([12ac005](https://github.com/cloudquery/cloudquery/commit/12ac005428a355d06a5939fbe06a82d49533e662))
* Rules for azure ([#7056](https://github.com/cloudquery/cloudquery/issues/7056)) ([5999548](https://github.com/cloudquery/cloudquery/commit/5999548516182019c4bff62ec2fe1a188ef10605))

## [3.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.3.1...plugins-source-azure-v3.4.0) (2023-01-17)


### Features

* Azure sql and postgresql resoruces for policies ([#6374](https://github.com/cloudquery/cloudquery/issues/6374)) ([97cbc2d](https://github.com/cloudquery/cloudquery/commit/97cbc2d733a148a12ec58ff40a35ac9064ed89f4))
* **azure-resources:** Add missing security tables ([#6905](https://github.com/cloudquery/cloudquery/issues/6905)) ([8a64414](https://github.com/cloudquery/cloudquery/commit/8a6441459039be190491de753f9ad23c8f07a3d4))


### Bug Fixes

* **azure:** Duplicate entries for `azure_subscription_subscriptions` ([#6887](https://github.com/cloudquery/cloudquery/issues/6887)) ([73cf12a](https://github.com/cloudquery/cloudquery/commit/73cf12a9821bef5a6c731b45bdf2fc5ba64ac5a8))

## [3.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.3.0...plugins-source-azure-v3.3.1) (2023-01-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.0 ([#6745](https://github.com/cloudquery/cloudquery/issues/6745)) ([9c41854](https://github.com/cloudquery/cloudquery/commit/9c418547c3bbff97449765e337182230fb5e40d5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.25.1 ([#6805](https://github.com/cloudquery/cloudquery/issues/6805)) ([9da0ce2](https://github.com/cloudquery/cloudquery/commit/9da0ce283f50410eb9274375ec1d22131a80d937))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.26.0 ([#6839](https://github.com/cloudquery/cloudquery/issues/6839)) ([6ccda8d](https://github.com/cloudquery/cloudquery/commit/6ccda8d0bc6e7ce75f4a64a18911e349ccaac277))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.27.0 ([#6856](https://github.com/cloudquery/cloudquery/issues/6856)) ([545799b](https://github.com/cloudquery/cloudquery/commit/545799bb0481087e187b5f27c88f5dde9c99f2f0))

## [3.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.2.0...plugins-source-azure-v3.3.0) (2023-01-11)


### Features

* **azure-resources:** Add Virtual Machine Scale Set VMs ([#6718](https://github.com/cloudquery/cloudquery/issues/6718)) ([00a096f](https://github.com/cloudquery/cloudquery/commit/00a096f1fde3a91ff3030473b5c80047140cd232))
* Log info message when namespace not registered for subscription ([#6705](https://github.com/cloudquery/cloudquery/issues/6705)) ([01a1e4c](https://github.com/cloudquery/cloudquery/commit/01a1e4c8d348c4411275b517790351aeaf8a645a))


### Bug Fixes

* **azure-resources:** Make `subscription_id, name` PKs of `azure_peering_service_providers` instead of `id` ([#6694](https://github.com/cloudquery/cloudquery/issues/6694)) ([d327465](https://github.com/cloudquery/cloudquery/commit/d327465a14defbde26356372b06348d2f5d38ddc))
* **azure-resources:** Pass `subscriptions/&lt;id&gt;` when listing role definitions ([#6669](https://github.com/cloudquery/cloudquery/issues/6669)) ([2983a4d](https://github.com/cloudquery/cloudquery/commit/2983a4d74affe0409a5053f330bf8059da4c4ca8))
* **azure-resources:** Remove non working `azure_windowsesu_multiple_activation_keys` ([#6680](https://github.com/cloudquery/cloudquery/issues/6680)) ([1fc3583](https://github.com/cloudquery/cloudquery/commit/1fc3583473d537f5bb31128f198b716186b889aa))
* **azure-resources:** Set APIVersion to "2019-05-10-preview" when fetching SQL Server Registrations ([#6672](https://github.com/cloudquery/cloudquery/issues/6672)) ([42942e0](https://github.com/cloudquery/cloudquery/commit/42942e03bbb59c7ca28e88c6b8156531fe656aae))
* **azure-resources:** Set APIVersion to "2022-08-08" when fetching Healthbots ([#6673](https://github.com/cloudquery/cloudquery/issues/6673)) ([ff79a3d](https://github.com/cloudquery/cloudquery/commit/ff79a3d2e28d574717780b85c584a28d9e6f7005))
* **azure-resources:** Update `armnginx` to v2 ([#6682](https://github.com/cloudquery/cloudquery/issues/6682)) ([aef47b5](https://github.com/cloudquery/cloudquery/commit/aef47b5e74996ac5d8242139a4b985244a2a4476))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.2 ([#6695](https://github.com/cloudquery/cloudquery/issues/6695)) ([694ab9f](https://github.com/cloudquery/cloudquery/commit/694ab9f3e20473146e3620d7b03bb17eb259d697))

## [3.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.1.0...plugins-source-azure-v3.2.0) (2023-01-10)


### Features

* Move azure to avoid codegen. ([#6276](https://github.com/cloudquery/cloudquery/issues/6276)) ([a0a8f0a](https://github.com/cloudquery/cloudquery/commit/a0a8f0aed52b9ae49b66ac564fbd4f8c4430c3b8))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.2 ([#6260](https://github.com/cloudquery/cloudquery/issues/6260)) ([805972a](https://github.com/cloudquery/cloudquery/commit/805972aa67ce54e3358501c6b7ee5d85e5f65cac))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.19.0 ([#6363](https://github.com/cloudquery/cloudquery/issues/6363)) ([ae6967c](https://github.com/cloudquery/cloudquery/commit/ae6967c22002c554a083f444eb611ac3e6d2698f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.20.0 ([#6376](https://github.com/cloudquery/cloudquery/issues/6376)) ([d6187ec](https://github.com/cloudquery/cloudquery/commit/d6187ec584f13be4fe9362dd393385b19d386113))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.21.0 ([#6382](https://github.com/cloudquery/cloudquery/issues/6382)) ([5baea40](https://github.com/cloudquery/cloudquery/commit/5baea40d2aec4e807db839c928be2e037d572bef))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.22.0 ([#6516](https://github.com/cloudquery/cloudquery/issues/6516)) ([b7e4e73](https://github.com/cloudquery/cloudquery/commit/b7e4e737a5f4d8f254960426ea8ba555d8f9b944))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.23.0 ([#6522](https://github.com/cloudquery/cloudquery/issues/6522)) ([ce24f1d](https://github.com/cloudquery/cloudquery/commit/ce24f1d64394cbb5ab07dcaa4af66c53f77f700f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.24.1 ([#6553](https://github.com/cloudquery/cloudquery/issues/6553)) ([392b848](https://github.com/cloudquery/cloudquery/commit/392b848b3124f9cf28f6234fdb9a43d671069879))

## [3.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v3.0.0...plugins-source-azure-v3.1.0) (2023-01-03)


### Features

* Add azure security pricings resource ([#6023](https://github.com/cloudquery/cloudquery/issues/6023)) ([9648f26](https://github.com/cloudquery/cloudquery/commit/9648f2690cd070803b5c5aa33683ceab5680b358))
* **azure:** Add storage_blob_services ([#6245](https://github.com/cloudquery/cloudquery/issues/6245)) ([0620ccd](https://github.com/cloudquery/cloudquery/commit/0620ccd6059818a9e560fd6a8657cb84e7da0ae3))
* **azure:** Cost management view queries, query views by subscription scope ([#5898](https://github.com/cloudquery/cloudquery/issues/5898)) ([ace315f](https://github.com/cloudquery/cloudquery/commit/ace315ff62818449889ab7d2f927dc06b4c7f045))


### Bug Fixes

* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/internal to v1.1.2 ([#6205](https://github.com/cloudquery/cloudquery/issues/6205)) ([154fa6f](https://github.com/cloudquery/cloudquery/commit/154fa6fccf41278ac395a1c8287634c4d65926d1))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift to v1.1.0 ([#6183](https://github.com/cloudquery/cloudquery/issues/6183)) ([3395c09](https://github.com/cloudquery/cloudquery/commit/3395c0920d5425c74f3dbf3ad4efeab6e6a6cda4))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage to v1.2.0 ([#6184](https://github.com/cloudquery/cloudquery/issues/6184)) ([adc96f3](https://github.com/cloudquery/cloudquery/commit/adc96f37379ca480dc5599c882415b8a638412ab))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.14.0 ([#6025](https://github.com/cloudquery/cloudquery/issues/6025)) ([35b2cfc](https://github.com/cloudquery/cloudquery/commit/35b2cfc7fc7bcdaceb7ee674e3a17f0f5673b366))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.0 ([#6071](https://github.com/cloudquery/cloudquery/issues/6071)) ([684b525](https://github.com/cloudquery/cloudquery/commit/684b525aaa285fcae70dd87af56679c1205adebe))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.15.1 ([#6079](https://github.com/cloudquery/cloudquery/issues/6079)) ([650659c](https://github.com/cloudquery/cloudquery/commit/650659c3c6766df571868e2ec3a2007cb76696eb))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.0 ([#6098](https://github.com/cloudquery/cloudquery/issues/6098)) ([7bacdf3](https://github.com/cloudquery/cloudquery/commit/7bacdf3364716eab08fa1a84ae4047b42edeee7e))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.16.1 ([#6214](https://github.com/cloudquery/cloudquery/issues/6214)) ([53b2415](https://github.com/cloudquery/cloudquery/commit/53b241508d7511d4b5fa74cc4262d180c1e6df66))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.17.0 ([#6256](https://github.com/cloudquery/cloudquery/issues/6256)) ([b19f6cd](https://github.com/cloudquery/cloudquery/commit/b19f6cd8e2c39994aeb19d78e78e927d6c3cf580))

## [3.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v2.2.0...plugins-source-azure-v3.0.0) (2022-12-27)


### ⚠ BREAKING CHANGES

* **azure:** Rename e_tag to etag ([#5902](https://github.com/cloudquery/cloudquery/issues/5902))

### Features

* Add instance_view to azure_compute_virtual_machines table ([#5941](https://github.com/cloudquery/cloudquery/issues/5941)) ([d65d80f](https://github.com/cloudquery/cloudquery/commit/d65d80fc24bf2ae1fcda502da5de744f86491afd))
* **azure:** Add search,logic,monitor,redis tables ([#5952](https://github.com/cloudquery/cloudquery/issues/5952)) ([9e16822](https://github.com/cloudquery/cloudquery/commit/9e168220a0e9b4883f6c182b5e4cc4b02f483801))


### Bug Fixes

* **azure:** Rename e_tag to etag ([#5902](https://github.com/cloudquery/cloudquery/issues/5902)) ([3c9daf5](https://github.com/cloudquery/cloudquery/commit/3c9daf58fccfc76f17aa1d49c04bfe935eb334f9))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.13.1 ([#5897](https://github.com/cloudquery/cloudquery/issues/5897)) ([ad15915](https://github.com/cloudquery/cloudquery/commit/ad15915f2951a75729859f6f1377ed789f8ba115))
* Update Azure auto provisioning queries ([#5888](https://github.com/cloudquery/cloudquery/issues/5888)) ([cf838ed](https://github.com/cloudquery/cloudquery/commit/cf838ed2e2a7bc2f356dfeb8fbf728241795e2cd))
* Update Azure Storage Queries ([#5908](https://github.com/cloudquery/cloudquery/issues/5908)) ([38690c6](https://github.com/cloudquery/cloudquery/commit/38690c617ddfb3b85b9e2ccb1814d407a7556e6c))

## [2.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v2.1.1...plugins-source-azure-v2.2.0) (2022-12-20)


### Features

* **azure:** Add storage_containers ([#5759](https://github.com/cloudquery/cloudquery/issues/5759)) ([18003e9](https://github.com/cloudquery/cloudquery/commit/18003e9cf6ac036a33e88e6cdcd9b626a792c7de))


### Bug Fixes

* **azure:** Remove extra `fmt.Println` ([#5756](https://github.com/cloudquery/cloudquery/issues/5756)) ([4c588b1](https://github.com/cloudquery/cloudquery/commit/4c588b1afe5bf0b71b4c77884994a02cf4edca56))
* **azure:** Use lowercase namespaces ([#5789](https://github.com/cloudquery/cloudquery/issues/5789)) ([b43e1bd](https://github.com/cloudquery/cloudquery/commit/b43e1bd3ca8e0e2e35c363385838fe170d7af59f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.5 ([#5661](https://github.com/cloudquery/cloudquery/issues/5661)) ([b354b8a](https://github.com/cloudquery/cloudquery/commit/b354b8a3683fa2bc918c1002afac487427d65a5f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.6 ([#5790](https://github.com/cloudquery/cloudquery/issues/5790)) ([8e2663c](https://github.com/cloudquery/cloudquery/commit/8e2663c17c3347afd5e53f665462adc3e709c96c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.7 ([#5797](https://github.com/cloudquery/cloudquery/issues/5797)) ([15da529](https://github.com/cloudquery/cloudquery/commit/15da5294786fa2656228ca5bbc48ef1fc44e486b))

## [2.1.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v2.1.0...plugins-source-azure-v2.1.1) (2022-12-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.4 ([#5649](https://github.com/cloudquery/cloudquery/issues/5649)) ([b4aa889](https://github.com/cloudquery/cloudquery/commit/b4aa889e396db3b0887d1684e4bc07da6050af43))

## [2.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v2.0.0...plugins-source-azure-v2.1.0) (2022-12-14)


### Features

* **azure:** Add azure_compute_skus ([#5629](https://github.com/cloudquery/cloudquery/issues/5629)) ([f169b9a](https://github.com/cloudquery/cloudquery/commit/f169b9a24a6ead4798296a3a4d7170192727d19e))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.3 ([#5639](https://github.com/cloudquery/cloudquery/issues/5639)) ([6452d0e](https://github.com/cloudquery/cloudquery/commit/6452d0ed5a44abad9d7530af6e79cde6504d0c4c))

## [2.0.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.10...plugins-source-azure-v2.0.0) (2022-12-13)


### ⚠ BREAKING CHANGES

* **azure:** Move to new SDK. Many new resources were added, please see https://www.cloudquery.io/docs/plugins/sources/azure/tables for the new list. The main difference in columns is that we don't unwrap Azure properties into separate columns. Instead `properties` are stored as a JSON column

### Features

* **azure:** Add keyvaults  and network_gateways ([#5564](https://github.com/cloudquery/cloudquery/issues/5564)) ([147becb](https://github.com/cloudquery/cloudquery/commit/147becb2a746f951169529437849e50ecedc43df))
* **azure:** Move to new sdk ([8657395](https://github.com/cloudquery/cloudquery/commit/8657395da26d34b68328096b0d8cdf5f7d0cc565))
* **azure:** New tables for mariadb,mysql,cdn ([#5549](https://github.com/cloudquery/cloudquery/issues/5549)) ([ae8bb85](https://github.com/cloudquery/cloudquery/commit/ae8bb85ea816c47661eb3168fdc8f71f92a47754))
* **website:** Add plugins tables ([#5259](https://github.com/cloudquery/cloudquery/issues/5259)) ([c336f4e](https://github.com/cloudquery/cloudquery/commit/c336f4e25e192ffdd4c211d4a35b67b71d01d1f8))


### Bug Fixes

* **azure:** Remove azure_storage_deleted_accounts ([#5551](https://github.com/cloudquery/cloudquery/issues/5551)) ([821d63a](https://github.com/cloudquery/cloudquery/commit/821d63ab46b857340a5aac6bbc9703a4ed564849))
* **deps:** Update golang.org/x/exp digest to 6ab00d0 ([#5200](https://github.com/cloudquery/cloudquery/issues/5200)) ([66a8ae4](https://github.com/cloudquery/cloudquery/commit/66a8ae439b643f01bd1d72f091b9abe04ab1013b))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.2.0 ([#5211](https://github.com/cloudquery/cloudquery/issues/5211)) ([46b470e](https://github.com/cloudquery/cloudquery/commit/46b470e93e24bd37789b8d85e6c5e0251629517b))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azidentity to v1.2.0 ([#5212](https://github.com/cloudquery/cloudquery/issues/5212)) ([64c1802](https://github.com/cloudquery/cloudquery/commit/64c18026118ff2b926229ec3f50e54fd82956838))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.11.0 ([#5416](https://github.com/cloudquery/cloudquery/issues/5416)) ([2e7ca35](https://github.com/cloudquery/cloudquery/commit/2e7ca35922fdb14fd717f582aaaa9693dae2ef4c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.0 ([#5539](https://github.com/cloudquery/cloudquery/issues/5539)) ([fb71293](https://github.com/cloudquery/cloudquery/commit/fb71293d5cfe1b2ef32ba83d604ac3c48e662bce))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.12.2 ([#5583](https://github.com/cloudquery/cloudquery/issues/5583)) ([d721c4e](https://github.com/cloudquery/cloudquery/commit/d721c4e06b8a97b5373215aca0e4ed64942ac489))
* **deps:** Update module github.com/gofrs/uuid to v4.3.1 ([#5204](https://github.com/cloudquery/cloudquery/issues/5204)) ([1ca1bdc](https://github.com/cloudquery/cloudquery/commit/1ca1bdceeeef21cbc1256b58effe0c103580dee0))

## [1.4.10](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.9...plugins-source-azure-v1.4.10) (2022-11-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v1.10.0 ([#5153](https://github.com/cloudquery/cloudquery/issues/5153)) ([ea1f77e](https://github.com/cloudquery/cloudquery/commit/ea1f77e910f430287600e74cedd7d3f4ae79eb18))
* **deps:** Update plugin-sdk for azure to v1.8.1 ([#5033](https://github.com/cloudquery/cloudquery/issues/5033)) ([3c40cb5](https://github.com/cloudquery/cloudquery/commit/3c40cb5942692f7d808834950409cee89fe49fd7))
* **deps:** Update plugin-sdk for azure to v1.8.2 ([#5075](https://github.com/cloudquery/cloudquery/issues/5075)) ([029b138](https://github.com/cloudquery/cloudquery/commit/029b13867c2d0cd1e119e6a8a1ed4fa8c30c6bdc))
* **deps:** Update plugin-sdk for azure to v1.9.0 ([#5093](https://github.com/cloudquery/cloudquery/issues/5093)) ([ff9427f](https://github.com/cloudquery/cloudquery/commit/ff9427f8514615d611e09a2f3a4e2bf1b46170fc))

## [1.4.9](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.8...plugins-source-azure-v1.4.9) (2022-11-23)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.7.0 ([#4904](https://github.com/cloudquery/cloudquery/issues/4904)) ([5cf943d](https://github.com/cloudquery/cloudquery/commit/5cf943da8672228e53e14c1f81100bbc99cb66d9))
* **deps:** Update plugin-sdk for azure to v1.8.0 ([#4967](https://github.com/cloudquery/cloudquery/issues/4967)) ([4adcc8f](https://github.com/cloudquery/cloudquery/commit/4adcc8f1f41bfab68a42b87bb314b7dede92137f))

## [1.4.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.7...plugins-source-azure-v1.4.8) (2022-11-21)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.6.0 ([#4842](https://github.com/cloudquery/cloudquery/issues/4842)) ([e341cc3](https://github.com/cloudquery/cloudquery/commit/e341cc3dc76bccef37cf85e65f4632373adcc07d))

## [1.4.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.6...plugins-source-azure-v1.4.7) (2022-11-15)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.5.3 ([#4641](https://github.com/cloudquery/cloudquery/issues/4641)) ([67a7aa0](https://github.com/cloudquery/cloudquery/commit/67a7aa0f8fd18ba074fd7e8746c12ae80ae0e653))

## [1.4.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.5...plugins-source-azure-v1.4.6) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.5.2 ([#4556](https://github.com/cloudquery/cloudquery/issues/4556)) ([034e8f7](https://github.com/cloudquery/cloudquery/commit/034e8f7668b7272aa9ad7299ee00eb6100e83bea))

## [1.4.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.4...plugins-source-azure-v1.4.5) (2022-11-14)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.5.1 ([#4497](https://github.com/cloudquery/cloudquery/issues/4497)) ([b0a2a7d](https://github.com/cloudquery/cloudquery/commit/b0a2a7d102b1b970163743cca3865683196c4bb5))

## [1.4.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.3...plugins-source-azure-v1.4.4) (2022-11-11)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.5.0 ([#4386](https://github.com/cloudquery/cloudquery/issues/4386)) ([26fa93f](https://github.com/cloudquery/cloudquery/commit/26fa93fb9ff3ada21d6dfc1ba75df423ffd7d176))
* Fix links in Grafana compliance dashboards ([#4338](https://github.com/cloudquery/cloudquery/issues/4338)) ([e71ba56](https://github.com/cloudquery/cloudquery/commit/e71ba567fdd21ae9cf059023795c6765d1766848))

## [1.4.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.2...plugins-source-azure-v1.4.3) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.4.1 ([#4289](https://github.com/cloudquery/cloudquery/issues/4289)) ([f91d03d](https://github.com/cloudquery/cloudquery/commit/f91d03dcee936fb3edcce49ec7dcbb43d3c47d8b))

## [1.4.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.1...plugins-source-azure-v1.4.2) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.4.0 ([#4227](https://github.com/cloudquery/cloudquery/issues/4227)) ([d9b582a](https://github.com/cloudquery/cloudquery/commit/d9b582a9dd153a5e15228c8ea4b4c56da4758001))

## [1.4.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.4.0...plugins-source-azure-v1.4.1) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.3.2 ([#4194](https://github.com/cloudquery/cloudquery/issues/4194)) ([792f8b4](https://github.com/cloudquery/cloudquery/commit/792f8b4ddeef7ecc5f67e1f5baf5ce43cd50a063))

## [1.4.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.12...plugins-source-azure-v1.4.0) (2022-11-10)


### Features

* **azure:** Add description to tables ([#4189](https://github.com/cloudquery/cloudquery/issues/4189)) ([a5fb33d](https://github.com/cloudquery/cloudquery/commit/a5fb33d1b2ee51f7abfa11eddf24d333668590d2))

## [1.3.12](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.11...plugins-source-azure-v1.3.12) (2022-11-10)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.3.1 ([#4145](https://github.com/cloudquery/cloudquery/issues/4145)) ([d9462e6](https://github.com/cloudquery/cloudquery/commit/d9462e6c66c0371f75fd6917d2668096bce661c7))

## [1.3.11](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.10...plugins-source-azure-v1.3.11) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.3.0 ([#4068](https://github.com/cloudquery/cloudquery/issues/4068)) ([30d5543](https://github.com/cloudquery/cloudquery/commit/30d55433a61e54758ba36173f0e626319b7afdab))

## [1.3.10](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.9...plugins-source-azure-v1.3.10) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.2.0 ([#4037](https://github.com/cloudquery/cloudquery/issues/4037)) ([b96a731](https://github.com/cloudquery/cloudquery/commit/b96a731bd9590f76c4d08518d2a9fa1e3747f137))

## [1.3.9](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.8...plugins-source-azure-v1.3.9) (2022-11-09)


### Bug Fixes

* **deps:** Update plugin-sdk for csv to v1.1.0 ([#3918](https://github.com/cloudquery/cloudquery/issues/3918)) ([f1acd68](https://github.com/cloudquery/cloudquery/commit/f1acd688fcd90011cc9be1be2285e3fe9369e341))

## [1.3.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.7...plugins-source-azure-v1.3.8) (2022-11-08)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1.1.0 ([#3915](https://github.com/cloudquery/cloudquery/issues/3915)) ([6b240e7](https://github.com/cloudquery/cloudquery/commit/6b240e758ea2ee6b72d68afffcb792b4117d93ba))

## [1.3.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.6...plugins-source-azure-v1.3.7) (2022-11-08)


### Bug Fixes

* **deps:** Update dependency cloudquery/cloudquery to v1.6.6 ([#3830](https://github.com/cloudquery/cloudquery/issues/3830)) ([2b30af3](https://github.com/cloudquery/cloudquery/commit/2b30af3b6269e827d4744748c898046330648521))
* **deps:** Update plugin-sdk for azure to v1.0.3 ([#3846](https://github.com/cloudquery/cloudquery/issues/3846)) ([5830730](https://github.com/cloudquery/cloudquery/commit/583073042b405463bf1176b9dcc9b2ad7f5dfa94))
* **deps:** Upgrade plugin-sdk to v1.0.4 for plugins ([#3889](https://github.com/cloudquery/cloudquery/issues/3889)) ([6767243](https://github.com/cloudquery/cloudquery/commit/6767243ec70bfae7a4c457bf4b5edf013c54c392))

## [1.3.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.5...plugins-source-azure-v1.3.6) (2022-11-07)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v1 ([#3775](https://github.com/cloudquery/cloudquery/issues/3775)) ([49928d6](https://github.com/cloudquery/cloudquery/commit/49928d62666f187417800c9628f5408c40359732))

## [1.3.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.4...plugins-source-azure-v1.3.5) (2022-11-07)


### Bug Fixes

* **deps:** Update SDK to v0.13.23 ([#3742](https://github.com/cloudquery/cloudquery/issues/3742)) ([eed6590](https://github.com/cloudquery/cloudquery/commit/eed6590a518340a96f0705ac9f0ff53344f57e88))

## [1.3.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.3...plugins-source-azure-v1.3.4) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.22 ([#3678](https://github.com/cloudquery/cloudquery/issues/3678)) ([7daf8aa](https://github.com/cloudquery/cloudquery/commit/7daf8aac83719199d64a44fe9ae82bd37e32b796))

## [1.3.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.2...plugins-source-azure-v1.3.3) (2022-11-06)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.21 ([#3630](https://github.com/cloudquery/cloudquery/issues/3630)) ([722057d](https://github.com/cloudquery/cloudquery/commit/722057d4d2a3fe7df8993fc3ecd23e015f626972))

## [1.3.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.1...plugins-source-azure-v1.3.2) (2022-11-04)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.20 ([#3570](https://github.com/cloudquery/cloudquery/issues/3570)) ([f8c0106](https://github.com/cloudquery/cloudquery/commit/f8c01061120a8d984eb9b6989b9f88f844131b24))

## [1.3.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.3.0...plugins-source-azure-v1.3.1) (2022-11-03)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.19 ([#3501](https://github.com/cloudquery/cloudquery/issues/3501)) ([fb6616e](https://github.com/cloudquery/cloudquery/commit/fb6616e34b44e664f75964c2ee15e52f2f6b85ce))

## [1.3.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.2.0...plugins-source-azure-v1.3.0) (2022-11-02)


### Features

* Add Azure functions ([#3182](https://github.com/cloudquery/cloudquery/issues/3182)) ([1020ef9](https://github.com/cloudquery/cloudquery/commit/1020ef98e8d2d65b82d215afecdb0d9339911451))
* **Azure:** Add site auth settings v2 ([#3269](https://github.com/cloudquery/cloudquery/issues/3269)) ([fd54aff](https://github.com/cloudquery/cloudquery/commit/fd54affea4b10c93a79bfb2f7213a87133ff85f5))


### Bug Fixes

* **azure:** Handle nil `VnetName` ([#3181](https://github.com/cloudquery/cloudquery/issues/3181)) ([0006490](https://github.com/cloudquery/cloudquery/commit/00064906c202b35e4a1577b9db27c1de52ac4b9c))

## [1.2.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.1.0...plugins-source-azure-v1.2.0) (2022-11-01)


### Features

* Migrate cli, plugins and destinations to new type system ([#3323](https://github.com/cloudquery/cloudquery/issues/3323)) ([f265a94](https://github.com/cloudquery/cloudquery/commit/f265a94448ad55c968b26ba8a19681bc81086c11))


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to c99f073 ([#3372](https://github.com/cloudquery/cloudquery/issues/3372)) ([c64bc54](https://github.com/cloudquery/cloudquery/commit/c64bc5410f20aba71e54308b39017dbf102fdead))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.1.4 ([#3376](https://github.com/cloudquery/cloudquery/issues/3376)) ([294708e](https://github.com/cloudquery/cloudquery/commit/294708e2c36649e23108b1a5985b9cae225fd54d))
* **deps:** Update plugin-sdk for azure to v0.13.17 ([#3400](https://github.com/cloudquery/cloudquery/issues/3400)) ([a244d6c](https://github.com/cloudquery/cloudquery/commit/a244d6c1421858b2829a8049ad6dfc5bbe655d6d))
* **deps:** Update plugin-sdk for azure to v0.13.18 ([#3410](https://github.com/cloudquery/cloudquery/issues/3410)) ([36ea67c](https://github.com/cloudquery/cloudquery/commit/36ea67c78e0985309924107b3b33877c800cc937))

## [1.1.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.10...plugins-source-azure-v1.1.0) (2022-10-31)


### Features

* Update all plugins to SDK with metrics and DFS scheduler ([#3286](https://github.com/cloudquery/cloudquery/issues/3286)) ([a35b8e8](https://github.com/cloudquery/cloudquery/commit/a35b8e89d625287a9b9406ff18cfac78ffdb1241))

## [1.0.10](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.9...plugins-source-azure-v1.0.10) (2022-10-27)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.14 ([#3212](https://github.com/cloudquery/cloudquery/issues/3212)) ([7cfc2ad](https://github.com/cloudquery/cloudquery/commit/7cfc2adfb91b3a94fd44703e29f208d37aa32bed))

## [1.0.9](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.8...plugins-source-azure-v1.0.9) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.12 ([#3100](https://github.com/cloudquery/cloudquery/issues/3100)) ([a517055](https://github.com/cloudquery/cloudquery/commit/a517055d1ad29a91b30e078bac316f33d16377c5))

## [1.0.8](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.7...plugins-source-azure-v1.0.8) (2022-10-20)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.12 ([#3100](https://github.com/cloudquery/cloudquery/issues/3100)) ([a517055](https://github.com/cloudquery/cloudquery/commit/a517055d1ad29a91b30e078bac316f33d16377c5))

## [1.0.7](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.6...plugins-source-azure-v1.0.7) (2022-10-19)


### Bug Fixes

* **deps:** Update plugin-sdk to v0.13.11 ([#3030](https://github.com/cloudquery/cloudquery/issues/3030)) ([9909c4a](https://github.com/cloudquery/cloudquery/commit/9909c4a0715a06b7c1d69c9bd23c500ac7b4adc1))

## [1.0.6](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.5...plugins-source-azure-v1.0.6) (2022-10-18)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.9 ([#2927](https://github.com/cloudquery/cloudquery/issues/2927)) ([4a556a5](https://github.com/cloudquery/cloudquery/commit/4a556a5fe73d8094a80afc09077232d21d4d1531))

## [1.0.5](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.4...plugins-source-azure-v1.0.5) (2022-10-14)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.8 ([#2849](https://github.com/cloudquery/cloudquery/issues/2849)) ([2eb9eb1](https://github.com/cloudquery/cloudquery/commit/2eb9eb14c273ace46562749db8bd78f642762a76))

## [1.0.4](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.3...plugins-source-azure-v1.0.4) (2022-10-13)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.7 ([#2779](https://github.com/cloudquery/cloudquery/issues/2779)) ([4d6b373](https://github.com/cloudquery/cloudquery/commit/4d6b373d4fa1525c2ae2bda107e2770e308264b4))

## [1.0.3](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.2...plugins-source-azure-v1.0.3) (2022-10-12)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.13.6 ([#2718](https://github.com/cloudquery/cloudquery/issues/2718)) ([8e42d5f](https://github.com/cloudquery/cloudquery/commit/8e42d5fbfb0b9c8352db876e56f8f1a91a91557b))

## [1.0.2](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.1...plugins-source-azure-v1.0.2) (2022-10-12)


### Bug Fixes

* **azure:** Migrate Grafana dashboards ([#2628](https://github.com/cloudquery/cloudquery/issues/2628)) ([d2436b9](https://github.com/cloudquery/cloudquery/commit/d2436b97ac1449fbefcf8bd1dee042bc4eb2b162))
* **azure:** Skip resource if SiteConfig is nil ([#2578](https://github.com/cloudquery/cloudquery/issues/2578)) ([1ccab59](https://github.com/cloudquery/cloudquery/commit/1ccab597d24c67dee95dbe6460d8e0cc4ccdcbdb))
* **deps:** Update plugin-sdk for azure to v0.12.10 ([#2545](https://github.com/cloudquery/cloudquery/issues/2545)) ([e68ee23](https://github.com/cloudquery/cloudquery/commit/e68ee23bbd9083993acbdf05bc22db58112e1011))
* Update Azure plugin to SDK v0.13.5 ([#2662](https://github.com/cloudquery/cloudquery/issues/2662)) ([50160e6](https://github.com/cloudquery/cloudquery/commit/50160e60c3b90ca2dfa3ab4b5f62a149d0cea5da))

## [1.0.1](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.0...plugins-source-azure-v1.0.1) (2022-10-09)


### Bug Fixes

* **deps:** Update plugin-sdk for azure to v0.12.3 ([#2278](https://github.com/cloudquery/cloudquery/issues/2278)) ([2035607](https://github.com/cloudquery/cloudquery/commit/2035607509c28c51bfdb2ae1fbde12f42d3151ab))
* **deps:** Update plugin-sdk for azure to v0.12.4 ([#2395](https://github.com/cloudquery/cloudquery/issues/2395)) ([730471e](https://github.com/cloudquery/cloudquery/commit/730471e44642d4e4daac5358a0fbfc59dab5c79b))
* **deps:** Update plugin-sdk for azure to v0.12.5 ([#2417](https://github.com/cloudquery/cloudquery/issues/2417)) ([b7ff62e](https://github.com/cloudquery/cloudquery/commit/b7ff62ea449673922d9a120a173009ded7946d15))
* **deps:** Update plugin-sdk for azure to v0.12.6 ([#2433](https://github.com/cloudquery/cloudquery/issues/2433)) ([4caccc9](https://github.com/cloudquery/cloudquery/commit/4caccc9be4abee1a78b86a68f8ef9fdb6d70fa6c))
* **deps:** Update plugin-sdk for azure to v0.12.7 ([#2446](https://github.com/cloudquery/cloudquery/issues/2446)) ([d0bf9d7](https://github.com/cloudquery/cloudquery/commit/d0bf9d7a77aab5e708c03677c055a9be77332bd9))
* **deps:** Update plugin-sdk for azure to v0.12.8 ([#2496](https://github.com/cloudquery/cloudquery/issues/2496)) ([9eddb8b](https://github.com/cloudquery/cloudquery/commit/9eddb8be9e80486427f3493cd3bba839a1508ce4))
* **deps:** Update plugin-sdk for azure to v0.12.9 ([#2510](https://github.com/cloudquery/cloudquery/issues/2510)) ([8a0161e](https://github.com/cloudquery/cloudquery/commit/8a0161eb610ef2ac68685bd47f572c34b0822683))

## [1.0.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.13.4...plugins-source-azure-v1.0.0) (2022-10-04)


### ⚠ BREAKING CHANGES

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

### Features

* [Official v1 release](https://www.cloudquery.io/blog/cloudquery-v1-release)

## [1.0.3-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.2-pre.0...plugins-source-azure-v1.0.3-pre.0) (2022-10-03)


### Bug Fixes

* Azure policies ([#1861](https://github.com/cloudquery/cloudquery/issues/1861)) ([062907a](https://github.com/cloudquery/cloudquery/commit/062907a9684a879c7ed2c8b1ab80d752993c8d15))
* **deps:** Update plugin-sdk for azure to v0.11.6 ([#2252](https://github.com/cloudquery/cloudquery/issues/2252)) ([3092acb](https://github.com/cloudquery/cloudquery/commit/3092acb1e14fe63c417187588e9b67db8b2fb82f))

## [1.0.2-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.1-pre.0...plugins-source-azure-v1.0.2-pre.0) (2022-10-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.2 ([#2162](https://github.com/cloudquery/cloudquery/issues/2162)) ([5701aa5](https://github.com/cloudquery/cloudquery/commit/5701aa5b0a8d04e9e99e3efe6e27d5f7ff29b216))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.4 ([#2210](https://github.com/cloudquery/cloudquery/issues/2210)) ([760d0a6](https://github.com/cloudquery/cloudquery/commit/760d0a6e7983cfb08fa4b519a908fcda91abbdc0))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.5 ([#2227](https://github.com/cloudquery/cloudquery/issues/2227)) ([7db2dde](https://github.com/cloudquery/cloudquery/commit/7db2dde8e14f370627451d8494f9a3b7fb20c61a))
* Don't multiplex subscriptions ([#2018](https://github.com/cloudquery/cloudquery/issues/2018)) ([94d43b3](https://github.com/cloudquery/cloudquery/commit/94d43b3db6ea284ec6d01047f07397b5404ef96f))

## [1.0.1-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v1.0.0-pre.0...plugins-source-azure-v1.0.1-pre.0) (2022-10-02)


### Bug Fixes

* **deps:** Update golang.org/x/exp digest to 540bb73 ([#2169](https://github.com/cloudquery/cloudquery/issues/2169)) ([d183fea](https://github.com/cloudquery/cloudquery/commit/d183feabb803cc5516cf23f651aa9ca33d13bfba))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.1.3 ([#2173](https://github.com/cloudquery/cloudquery/issues/2173)) ([d707d60](https://github.com/cloudquery/cloudquery/commit/d707d6061b2ea53a5c746d6dd9ea0d851fab4368))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.10.2 ([#2048](https://github.com/cloudquery/cloudquery/issues/2048)) ([e407991](https://github.com/cloudquery/cloudquery/commit/e4079914772d8191639b9935aa5970b8e27b082f))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.11.0 ([#2135](https://github.com/cloudquery/cloudquery/issues/2135)) ([1729467](https://github.com/cloudquery/cloudquery/commit/1729467b2119555e18b15d73c91cd501ccf7ecb8))
* Use ParentResourceFieldResolver instead of ParentIDResolver ([#2125](https://github.com/cloudquery/cloudquery/issues/2125)) ([5039788](https://github.com/cloudquery/cloudquery/commit/50397883131f6652cfbcccadcf2817784a46e199))
* Use TypeString for parent ids ([#2136](https://github.com/cloudquery/cloudquery/issues/2136)) ([a62f1a0](https://github.com/cloudquery/cloudquery/commit/a62f1a07e9a35198a984b081f4fb83ecaf79c4e8))

## [1.0.0-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v0.14.2-pre.0...plugins-source-azure-v1.0.0-pre.0) (2022-09-26)


### ⚠ BREAKING CHANGES

* Migrate Azure plugin to v2 (#1754)
* Fix Azure credential chain (#1283)

### Features

* Add website, docs and blog to our main repo ([#1159](https://github.com/cloudquery/cloudquery/issues/1159)) ([dd69948](https://github.com/cloudquery/cloudquery/commit/dd69948feced004497f127d284f2604de0354a1f))
* Added azure cdn profiles ([#1460](https://github.com/cloudquery/cloudquery/issues/1460)) ([cc154c5](https://github.com/cloudquery/cloudquery/commit/cc154c5128d58474958ffd8330ebfdf281ebbe94))
* Migrate Azure plugin to v2 ([#1754](https://github.com/cloudquery/cloudquery/issues/1754)) ([ee9bef2](https://github.com/cloudquery/cloudquery/commit/ee9bef21910890b2e81e00c4aed598e400ad5f85))


### Bug Fixes

* Add missing `azure_keyvault_secrets` tables ([#1937](https://github.com/cloudquery/cloudquery/issues/1937)) ([491aa66](https://github.com/cloudquery/cloudquery/commit/491aa6665d973a29bd8d95042df1d5082edb3770))
* **deps:** Update golang.org/x/sync digest to 7fc1605 ([#1652](https://github.com/cloudquery/cloudquery/issues/1652)) ([daafae1](https://github.com/cloudquery/cloudquery/commit/daafae1c60c14c90b70c3338a8ff6dc25ba84290))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.1.2 ([#1664](https://github.com/cloudquery/cloudquery/issues/1664)) ([5390e13](https://github.com/cloudquery/cloudquery/commit/5390e1350854a74b5431ebaa18cb230687481819))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.3 ([#1858](https://github.com/cloudquery/cloudquery/issues/1858)) ([9e3ace7](https://github.com/cloudquery/cloudquery/commit/9e3ace775da2d600968ef4275e9e0013d4dfd825))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.4 ([#1862](https://github.com/cloudquery/cloudquery/issues/1862)) ([5d141cf](https://github.com/cloudquery/cloudquery/commit/5d141cf6006e26cf240ddf295dda53c16f7386a4))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.1 ([#1865](https://github.com/cloudquery/cloudquery/issues/1865)) ([474bb70](https://github.com/cloudquery/cloudquery/commit/474bb7081b6e9b6ffc5ac949ed3a664f92083c82))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.2 ([#1872](https://github.com/cloudquery/cloudquery/issues/1872)) ([49ed26d](https://github.com/cloudquery/cloudquery/commit/49ed26d231c91ac1b5b00cc55d3d0a8a5a6306f7))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.3 ([#1886](https://github.com/cloudquery/cloudquery/issues/1886)) ([7435d59](https://github.com/cloudquery/cloudquery/commit/7435d593e51ca829d3a328eebc9517e9cb2a4ef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.4 ([#1889](https://github.com/cloudquery/cloudquery/issues/1889)) ([63a5362](https://github.com/cloudquery/cloudquery/commit/63a5362995aa680b291f2411d01e776e884896d4))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.9 ([#1891](https://github.com/cloudquery/cloudquery/issues/1891)) ([3469f20](https://github.com/cloudquery/cloudquery/commit/3469f20e76e9dcbf48b9c6e3e7c0c2224c5b8ad3))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))
* Fix Azure credential chain ([#1283](https://github.com/cloudquery/cloudquery/issues/1283)) ([c2aadf7](https://github.com/cloudquery/cloudquery/commit/c2aadf78533a65679ef40ea32c1b899724ab6d69))
* Generate Azure date.time as Timestamps ([#1885](https://github.com/cloudquery/cloudquery/issues/1885)) ([92d41a1](https://github.com/cloudquery/cloudquery/commit/92d41a1df754bdcb75d40652e438d66352db435c))
* Regenerate Azure resources ([#1875](https://github.com/cloudquery/cloudquery/issues/1875)) ([7411a27](https://github.com/cloudquery/cloudquery/commit/7411a2742a9bc4eb0b1bf1cf490d92a3a41c390f))
* Update Azure codegen ([#1936](https://github.com/cloudquery/cloudquery/issues/1936)) ([4b739db](https://github.com/cloudquery/cloudquery/commit/4b739db73b21a54320a49dc7c421b6eeb92a6a4a))

## [0.14.2-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins-source-azure-v0.14.1-pre.0...plugins-source-azure-v0.14.2-pre.0) (2022-09-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.0 ([#1997](https://github.com/cloudquery/cloudquery/issues/1997)) ([4fa40da](https://github.com/cloudquery/cloudquery/commit/4fa40da04b427f864d2dc11f133e5c83e53ce4b6))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.1 ([#2024](https://github.com/cloudquery/cloudquery/issues/2024)) ([8f88de4](https://github.com/cloudquery/cloudquery/commit/8f88de4b4eaeabae7369ba309e765a252392ee8c))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.8.2 ([#2044](https://github.com/cloudquery/cloudquery/issues/2044)) ([9b69b46](https://github.com/cloudquery/cloudquery/commit/9b69b468536521b20b77ec1fc180fc85aeeba376))

## [0.14.1-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.14.0-pre.0...plugins/source/azure/v0.14.1-pre.0) (2022-09-22)


### Bug Fixes

* Add missing `azure_keyvault_secrets` tables ([#1937](https://github.com/cloudquery/cloudquery/issues/1937)) ([491aa66](https://github.com/cloudquery/cloudquery/commit/491aa6665d973a29bd8d95042df1d5082edb3770))
* Update Azure codegen ([#1936](https://github.com/cloudquery/cloudquery/issues/1936)) ([4b739db](https://github.com/cloudquery/cloudquery/commit/4b739db73b21a54320a49dc7c421b6eeb92a6a4a))

## [0.14.0-pre.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure-v0.13.4-pre.0...plugins/source/azure/v0.14.0-pre.0) (2022-09-21)


### ⚠ BREAKING CHANGES

* Migrate Azure plugin to v2 (#1754)
* Fix Azure credential chain (#1283)

### Features

* Add website, docs and blog to our main repo ([#1159](https://github.com/cloudquery/cloudquery/issues/1159)) ([dd69948](https://github.com/cloudquery/cloudquery/commit/dd69948feced004497f127d284f2604de0354a1f))
* Added azure cdn profiles ([#1460](https://github.com/cloudquery/cloudquery/issues/1460)) ([cc154c5](https://github.com/cloudquery/cloudquery/commit/cc154c5128d58474958ffd8330ebfdf281ebbe94))
* Migrate Azure plugin to v2 ([#1754](https://github.com/cloudquery/cloudquery/issues/1754)) ([ee9bef2](https://github.com/cloudquery/cloudquery/commit/ee9bef21910890b2e81e00c4aed598e400ad5f85))


### Bug Fixes

* **deps:** Update golang.org/x/sync digest to 7fc1605 ([#1652](https://github.com/cloudquery/cloudquery/issues/1652)) ([daafae1](https://github.com/cloudquery/cloudquery/commit/daafae1c60c14c90b70c3338a8ff6dc25ba84290))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.1.2 ([#1664](https://github.com/cloudquery/cloudquery/issues/1664)) ([5390e13](https://github.com/cloudquery/cloudquery/commit/5390e1350854a74b5431ebaa18cb230687481819))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.9 ([#1286](https://github.com/cloudquery/cloudquery/issues/1286)) ([67ac422](https://github.com/cloudquery/cloudquery/commit/67ac422f392387e674cb70386e612befa5b455f0))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.3 ([#1858](https://github.com/cloudquery/cloudquery/issues/1858)) ([9e3ace7](https://github.com/cloudquery/cloudquery/commit/9e3ace775da2d600968ef4275e9e0013d4dfd825))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.6.4 ([#1862](https://github.com/cloudquery/cloudquery/issues/1862)) ([5d141cf](https://github.com/cloudquery/cloudquery/commit/5d141cf6006e26cf240ddf295dda53c16f7386a4))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.1 ([#1865](https://github.com/cloudquery/cloudquery/issues/1865)) ([474bb70](https://github.com/cloudquery/cloudquery/commit/474bb7081b6e9b6ffc5ac949ed3a664f92083c82))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.2 ([#1872](https://github.com/cloudquery/cloudquery/issues/1872)) ([49ed26d](https://github.com/cloudquery/cloudquery/commit/49ed26d231c91ac1b5b00cc55d3d0a8a5a6306f7))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.3 ([#1886](https://github.com/cloudquery/cloudquery/issues/1886)) ([7435d59](https://github.com/cloudquery/cloudquery/commit/7435d593e51ca829d3a328eebc9517e9cb2a4ef0))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.4 ([#1889](https://github.com/cloudquery/cloudquery/issues/1889)) ([63a5362](https://github.com/cloudquery/cloudquery/commit/63a5362995aa680b291f2411d01e776e884896d4))
* **deps:** Update module github.com/cloudquery/plugin-sdk to v0.7.9 ([#1891](https://github.com/cloudquery/cloudquery/issues/1891)) ([3469f20](https://github.com/cloudquery/cloudquery/commit/3469f20e76e9dcbf48b9c6e3e7c0c2224c5b8ad3))
* Fix Azure credential chain ([#1283](https://github.com/cloudquery/cloudquery/issues/1283)) ([c2aadf7](https://github.com/cloudquery/cloudquery/commit/c2aadf78533a65679ef40ea32c1b899724ab6d69))
* Generate Azure date.time as Timestamps ([#1885](https://github.com/cloudquery/cloudquery/issues/1885)) ([92d41a1](https://github.com/cloudquery/cloudquery/commit/92d41a1df754bdcb75d40652e438d66352db435c))
* Regenerate Azure resources ([#1875](https://github.com/cloudquery/cloudquery/issues/1875)) ([7411a27](https://github.com/cloudquery/cloudquery/commit/7411a2742a9bc4eb0b1bf1cf490d92a3a41c390f))

## [0.13.4](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.13.3...plugins/source/azure/v0.13.4) (2022-09-01)


### Bug Fixes

* **deps:** Update golang.org/x/sync digest to 7fc1605 ([#1652](https://github.com/cloudquery/cloudquery/issues/1652)) ([daafae1](https://github.com/cloudquery/cloudquery/commit/daafae1c60c14c90b70c3338a8ff6dc25ba84290))
* **deps:** Update module github.com/Azure/azure-sdk-for-go/sdk/azcore to v1.1.2 ([#1664](https://github.com/cloudquery/cloudquery/issues/1664)) ([5390e13](https://github.com/cloudquery/cloudquery/commit/5390e1350854a74b5431ebaa18cb230687481819))

## [0.13.3](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.13.2...plugins/source/azure/v0.13.3) (2022-08-21)


### Features

* Added azure cdn profiles ([#1460](https://github.com/cloudquery/cloudquery/issues/1460)) ([cc154c5](https://github.com/cloudquery/cloudquery/commit/cc154c5128d58474958ffd8330ebfdf281ebbe94))

## [0.13.2](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.13.1...plugins/source/azure/v0.13.2) (2022-08-21)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.11 ([#1491](https://github.com/cloudquery/cloudquery/issues/1491)) ([5140bef](https://github.com/cloudquery/cloudquery/commit/5140bef4aa7c50a97a604db1e92df75ead2893fc))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.12 ([#1503](https://github.com/cloudquery/cloudquery/issues/1503)) ([a740719](https://github.com/cloudquery/cloudquery/commit/a7407199c9617784a1834b9d0c42788e03301de5))

## [0.13.1](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.13.0...plugins/source/azure/v0.13.1) (2022-08-18)


### Features

* Add website, docs and blog to our main repo ([#1159](https://github.com/cloudquery/cloudquery/issues/1159)) ([dd69948](https://github.com/cloudquery/cloudquery/commit/dd69948feced004497f127d284f2604de0354a1f))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.10 ([#1474](https://github.com/cloudquery/cloudquery/issues/1474)) ([b142e13](https://github.com/cloudquery/cloudquery/commit/b142e135172b1eed1abb2cbec85054ea7f66199d))

## [0.13.0](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure/v0.12.5...plugins/source/azure/v0.13.0) (2022-08-15)


### ⚠ BREAKING CHANGES

* Fix Azure credential chain (#1283)

### Bug Fixes

* Fix Azure credential chain ([#1283](https://github.com/cloudquery/cloudquery/issues/1283)) ([c2aadf7](https://github.com/cloudquery/cloudquery/commit/c2aadf78533a65679ef40ea32c1b899724ab6d69))

## [0.12.5](https://github.com/cloudquery/cloudquery/compare/plugins/source/azure-v0.12.4...plugins/source/azure/v0.12.5) (2022-08-15)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.9 ([#1286](https://github.com/cloudquery/cloudquery/issues/1286)) ([67ac422](https://github.com/cloudquery/cloudquery/commit/67ac422f392387e674cb70386e612befa5b455f0))

## [0.12.4](https://github.com/cloudquery/cq-provider-azure/compare/v0.12.3...v0.12.4) (2022-08-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.7 ([#458](https://github.com/cloudquery/cq-provider-azure/issues/458)) ([39c0907](https://github.com/cloudquery/cq-provider-azure/commit/39c090712648ad4989deb10e104f943736dd427c))

## [0.12.3](https://github.com/cloudquery/cq-provider-azure/compare/v0.12.2...v0.12.3) (2022-08-07)


### Features

* Add support for Tenants ([#412](https://github.com/cloudquery/cq-provider-azure/issues/412)) ([940af7f](https://github.com/cloudquery/cq-provider-azure/commit/940af7fdf18ef8e96a178177095947edca21de1c))


### Bug Fixes

* **deps:** Update module github.com/Azure/go-autorest/autorest to v0.11.28 ([#443](https://github.com/cloudquery/cq-provider-azure/issues/443)) ([bc87594](https://github.com/cloudquery/cq-provider-azure/commit/bc8759439a0123aaf7d9359e6970abe3c8d3404e))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.7 ([#453](https://github.com/cloudquery/cq-provider-azure/issues/453)) ([361d908](https://github.com/cloudquery/cq-provider-azure/commit/361d90883c42331304fcbf9f5ffa9579a94a5bb1))
* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.8 ([#457](https://github.com/cloudquery/cq-provider-azure/issues/457)) ([7d61851](https://github.com/cloudquery/cq-provider-azure/commit/7d618516c3ac4a945308e248eebeb7fba9a442c7))
* **deps:** Update module github.com/hashicorp/go-hclog to v1.2.2 ([#444](https://github.com/cloudquery/cq-provider-azure/issues/444)) ([6c769c9](https://github.com/cloudquery/cq-provider-azure/commit/6c769c955a9b8874f4b3cc297584851f079c2059))
* **deps:** Update tubone24/update_release digest to 2146f15 ([#389](https://github.com/cloudquery/cq-provider-azure/issues/389)) ([e5682fd](https://github.com/cloudquery/cq-provider-azure/commit/e5682fd6b3a169c7ff47483c0727f079e4cf92d4))
* **docs:** Some minor docs fixes ([#454](https://github.com/cloudquery/cq-provider-azure/issues/454)) ([6feb015](https://github.com/cloudquery/cq-provider-azure/commit/6feb01553b2d3d14c4c168ed57d19daa25fd32f5))
* **tests:** Remove terraform for Azure ExpressRoute port ([#440](https://github.com/cloudquery/cq-provider-azure/issues/440)) ([7413740](https://github.com/cloudquery/cq-provider-azure/commit/7413740d6e0f5378642aeb903086a7db8a17db53))

## [0.12.2](https://github.com/cloudquery/cq-provider-azure/compare/v0.12.1...v0.12.2) (2022-07-28)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.6 ([#438](https://github.com/cloudquery/cq-provider-azure/issues/438)) ([b4ff9c5](https://github.com/cloudquery/cq-provider-azure/commit/b4ff9c572f4868d3a056d92454ee2839b74adc37))

## [0.12.1](https://github.com/cloudquery/cq-provider-azure/compare/v0.12.0...v0.12.1) (2022-07-28)


### Features

* Add Azure Front Door CDN resource ([#416](https://github.com/cloudquery/cq-provider-azure/issues/416)) ([f5cec30](https://github.com/cloudquery/cq-provider-azure/commit/f5cec307aa3081125c92e8c6f57465f8da52148e))
* Add dashboards ([#430](https://github.com/cloudquery/cq-provider-azure/issues/430)) ([1ca6e0c](https://github.com/cloudquery/cq-provider-azure/commit/1ca6e0c27e213f1e0120d2c03f75868b5c351e49))
* Added servicebus topics ([#419](https://github.com/cloudquery/cq-provider-azure/issues/419)) ([00accb7](https://github.com/cloudquery/cq-provider-azure/commit/00accb7e74a667046c85f42ac626f1fe048d8f39))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-gen to v0.0.6 ([#427](https://github.com/cloudquery/cq-provider-azure/issues/427)) ([25baa83](https://github.com/cloudquery/cq-provider-azure/commit/25baa83188d59eb909f542e6bd94b44f9920dd24))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.2 ([#423](https://github.com/cloudquery/cq-provider-azure/issues/423)) ([88db3a1](https://github.com/cloudquery/cq-provider-azure/commit/88db3a16e5829c9eebbd0d6ab5d88d15cd46d61a))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.3 ([#429](https://github.com/cloudquery/cq-provider-azure/issues/429)) ([ba4ea62](https://github.com/cloudquery/cq-provider-azure/commit/ba4ea625052a66dc9f224e7cac76bd5ab67e014f))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.4 ([#431](https://github.com/cloudquery/cq-provider-azure/issues/431)) ([90dc83c](https://github.com/cloudquery/cq-provider-azure/commit/90dc83cd08399653667510eaf7fbc9a6706fb2ee))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.14.5 ([#436](https://github.com/cloudquery/cq-provider-azure/issues/436)) ([dfe72e6](https://github.com/cloudquery/cq-provider-azure/commit/dfe72e6aec1f8601ae7d015dec3a00d139f4c360))
* **deps:** Update myrotvorets/set-commit-status-action digest to 9d8a3c7 ([#388](https://github.com/cloudquery/cq-provider-azure/issues/388)) ([40dc1c2](https://github.com/cloudquery/cq-provider-azure/commit/40dc1c2b5687b7b8d7d4954a96617c85bcdeacd4))

## [0.12.0](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.23...v0.12.0) (2022-07-20)


### ⚠ BREAKING CHANGES

* Update SDK to v0.14.1 (#422)

### Features

* Policies ([#415](https://github.com/cloudquery/cq-provider-azure/issues/415)) ([70187ad](https://github.com/cloudquery/cq-provider-azure/commit/70187ad84b55306bc9bddf28989945760b096ca1))


### Bug Fixes

* **terraform:** Front Door Terraform resource specification fix ([#424](https://github.com/cloudquery/cq-provider-azure/issues/424)) ([347bb8e](https://github.com/cloudquery/cq-provider-azure/commit/347bb8eb14ea837814f122f7d6133b55f6e2f9ae))


### Miscellaneous Chores

* Update SDK to v0.14.1 ([#422](https://github.com/cloudquery/cq-provider-azure/issues/422)) ([92afc51](https://github.com/cloudquery/cq-provider-azure/commit/92afc5147d69c4d1225bf1f78885f2f621104ece))

## [0.11.23](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.22...v0.11.23) (2022-07-12)


### Features

* Set location to 'unavailable' rather than null in view ([#407](https://github.com/cloudquery/cq-provider-azure/issues/407)) ([7ff8c50](https://github.com/cloudquery/cq-provider-azure/commit/7ff8c506f140139326cedecf6f795d367f9f84df))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.5 ([#411](https://github.com/cloudquery/cq-provider-azure/issues/411)) ([4d7a95c](https://github.com/cloudquery/cq-provider-azure/commit/4d7a95cc81038ed2bbef4f327e3dbbb497dcb58a))
* move mock generation to separate files ([#409](https://github.com/cloudquery/cq-provider-azure/issues/409)) ([9f721e9](https://github.com/cloudquery/cq-provider-azure/commit/9f721e99139dcbe7e80f8c709a1987ba250fe62a))

## [0.11.22](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.21...v0.11.22) (2022-07-06)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.4 ([#403](https://github.com/cloudquery/cq-provider-azure/issues/403)) ([bbbb646](https://github.com/cloudquery/cq-provider-azure/commit/bbbb6468b927aaccfd4fe5ae337ddb97ac52a0f2))
* **deps:** Update module github.com/cloudquery/faker/v3 to v3.7.6 ([#401](https://github.com/cloudquery/cq-provider-azure/issues/401)) ([dd71206](https://github.com/cloudquery/cq-provider-azure/commit/dd712069fb38681f434edf2787ed003e13325cf3))
* Remove relation tables PK ([#404](https://github.com/cloudquery/cq-provider-azure/issues/404)) ([34cf0cc](https://github.com/cloudquery/cq-provider-azure/commit/34cf0cc6b0118a3259c8c317c9035919a0e1ef9f))

## [0.11.21](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.20...v0.11.21) (2022-07-04)


### Bug Fixes

* **deps:** Update module github.com/Azure/go-autorest/autorest to v0.11.27 ([#390](https://github.com/cloudquery/cq-provider-azure/issues/390)) ([5610181](https://github.com/cloudquery/cq-provider-azure/commit/5610181aaf284cf8bfac7b4222c351533976aa9d))
* **deps:** Update module github.com/Azure/go-autorest/autorest/azure/auth to v0.5.11 ([#391](https://github.com/cloudquery/cq-provider-azure/issues/391)) ([5117b28](https://github.com/cloudquery/cq-provider-azure/commit/5117b28e267085012f38cb4c4a31f2fc8372c642))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.3 ([#398](https://github.com/cloudquery/cq-provider-azure/issues/398)) ([7af200f](https://github.com/cloudquery/cq-provider-azure/commit/7af200f7fbeb0d5f7e3e06b0acbe9b1d5dc3787d))
* **deps:** Update module github.com/stretchr/testify to v1.8.0 ([#392](https://github.com/cloudquery/cq-provider-azure/issues/392)) ([0efa132](https://github.com/cloudquery/cq-provider-azure/commit/0efa132013ca4008019c2a122ccc4606847747a4))
* **deps:** Update module github.com/tombuildsstuff/giovanni to v0.20.0 ([#393](https://github.com/cloudquery/cq-provider-azure/issues/393)) ([6624f09](https://github.com/cloudquery/cq-provider-azure/commit/6624f09e3eb821db0d33ac59f283783b4dcf8b7c))

## [0.11.20](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.19...v0.11.20) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.2 ([#385](https://github.com/cloudquery/cq-provider-azure/issues/385)) ([2acb762](https://github.com/cloudquery/cq-provider-azure/commit/2acb7626c60ef2e97225be442fe36cb975b35197))

## [0.11.19](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.18...v0.11.19) (2022-07-03)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.1 ([#379](https://github.com/cloudquery/cq-provider-azure/issues/379)) ([98c4949](https://github.com/cloudquery/cq-provider-azure/commit/98c49490f0d8380a1565bafe4615f504a713e274))

## [0.11.18](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.17...v0.11.18) (2022-06-30)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.13.0 ([#377](https://github.com/cloudquery/cq-provider-azure/issues/377)) ([a183b5b](https://github.com/cloudquery/cq-provider-azure/commit/a183b5b77b05345abafecb775ec0256e7cfde079))

## [0.11.17](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.16...v0.11.17) (2022-06-29)


### Features

* Added Route Tables ([#358](https://github.com/cloudquery/cq-provider-azure/issues/358)) ([71a14c9](https://github.com/cloudquery/cq-provider-azure/commit/71a14c9842fc0bd2dbfb51b082ef0658e36fc4e9))


### Bug Fixes

* Docs to Yaml ([#371](https://github.com/cloudquery/cq-provider-azure/issues/371)) ([e36569a](https://github.com/cloudquery/cq-provider-azure/commit/e36569a055d5f33a4208bc7c5c75d14ce625bf4d))
* Resources view id ([#375](https://github.com/cloudquery/cq-provider-azure/issues/375)) ([7b2053c](https://github.com/cloudquery/cq-provider-azure/commit/7b2053c4ba9bd0689c999b1e87574347130ce539))

## [0.11.16](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.15...v0.11.16) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.5 ([#368](https://github.com/cloudquery/cq-provider-azure/issues/368)) ([b80323a](https://github.com/cloudquery/cq-provider-azure/commit/b80323a9b2f7baf0ff464d6151e8c625f8c6c9c2))

## [0.11.15](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.14...v0.11.15) (2022-06-27)


### Bug Fixes

* **deps:** fix(deps): Update module github.com/cloudquery/cq-provider-sdk to v0.12.4 ([#366](https://github.com/cloudquery/cq-provider-azure/issues/366)) ([1d0ccea](https://github.com/cloudquery/cq-provider-azure/commit/1d0ccea3018fc793800ba5e9798755d473fc049f))

## [0.11.14](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.13...v0.11.14) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.3 ([#361](https://github.com/cloudquery/cq-provider-azure/issues/361)) ([8ff9ead](https://github.com/cloudquery/cq-provider-azure/commit/8ff9ead6f60a1204b373d532071b168fef107dd8))

## [0.11.13](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.12...v0.11.13) (2022-06-26)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.12.2 ([#352](https://github.com/cloudquery/cq-provider-azure/issues/352)) ([f3f36ad](https://github.com/cloudquery/cq-provider-azure/commit/f3f36adeeb481a189009377956b4b1feac911d64))

## [0.11.12](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.11...v0.11.12) (2022-06-23)


### Bug Fixes

* Use errgroup SetLimit ([#351](https://github.com/cloudquery/cq-provider-azure/issues/351)) ([7932589](https://github.com/cloudquery/cq-provider-azure/commit/79325899aca72dfc3b31b3248bb0608e38ac0fbe))

## [0.11.11](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.10...v0.11.11) (2022-06-22)


### Features

* YAML config support ([#353](https://github.com/cloudquery/cq-provider-azure/issues/353)) ([045e92f](https://github.com/cloudquery/cq-provider-azure/commit/045e92fbe9110d9511ed9c5784406431aeea4b10))

## [0.11.10](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.9...v0.11.10) (2022-06-20)


### Features

* Add Azure resources view ([#347](https://github.com/cloudquery/cq-provider-azure/issues/347)) ([2b287af](https://github.com/cloudquery/cq-provider-azure/commit/2b287af1d993eaac53ed5bfaf9eb4e778235928b))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.4 ([#349](https://github.com/cloudquery/cq-provider-azure/issues/349)) ([d5eae61](https://github.com/cloudquery/cq-provider-azure/commit/d5eae61ac9360be394bb992bd9d046229bbe77fa))

## [0.11.9](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.8...v0.11.9) (2022-06-15)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.2 ([#343](https://github.com/cloudquery/cq-provider-azure/issues/343)) ([8688999](https://github.com/cloudquery/cq-provider-azure/commit/86889997f2ac9d50e5ab1628a63ad9e055244591))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.3 ([#345](https://github.com/cloudquery/cq-provider-azure/issues/345)) ([415100c](https://github.com/cloudquery/cq-provider-azure/commit/415100c322a8c068577daba27f0e5af1052a09ff))

## [0.11.8](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.7...v0.11.8) (2022-06-14)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.1 ([#340](https://github.com/cloudquery/cq-provider-azure/issues/340)) ([77c7e50](https://github.com/cloudquery/cq-provider-azure/commit/77c7e50e9209b36b39517066066119c2b38f5574))
* Panic in `security.jit_network_access_policies` ([#342](https://github.com/cloudquery/cq-provider-azure/issues/342)) ([3604559](https://github.com/cloudquery/cq-provider-azure/commit/36045590030e6431573d710a03719b6966968f09))

## [0.11.7](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.6...v0.11.7) (2022-06-14)


### Features

* Added Virtual Network Gateways and Connections ([#331](https://github.com/cloudquery/cq-provider-azure/issues/331)) ([76d335c](https://github.com/cloudquery/cq-provider-azure/commit/76d335cac37e42f134b4b5a2e71baab7deaa1b06))

## [0.11.6](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.5...v0.11.6) (2022-06-08)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.11.0 ([#332](https://github.com/cloudquery/cq-provider-azure/issues/332)) ([dca431d](https://github.com/cloudquery/cq-provider-azure/commit/dca431d6b59441a2738151cc7fb6cf13df2de8f9))

## [0.11.5](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.4...v0.11.5) (2022-06-07)


### Bug Fixes

* Correctly process IPs in ip_address_or_range ([#310](https://github.com/cloudquery/cq-provider-azure/issues/310)) ([1c7afb9](https://github.com/cloudquery/cq-provider-azure/commit/1c7afb93fcefb2d05056fed93bb179343b077aed))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.10 ([#329](https://github.com/cloudquery/cq-provider-azure/issues/329)) ([b57c3ef](https://github.com/cloudquery/cq-provider-azure/commit/b57c3ef78f195d0b09e4d5a6ef8139c3845140fb))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.11 ([#330](https://github.com/cloudquery/cq-provider-azure/issues/330)) ([8afdef5](https://github.com/cloudquery/cq-provider-azure/commit/8afdef541f7c5e52ee05b37a19ce2d3c48c0e816))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.9 ([#327](https://github.com/cloudquery/cq-provider-azure/issues/327)) ([bded97f](https://github.com/cloudquery/cq-provider-azure/commit/bded97f830c739e0d692d200fab429686f19ef2e))

## [0.11.4](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.3...v0.11.4) (2022-06-07)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.8 ([#325](https://github.com/cloudquery/cq-provider-azure/issues/325)) ([14a4c68](https://github.com/cloudquery/cq-provider-azure/commit/14a4c68d03268726bf4dc0fe96b384134551a981))

### [0.11.3](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.2...v0.11.3) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.7 ([#318](https://github.com/cloudquery/cq-provider-azure/issues/318)) ([d071d75](https://github.com/cloudquery/cq-provider-azure/commit/d071d7571ad5958ef051a6ac2282ae31ce0b6719))
* Wrap provider errors ([#320](https://github.com/cloudquery/cq-provider-azure/issues/320)) ([e0fd4de](https://github.com/cloudquery/cq-provider-azure/commit/e0fd4deffbb965181a0c2650c7aac5603468ede1))

### [0.11.2](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.1...v0.11.2) (2022-06-01)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.6 ([#314](https://github.com/cloudquery/cq-provider-azure/issues/314)) ([758929c](https://github.com/cloudquery/cq-provider-azure/commit/758929cbb338efca4174c6e7357aa632df3b499d))

### [0.11.1](https://github.com/cloudquery/cq-provider-azure/compare/v0.11.0...v0.11.1) (2022-05-31)


### Bug Fixes

* Added IgnoreError filter for diagnostic setting resource ([#284](https://github.com/cloudquery/cq-provider-azure/issues/284)) ([e0c4330](https://github.com/cloudquery/cq-provider-azure/commit/e0c43301bde6a05bd0764d49f90a56f64a0abe75))
* Clasify Subscription Not Registered ([#305](https://github.com/cloudquery/cq-provider-azure/issues/305)) ([ac41418](https://github.com/cloudquery/cq-provider-azure/commit/ac414184a6fe9eee9a3649db61da9de2276d8e8c))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.2 ([#300](https://github.com/cloudquery/cq-provider-azure/issues/300)) ([0cc7511](https://github.com/cloudquery/cq-provider-azure/commit/0cc751136459fd8eb77cfb97ed606488729bfe03))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.4 ([#301](https://github.com/cloudquery/cq-provider-azure/issues/301)) ([d7b1d85](https://github.com/cloudquery/cq-provider-azure/commit/d7b1d85a55539ba4561d8bedc7d152d9444d70f8))
* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.5 ([#311](https://github.com/cloudquery/cq-provider-azure/issues/311)) ([78bd6ee](https://github.com/cloudquery/cq-provider-azure/commit/78bd6eed676c0ccfc3553393014b6fd109ef1e30))
* Remove relation tables PK ([#286](https://github.com/cloudquery/cq-provider-azure/issues/286)) ([f5e09e8](https://github.com/cloudquery/cq-provider-azure/commit/f5e09e87d3145a39baa7c3f6cc688ba9a0b11f96))

## [0.11.0](https://github.com/cloudquery/cq-provider-azure/compare/v0.10.4...v0.11.0) (2022-05-24)


### ⚠ BREAKING CHANGES

* Remove migrations (#290)

### Features

* Remove migrations ([#290](https://github.com/cloudquery/cq-provider-azure/issues/290)) ([2bbca07](https://github.com/cloudquery/cq-provider-azure/commit/2bbca07ef5e737cee530e5df9e5444e688a38f33))


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.10.1 ([#296](https://github.com/cloudquery/cq-provider-azure/issues/296)) ([30edd91](https://github.com/cloudquery/cq-provider-azure/commit/30edd91bfbd7be55ed66e19beeddb06b69d8346f))

### [0.10.4](https://github.com/cloudquery/cq-provider-azure/compare/v0.10.3...v0.10.4) (2022-05-23)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.5 ([#292](https://github.com/cloudquery/cq-provider-azure/issues/292)) ([755ba4f](https://github.com/cloudquery/cq-provider-azure/commit/755ba4fd8c59fb781f78a6c8b835b860c0a02d5a))

### [0.10.3](https://github.com/cloudquery/cq-provider-azure/compare/v0.10.2...v0.10.3) (2022-05-17)


### Bug Fixes

* **deps:** Update module github.com/cloudquery/cq-provider-sdk to v0.9.4 ([#281](https://github.com/cloudquery/cq-provider-azure/issues/281)) ([9d9da12](https://github.com/cloudquery/cq-provider-azure/commit/9d9da12d3759761fcbdc935cee0b33f683c8bcf4))

### [0.10.2](https://github.com/cloudquery/cq-provider-azure/compare/v0.10.1...v0.10.2) (2022-05-10)


### Miscellaneous Chores

* Release 0.10.2 ([#270](https://github.com/cloudquery/cq-provider-azure/issues/270)) ([35fd989](https://github.com/cloudquery/cq-provider-azure/commit/35fd989a2d4fd57abcc45c58cab7e2d4e7750e83))

## [v0.3.9] - 2022-01-03
###### SDK Version: 0.6.1
### 🚀 Added
* Add `keyvault.vault` resource back (this result requires special permissions) [#111](https://github.com/cloudquery/cq-provider-azure/pull/111)

## [v0.3.9] - 2022-01-03
###### SDK Version: 0.6.1
### :gear: Changed
* Update to SDK Version [v0.6.1](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md#v061---2021-01-03)
* Remove ad resources (deprecated and will be migrate to msgraph)
### 💥 Breaking Changes
* Renamed table `azure_container_managed_cluster_pip_user_assigned_identity_exceptions` -> `azure_container_managed_cluster_pip_user_assigned_id_exceptions` [#97](https://github.com/cloudquery/cq-provider-azure/pull/97)
### :spider: Fixed
* Fixed disabled migrations [#104](https://github.com/cloudquery/cq-provider-azure/pull/104)

## [v0.3.8] - 2021-11-23
###### SDK Version: 0.4.9

### :rocket: Added
* Added support for ARM binary fixed [#92](https://github.com/cloudquery/cq-provider-azure/pull/92)

### :spider: Fixed
* Fixed names of `azure_network_virtual_network_subnets`, `azure_network_virtual_network_peerings`, `azure_network_virtual_network_ip_allocations` tables according to naming convention [#76](https://github.com/cloudquery/cq-provider-azure/issues/76)
### :gear: Changed
* `azure_network_virtual_network_ip_allocations` is now a string array column of `azure_network_virtual_networks`  


## [v0.3.7] - 2021-10-07
###### SDK Version: 0.4.9

### :gear: Changed
* Upgraded to SDK Version [v0.4.9](https://github.com/cloudquery/cq-provider-sdk/blob/main/CHANGELOG.md)

## [v0.3.6] - 2021-10-03
###### SDK Version: 0.4.7

Base version at which changelog was introduced.

### Supported Resources
- ad.applications
- ad.groups
- ad.service_principals
- ad.users
- authorization.role_assignments
- authorization.role_definitions
- compute.disks
- compute.virtual_machines
- container.managed_clusters
- keyvault.vaults
- monitor.activity_log_alerts
- monitor.activity_logs
- monitor.diagnostic_settings
- monitor.log_profiles
- mysql.servers
- network.public_ip_addresses
- network.security_groups
- network.virtual_networks
- network.watchers
- postgresql.servers
- resources.groups
- resources.policy_assignments
- security.auto_provisioning_settings
- security.contacts
- security.pricings
- security.settings
- sql.servers
- storage.accounts
- subscription.subscriptions
- web.apps
