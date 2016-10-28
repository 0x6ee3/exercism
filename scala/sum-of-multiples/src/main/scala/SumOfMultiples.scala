object SumOfMultiples {
  def sumOfMultiples(params: List[Int], bound: Int) = {
    params.flatMap((p: Int) => {
      for {
        x <- 1 to bound
        z = p * x
        if z < bound
      } yield z
    }).toSet.sum
  }
}
