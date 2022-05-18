package main

import (
	"fmt"
	xerrors "github.com/pkg/errors"
	"os"
)

var AuthenticateErr = xerrors.New("authenticate failed")

func main() {
	err := test()
	if err != nil {
		//if xerrors.Is(err, AuthenticateErr) {
		//	//if xerrors.As(err, &AuthenticateErr) {
		//	fmt.Printf("%+v", err)
		//}
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}

func test() error {
	return xerrors.Wrap(test1(), "from test1")
}

func test1() error {
	return fmt.Errorf("from test2: %v", test2())
}

func test2() error {
	return AuthenticateErr
}
