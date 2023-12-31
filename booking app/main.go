package main

import ( 
	"fmt"
	"strings"
)

func main()  {

	conferenceName := "Go Conference" // alternative declaration for variables
	const conferenceTickets = 50
	var remainingTickets = 50
	bookings := []string{}


	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter no of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)


			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation mail at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else {

			if !isValidName{
				fmt.Println("first name or last name you entered is too short.")
			} 
			
			if !isValidEmail {
				fmt.Println("The email you entered is invalid.")
			} 
			
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

		
	}

	
}
