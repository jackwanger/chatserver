package config

import (
	"fmt"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	err := ReadIniFile("../config.ini")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%q\n", Addr)
	fmt.Printf("%v\n", NumCpu)
	fmt.Printf("%v\n", AcceptTimeout)
	fmt.Printf("%v\n", ReadTimeout)
	fmt.Printf("%v\n", WriteTimeout)
	fmt.Printf("%q\n", UuidDB)
	fmt.Printf("%q\n", OfflineMsgidsDB)
	fmt.Printf("%q\n", IdToMsgDB)
	fmt.Printf("%v\n", TimedUpdateDB)
	fmt.Printf("%q\n", LogFile)
	fmt.Printf("%q\n", EmailServerAddr)
	fmt.Printf("%q\n", EmailServerPort)
	fmt.Printf("%q\n", EmailAccount)
	fmt.Printf("%q\n", EmailPassword)
	fmt.Printf("%v\n", EmailDuration)
	fmt.Printf("%q\n", EmailToList)

	s := strings.Split(EmailToList, " ")
	for _, v := range s {
		if v != "" {
			fmt.Printf("%q\n", v)
		}
	}
}
