package main

import (
  "fmt"
  "bitbucket.org/ryde/go-test-models.git"
)

func main() {
  fmt.Println("Starting up...")
  user := gotestmodels.User.NewUser("Sourav Upadhyay", "souravupadhyay", "sourav")
  fmt.Println(user)
}

