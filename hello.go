import fmt


func main (){
	fmt.Println("Hello")
	c := add(3,4)
	d := sub(6,2)
}


func add(a int, b int) {
	return a + b
}

func sub(a int, b int) {
	return a - b
}