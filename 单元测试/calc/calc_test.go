package calc

import (
	"testing"
)

//测试命令 go test -run TestAdd
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3{
		t.Errorf("1 + 2 expected be 3, but %d got", ans)  //只会打印错误信息，不会停止执行
	}
}

func TestMul(t *testing.T) {
	if ans := Mul(2, 3); ans != 6{
		t.Errorf("2 * 3 expected be 6, bug %d got", ans)
	}
}

//测试命令 go test -run TestMul2/pos -v
func TestMul2(t *testing.T) {
	t.Run("pos", func(t *testing.T){
		if Mul(2, 3) != 6{
			t.Fatal("fail")
		}
	})
	t.Run("neg", func(t *testing.T){
		if Mul(2, -3) != -6{
			t.Fatal("fail")
		}
	})
}

//测试命令 go test -run TestMul3 -v
func TestMul3(t *testing.T) {
	cases := []struct{
		Name string
		A, B,  Expected int
	}{
		{"pos", 2 ,3 ,6},
		{"neg", 2 ,-3, -7},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases{
		t.Run(c.Name, func(t *testing.T){
			if ans := Mul(c.A, c.B); ans != c.Expected{
				t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
			}
		})
	}
}


type calcCase struct{A, B, Expected int}

func createMulTestCase(t *testing.T, c *calcCase){
	t.Helper()      //使用helper函数可以打印出报错的行数
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
	}
}

func TestMul4(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, 0, 1})
}



