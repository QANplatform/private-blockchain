package main

import (
	"fmt"
	"os"
	"strconv"
)

var maxUser int = 100

// $0 construct
func constructor() int {
	if os.Getenv("DB_QVM_INITIALIZED") == "true" {
		os.Stderr.WriteString("contract is already initialized\n")
		return -1
	}
	os.Stdout.WriteString(fmt.Sprintf("DBW=QVM_INIT_MAXUSER=%d\n", maxUser))
	return 0
}

func initialize() int {
	if os.Getenv("DB_QVM_INITIALIZED") != "true" {
		os.Stderr.WriteString("contract is not initialized\n")
		return -1
	}
	maxUser, _ = strconv.Atoi(os.Getenv("DB_QVM_INIT_MAXUSER"))
	return 0
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "construct" {
		os.Exit(constructor())
	} else {
		ret := initialize()
		if ret < 0 {
			os.Exit(ret)
		}
	}

	if len(os.Args) == 3 && os.Args[1] == "register" {

		// GET THE CURRENT USER'S NAME
		previousName := os.Getenv("DB_USER_CURRENT")

		// OR DEFAULT TO "unknown" IF THIS IS THE FIRST CALL
		if len(previousName) == 0 {
			previousName = "unknown"
		}

		// GET THE TOTAL USER COUNT
		totalUserCount, _ := strconv.Atoi(os.Getenv("DB_TOTALUSERS"))
		if totalUserCount+1 > maxUser {
			os.Stderr.WriteString("exceeded max user\n")
			os.Exit(1)
		}

		// WRITE PREVIOUS USER NAME TO STDOUT
		os.Stdout.WriteString(fmt.Sprintf("OUT=prevname: %s\n", previousName))

		// UPDATE CURRENT USER NAME BY WRITING IT TO DB
		os.Stdout.WriteString(fmt.Sprintf("DBW=USER_CURRENT=%s\n", os.Args[2]))

		// STORE USER NAME UNDER A STORAGE SLOT FOR PERSISTENCE (CURRENT GETS OVERWRITTEN ON EACH CALL)
		os.Stdout.WriteString(fmt.Sprintf("DBW=USER_%d=%s\n", totalUserCount, os.Args[2]))

		// INCREMENT THE TOTAL USER COUNT
		os.Stdout.WriteString(fmt.Sprintf("DBW=TOTALUSERS=%d\n", totalUserCount+1))
		os.Exit(0)
	}
	if len(os.Args) >= 2 {
		os.Stderr.WriteString(fmt.Sprintf("Wrong CMD: %s\n", os.Args[1]))
		os.Exit(1)
	}
	os.Stderr.WriteString("Wrong args!\n")
	os.Exit(1)
}
