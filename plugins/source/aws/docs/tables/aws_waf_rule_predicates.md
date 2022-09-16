
# Table: aws_waf_rule_predicates
This is AWS WAF Classic documentation
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|rule_cq_id|uuid|Unique CloudQuery ID of aws_waf_rules table (FK)|
|data_id|text|A unique identifier for a predicate in a Rule, such as ByteMatchSetId or IPSetId|
|negated|boolean|Set Negated to False if you want AWS WAF to allow, block, or count requests based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, RegexMatchSet, GeoMatchSet, or SizeConstraintSet|
|type|text|The type of predicate in a Rule, such as ByteMatch or IPSet.|
