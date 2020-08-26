// Code generated by github.com/jim-minter/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"context"
	"net/http"

	pkg "github.com/jim-minter/go-cosmosdb/example/types"
)

// TODO: we should be doing lots of deep copying in this file - perhaps m should
// be of type map[string][]byte and we should serialise?

type fakePersonClient struct {
	m map[string]*pkg.Person
}

type fakePersonListIterator struct {
	c       *fakePersonClient
	options *Options
}

type fakePersonNotImplementedIterator struct {
}

// NewFakePersonClient returns a new fake person client
func NewFakePersonClient() PersonClient {
	return &fakePersonClient{
		m: map[string]*pkg.Person{},
	}
}

func (c *fakePersonClient) Create(ctx context.Context, partitionkey string, newperson *pkg.Person, options *Options) (*pkg.Person, error) {
	if options != nil {
		return nil, ErrNotImplemented
	}

	if _, ok := c.m[newperson.ID]; ok {
		return nil, &Error{StatusCode: http.StatusConflict} // TODO: check this
	}

	c.m[newperson.ID] = newperson

	return newperson, nil
}

func (c *fakePersonClient) List(options *Options) PersonIterator {
	return &fakePersonListIterator{c: c, options: options}
}

func (c *fakePersonClient) ListAll(ctx context.Context, options *Options) (*pkg.People, error) {
	if options != nil {
		return nil, ErrNotImplemented
	}

	people := &pkg.People{
		Count:     len(c.m),
		People: make([]*pkg.Person, 0, len(c.m)),
	}

	for _, person := range c.m {
		people.People = append(people.People, person)
	}

	return people, nil
}

func (c *fakePersonClient) Get(ctx context.Context, partitionkey, personid string, options *Options) (*pkg.Person, error) {
	if options != nil {
		return nil, ErrNotImplemented
	}

	if _, ok := c.m[personid]; !ok {
		return nil, &Error{StatusCode: http.StatusNotFound}
	}

	return c.m[personid], nil
}

func (c *fakePersonClient) Replace(ctx context.Context, partitionkey string, newperson *pkg.Person, options *Options) (*pkg.Person, error) {
	if options != nil {
		return nil, ErrNotImplemented
	}

	c.m[newperson.ID] = newperson

	return newperson, nil
}

func (c *fakePersonClient) Delete(ctx context.Context, partitionkey string, person *pkg.Person, options *Options) error {
	if options != nil {
		return ErrNotImplemented
	}

	if _, ok := c.m[person.ID]; !ok {
		return &Error{StatusCode: http.StatusNotFound}
	}

	delete(c.m, person.ID)

	return nil
}

func (c *fakePersonClient) Query(partitionkey string, query *Query, options *Options) PersonRawIterator {
	return &fakePersonNotImplementedIterator{}
}

func (c *fakePersonClient) QueryAll(ctx context.Context, partitionkey string, query *Query, options *Options) (*pkg.People, error) {
	return nil, ErrNotImplemented
}

func (c *fakePersonClient) ChangeFeed(options *Options) PersonIterator {
	return &fakePersonNotImplementedIterator{}
}

func (i *fakePersonListIterator) Next(ctx context.Context, maxItemCount int) (*pkg.People, error) {
	if i.options != nil {
		return nil, ErrNotImplemented
	}

	people := &pkg.People{
		Count:     len(i.c.m),
		People: make([]*pkg.Person, 0, len(i.c.m)),
	}

	for _, person := range i.c.m {
		people.People = append(people.People, person)
	}

	return people, nil
}

func (i *fakePersonListIterator) Continuation() string {
	return ""
}

func (i *fakePersonNotImplementedIterator) Next(ctx context.Context, maxItemCount int) (*pkg.People, error) {
	return nil, ErrNotImplemented
}

func (i *fakePersonNotImplementedIterator) NextRaw(context.Context, int, interface{}) error {
	return ErrNotImplemented
}

func (i *fakePersonNotImplementedIterator) Continuation() string {
	return ""
}
