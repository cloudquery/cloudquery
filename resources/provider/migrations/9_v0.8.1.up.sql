ALTER TABLE IF EXISTS "aws_cloudfront_distribution_default_cache_behavior_lambda_functions"
    RENAME TO "aws_cloudfront_distribution_default_cache_behavior_functions";


INSERT INTO aws_cloudfront_distribution_default_cache_behavior_functions (cq_id, distribution_cq_id, event_type,
                                                                          lambda_function_arn, include_body)
SELECT cq_id, distribution_cq_id, event_type, lambda_function_arn, include_body
FROM aws_cloudfront_distribution_default_behaviour_lambda_functions;
