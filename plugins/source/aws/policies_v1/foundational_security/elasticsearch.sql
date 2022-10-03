\set check_id 'Elasticsearch.1'
\echo "Executing check Elasticsearch.1"
\ir ../queries/elasticsearch/elasticsearch_domains_should_have_encryption_at_rest_enabled.sql

\set check_id 'Elasticsearch.2'
\echo "Executing check Elasticsearch.2"
\ir ../queries/elasticsearch/elasticsearch_domains_should_be_in_vpc.sql

\set check_id 'Elasticsearch.3'
\echo "Executing check Elasticsearch.3"
\ir ../queries/elasticsearch/elasticsearch_domains_should_encrypt_data_sent_between_nodes.sql

\set check_id 'Elasticsearch.4'
\echo "Executing check Elasticsearch.4"
\ir ../queries/elasticsearch/elasticsearch_domain_error_logging_to_cloudwatch_logs_should_be_enabled.sql

\set check_id 'Elasticsearch.5'
\echo "Executing check Elasticsearch.5"
\ir ../queries/elasticsearch/elasticsearch_domains_should_have_audit_logging_enabled.sql

\set check_id 'Elasticsearch.6'
\echo "Executing check Elasticsearch.6"
\ir ../queries/elasticsearch/elasticsearch_domains_should_have_at_least_three_data_nodes.sql

\set check_id 'Elasticsearch.7'
\echo "Executing check Elasticsearch.7"
\ir ../queries/elasticsearch/elasticsearch_domains_should_be_configured_with_at_least_three_dedicated_master_nodes.sql

\set check_id 'Elasticsearch.8'
\echo "Executing check Elasticsearch.8"
\ir ../queries/elasticsearch/connections_to_elasticsearch_domains_should_be_encrypted_using_tls_1_2.sql
