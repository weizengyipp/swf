package engine

import (
	"testing"
)

func TestGroup(t *testing.T) {
	g := New()
	v1 := g.Group("/v1")
	v2 := g.Group("/v2")
	v3 := g.Group("/v3")
	if v1.prefix != "/v1" {
		t.Fatal("prefix should be /v1")
	}
	if v2.prefix != "/v2" {
		t.Fatal("prefix should be /v1/v2, got", v2.prefix)
	}
	if v3.prefix != "/v3" {
		t.Fatal("prefix should be /v3,got", v3.prefix)
	}

}
