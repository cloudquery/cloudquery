# Creating a New Provider

In this tutorial, you will build a new CloudQuery Provider that will interact with GitHub API.

## Bootstrap New Provider

The best way to bootstrap a new provider is to use the [cq-provider-template](https://github.com/cloudquery/cq-provider-template).

You can either click `Use this template` in GitHub UI or clone the repository and reinitialize the git (as you don't need the template history).

```bash
git clone https://github.com/cloudquery/cq-provider-template cq-provider-github
cd cq-provider-github
rm -rf .git
git init -b main
```

Initial commit for this tutorial is at [cloudquery/cq-provider-github/tree/195f](https://github.com/cloudquery/cq-provider-github/tree/195f13b1d0a4ce7aa6b9b382feeef76c84b1d162).

### Update cq-provider-template reference to your new provider name

There are a few places where you will need to update the template stubs (you can also `grep` or search for `CHANGEME` comment in the code):

`go.mod`, `main.go`

```ini
module github.com/cloudquery/cq-provider-template
# Change to
module github.com/your_org_or_user/cq-provider-github
# we will use github.com/cloudquery/cq-provider-github in this tutorial
```

### Update Provider Name

The provider name is the name you will use when you will call `cloudquery init [provider]`.

Change `resources.go`:

```go
func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "github",
```

### Choose Go API Client

Usually, each provider will use one Go Client to interact with the service. As we need to load the data to relational database, we will go with [google/go-github](https://github.com/google/go-github) that implements all GitHub RestAPIs.

Install `go-github`

```bash
go get github.com/google/go-github/v40
```

### Define Provider Configuration

Each provider defines set of `required` or `optional` arguments that can be passed by a user in the `cloudquery.yml` block.

In our case, to initialize an [authenticated](https://github.com/google/go-github#authentication) GitHub API client you will need an AccessToken provided by the user.

`client/config.go`

```go
type Config struct {
    // ADD THIS LINE:
	GitHubToken string `hcl:"github_token,optional"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources []struct {
		Name  string
		Other map[string]interface{} `hcl:",inline"`
	}
}
func (c Config) Example() string {
	return `
  // Add this line
	// api_key: ${your_env_variable}
  // api_key: static_api_key
`
```

`Config` will be automatically parsed by CloudQuery SDK, you just need to define the name via the `yml` tag. We will also make it optional as we want also to support default environment variables.

The `Example` function returns an example `cloudquery.yml` when you will run `cloudquery init github`. The `configuration` is mostly commented out as the `api_key` is optional (we want to support `GITHUB_TOKEN` env variable by default).

#### Initialize GitHub API Client

Following the example in [go-github](https://github.com/google/go-github#authentication) we need to initialize the client in `client/client.go`

```go
type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

    // Add this line
	GithubClient *github.Client
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
    // Start Snippet
	ctx := context.Background()
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		token = providerConfig.GitHubToken
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
    // End Snippet

	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	client := Client{
		logger: logger,
		// Add this line: pass the initialized third pard client
		GithubClient: githubClient,
	}

	// Return the initialized client and it will be passed to your resources
	return &client, nil
}
```

Configure is called once before starting an operation such as `fetch`. This is usually the place where you need to parse the user `configuration` and initialize the API Client.

In this case we first check if the token is available in `GITHUB_TOKEN` and if not we read what is available in the parsed configuration.

### Adding a Resource

Now we are set to implement our first resource which will extract, transform and load configuration from GitHub to PostgreSQL.

Our first resource will be GitHub organizations which is available via [List API](https://pkg.go.dev/github.com/google/go-github/v41/github#RepositoriesService.List).

For every resource you need to create a new file under `resources/` and implement a function that returns `*schema.Table`.

```go
func Repositories() *schema.Table {
	return &schema.Table{
		Name:     "repositories",
		Resolver: fetchRepositories,
		Options: schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("node_id"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "full_name",
				Type: schema.TypeString,
			},
			...

func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	opts := github.RepositoryListOptions{}
	repositories, response, err := c.GithubClient.Repositories.List(ctx, "cloudquery", &opts)
	if err != nil {
		return err
	}
	_ = response
	res <- repositories
	return nil
}
```

The key things here are the table definition and the `fetchRepositories` function.

Each table defines list of columns and their type. CloudQuery SDK by default reads the value from the corresponding [Repository](https://pkg.go.dev/github.com/google/go-github/v41/github#Repository) `struct` by turning CamelCase fields into `snake_case`. If a custom transformation is needed, you can use `schema.PathResolver`.

To run this in debug mode while you develop, checkout [debugging a provider](https://www.cloudquery.io/docs/developers/debugging)

Congratulations! This is it, you have your first custom provider!
