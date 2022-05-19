package organization

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-github/organizations"
	_ "github.com/whosonfirst/go-whosonfirst-iterate-git/v2"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/emitter"
	"github.com/whosonfirst/go-whosonfirst-iterate/v2/iterator"
	"net/url"
)

func init() {
	ctx := context.Background()
	emitter.RegisterEmitter(ctx, "org", NewOrganizationEmitter)
}

// type OrganizationEmitter implements the `emitter.Emitter` interface for iterating multiple repositories in a GitHub organization.
type OrganizationEmitter struct {
	emitter.Emitter
	target string
}

func NewOrganizationEmitter(ctx context.Context, uri string) (emitter.Emitter, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}
	
	em := &OrganizationEmitter{
		target: u.Path,
	}
	
	return em, nil
}

func (em *OrganizationEmitter) WalkURI(ctx context.Context, cb emitter.EmitterCallbackFunc, uri string) error {

	u, err := url.Parse(uri)

	if err != nil {
		return fmt.Errorf("Failed to parse URI, %w", err)
	}

	organization := u.Scheme

	q := u.Query()

	prefix := q["prefix"]
	exclude := q["exclude"]

	access_token := q.Get("access_token")

	list_opts := organizations.NewDefaultListOptions()
	list_opts.Prefix = prefix
	list_opts.Exclude = exclude
	list_opts.AccessToken = access_token

	repos, err := organizations.ListRepos(organization, list_opts)

	if err != nil {
		return fmt.Errorf("Failed to list repos, %w", err)
	}

	iterator_sources := make([]string, len(repos))

	for idx, repo := range repos {
		iterator_sources[idx] = fmt.Sprintf("https://github.com/%s/%s.git", organization, repo)
	}

	iterator_uri := url.URL{}
	iterator_uri.Scheme = "git"
	iterator_uri.Path = em.target

	iter, err := iterator.NewIterator(ctx, iterator_uri.String(), cb)

	if err != nil {
		return fmt.Errorf("Failed to create new iterator, %w", err)
	}

	err = iter.IterateURIs(ctx, iterator_sources...)

	if err != nil {
		return fmt.Errorf("Failed to iterate URIs, %w", err)
	}

	return nil
}
