package bitbucket;

import io.cloudquery.server.PluginServe;

public class MainClass {

  public static void main(String[] args) {
    BitbucketPlugin plugin = new BitbucketPlugin();
    PluginServe pluginServe = PluginServe.builder().args(args).plugin(plugin).build();
    int exitCode = pluginServe.Serve();
    System.exit(exitCode);
  }
}
