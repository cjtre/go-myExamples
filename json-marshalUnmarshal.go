package main

import(
	"fmt"
	"encoding/json"
)

type Hobby struct {
	MyField string
}

type Person struct{
	Fname string
	Lname string
	Age int
	MyHobbies []Hobby
}

func main(){


	p1:= Person{
		Fname: "Chris", Lname: "Tregear", Age: 35, MyHobbies: []Hobby{
			{"Jumping"},
			{"Running"}
		},
	}
	
	p2:= Person{
		Fname: "George", Lname: "Geller", Age: 22,
	}
	p3:= Person{
		Fname: "Penny", Lname: "Polo", Age: 23,
	}
	p4:= Person{
		Fname: "Rachel", Lname: "Rhino", Age: 27,
	}
	p5:= Person{
		Fname: "Roger", Lname: "Rainbow", Age: 33,
	}
	p6:= Person{
		Fname: "Mary", Lname: "Miller", Age: 45,
	}
	p7:= Person{
		Fname: "Lily", Lname: "Flower", Age: 35,
	}
	p8:= Person{
		Fname: "Katie", Lname: "Fleming", Age: 39,
	}

	people := []Person{p1, p2, p3, p4, p5, p6, p7, p8}

	val, err := json.Marshal(people)
	if err != nil{
		fmt.Println(err)
	}

	dataB := string(val)
	fmt.Println(dataB)

	bs := []byte(dataB)

	fmt.Printf("%T\n", dataB)
	fmt.Printf("%T\n", bs)

	var ppl[]Person

	err = json.Unmarshal(bs, &ppl)
	if err != nil{
		 fmt.Println(err)
	}

	fmt.Println(ppl)

	for i, v := range ppl{
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.Fname, v.Lname, v.Age)
	}
}