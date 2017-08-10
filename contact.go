package main

import (
	"fmt"
	"os"
	"sort"
	"log"
)

type Phone struct {
	Name string
	Family string
	Tell string
}

var obj []*Phone

type ByName []*Phone
func (n ByName) Len() int { return len(n)}
func (n ByName) Swap(i,j int) { n[i],n[j]= n[j],n[i]}
func (n ByName) Less(i,j int) bool { return n[i].Name < n[j].Name}

type ByFamily []*Phone
func (f ByFamily) Len() int { return len(f)}
func (f ByFamily) Swap(i,j int) { f[i],f[j]= f[j],f[i]}
func (f ByFamily) Less(i,j int) bool { return f[i].Family < f[j].Family}


type ByTell []*Phone
func (t ByTell) Len() int { return len(t)}
func (t ByTell) Swap(i,j int) { t[i],t[j]= t[j],t[i]}
func (t ByTell) Less(i,j int) bool {return t[i].Tell < t[j].Tell}


func main() {
	file, err := os.Create("input.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintln(file,"Contact Book")
	var n int
	for n!=7 {
		Menu()
		Select(n)
	}
}
//*************************Menu*************************
func Menu() {
	fmt.Println("1.Add New Phone\n","2.Delete Phone\n",
		"3.Search\n","4.Sort\n","5.Show List\n",
		"6.Erase All Contact\n","7.Exit\n")
}
//*************************Select In Menu*************************
func Select(n int) {
	fmt.Println("Choose A Num In List : ")
	fmt.Scanln(&n)

	switch n {
	case 1:
		AddPhone()
	case 2:
		DeletePhone()
	case 3:
		Search()
	case 4:
		Sort()
	case 5:
		ShowAll()
	case 6:
		EraseAll()
	case 7:
		ExitPro()
	default:
		fmt.Println("None Exist Number! Try Again PLZ\n")
	}
}
//*****************************Add phone******************************
func AddPhone() {
	var temp *Phone=new(Phone)

	//take name
	fmt.Println("Enter The Name :")
	fmt.Scanln(&temp.Name)

	//take family name
	fmt.Println("Enter The FamilyName :")
	fmt.Scanln(&temp.Family)

	//take telephone
	fmt.Println("Enter The TelephoneNum :")
        fmt.Scanln(&temp.Tell)

	//adding
	obj=append(obj,temp)
	fmt.Println()
}
//*****************************Delete phone******************************
func DeletePhone() {

	//	var temp *Phone = new(Phone)
	fmt.Println("which phone want delete?:\n",
		"First Search Ur Contact")
	/*out:=Search()
	obj=obj[0:out]
	obj=obj[out:]*/
}
//*****************************Search******************************
func Search() int {

	var temp *Phone = new(Phone)

	fmt.Println("1.Search By Name\n","2.Search By FamilyName\n",
		"3.Search By PhoneNumber")

	var n int
	fmt.Scanln(&n)
	switch n {
	case 1:
		//take name
		fmt.Println("Enter The Name :")
		fmt.Scanln(&temp.Name)
		for i:=0;i<len(obj);i++ {
			if obj[i].Name==temp.Name {
				fmt.Println(i+1,":",obj[i].Name,"\t\t",
					obj[i].Family,"\t\t",obj[i].Tell,"\n")
				return i
			}else {
				fmt.Println("404 Not Found\n")
			}
		}
	case 2:
		//take familyname
		fmt.Println("Enter The FamilyName :")
		fmt.Scanln(&temp.Family)
		for i:=0;i<len(obj);i++ {
			if obj[i].Family==temp.Family {
				fmt.Println(i+1,":",obj[i].Name,"\t\t",
					obj[i].Family,"\t\t",obj[i].Tell,"\n")
				return i
			}else {
				fmt.Println("404 Not Found\n")
			}
		}
	case 3:
		//take phonenumber
		fmt.Println("Enter The PhoneNumber :")
		fmt.Scanln(&temp.Tell)
		for i:=0;i<len(obj);i++ {
			if obj[i].Tell==temp.Tell {
				fmt.Println(i+1,":",obj[i].Name,"\t\t",
					obj[i].Family,"\t\t",obj[i].Tell,"\n")
				return i
			}else {
				fmt.Println("404 Not Found\n")
			}
		}
	default:
		fmt.Println("Non Exist Number!Try Again PLZ\n")
	}
	return 1
}
//*****************************Sort By Name******************************
func Sort() {
	fmt.Println("1.Sort By Name\n","2.Sort By Family\n",
		"3.Sort By PhoneNumber")

	var n int
	fmt.Scanln(&n)
	switch n {
	case 1 :
		for i:=0;i<len(obj);i++ {
			sort.Sort(ByName(obj))
			fmt.Println(i+1,":",obj[i].Name,"\t\t",obj[i].Family,
			"\t\t",obj[i].Tell,"\n")
		}
	case 2:
		for i:=0;i<len(obj);i++ {
			sort.Sort(ByFamily(obj))
			fmt.Println(i+1,":",obj[i].Name,"\t\t",obj[i].Family,
				"\t\t",obj[i].Tell,"\n")
		}
	case 3:
		for i:=0;i<len(obj);i++ {
			sort.Sort(ByTell(obj))
			fmt.Println(i+1,":",obj[i].Name,"\t\t",obj[i].Family,
				"\t\t",obj[i].Tell,"\n")
		}
	default:
		fmt.Println("Non Exist Number! Try Again PLZ")
	}
}
//*****************************Show All******************************
func ShowAll() {

	fmt.Println("Name \t\t FamilyName \t\t PhoneNumber")
	fmt.Println("....... \t ............ \t\t .............")

	for i:=0;i<len(obj);i++ {
		fmt.Println(i+1,":",obj[i].Name,"\t\t",obj[i].Family,
			"\t\t",obj[i].Tell,"\n")
		}
}
//*****************************Erase All******************************
func EraseAll() {
	if obj[0:len(obj)]==nil {
		fmt.Println("Empty List")
	}else {
		for i:=0;i<len(obj);i++ {
			obj[i]=nil
		}
		fmt.Println("Cleared List\n")
	}
}
//*****************************Exit******************************
func ExitPro() {
	var key string
	fmt.Println("For Exit Press 'Q' ")
	fmt.Scanln(&key)
	if key=="q" {
		os.Exit(0)
	}
}
