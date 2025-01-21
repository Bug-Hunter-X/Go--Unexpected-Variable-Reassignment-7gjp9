func main() {



  x := 10

  fmt.Println(x)



  // To avoid reassignment and maintain the original value of x, use a new variable:

  y := 11

  fmt.Println(y)

  fmt.Println(x)

} 

//Alternative solution: If you intend to modify x in place:

//x += 1 //This explicitly modifies the existing value of x 