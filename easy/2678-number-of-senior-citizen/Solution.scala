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

  def main(args: Array[String]): Unit = {
    val details = Array("9751302862F0693", "3888560693F7262", "5485983835F0649", "2580974299F6042", "9976672161M6561", "0234451011F8013", "4294552179O6482")
    println(countSeniors(details))
  }
}
