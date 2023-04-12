## Program for processing files 

<p>This is a command line application written in Golang for processing files and returning data in a standard format.</p>

<p>Currently, this app only supports parsing orders (orders.txt) text files in comma separate value (CSV) format and printing data as JSON.</p>

## How to run this client library

To run this code locally, use the command `go run main.go csv orders.txt`

Alternatively, this code can also be run via [Docker](https://www.docker.com/). To build the docker image use: `docker build -t file-processor . ` 

To run the docker image: `docker run file-processor csv orders.txt`

Finally, to run the tests use: `go test ./...` and to check code coverage: `go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out`

Example input format:
```
A,red,80,20
B,red,120,20
C,red,100,30
D,red,120,10
```

Example output format:
```
{
  "120":
        [
          {"client":"B","quantity":20},
          {"client":"D","quantity":10}
        ],
  "100":
        [
          {"client":"C","quantity":30}
        ],
  "80":
      [
        {"client":"A","quantity":20}
      ]
}
```

## Possible improvements to discuss

- Reduce number of loops required to construct the data for JSON marshalling.
- Find better solution for MarshalJSON function in order.go. I only did this because the ordering of the maps (price) kept changing.
- Make processors more polymorphic. I would like to chose a processor using some kind of generics so they're easier to swap out.