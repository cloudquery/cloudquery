package bitbucket.client;

import static org.junit.jupiter.api.Assertions.*;

import bitbucket.client.model.Repository;
import bitbucket.client.model.Workspace;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpServer;
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.charset.StandardCharsets;
import java.util.List;
import java.util.concurrent.CopyOnWriteArrayList;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

public class BitbucketClientTest {
  private HttpServer bitbucketStub;
  private HttpServer externalStub;
  private final List<String> bitbucketRequests = new CopyOnWriteArrayList<>();
  private final List<String> bitbucketAuthHeaders = new CopyOnWriteArrayList<>();
  private final List<String> externalRequests = new CopyOnWriteArrayList<>();
  private String nextPageUrl;

  @BeforeEach
  public void startStubs() throws IOException {
    externalStub = HttpServer.create(new InetSocketAddress("127.0.0.1", 0), 0);
    externalStub.createContext(
        "/",
        exchange -> {
          externalRequests.add(exchange.getRequestURI().toString());
          respond(exchange, "{\"values\":[]}");
        });
    externalStub.start();

    bitbucketStub = HttpServer.create(new InetSocketAddress("127.0.0.1", 0), 0);
    bitbucketStub.createContext(
        "/",
        exchange -> {
          bitbucketRequests.add(exchange.getRequestURI().getRawPath());
          bitbucketAuthHeaders.add(exchange.getRequestHeaders().getFirst("Authorization"));
          if (exchange.getRequestURI().getRawPath().endsWith("/page2")) {
            respond(exchange, "{\"values\":[{\"slug\":\"second\"}]}");
            return;
          }
          respond(exchange, "{\"values\":[{\"slug\":\"first\"}],\"next\":\"" + nextPageUrl + "\"}");
        });
    bitbucketStub.start();
  }

  @AfterEach
  public void stopStubs() {
    bitbucketStub.stop(0);
    externalStub.stop(0);
  }

  private static void respond(HttpExchange exchange, String body) throws IOException {
    byte[] bytes = body.getBytes(StandardCharsets.UTF_8);
    exchange.getResponseHeaders().add("Content-Type", "application/json");
    exchange.sendResponseHeaders(200, bytes.length);
    exchange.getResponseBody().write(bytes);
    exchange.close();
  }

  private String bitbucketBaseUrl() {
    return "http://127.0.0.1:" + bitbucketStub.getAddress().getPort() + "/2.0";
  }

  private String externalUrl() {
    return "http://127.0.0.1:" + externalStub.getAddress().getPort() + "/2.0/workspaces?page=2";
  }

  @Test
  public void followsNextPageOnTheSameHostWithCredentials() {
    nextPageUrl = bitbucketBaseUrl() + "/workspaces/page2";
    BitbucketClient client = new BitbucketClient("user", "app-password", bitbucketBaseUrl());

    List<Workspace> workspaces = client.listWorkspaces();

    assertEquals(List.of("first", "second"), workspaces.stream().map(Workspace::getSlug).toList());
    assertEquals(2, bitbucketAuthHeaders.size());
    bitbucketAuthHeaders.forEach(Assertions::assertNotNull);
  }

  @Test
  public void doesNotSendCredentialsToAnExternalNextPage() {
    nextPageUrl = externalUrl();
    BitbucketClient client = new BitbucketClient("user", "app-password", bitbucketBaseUrl());

    assertThrows(IllegalStateException.class, client::listWorkspaces);
    assertTrue(externalRequests.isEmpty());
  }

  @Test
  public void doesNotSendCredentialsToAnExternalRepositoriesNextPage() {
    nextPageUrl = externalUrl();
    BitbucketClient client = new BitbucketClient("user", "app-password", bitbucketBaseUrl());

    assertThrows(IllegalStateException.class, () -> client.listRepositoriesForWorkspace("ws"));
    assertTrue(externalRequests.isEmpty());
  }

  @Test
  public void encodesTheWorkspaceSlugAsOnePathSegment() {
    nextPageUrl = bitbucketBaseUrl() + "/repositories/page2";
    BitbucketClient client = new BitbucketClient("user", "app-password", bitbucketBaseUrl());

    List<Repository> repositories = client.listRepositoriesForWorkspace("a/../b c");

    assertEquals("/2.0/repositories/a%2F..%2Fb%20c", bitbucketRequests.get(0));
    assertEquals(2, repositories.size());
  }
}
