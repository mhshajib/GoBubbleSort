package main

import "fmt"

func main() {

	a:= [6]int{5, 1, 6, 2, 4, 3} //Setting values to array

	fmt.Println("Array Before Sorting : ", a)

	for i:=0; i<6; i++ {
	  flag := 0;        		//taking a flag variable
	  for j:=0; j<6-i-1; j++ {
	  	if a[j] > a[j+1] {
			temp := a[j];		//
			a[j] = a[j+1];		//Swapping values
			a[j+1] = temp;		//
			flag = 1;			//setting flag as 1, if swapping occurs
		} 
	  }
	  if flag == 0 {			//breaking out of for loop to reduce extra loops
	    break
	  }
	}
	
	fmt.Println("Sorted Array : ", a)
}