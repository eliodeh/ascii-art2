package main 

import ("fmt")

type human struct {
   name  string
   age int
   salary int
   occupation string
}

func main() {
var human1 human
var human2 human


human1.name = "Blessing"
human1.age = 25
human1.salary =70000
human1.occupation = "nurse"


human2.name = "Patrick"
human2.age = 38
human2.salary = 150000
human2.occupation = "engineer"

printhuman(human1)

printhuman(human2)

}
func printhuman(hum human) {
	fmt.Println("name:", hum.name)
	fmt.Println("age:", hum.age)
	fmt.Println("salary:", hum.salary)
	fmt.Println("occupation:", hum.occupation)
}

