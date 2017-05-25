# IsDocker [![Build Status](https://travis-ci.org/dlion/IsDocker.svg?branch=master)](https://travis-ci.org/dlion/IsDocker)
Check if the process is running inside a Docker container

## How to use it

```go
package main

import (
	"fmt"
	. "github.com/dlion/IsDocker"
)

func main() {
	if IsDocker() {
		fmt.Println("Info: Running inside a docker container")
	}
}
```

## Test

```
go test
```

## Author

Domenico Luciani

https://domenicoluciani.com

## License

MIT
