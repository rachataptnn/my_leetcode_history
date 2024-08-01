import scala.collection.mutable

object Solution {

  def findDuplicates(nums: Array[Int]): Array[Int] = {
    val someMap: mutable.Map[Int, Int] = mutable.Map()
    val res = mutable.ListBuffer[Int]()

    // Populate the map with counts
    for (num <- nums) {
      someMap(num) = someMap.getOrElse(num, 0) + 1
    }

    // Find duplicates and add to result
    for ((k, v) <- someMap) {
      if (v > 1) {
        res += k
      }
    }

    res.toArray
  }

  def main(args: Array[String]): Unit = {
    val nums = Array(4, 3, 2, 7, 8, 2, 3, 1)
    val duplicates = findDuplicates(nums)
    println(duplicates.mkString(", ")) // Output: 2, 3
  }
}
