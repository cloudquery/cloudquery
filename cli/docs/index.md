# Cloudquery

This is the CLI that interacts with providers and database 


## Download

You can download pre-compiled binaries from Github:

``` shell
export OS=Darwin # Possible values: Linux, Windows, Darwin
export ARCH=x86_64 # Possible values x86_64, arm64
curl -L https://github.com/cloudquery/cloudquery/releases/latest/download/cloudquery_${OS}_${ARCH} -o cloudquery
chmod a+x cloudquery
```

Or you can build from source:

```
git clone https://github.com/cloudquery/cloudquery.git
cd cloudquery
```


## Setup

```shell
cloudquery init aws
```

or 

``` shell
go run main.go init aws
```