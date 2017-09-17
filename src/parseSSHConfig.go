package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const USER_POSITION, NAME_POSITION, HOST_POSITION = 0, 1, 2

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readConfig() ([3]int, [][3]string) {
	longestValues := [3]int{0, 0, 0}
	var prevHost, name, user, host = "", "", "", ""
	var rslt [][3]string

	data, err := ioutil.ReadFile("/Users/rjf/.ssh/config")
	check(err)

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) > 10 && line[0:11] == "  HostName " {
			name = line[11:len(line)]
			longestValues[NAME_POSITION] = max(longestValues[NAME_POSITION], len(name))
		} else if len(line) > 6 && line[0:7] == "  User " {
			user = line[7:len(line)]
			longestValues[USER_POSITION] = max(longestValues[USER_POSITION], len(user))
		} else if len(line) > 3 && line[0:4] == "Host" {
			host = line[5:len(line)]
			longestValues[HOST_POSITION] = max(longestValues[HOST_POSITION], len(host))
			if prevHost != "" {
				rslt = append(rslt, [3]string{user, name, prevHost})
			}
			prevHost = host
			host, user, name = "", "", ""
		}
	}
	return longestValues, append(rslt, [3]string{user, name, prevHost})
}

func main() {
	longestValues, rslts := readConfig()
	userLen := longestValues[USER_POSITION]
	nameLen := longestValues[NAME_POSITION]
	hostLen := longestValues[HOST_POSITION]
	formatStr := "%" + strconv.Itoa(userLen) + "s@%-" + strconv.Itoa(nameLen) + "s - Aliases [%-" + strconv.Itoa(hostLen) + "s]\n"
	for _, rslt := range rslts {
		fmt.Printf(formatStr, rslt[USER_POSITION], rslt[NAME_POSITION], rslt[HOST_POSITION])
	}
}
