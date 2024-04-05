package main

import (
	"fmt"

	"github.com/RafaelPereiraSantos/grpc-test/cmd/server/internal"
)

func main() {
	cdpImplementation := internal.NewCDPImplementation()

	errChan := make(chan error)

	go func() {
		internal.StartProtoServer(cdpImplementation)

		errChan <- nil
	}()

	go func() {
		internal.StartHTTPServer(cdpImplementation)

		errChan <- nil
	}()

	fmt.Println(<-errChan)

}
