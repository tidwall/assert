package assert

import (
	"bytes"
	"strings"
	"testing"
)

type vv struct {
}

func (v *vv) do() {
	Assert(false)
}

var alt = func() string {
	var buf bytes.Buffer
	stderr = &buf
	exit = func(code int) {}
	Assert(false)
	if !strings.Contains(buf.String(), `function [anonymous]`) {
		panic("missing function")
	}
	return buf.String()
}

func TestAssert(t *testing.T) {
	var buf bytes.Buffer
	stderr = &buf
	exit = func(code int) {}
	var myval bool
	Assert(
		(myval == true ||
			true == false ||
			"h(el)\"" == "hello"))
	if !strings.HasPrefix(buf.String(), "Assertion failed:") {
		t.Fatal("missing failed text")
	}
	if !strings.Contains(buf.String(),
		`((myval == true || true == false || "h(el)\"" == "hello"))`) {
		t.Fatal("missing condition text")
	}
	buf.Reset()
	var v vv
	v.do()
	if !strings.Contains(buf.String(), `function vv.do`) {
		t.Fatal("missing function")
	}
	alt()
}
