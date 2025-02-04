package main

import (
	"context"
	"fmt"
	"github.com/Integrio/biztalk-server-go/client"
)

func main() {
	x, err := client.NewClientBuilder("http://<host>").UseNtlmAuth("<username>", "<password>").Build()
	if err != nil {
		panic(err)
	}

	y, err := x.ReceiveLocationsGet(context.Background())
	if err != nil {
		panic(err)
	}
	if y == nil {
		_ = fmt.Errorf("empty response")
		return
	}

	receiveLocations := *y

	fmt.Println(receiveLocations[0].Name)

	z, err := x.SendPortsGet(context.Background())
	if err != nil {
		panic(err)
	}
	if z == nil {
		_ = fmt.Errorf("empty response")
		return
	}

	sendPorts := *z

	fmt.Println(sendPorts[0].Name)
}
