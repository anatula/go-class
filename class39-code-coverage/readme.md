## class 39 Code coverage
- a powerful tool for getting a feeling for what are your unit tests helping you to bind and how you can improve them
- run `go test` with `-cover` option
- `coverprofile=c.out` `covermode=count` how many times were certain statements hit? not only where but how many times
- we'll see a heatmap of code coverage
- `go tool cover -html=c.out` use in the browser and see it graphically
- code coverage will help you improve testing