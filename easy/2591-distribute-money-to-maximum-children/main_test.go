package main

import "testing"

func TestDistMoney(t *testing.T) {
	type testCase struct {
		money     int
		children  int
		childGot8 int
	}

	tests := []testCase{
		// {
		// 	money:     20,
		// 	children:  3,
		// 	childGot8: 1,
		// },
		// {
		// 	money:     16,
		// 	children:  2,
		// 	childGot8: 2,
		// },
		// {
		// 	money:     12,
		// 	children:  2,
		// 	childGot8: 0,
		// },
		{
			money:     8,
			children:  2,
			childGot8: 0,
		},
		{
			money:     12,
			children:  3,
			childGot8: 0,
		},
	}

	for _, v := range tests {
		actual := distMoney(v.money, v.children)
		if actual != v.childGot8 {
			t.Errorf(`
				money: %d, child: %d, 
				actual: %d
				want: %d`, v.money, v.children, actual, v.childGot8)
		}
	}
}
