
# Table: aws_wafregional_rate_based_rule_match_predicates
Contains one Predicate element for each ByteMatchSet, IPSet, or SqlInjectionMatchSet object that you want to include in a RateBasedRule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rate_based_rule_cq_id|uuid|Unique CloudQuery ID of aws_wafregional_rate_based_rules table (FK)|
|data_id|text|A unique identifier for a predicate in a Rule, such as ByteMatchSetId or IPSetId|
|negated|boolean|Set Negated to False if you want AWS WAF to allow, block, or count requests based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, RegexMatchSet, GeoMatchSet, or SizeConstraintSet|
|type|text|The type of predicate in a Rule, such as ByteMatch or IPSet.  This member is required.|
