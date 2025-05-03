package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func AskForOneEntry() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📖 Enter your service name : ")
	service, _ := reader.ReadString('\n')

	fmt.Print("👤 Enter your login : ")
	login, _ := reader.ReadString('\n')

	password := AskPassword("Enter your password : ")

	return Clear(service), Clear(login), Clear(password)
}

func AskPassword(prompt string) string {
	fmt.Print("🔑 " + prompt)
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return string(bytePassword)
}

func Clear(s string) string {
	return strings.TrimSpace(s)
}

func ExitApp() {
	fmt.Println("👋 Bye ! ")
	os.Exit(0)
}
