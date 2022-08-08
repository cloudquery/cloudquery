
# Table: digitalocean_balance
Balance represents a DigitalOcean Balance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|month_to_date_balance|text|Balance as of the `generated_at` time.  This value includes the `account_balance` and `month_to_date_usage`.|
|account_balance|text|Current balance of the customer's most recent billing activity.  Does not reflect `month_to_date_usage`.|
|month_to_date_usage|text|Amount used in the current billing period as of the `generated_at` time.|
|generated_at|timestamp without time zone|The time at which balances were most recently generated.|
