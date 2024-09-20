[my public solution on leetcode](https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/solutions/5806929/time-on-space-o1-visualizing-the-input-for-easier-understanding-go-python-javascript/)

# Intuition

This is the `example input` for the best intuition:

```
position := []int{6, 4, 7, 8, 2, 10, 2, 7, 9, 7}
```

However, this input is not the best visualization to understand this problem. We need to transform it into a **Frequency map**.

## Frequency map of Coins on Positions

| 1th | 2nd | 3rd | 4th | 5th | 6th | 7th | 8th | 9th | 10th |
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:----:|
|     |  o  |     |  o  |     |  o  |  o  |  o  |  o  |  o   |
|     |  o  |     |     |     |     |  o  |     |     |      |
|     |     |     |     |     |     |  o  |     |     |      |

## Naive Algorithm
1. Identify the highest stack of coins for both even and odd positions:
    - The `highest even stack` is at the `2nd` position.
    - The `highest odd stack` is at the `7th` position.
2. Next, move the remaining coins to the highest stack.
3. We can calculate the cost of the moves as follows:
    3.1 Suppose we pick the `highest even stack`:
    3.2 When we choose even positions, there is no cost to move coins between even positions (`position[i] + 2 or position[i] - 2 with cost = 0.`)
    3.3 Therefore, we only focus on moving the odd coins (`7th`, `9th`).
    3.4 The **`totalCost`** is 3 + 1 = **4** (3 moves from the `7th` position and 1 move from the `9th` position).

4. Now, let’s repeat the process, but this time pick the `highest odd stack`.
    - The **`totalCost`** is calculated as: 2 + 1 + 1 + 1 + 1 = **6** (2 moves from the `2nd` position, and 1 move each from the `4th`, `6th`, `8th`, and `10th` positions).
5. The minimum cost is: `min(4, 6)`, which equals 4.

> I included the Naive approach first because it closely relates to the problem description.
And it’s a great way to develop problem-solving skills

# Final Algorithm
1. Count the number of `odd coins` and `even coins`.
2. Compare them to determine the minimum.

The result will be the same as the `Naive Algorithm`. 

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code

```golang []
func minCostToMoveChips(position []int) int {
	oddCoins, evenCoins := 0, 0
	for _, v := range position {
		if v%2 != 0 {
			oddCoins++
		} else {
			evenCoins++
		}
	}

	if oddCoins < evenCoins {
		return oddCoins
	}

	return evenCoins
}
```
```python []
class Solution(object):
    def minCostToMoveChips(self, position):
        odd_coins, even_coins = 0, 0
        for v in position:
            if v % 2 != 0:
                odd_coins += 1
            else:
                even_coins += 1
        return min(odd_coins, even_coins)
```
```javascript []
var minCostToMoveChips = function(position) {
    let oddCoins = 0, evenCoins = 0;
    for (let v of position) {
        if (v % 2 !== 0) {
            oddCoins++;
        } else {
            evenCoins++;
        }
    }

    return Math.min(oddCoins, evenCoins);
};

```