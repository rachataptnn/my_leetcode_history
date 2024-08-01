object Solution {

    import scala.collection.mutable

    def findDuplicates(nums Array[Int]): Array[Int] = {
        val someMap: mutable.Map[Int, Int] = mutable.Map()
        val res: Array[Int]

        for (num <- nums) {
            someMap[num]++
        }

        for ((k,v) <- someMap) {
            if (v > 1) {
                val res = res :+ k
            }
        }

    }

    def main(args: Array[Int]): Unit = {
        val nums = Array(4, 3, 2, 7, 8, 2, 3, 1)
        val dup = findDuplicates(num)
        println(dup)
    }
}