package bitbucket;

import bitbucket.client.BitbucketClient;
import bitbucket.client.configuration.Spec;
import bitbucket.resources.Workspaces;
import io.cloudquery.messages.WriteMessage;
import io.cloudquery.plugin.BackendOptions;
import io.cloudquery.plugin.ClientNotInitializedException;
import io.cloudquery.plugin.NewClientOptions;
import io.cloudquery.plugin.Plugin;
import io.cloudquery.plugin.PluginKind;
import io.cloudquery.plugin.v3.Sync;
import io.cloudquery.scheduler.Scheduler;
import io.cloudquery.schema.ClientMeta;
import io.cloudquery.schema.NullClient;
import io.cloudquery.schema.SchemaException;
import io.cloudquery.schema.Table;
import io.cloudquery.transformers.Tables;
import io.grpc.stub.StreamObserver;
import java.util.List;

public class BitbucketPlugin extends Plugin {
  public static final String PLUGIN_VERSION = "0.0.1"; // {x-release-please-version}

  private Spec spec;
  private List<Table> allTables;

  public BitbucketPlugin() {
    super("bitbucket", PLUGIN_VERSION);
    this.setJsonSchema(Spec.jsonSchema);
    this.setTeam("cloudquery");
    this.setKind(PluginKind.Source);
  }

  @Override
  public ClientMeta newClient(String spec, NewClientOptions options) throws Exception {
    this.allTables = getTables();
    Tables.transformTables(this.allTables);
    for (Table table : this.allTables) {
      table.addCQIDs();
    }

    if (options.isNoConnection()) {
      return new NullClient();
    }

    this.spec = Spec.fromJSON(spec);
    this.spec.validate();

    return new BitbucketClient(this.spec.getUsername(), this.spec.getPassword());
  }

  @Override
  public List<Table> tables(
      List<String> includeList, List<String> skipList, boolean skipDependentTables)
      throws SchemaException, ClientNotInitializedException {
    if (client == null) {
      throw new ClientNotInitializedException();
    }
    return Table.filterDFS(allTables, includeList, skipList, skipDependentTables);
  }

  @Override
  public void sync(
      List<String> includeList,
      List<String> skipList,
      boolean skipDependentTables,
      boolean deterministicCqId,
      BackendOptions backendOptions,
      StreamObserver<Sync.Response> syncStream)
      throws SchemaException, ClientNotInitializedException {
    if (this.client == null) {
      throw new ClientNotInitializedException();
    }

    List<Table> filtered = Table.filterDFS(allTables, includeList, skipList, skipDependentTables);
    Scheduler.builder()
        .client(client)
        .tables(filtered)
        .syncStream(syncStream)
        .deterministicCqId(deterministicCqId)
        .logger(getLogger())
        .concurrency(spec.getConcurrency())
        .build()
        .sync();
  }

  @Override
  public void read() {}

  @Override
  public void write(WriteMessage message) {}

  @Override
  public void close() {}

  private static List<Table> getTables() {
    return List.of(Workspaces.getTable());
  }
}
