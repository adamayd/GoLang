package main

func main() {
  slice := []int{1, 2, 3}
  for i := 0; i < len(slice); i++ {
    println(slice[i])
  }
  for idx, v := range slice {
    println(idx, v)
  }
  wellKnownPorts := map[string]int{"http": 80, "https": 443}
  for key, value := range wellKnownPorts {
    println(key, value)
  }
  for key :=  range wellKnownPorts {
    println(key)
  }
  for _, value := range wellKnownPorts {
    println(value)
  }
}
