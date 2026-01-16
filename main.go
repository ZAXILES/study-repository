package main

import (
	"awesomeProject/feature1"
	"awesomeProject/feature2"
	"awesomeProject/feature_postrges/simple_connection"
	"fmt"
)

func main() {
	fmt.Println("Hello git")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.TestDb()

}
