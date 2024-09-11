package main
import "fmt"
type Todo struct {
  task string
  isDone bool
  priority int
}
func main() {
  t := Todo{
    "Push git",
    true,
    0,
  }
  fmt.Println(t)
}
