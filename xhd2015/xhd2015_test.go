package xhd2015

import (
	"testing"

	"github.com/xhd2015/xgo/runtime/mock"
)

// xgo test -v ./
func TestFuncMock(t *testing.T) {
	mock.Patch(MyFunc, func() string {
		return "mock func"
	})
	text := MyFunc()
	if text != "mock func" {
		t.Fatalf("expect MyFunc() to be 'mock func', actual: %s", text)
	}
}
