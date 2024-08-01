# Intuition
&nbsp;&nbsp;&nbsp;&nbsp;Each string has a part that shows the age. We can take this part, turn it into a number, and check if it's over 60.


<br/>

# Approach
- Use for loop through each string in the array
- get the age part 
- convert it to a number
- and check if it's over 60. If it is, increase the count

<br/>

# Complexity
##### Time complexity: $$O(n)$$ `We check each string once.`

##### Space complexity: $$O(1)$$ `Since the function does not allocate additional space that grows with the input size, the space complexity remains O(1)`

<br/>

# Code
```go []
func countSeniors(details []string) int {
	cnt := 0

	for _, person := range details {
		ageInt, _ := strconv.Atoi(person[11:13])
		if ageInt > 60 {
			cnt++
		}
	}

	return cnt
}

```
```scala []
object Solution {
  def countSeniors(details: Array[String]): Int = {
    var cnt = 0
    
    for (person <- details) {
      val ageStr = person.substring(11, 13)
      val ageInt = ageStr.toInt
      if (ageInt > 60) {
        cnt += 1
      }
    }

    cnt
  }
}
```