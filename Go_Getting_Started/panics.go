package main

func main() {
  /* A panic is similiar to an exception in other languages, but in Go when 
   * the application hits a condition that is absolutely cannot proceed with no
   * way to recover, then introduce a panic.
   */
  println("Starting server...")

  // do important things
  panic("Something bad just happened")

  println("Server started")
}

