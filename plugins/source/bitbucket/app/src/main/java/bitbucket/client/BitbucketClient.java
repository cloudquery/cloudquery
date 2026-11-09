package bitbucket.client;

import bitbucket.client.model.Repository;
import bitbucket.client.model.Workspace;
import com.google.common.net.UrlEscapers;
import io.cloudquery.messages.WriteMessage;
import io.cloudquery.schema.ClientMeta;
import java.net.URI;
import java.util.ArrayList;
import java.util.List;
import kong.unirest.GenericType;
import kong.unirest.HttpResponse;
import kong.unirest.Unirest;

public class BitbucketClient implements ClientMeta {
  public static final String BASE_URL = "https://api.bitbucket.org/2.0";

  private final String userName;
  private final String password;
  private final String baseUrl;
  private final String trustedUrlPrefix;

  public BitbucketClient(String userName, String password) {
    this(userName, password, BASE_URL);
  }

  BitbucketClient(String userName, String password, String baseUrl) {
    this.userName = userName;
    this.password = password;
    this.baseUrl = baseUrl;
    this.trustedUrlPrefix = originOf(baseUrl) + "/";
  }

  public List<Workspace> listWorkspaces() {
    return listAllPages(baseUrl + "/workspaces", new GenericType<>() {});
  }

  public List<Repository> listRepositoriesForWorkspace(String workspaceSlug) {
    String encodedSlug = UrlEscapers.urlPathSegmentEscaper().escape(workspaceSlug);
    return listAllPages(baseUrl + "/repositories/" + encodedSlug, new GenericType<>() {});
  }

  private <T> List<T> listAllPages(String firstPageUrl, GenericType<ResponseResult<T>> pageType) {
    List<T> items = new ArrayList<>();

    String nextURL = firstPageUrl;
    do {
      requireTrustedUrl(nextURL);
      HttpResponse<ResponseResult<T>> response =
          Unirest.get(nextURL).basicAuth(userName, password).asObject(pageType);
      ResponseResult<T> body = response.getBody();
      items.addAll(body.getValues());
      nextURL = body.getNext();
    } while (nextURL != null);

    return items;
  }

  private void requireTrustedUrl(String url) {
    if (!url.startsWith(trustedUrlPrefix)) {
      throw new IllegalStateException(
          "refusing to request a page outside " + trustedUrlPrefix + ": " + url);
    }
  }

  private static String originOf(String url) {
    URI uri = URI.create(url);
    String origin = uri.getScheme() + "://" + uri.getHost();
    return uri.getPort() == -1 ? origin : origin + ":" + uri.getPort();
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
