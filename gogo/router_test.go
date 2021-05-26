package gogo

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)

	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/assets/gogo/memem")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/assets/*filepath" {
		t.Fatal("shouldn't match /hello/:name")
	}

	if ps["filepath"] != "gogo/memem" {
		t.Fatal("name should be equal to 'gogo'")
	}

	fmt.Printf("matched path: %s, param['filepath']: %s\n", n.pattern, ps["filepath"])
}
