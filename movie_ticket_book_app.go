package main
import (
"fmt"
"strings"

)

func main(){



 var conference_name="Go Movies Ticket Application";
 const conference_tickets uint=50;
 var remain_tickets uint=50;

 var bookings []string

 fmt.Println("")

 fmt.Printf("Welcome to  %v Movie Booking Application \n", conference_name);
 fmt.Println("")
 
 fmt.Printf("We have total of %v and %v Available.", conference_tickets, remain_tickets);	
 fmt.Println("")
 
 fmt.Println("Get your tickets here to attend")
 fmt.Println("")


 for {
	

	//var user_name string
	//Ask for user input
	var first_name string
	var last_name string
	var email string 
	var user_tickets uint
	bookings =append(bookings, first_name + " " + last_name)

	//user_name= first_name+ " " + "  "+ last_name

	fmt.Println("Please Enter your first name:")
	fmt.Scan(&first_name);

	//if (first_name== " " | first_name == int(number)){
		//fmt.Println("Please Enter only alphabets")
	
	fmt.Println("")

	fmt.Println("Please Enter your last name:")
	fmt.Scan(&last_name);
	fmt.Println("")

	fmt.Println("Please Enter your email:")
	fmt.Scan(&email);
	fmt.Println("")

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&user_tickets)
	fmt.Println("")

	if user_tickets <= remain_tickets{
		
		continue
	
	} 
	
	remain_tickets=conference_tickets - user_tickets
	

	bookings[0]=first_name  + " " + last_name
	fmt.Printf("The whole array : %v\n", bookings)
	fmt.Println("")

	fmt.Printf("The first value is: %v\n", bookings[0])
	fmt.Println("")

	fmt.Printf("Array type is : %T\n", bookings)
	fmt.Println("")

	fmt.Printf("Array length is : %v\n", len(bookings))



	fmt.Printf("Thank you  %v  %v for booking %v tickets. You will receive a confirmation email at %v\n", first_name, last_name, user_tickets, email);
	fmt.Println("")
	fmt.Printf("The remaining tickets is:  %v", remain_tickets)

	first_names :=[] string{}

	for _, booking := range bookings{
		var names= strings.Fields(booking) 
		 first_names=append(first_names, names[0])

	}
	fmt.Printf("The first names of bookings are %v\n: ", first_names)
	fmt.Println("")

	if remain_tickets == 0{
		//Bye bye!
		fmt.Println("The movie tickets has finished! See you Week end")

		break
	}

	}




}