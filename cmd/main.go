// Sizning hamkasblaringizdan biri noyob main.go faylidan iborat dastur yaratdi
// Kodning barqarorligini yaxshilash uchun sizdan kodni qayta tahrirlashingiz so'raladi (refactor).
// Yangi kod tashkilotini taklif qiling:

// Qaysi paketlarni yaratishingiz kerak?

// Yangi katalog yaratish kerakmi?

// package main

// import "fmt"

// type Users struct{
// 	Name string
// 	Email string
// 	Night uint
// }

// func main() {
// 	// first reservation
// 	var us Users
// 	fmt.Printf("Name: ")
// 	fmt.Scanln(&us.Name)
// 	fmt.Printf("Email: ")
// 	fmt.Scanln(&us.Email)
// 	fmt.Printf("Night: ")
// 	fmt.Scanln(us.Night)
// 	emailContents := getEmailContents("Mr", us.Name, us.Night)
// 	sendEmail(emailContents, us.Email)
// 	createAndSaveInvoice(us.Name, us.Night, 145.32)
// }

// // send an email
// func sendEmail(contents string, to string) {
// 	// ...
// 	// ...
// }

// // prepare email template
// func getEmailContents(title string, name string, nights uint) string {
// 	text := "Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !"
// 	return fmt.Sprintf(text,
// 		title,
// 		name,
// 		nights)
// }

// // create the invoice for the reservation
// func createAndSaveInvoice(name string, nights uint, price float32) {
// 	// ...
// }
package main

import (
	"fmt"
	"os"
	"net/smtp"
)

type User struct {
	Name  string
	Email string
	Night uint
}

func main() {
	user := getUserInput()
	emailContents := getEmailContents("Mr", user.Name, user.Night)
	sendEmail(emailContents, user.Email)
	createAndSaveInvoice(user.Name, user.Night, 145.32)
}

func getUserInput() User {
	var user User
	fmt.Printf("Name: ")
	fmt.Scanln(&user.Name)
	fmt.Printf("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Printf("Night: ")
	fmt.Scanln(&user.Night)
	return user
}

func sendEmail(contents string, to string) {
	from := "your_email@example.com"
	password := "your_password"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Room Reservation Confirmation\n\n" +
		contents

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email sent successfully to:", to)
}

func getEmailContents(title string, name string, nights uint) string {
	text := "Dear %s %s,\nYour room reservation for %d night(s) is confirmed. Have a nice day!"
	return fmt.Sprintf(text, title, name, nights)
}

func createAndSaveInvoice(name string, nights uint, price float32) {
	// Faktura ma'lumotlarini tayyorlash
	invoice := fmt.Sprintf("Invoice for %s:\n"+
		"Name: %s\n"+
		"Night(s): %d\n"+
		"Price per night: $%.2f\n"+
		"Total Price: $%.2f\n",
		name, name, nights, price, float32(nights)*price)

	// Faylga faktura ma'lumotlarini saqlash
	err := saveToFile(name+"_invoice.txt", invoice)
	if err != nil {
		fmt.Println("Failed to save invoice:", err)
		return
	}

	fmt.Println("Invoice saved successfully for:", name)
}

// Faylga ma'lumotlar saqlash
func saveToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
