package main

import "fmt"
import "strings"

func main(){
	fmt.Println("hello this is my first go program!!!");

	var age = 33;
	var initial string = "a"
	z:= "Alex";
	y:= []string{"Matei"}

	

	fmt.Println(age);
	fmt.Println(initial)
	fmt.Println(z);
	fmt.Printf("Variabila z este de tipul %T\n ",z);

	for i:=0;i<len(z);i++ {
		fmt.Println(strings.Join(y, " "));
	}


	var numbers [5]int;
	var numbers2 = [3]int{1,2,3};
	var names =[4]string{"eu", "tu"};

	for i:=0;i<len(numbers2); i++{
		fmt.Println("numarule de pe pozitia %d este %d",i,numbers2[i]);
	}

	for i:=0;i<len(names);i++{
		fmt.Println("stringul de pe pozitia %d este %x",i,names[i]);
	}

	numbers[0]=200;

	for i:= range numbers2{
		fmt.Println("%d", numbers2[i]);
	}



	

}