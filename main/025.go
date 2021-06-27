package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var input int
	var dec int
FRE:
	fmt.Print("Enter Integer (1 - 100)")
	fmt.Scan(&input)
	if input < 1 || input > 100 {
		fmt.Print("\nOnly 1 - 100")
		goto FRE
	} else {
		clearscr()
	}
RE:
	fmt.Print("\nEnter Integer you think it should be : ")
	fmt.Scan(&dec)
	if dec < input {
		fmt.Print("too small.....")
		goto RE
	} else if dec > input {
		fmt.Print("Too big......")
		goto RE
	} else if dec == input {
		fmt.Print("Congratulation you win!!!!")
	}

}

func clearscr() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
