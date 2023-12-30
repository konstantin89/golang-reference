package main

import (
	"my-package/hello"
	"testing"
)

func TestHello(t *testing.T) {
	//DrawExampleGraph()
	//expected := "Hello, World!"

	hello.HelloPackage()

	//if result != expected {
	//	t.Errorf("Hello() returned %q, expected %q", result, expected)
	//}
}
