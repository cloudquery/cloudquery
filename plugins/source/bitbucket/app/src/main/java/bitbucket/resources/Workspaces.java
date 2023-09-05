package bitbucket.resources;

import bitbucket.client.BitbucketClient;
import bitbucket.client.model.Workspace;
import io.cloudquery.schema.Table;
import io.cloudquery.schema.TableResolver;
import io.cloudquery.transformers.TransformWithClass;
import java.util.List;

public class Workspaces {
  public static Table getTable() {
    return Table.builder()
        .name("bitbucket_workspaces")
        .description(
            "https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-group-workspaces")
        .transform(TransformWithClass.builder(Workspace.class).pkField("uuid").build())
        .relations(List.of(Repositories.getTable()))
        .resolver(resolveResources())
        .build();
  }

  public static TableResolver resolveResources() {
    return (clientMeta, parent, stream) -> {
      List<Workspace> workspaces = ((BitbucketClient) clientMeta).listWorkspaces();

      workspaces.forEach(stream::write);
    };
  }
}
