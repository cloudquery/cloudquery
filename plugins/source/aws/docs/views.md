# AWS Resources View

The AWS Resources view, called `aws_resources`, is a view that shows all the resources that are supported by the AWS plugin. It collects all resources into a single view and allows them to be queried by ARN, region, service or tag. With the view in place, selecting resources of different types in a single query becomes easy. For example:

```sql
SELECT * FROM aws_resources
WHERE
  region LIKE 'us-east%'
  AND service = 'ec2'
  AND (type = 'instance' OR type = 'network-interface');
````

The view is not created as part of `cloudquery sync`. Instead, it needs to be created with a SQL query after the sync completes. The SQL query to create the `aws_resources` view in PostgreSQL can be found [on GitHub](https://github.com/cloudquery/cloudquery/blob/main/plugins/source/aws/views/resources.sql). The SQL query may work for other destinations as well (with some tweaking), but this has not been tested. There are [separate instructions for creating the view in Athena](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/views/athena).  
