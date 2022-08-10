\set check_id 'ElasticBeanstalk.1'
\echo "Executing check ElasticBeanstalk.1"
\ir ../queries/elasticbeanstalk/advanced_health_reporting_enabled.sql

\set check_id 'ElasticBeanstalk.2'
\echo "Executing check ElasticBeanstalk.2"
\ir ../queries/elasticbeanstalk/elastic_beanstalk_managed_updates_enabled.sql
