package hello

import "testing"

/**
 * go 语言测试文件的编写
 * 1.源文件以_test.go结尾：xxx_test.go
 * 2.测试方法名以Test开头：func TestXXX(t *testing.T){...}
 */

func TestHello(t *testing.T) {
	t.Log("hello world")
}
