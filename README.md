# grpccmd

`grpccmd` is a CLI generator for any gRPC service in Go.
While the CLI is written in Go the CLI should be able to talk to any gRPC server in any language.
The grpccmd is implemented as a plugin to protoc.

## Install

To install the protoc plugin binary run:

```
git clone https://github.com/sshaman1101/grpccmd.git
cd grpccmd
make all
```

## Example

To generate code for a CLI run this command.

```
protoc -I proto proto/*.proto --grpccmd_out=proto/
```


Create a main.go file that references the generated code.

```go
package main

import (
    "fmt"
    "os"

    // import grpccmd generated code
    _ "github.com/sonm-io/core/proto"
    // import cobra's wrapper
    "github.com/sshaman1101/grpccmd"
)

func main() {
    if err := grpccmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
```

Make a few calls to the server

```
./cligen --remote=0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD@127.0.0.1:15020 --input=data.json locator resolve
./cligen --remote=0x733193d40B6F03c3da33Dbb2e0e070aCbBf8d91b@127.0.0.1:15095 hub status
```
