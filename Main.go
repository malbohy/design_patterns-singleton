package main

import (
	"./Singleton"
	"fmt"
)


func main(){
	firstInstance := Singleton.GetInstance()
	firstCount:= firstInstance.AddOne()

	fmt.Println("count from first instance is " , firstCount)

	secondInstance := Singleton.GetInstance()

	secondCount := secondInstance.AddOne()

	fmt.Println("count from second instance is " , secondCount)

	fmt.Println("count from first instance after increment second instance " , firstInstance)

}