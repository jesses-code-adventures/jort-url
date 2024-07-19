# jort-url

<img src="https://pbs.twimg.com/media/CjoPLolUUAAjPR2?format=jpg&name=medium" alt="the jorts" width=35% height=35%>

## what is it?

a url shortening service.

## running locally

clone the repo and run:

```bash
go run main.go
```

running the program will create an sqlite database and start the server on port 8080.

you can verify everything is running correctly by running:

```bash
./scripts/test_endpoints
```

nb: this runs a few bash scripts, so it may not work on windows.

these scripts will hit each of the application's endpoints with some dummy data to ensure everything is working smoothly. you will see errors printed if they occur.

if everything is all gravy, navigate to localhost:8080 on your browser and you should be good to go!

## running tests

```bash
go test ./...
```

please note running tests with -v may show identical output on randomized functions for consecutive runs of go test. this is due to go's test caching, and the cache can be cleared.

```bash
go clean -testcache
```

## limitations

- designed as a demo to be run locally at this point. a live production version would likely move over to postgres rather than using sqlite.
