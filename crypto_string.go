package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cryptoString string

func (cp *cryptoString) Scan() {
	fmt.Println("Enter string:")
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cp.Write(line)
}

func (cp *cryptoString) Read() string {
	return string(*cp)
}

func (cp *cryptoString) Write(str string) {
	*cp = cryptoString(str)
}

func (cp *cryptoString) Show(strs ...string) {
	msg := new(string)
	if len(strs) < 1 {
		*msg = "Your new string:"
	} else {
		*msg = strings.Join(strs, " ")
	}
	fmt.Printf("%s %c[35m%s%c[0m\n", *msg, 27, *cp, 27)
}

func (cp *cryptoString) Encode(method func(string) string) {
	cp.Scan()
	cp.Write(method(cp.Read()))
}

func (cp *cryptoString) EncodeToNumber() {
	cp.Write(encodeToNum(cp.Read()))
}

func (cp *cryptoString) EncodeToPigLng() {
	cp.Write(encodeToPigLng(cp.Read()))
}

func (cp *cryptoString) Decode(method func(string) (string, error)) {
	str, err := method(cp.Read())

	if err != nil {
		fmt.Println(err)
	} else {
		cp.Write(str)
	}
}

func (cp *cryptoString) DecodeFromPigLng() {
	cp.Decode(decodeFromPigLng)
}

func (cp *cryptoString) DecodeFromNumber() {
	cp.Decode(decodeFromNum)
}
