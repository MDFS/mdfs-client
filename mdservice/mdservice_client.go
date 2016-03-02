package main

import (
	"bufio"
	"fmt"
	//"github.com/CPSSD/MDFS/mdservice"
	//"github.com/CPSSD/MDFS/utils"
	"net"
	"os"
	"strings"
)

/*
var currentDir mdservice.DirNode
var user mdservice.UUID
*/
func main() {

	// config will be read locally later
	protocol := "tcp"
	socket := "localhost:8081"
	user := "jim"

	conn, _ := net.Dial(protocol, socket)
	defer conn.Close()

	// read and write buffer to the mdserv
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	var sendcode uint8

	currentDir := ""

	//rootDir := mdservice.MkRoot()
	//currentDir := rootDir
	//user.Initialise("jim")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(user + ":/" + currentDir + " >> ")
		cmd, _ := reader.ReadString('\n')
		// remove trailing newline character before splitting
		args := strings.Split(strings.TrimSpace(cmd), " ")

		switch args[0] {
		case "":
			continue

		case "ls":
			sendcode = 1

			err := w.WriteByte(sendcode)
			w.Flush()
			if err != nil {
				panic(err)
			}

			err = w.WriteByte(uint8(len(args)))
			w.Flush()
			if err != nil {
				panic(err)
			}

			for i := 1; i < len(args); i++ {
				w.WriteString(currentDir + args[i] + "\n")
				w.Flush()
			}

			msg, _ := r.ReadString(' ')

			files := strings.Split(msg, ",")

			for n, file := range files {
				if n != len(files)-1 {
					fmt.Println(file)
				}
			}

		case "mkdir":
			sendcode = 2

			err := w.WriteByte(sendcode)
			w.Flush()
			if err != nil {
				panic(err)
			}

			err = w.WriteByte(uint8(len(args)))
			w.Flush()
			if err != nil {
				panic(err)
			}

			for i := 1; i < len(args); i++ {
				w.WriteString(args[i] + "\n")
				w.Flush()
			}
		case "rmdir":
			sendcode = 3

			err := w.WriteByte(sendcode)
			w.Flush()
			if err != nil {
				panic(err)
			}

			err = w.WriteByte(uint8(len(args)))
			w.Flush()
			if err != nil {
				panic(err)
			}

			for i := 1; i < len(args); i++ {
				w.WriteString(args[i] + "\n")
				w.Flush()
			}

		case "pwd":
			continue
			//currentDir.Pwd(

		case "cd":
			continue
			/*
				err, next := mdservice.Cd(currentDir, args[1])
				if err != nil {
					fmt.Println(err)
				}
				currentDir = next
			*/

		case "exit":
			os.Exit(1)

		default:
			fmt.Println("Unrecognised command")
		}
	}
}
