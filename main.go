package main

import (
  "fmt"
  "github.com/souravupadhyay-ryde/go-test-models"
)

func main() {
  fmt.Println("Starting up...")
  user := gotestmodels.NewUser("Sourav Upadhyay", "souravupadhyay", "sourav")
  fmt.Println(user)
}

