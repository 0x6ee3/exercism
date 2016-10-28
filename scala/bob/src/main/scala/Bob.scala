
class Bob {
  def hey(line: String) = {
    def answer(line: String) =
      if (line.isEmpty) "Fine. Be that way!"
      else if (line.equals(line.toUpperCase) && line.exists(_.isLetter)) "Whoa, chill out!"
      else if (line.endsWith("?")) "Sure."
      else  "Whatever."

    answer(line.trim)
  }
}

