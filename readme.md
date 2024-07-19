# jort-url

<img src="https://pbs.twimg.com/media/CjoPLolUUAAjPR2?format=jpg&name=medium" alt="the jorts" width=35% height=35%>

## what is it?

a url shortening service.

## running tests

```bash
go test ./...
```

please note running tests with -v may show identical output on randomized functions for consecutive runs of go test. this is due to go's test caching, and the cache can be cleared.

```bash
go clean -testcache
```
