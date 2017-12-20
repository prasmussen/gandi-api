package record

import (
	"fmt"
	"strings"

	"github.com/prasmussen/gandi-api/client"
)

// Record holds the zone client structure
type Record struct {
	*client.Client
	Prefix string
}

// New instanciates a new instance of a Zone client
func New(c *client.Client, prefix string) *Record {
	return &Record{c, prefix}
}

func (r *Record) uri(pattern string, paths ...string) string {
	args := make([]interface{}, len(paths))
	for i, v := range paths {
		args[i] = v
	}
	return fmt.Sprintf("%s/%s",
		strings.TrimRight(r.Prefix, "/"),
		strings.TrimLeft(fmt.Sprintf(pattern, args...), "/"))
}

func (r *Record) formatCallError(function string, args ...string) error {
	format := "unexpected arguments for function %s." +
		" supported calls are: %s(), %s(<Name>), %s(<Name>, <Type>)" +
		" %s called with"
	a := []interface{}{
		function,
		function, function, function,
		function,
	}
	for _, v := range args {
		format = format + " %s"
		a = append(a, v)
	}
	return fmt.Errorf(format, a...)
}

func (r *Record) Create(recordInfo RecordInfo, args ...string) (status *Status, err error) {
	switch len(args) {
	case 0:
		_, err = r.Post(r.uri("/records"), recordInfo, &status)
	case 1:
		_, err = r.Post(r.uri("/records/%s", args...), recordInfo, &status)
	case 2:
		_, err = r.Post(r.uri("/records/%s/%s", args...), recordInfo, &status)
	default:
		err = r.formatCallError("Create", args...)
	}
	return
}

func (r *Record) Update(recordInfo RecordInfo, args ...string) (status *Status, err error) {
	switch len(args) {
	case 0:
		_, err = r.Put(r.uri("/records"), recordInfo, &status)
	case 1:
		_, err = r.Put(r.uri("/records/%s", args...), recordInfo, &status)
	case 2:
		_, err = r.Put(r.uri("/records/%s/%s", args...), recordInfo, &status)
	default:
		err = r.formatCallError("Update", args...)
	}
	return
}

func (r *Record) List(args ...string) (list []*RecordInfo, err error) {
	switch len(args) {
	case 0:
		_, err = r.Get(r.uri("/records"), &list)
	case 1:
		_, err = r.Get(r.uri("/records/%s", args...), &list)
	case 2:
		_, err = r.Get(r.uri("/records/%s/%s", args...), &list)
	default:
		err = r.formatCallError("List", args...)
	}
	return
}

// Delete deletes records matching the
func (r *Record) Delete(args ...string) (err error) {
	switch len(args) {
	case 0:
		_, err = r.Client.Delete(r.uri("/records"), nil)
	case 1:
		_, err = r.Client.Delete(r.uri("/records/%s", args...), nil)
	case 2:
		_, err = r.Client.Delete(r.uri("/records/%s/%s", args...), nil)
	default:
		err = r.formatCallError("Delete", args...)
	}
	return
}
