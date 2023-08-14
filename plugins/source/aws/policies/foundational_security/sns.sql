\set check_id 'SNS.1'
\echo "Executing check SNS.1"
\ir ../queries/sns/sns_topics_should_be_encrypted_at_rest_using_aws_kms.sql

\set check_id 'SNS.2'
\echo "Executing check SNS.2"
\ir ../queries/sns/sns_topics_should_have_message_delivery_notification_enabled.sql
