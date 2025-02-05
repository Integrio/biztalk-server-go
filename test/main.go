package main

import (
	"context"
	"fmt"
	"github.com/Integrio/biztalk-server-go/client"
)

func main() {
	x, err := client.NewClientBuilder("http://98.64.171.120/").UseNtlmAuth("btapiuser", "P1ngP0ng").Build()
	if err != nil {
		panic(err)
	}

	// Instances
	instances, err := x.GetOperationalDataInstances(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(instances[0].HostName)

	// Messages
	messages, err := x.GetOperationalDataMessages(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(messages[0].HostName)

	// Orchestrations
	orchestrations, err := x.GetOrchestrations(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(orchestrations[0].FullName)

	receivePorts, err := x.GetReceivePorts(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(receivePorts[0].Name)

	receiveLocations, err := x.GetReceiveLocations(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(receiveLocations[0].Name)

	sendPorts, err := x.GetSendPorts(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(sendPorts[0].Name)

	sendPortGroups, err := x.GetSendPortGroups(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(sendPortGroups[0].Name)

	hostInstances, err := x.GetHostInstances(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(hostInstances[0].HostName)
}
