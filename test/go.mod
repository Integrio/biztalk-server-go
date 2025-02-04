module main

go 1.23.5

require github.com/Integrio/biztalk-server-go/client v0.0.0-00010101000000-000000000000

require (
	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
	github.com/google/uuid v1.5.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
)

replace github.com/Integrio/biztalk-server-go/client => ./../client
