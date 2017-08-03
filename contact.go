package main

import (
	"fmt"
	"os"
)

type Phone struct {
	Name string
	Family string
	Tel int
	link *Phone
}

var first *Phone = nil
var current *Phone = nil

var text string

func main() {
	fmt.Println("Contact Book")
	var n int
	for n!=9 {
		Menu()
		Select(n)
	}
}
//*************************Menu*************************
func Menu() {
	fmt.Println("1.Add New Phone\n","2.Delete Phone\n",
		"3.Search By Name\n","4.Search By Number\n",
		"5.Sort By Name\n","6.Sort By Family\n",
		"7.Show List\n","8.Erase All Contact\n","9.Exit\n")
}
//*************************Select In Menu*************************
func Select(n int) {
	fmt.Println("Choose A Num In List : ")
	fmt.Scanln(&n)

	switch n {
	case 1:
		AddPhone()
	case 2:
		//DeletePhone()
	case 3:
		SearchName()
	case 4:
		SearchNum()
	case 5:
		SortName()
	case 6:
		SortFamily()
	case 7:
		ShowAll()
	case 8:
		//EraseAll()
	case 9:
		ExitPro()
	default:
		fmt.Println("None Exist Num! Try again")
	}
}
//*****************************Add phone******************************
func AddPhone() {
	var temp1 *Phone = first
	var temp2 *Phone = new(Phone)

	//take name
	fmt.Println("Enter The Name :")
	fmt.Scanln(&temp1.Name)

	//take family name
	fmt.Println("Enter The FamilyName :")
	fmt.Scanln(&temp1.Family)

	//take telephone
	fmt.Println("Enter The TelephoneNum :")
	fmt.Scanln(&temp1.Tel)

	//adding
	temp2.link = nil

	if first == nil {
		first = temp2
	}else {
		for temp1.link != nil {
			temp1 = temp1.link
		}
		temp1.link=temp2
	}
}
//*****************************Delete phone******************************
/*func DeletePhone() {
	fmt.Println("which phone want delete?(by name:1 or tel:2)  :")
	var temp *Phone = new(Phone)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The Name :")
	n,_ := reader1.ReadString('\n')
	//choose
	switch n {
	case "1":
		//take name
		fmt.Println("Enter The Name :")
		fmt.Scanln(&temp.Name)

		//deleting
		if first == nil {
			fmt.Println("empty list")
		}else {
			for current != nil  {
				current = current.link
				if current.Name == temp.Name {
					delete(current)
					current = first
				}else {
					continue
				}
			}
		}*/
	/*case "2":
		//take num
		fmt.Println("Enter The Name :")
		fmt.Scanln(&temp.Tel)

		//deleting
		if first == nil {
			fmt.Println("empty list")
		}else {
			for current != nil  {
				current = current.link
				if current.Tel == temp.Tel {
					delete(current)
					current = first
				}else {
					continue
				}
			}

		}
	default:
		fmt.Println("none exist number! try again ")
	}
}*/
//*****************************Search By Name******************************
func SearchName() *Phone {

	var temp *Phone = new(Phone)

	//take name
	fmt.Println("Enter The Name :")
	fmt.Scanln(&temp.Name)

	//searching
	if first == nil {
		fmt.Println("Empty List")
		return nil
	}else {
		for current != nil && current.Name != temp.Name {
			current = current.link
		}
	}
	if current != nil {
		fmt.Print("Found It\n",current.Name,"\t",
			"\t",current.Family,"\t",current.Tel)
		return current
	}else {
		fmt.Println("404 Not Found")
		current = first
		return nil
	}
}
//*****************************Search By Num******************************
func SearchNum() *Phone {

	var temp *Phone = new(Phone)

	//take num
	fmt.Println("Enter The Name :")
	fmt.Scanln(&temp.Tel)

	//searching
	if first == nil {
		fmt.Println("Empty List")
		return nil
	}else {
		for current != nil && current.Tel != temp.Tel {
			current = current.link
		}
	}
	if current != nil {
		fmt.Print("Found It\n",current.Name,"\t",
			"\t",current.Family,"\t",current.Tel)
		return current
	}else {
		fmt.Println("404 Not Found")
		current = first
		return nil
	}
}
//*****************************Sort By Name******************************
func SortName() {
	var temp1 *Phone = first
	countNode :=1
	if first == nil {
		fmt.Println("Empty List")
	}else {
		for temp1.link != nil {
			temp1 = temp1.link
			countNode++
		}
	}

	var s[100] *Phone
	s[0] = first
	var counter int

	for counter=0 ; counter < countNode ; counter++ {
		s[counter]=s[counter-1].link
	}
	s[counter-1].link=nil

	//checking
	for i:=0;i<counter-1;i++ {
		for j:=0;j<counter-i-1;j++ {
			if s[j].Name > s[j+1].Name {
				s[j+1],s[j] = s[j],s[j+1]
			}
		}
	}
	//show sorted list
}
//*****************************Sort By Family******************************
func SortFamily() {
	var temp1 *Phone = first
	countNode :=1
	if first == nil {
		fmt.Println("Empty List")
	}else {
		for temp1.link != nil {
			temp1 = temp1.link
			countNode++
		}
	}

	var s[100] *Phone
	s[0] = first
	var counter int

	for counter=0 ; counter < countNode ; counter++ {
		s[counter]=s[counter-1].link
	}
	s[counter-1].link=nil

	//checking
	for i:=0;i<counter-1;i++ {
		for j:=0;j<counter-i-1;j++ {
			if s[j].Family > s[j+1].Family {
				s[j+1],s[j] = s[j],s[j+1]
			}
		}
	}
	//show sorted list
}
//*****************************Show All******************************
func ShowAll() {

}
//*****************************Erase All******************************
/*func EraseAll() {
        var temp *Phone = first
	if first == nil {
		fmt.Println("list is empty")
	}else {
		for first != nil {
			temp = first
			first = first.link
			delete(temp)
		}
	}
}*/
//*****************************Exit******************************
func ExitPro() {
	var key string
	fmt.Println("For Exit Press 'Q' ")
	fmt.Scanln(&key)
	if key=="q" {
		os.Exit(0)
	}
}
