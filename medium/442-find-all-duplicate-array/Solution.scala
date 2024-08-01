import scala.collection.mutable

object Solution {
  def findDuplicates(nums: Array[Int]): List[Int] = {
    val seen = mutable.Set[Int]()
    val duplicates = mutable.ListBuffer[Int]()

    for (num <- nums) {
      if (seen.contains(num)) {
        duplicates += num
      } else {
        seen += num
      }
    }

    duplicates.toList
  }

  def main(args: Array[String]): Unit = {
    val nums = Array(4, 3, 2, 7, 8, 2, 3, 1)
    val duplicates = findDuplicates(nums)
    println(duplicates.mkString(", ")) // Output: 2, 3
  }
}
