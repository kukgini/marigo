package main

import (
	"log"
	/*
		"github.com/hyperledger/aries-framework-go-ext/component/vdr/indy"
		"github.com/hyperledger/indy-vdr/wrappers/golang/vdr"
	*/
	"github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/hyperledger/aries-framework-go/pkg/vdr/httpbinding"
)

func main() {
	log.Println("create httpbinding vdr")
	vdr, err := httpbinding.New("http://127.0.0.1:3000/did")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("check vdr is available")
	did, err := vdr.Read("21tDAKCERh95uGgKbJNHYp")
	if err != nil {
		log.Println("read did failed")
		log.Fatalln(err)
	}
	log.Println(did)

	log.Println("create aries framework")
	framework, err := aries.New(aries.WithVDR(vdr))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("get current aries context")
	ctx, _ := framework.Context()

	log.Println("create new didexchange client")
	client, _ := didexchange.New(ctx)

	log.Println("create invitation")
	invitation, _ := client.CreateInvitation("test")

	resout := `created invitation:
	             - id: %s 
	             - label: %s
	             - did: %s
	             - type: %s
               `
	log.Printf(resout,
		invitation.Invitation.ID,
		invitation.Invitation.Label,
		invitation.Invitation.DID,
		invitation.Invitation.Type)

	log.Println("create invitation with DID")
	invitation2, err := client.CreateInvitationWithDID("invitationPublic", "21tDAKCERh95uGgKbJNHYp")
	if err != nil {
		log.Fatal(err)
	}

	resout2 := `created invitation:
	             - id: %s 
	             - label: %s
	             - did: %s
	             - type: %s
               `
	log.Printf(resout2,
		invitation2.Invitation.ID,
		invitation2.Invitation.Label,
		invitation2.Invitation.DID,
		invitation2.Invitation.Type)
}
