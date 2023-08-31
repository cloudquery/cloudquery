package bitbucket.client;

import bitbucket.client.model.Repository;
import bitbucket.client.model.Workspace;
import io.cloudquery.messages.WriteMessage;
import io.cloudquery.schema.ClientMeta;
import java.util.ArrayList;
import java.util.List;
import kong.unirest.GenericType;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;

public class BitbucketClient implements ClientMeta {
  public static final String BASE_URL = "https://api.bitbucket.org/2.0";

  private final String userName;
  private final String password;

  public BitbucketClient(String userName, String password) {
    this.userName = userName;
    this.password = password;
  }

  public List<Workspace> listWorkspaces() {
    List<Workspace> workspaces = new ArrayList<>();

    String nextURL = BASE_URL + "/workspaces";
    do {
      HttpResponse<ResponseResult<Workspace>> response =
          Unirest.get(nextURL).basicAuth(userName, password).asObject(new GenericType<>() {});
      ResponseResult<Workspace> body = response.getBody();
      workspaces.addAll(body.getValues());
      nextURL = body.getNext();
    } while (nextURL != null);

    return workspaces;
  }

  public List<Repository> listRepositoriesForWorkspace(String workspace) {
    List<Repository> repositories = new ArrayList<>();

    String nextURL = BASE_URL + "/repositories/" + workspace;
    do {
      HttpResponse<ResponseResult<Repository>> response;
      response =
          Unirest.get(nextURL).basicAuth(userName, password).asObject(new GenericType<>() {});
      ResponseResult<Repository> body = response.getBody();
      repositories.addAll(body.getValues());
      nextURL = body.getNext();
    } while (nextURL != null);

    return repositories;
  }

  @Override
  public String getId() {
    return "bitbucket";
  }

  @Override
  public void write(WriteMessage message) {
    throw new UnsupportedOperationException();
  }
}
