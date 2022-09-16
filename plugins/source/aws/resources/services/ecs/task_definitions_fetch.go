package ecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type TaskDefinitionWrapper struct {
	*types.TaskDefinition
	Tags []types.Tag
}

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listEcsTaskDefinitions, ecsTaskDefinitionDetail)
}

func ecsTaskDefinitionDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, detail interface{}) {
	c := meta.(*client.Client)
	svc := c.Services().ECS
	taskArn := detail.(string)
	describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	})
	if err != nil {
		errorChan <- err
		return
	}
	if describeTaskDefinitionOutput.TaskDefinition == nil {
		return
	}
	resultsChan <- TaskDefinitionWrapper{
		TaskDefinition: describeTaskDefinitionOutput.TaskDefinition,
		Tags:           describeTaskDefinitionOutput.Tags,
	}
}

func listEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, res chan<- interface{}) error {
	var config ecs.ListTaskDefinitionsInput
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().ECS
	for {
		listClustersOutput, err := svc.ListTaskDefinitions(ctx, &config, func(o *ecs.Options) {
			o.Region = region
		})
		if err != nil {
			return err
		}
		for _, taskDefinitionArn := range listClustersOutput.TaskDefinitionArns {
			res <- taskDefinitionArn
		}
		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}

func resolveEcsTaskDefinitionsInferenceAccelerators(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]interface{}{}
	for _, a := range r.InferenceAccelerators {
		if a.DeviceName == nil {
			continue
		}
		j[*a.DeviceName] = aws.ToString(a.DeviceType)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionsPlacementConstraints(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]interface{}{}
	for _, p := range r.PlacementConstraints {
		if p.Expression == nil {
			continue
		}
		j[*p.Expression] = p.Type
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionsProxyConfigurationProperties(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]interface{}{}
	if r.ProxyConfiguration == nil {
		return nil
	}
	for _, p := range r.ProxyConfiguration.Properties {
		if p.Name == nil {
			continue
		}
		j[*p.Name] = aws.ToString(p.Value)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsDependsOn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, p := range r.DependsOn {
		if p.ContainerName == nil {
			continue
		}
		j[*p.ContainerName] = string(p.Condition)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsEnvironment(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, p := range r.Environment {
		if p.Name == nil {
			continue
		}
		j[*p.Name] = aws.ToString(p.Value)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsEnvironmentFiles(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, p := range r.EnvironmentFiles {
		j[string(p.Type)] = aws.ToString(p.Value)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsExtraHosts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]interface{}{}
	for _, h := range r.ExtraHosts {
		if h.Hostname == nil {
			continue
		}
		j[*h.Hostname] = h.IpAddress
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]interface{}{}
	if r.LogConfiguration == nil {
		return nil
	}
	for _, s := range r.LogConfiguration.SecretOptions {
		j[*s.Name] = *s.ValueFrom
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsResourceRequirements(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, s := range r.ResourceRequirements {
		j[string(s.Type)] = aws.ToString(s.Value)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsSecrets(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, s := range r.Secrets {
		if s.Name == nil {
			continue
		}
		j[*s.Name] = aws.ToString(s.ValueFrom)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsSystemControls(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]string{}
	for _, s := range r.SystemControls {
		if s.Namespace == nil {
			continue
		}
		j[*s.Namespace] = aws.ToString(s.Value)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionContainerDefinitionsVolumesFrom(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ContainerDefinition)
	j := map[string]interface{}{}
	for _, s := range r.VolumesFrom {
		if s.SourceContainer == nil {
			continue
		}
		j[*s.SourceContainer] = aws.ToBool(s.ReadOnly)
	}
	return resource.Set(c.Name, j)
}
func resolveEcsTaskDefinitionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(TaskDefinitionWrapper)
	j := map[string]string{}
	for _, a := range r.Tags {
		if a.Key == nil {
			continue
		}
		j[*a.Key] = aws.ToString(a.Value)
	}
	return resource.Set(c.Name, j)
}
