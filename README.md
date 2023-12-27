# Remote store and binary search gRPC service
This gRPC service is used and called by the [RESTful Gateway project](https://github.com/Rashad-j/gateway-grpc).

This microservice provides three methods: `insert`, `delete`, and `search` to maintain an internal data cache. This service store data in an asc array, where data guaranteed to be in an asc order. It uses binary search algorithm for search, whereas for insert it combines binary search to find the insertion position and array shifting to make room for the new element, resulting in an efficient way to insert elements into a sorted slice. Same pattern is applied on delete.Bin

I used binary as it is a preferred choice for searching in large, sorted datasets where minimizing both time and space requirements is crucial.
Binary search is a time-efficient algorithm with a logarithmic time complexity of `O(log n)`` and a space-efficient algorithm with a constant space complexity of `O(1)``.

## Design patterns
This services applies the decorator design pattern for the search service, achieving by that the second principle of the SOLID: open/closed. Other parts of the code also follow the SOLID principle where single responsibility and interface segregation is used extensively.

## Unit test
Simple unit test cases provided to test different cases for `insert`, `delete` and `search`. 

## Docker
This service is built and pushed to docker hub in order to be used in the RESTful gateway. 

## How to test?
The default config has a port of `8082`. Have a look on the config package. If you want to change this port, add the env variable to a file called `.env`. To run a docker container, make sure to have `.env` file in your main dir and run `make dockerBuildRun`. 

Note, when the service starts, it will generate some random data up to `100_000`, duplicate elements will be dropped when inserting these random numbers. The aim of this bootstrapping step is to enable you play with the service and test it. 

### Test using grpcurl
To test this service you can try calling different methods using `grpcurlSearch`, `grpcurlInsert`, or `grpcurlDelete`. Have a look on the `make` file to learn more.

### Test using the client
Alternatively, you can run the client in `cmd/client`:
```go
$ go run .
```

## Tools and frameworks used
* Zerolog for logging