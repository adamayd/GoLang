package main

func main() {
  // Standard Loop
  var i int
  for i < 5 {
    println("i:", i)
    i++
    if i == 2 {
      continue
    }
    println("continuing...")
    if i == 3 {
      break
    }
  }
  // Post Clause Loop
  for j :=0; j < 5; j++ {
    println("j:", j)
  }
  println("i:", i, "j: is undeclared outside of the for loop")
  // Loop set up for infinite loop
  var k int
  // for ; ; { // ugly syntax, same as below
  for {
    if k == 5 {
      break
    }
    println("k:", k)
    k++
  }
}
