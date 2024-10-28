package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a, b, expected := 1, 2, 3
	result := Add(a, b)
	if result != expected {
		t.Fatalf("%d add %d eq %d, but expect %d", a, b, result, expected)
	}
	t.Log("test pass")
}

func TestSubtests(t *testing.T) {
	// setup code
	a, b := 100, 10
	t.Run("Add", func(t *testing.T) {
		expected := 110
		result := Add(a, b)
		if result != expected {
			t.Fatalf("%d add %d eq %d, but expect %d", a, b, result, expected)
		}
		t.Log("test pass")
	})
	t.Run("Sub", func(t *testing.T) {
		expected := 90
		result := Sub(a, b)
		if result != expected {
			t.Fatalf("%d sub %d eq %d, but expect %d", a, b, result, expected)
		}
		t.Log("test pass")
	})
	t.Run("Mul", func(t *testing.T) {
		expected := 1000
		result := Mul(a, b)
		if result != expected {
			t.Fatalf("%d mul %d eq %d, but expect %d", a, b, result, expected)
		}
		t.Log("test pass")
	})
	t.Run("Div", func(t *testing.T) {
		expected := 10
		result := Div(a, b)
		if result != expected {
			t.Fatalf("%d div %d eq %d, but expect %d", a, b, result, expected)
		}
		t.Log("test pass")
	})
	// tear-down code
}

func TestTableDriven(t *testing.T) {
	var tests = []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 5, 8},
	}
	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d): got %d, want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

//func TestMain(m *testing.M) {
//	fmt.Println("有些场景需要在测试用例运行之前做一些初始化工作如: 读取配置、建立数据库连接...")
//	m.Run()
//	fmt.Println("有些场景需要在测试用例运行之后做一些清理工作如: 关闭文件、关闭数据连接...")
//}
