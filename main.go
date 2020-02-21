package main

import "fmt"

func main() {
  var input int
  var count int
  count = 0

 //get user input for the ammount of times to loop
  fmt.Println("Please enter a number.")
  fmt.Scanln(&input)

  fmt.Println("Well the years start coming")
 
 for input != count{
 fmt.Println("and they don't stop coming")
 count++
 }
}