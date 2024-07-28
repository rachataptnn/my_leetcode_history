object Solution {
  def main(args: Array[String]): Unit = {
    val n = 5
    val edges = Array(
      Array(0, 1, 2), Array(0, 4, 8), Array(1, 2, 3),
      Array(1, 4, 2), Array(2, 3, 1), Array(3, 4, 1)
    )
    val distanceThreshold = 2

    println(findTheCity(n, edges, distanceThreshold))
  }

  def findTheCity(n: Int, edges: Array[Array[Int]], distanceThreshold: Int): Int = {
    val summaryDistances = floydWarshallAlgorithm(edges, n)
    getTheRightCity(summaryDistances, distanceThreshold, n)
  }

  def floydWarshallAlgorithm(edges: Array[Array[Int]], n: Int): Array[Array[Int]] = {
    val summaryDistances = Array.fill(n, n)(10001)
    for (i <- 0 until n) summaryDistances(i)(i) = 0

    for (edge <- edges) {
      val srcNode = edge(0)
      val dstNode = edge(1)
      val distance = edge(2)
      summaryDistances(srcNode)(dstNode) = distance
      summaryDistances(dstNode)(srcNode) = distance
    }

    for {
      k <- 0 until n
      i <- 0 until n
      j <- 0 until n
    } {
      if (summaryDistances(i)(k) + summaryDistances(k)(j) < summaryDistances(i)(j)) {
        summaryDistances(i)(j) = summaryDistances(i)(k) + summaryDistances(k)(j)
      }
    }

    summaryDistances
  }

  def getTheRightCity(dist: Array[Array[Int]], distanceThreshold: Int, n: Int): Int = {
    var minReachableCities = n
    var result = -1

    for (i <- 0 until n) {
      val reachableCities = dist(i).count(_ <= distanceThreshold)
      if (reachableCities <= minReachableCities) {
        minReachableCities = reachableCities
        result = i
      }
    }

    result
  }
}