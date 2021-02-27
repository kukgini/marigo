package main

import (
	"log"

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
	did := "21tDAKCERh95uGgKbJNHYp"
	diddoc, err := vdr.Read(did)
	if err != nil {
		log.Println("read did failed")
		log.Fatalln(err)
	}
	log.Printf("did=%s, diddoc=%+v\n", did, *diddoc)

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
	createInvitation(client, "test", "")

	log.Println("create invitation with DID")
	createInvitation(client, "invitationPublic", "21tDAKCERh95uGgKbJNHYp")
}

func createInvitation(client *didexchange.Client, label string, did string) (*didexchange.Invitation, error) {
	var invitation *didexchange.Invitation
	var err error
	if len(did) > 0 {
		invitation, err = client.CreateInvitationWithDID(label, did)
	} else {
		invitation, err = client.CreateInvitation(label)
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" - type: %s", invitation.Type)
	log.Printf(" - id: %s", invitation.ID)
	log.Printf(" - label: %s", invitation.Label)
	log.Printf(" - did: %s", invitation.DID)
	log.Printf(" - service endpoint: %s", invitation.ServiceEndpoint)
	log.Printf(" - recipient keys: %s", invitation.RecipientKeys)
	log.Printf(" - thread: %+v", invitation.Thread)

	return invitation, err
}
