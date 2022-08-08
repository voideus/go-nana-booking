package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/voideus/go-nana-booking/helper"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName {
			fmt.Println("First name or Last name must be at least two characters long.")
		}
		if !isValidEmail {
			fmt.Println("Email address is invalid.")
		}
		if !isValidTickets {
			fmt.Printf("Number of tickets is invalid.\n")
		}

		if !isValidName || !isValidTickets || !isValidEmail {
			fmt.Printf("Try again.\n")
			continue
		}
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		// print first names
		firstNames := getFirstNames()
		fmt.Printf("List of ticket holders: %v \n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Conference is booked out.")
			break
		}
	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here:")
	fmt.Println("----------------------")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#######################")
	wg.Done()
}
