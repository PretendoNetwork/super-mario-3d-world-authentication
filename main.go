package main

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-common-go/authentication"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPrudpVersion(1)
	nexServer.SetPRUDPProtocolMinorVersion(2)
	nexServer.SetNexVersion(30400)
	nexServer.SetKerberosPassword(os.Getenv("KERBEROS_PASSWORD"))
	nexServer.SetAccessKey("86a89872")

	nexServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==SM3DW - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("==================")
	})

	authenticationProtocol := authentication.NewCommonAuthenticationProtocol(nexServer)

	secureStationURL := nex.NewStationURL("")
	secureStationURL.SetScheme("prudps")
	secureStationURL.SetAddress(os.Getenv("SECURE_SERVER_LOCATION"))
	secureStationURL.SetPort(os.Getenv("SECURE_SERVER_PORT"))
	secureStationURL.SetCID("1")
	secureStationURL.SetPID("2")
	secureStationURL.SetSID("1")
	secureStationURL.SetStream("10")
	secureStationURL.SetType("2")

	authenticationProtocol.SetSecureStationURL(secureStationURL)
	authenticationProtocol.SetBuildName("Pretendo SM3DW")
	authenticationProtocol.SetPasswordFromPIDFunction(passwordFromPID)

	nexServer.Listen(":60030")
}
