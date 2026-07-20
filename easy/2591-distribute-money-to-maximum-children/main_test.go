package main

import "testing"

func TestDistMoney(t *testing.T) {
	type input struct {
		money    int
		children int
	}
	type output struct {
		cntChildGot8 int
	}
	type testCase struct {
		in  input
		out output
	}

	tests := []testCase{
		{
			input{
				money:    20,
				children: 3,
			},
			output{
				cntChildGot8: 1,
			},
		},
	}

	for _, v := range tests {
		actual := distMoney(v.in.money, v.in.children)
		if actual != v.out.cntChildGot8 {
			t.Errorf("distMoney(%d, %d) = %d, want %d", v.in.money, v.in.children, actual, v.out.cntChildGot8)
		}
	}
}
