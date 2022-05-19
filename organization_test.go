package organization

import (
	"context"
	"io"
	_ "fmt"
	"testing"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"sync/atomic"
)

func TestIterateOrganization(t *testing.T) {

	ctx := context.Background()

	iter_uri := "org://"

	counter := int32(0)
	expected := int32(37)
	
	iter_cb := func(ctx context.Context, path string, r io.ReadSeeker, args ...interface{}) error {
		atomic.AddInt32(&counter, 1)
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
		t.Fatalf("Unexpected count: %d", counter)
	}
}
