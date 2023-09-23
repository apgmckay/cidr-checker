package errors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	n := New()
	testCase := fmt.Sprintf("%+v", reflect.TypeOf(n))

	if testCase != "*errors.ErrorStack" {
		t.Logf("Type returned by New() should be of type %s, got: %s\n", "ErrorStack", testCase)
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	n := New()
	n.Push(fmt.Errorf("err 1"))
	if n.Size() != 1 {
		t.Logf("Size of error stack is %d, expected: 1", n.errors)
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	stackSize := 400

	n := New()
	for i := 0; i < stackSize; i++ {
		n.Push(fmt.Errorf("err %d", i))
	}
	for i := stackSize; i == 0; i-- {
		n.Pop()
		if n.Size() != i {
			t.Logf("Size of error stack is %d, expected: %d", n.errors, i)
			t.Fail()
		}
	}
}
