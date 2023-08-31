package bitbucket.resources;

import bitbucket.client.BitbucketClient;
import bitbucket.client.model.Repository;
import bitbucket.client.model.Workspace;
import io.cloudquery.schema.Table;
import io.cloudquery.schema.TableResolver;
import io.cloudquery.transformers.TransformWithClass;

public class Repositories {
  public static Table getTable() {
    return Table.builder()
        .name("bitbucket_repositories")
        .description(
            "https://developer.atlassian.com/cloud/bitbucket/rest/api-group-repositories/#api-group-repositories")
        .transform(TransformWithClass.builder(Repository.class).pkField("uuid").build())
        .resolver(resolveResources())
        .build();
  }

  public static TableResolver resolveResources() {
    return (clientMeta, parent, stream) -> {
      BitbucketClient bitbucketClient = (BitbucketClient) clientMeta;

      String workspaceName = ((Workspace) parent.getItem()).getName();

      bitbucketClient.listRepositoriesForWorkspace(workspaceName).forEach(stream::write);
    };
  }
}
