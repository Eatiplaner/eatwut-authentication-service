# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).

### Setup Git hooks:

```bash
pnpm install
pnpm run prepare
```

### Setup Environment:

```bash
cp .env.example .env
```

### Sync Proto:

```bash
brew install yq # For Mac
./scripts/sync-proto
```

### Start the web server:

```bash
revel run myapp
```

### Go to http://localhost:9000/ and you'll see:

    "It works"

### Start the web server with docker:

docker-compose build
docker-compose up

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

## Help

- The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
- The [Revel guides](http://revel.github.io/manual/index.html).
- The [Revel sample apps](http://revel.github.io/examples/index.html).
- The [API documentation](https://godoc.org/github.com/revel/revel).
