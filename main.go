//I wanted to create a little reverse shell maker
//playing a little with the "fmt" package
//have fun reading my code lel
package main

import (
	"fmt"
	"log"
	"net"
)

/*func lstOfRevShells() {
const bashS =
const phpS =
const javaS =
const perlS =
const pythonS =
const rubyS =
const ncS =
}*/

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.IP)
	return localAddr.IP
}

func a() {
	const intMessage = "Hello user, let us create so reverse shells! What we going to choose today?"
	const listOfShells = "Bash Reverse Shell\nPHP Reverse Shell.\nJava Reverse Shell.\n" +
		"Perl Reverse Shell.\nPython Reverse Shell.\nRuby Reverse Shell.\nNetcat Reverse Shell."
	fmt.Println(intMessage, "\n"+listOfShells)
	fmt.Println("So what shell we going to take? ")
	// var then variable name then variable type
	var choosenS string
	// Taking input from user
	fmt.Scanln(&choosenS)
	fmt.Println("you took", choosenS, "shell")
}

func main() {
	a()
	getOutboundIP()
}
