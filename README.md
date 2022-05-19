# go-whosonfirst-iterate-organization

Go package to implement the `whosonfirst/go-whosonfirst-iterate/v2/emitter.Emitter` interface for iterating multiple repositories in a GitHub organization.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-whosonfirst-iterate-organization.svg)](https://pkg.go.dev/github.com/whosonfirst/go-whosonfirst-iterate-organization)

## Important

This is work in progress and may still change. Documentation is incomplete.

## Example

```
package main

import (
	"context"
	"io"
	"fmt"
	"testing"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
)

func main()

	ctx := context.Background()

	iter_uri := "org:///tmp"
	
	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {
		fmt.Println(path)
		return nil
	}

	iter, _ := iterator.NewIterator(ctx, iter_uri, iter_cb)

	iter.IterateURIs(ctx, "sfomuseum-data://?prefix=sfomuseum-data-flights-&exclude=sfomuseum-data-flights-YYYY-MM")
}
```

_Error handling omitted for the sake of brevity._

## See also

* https://github.com/whosonfirst/go-whosonfirst-iterate
* https://github.com/whosonfirst/go-whosonfirst-iterate-git