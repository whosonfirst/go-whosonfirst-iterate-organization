package organization

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"
	"testing"

	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"	
)

func TestIterateOrganization(t *testing.T) {

	ctx := context.Background()

	iter_uri := "org://"

	counter := int32(0)
	expected := int32(43)

	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {
		new_count := atomic.AddInt32(&counter, 1)
		fmt.Printf("%s %d\n", path, new_count)
		return nil
	}

	iter, err := iterator.NewIterator(ctx, iter_uri, iter_cb)

	if err != nil {
		t.Fatalf("Failed to create iterator, %v", err)
	}

	err = iter.IterateURIs(ctx, "sfomuseum-data://?prefix=sfomuseum-data-map")

	if err != nil {
		t.Fatalf("Failed to iterate URIs, %v", err)
	}

	if counter != expected {
		t.Fatalf("Unexpected %d count but got: %d", expected, counter)
	}
}
