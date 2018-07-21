package main

import (
	"gotalk/gotalkcomms/gotalkmessages"
	//hydraproto "Hydra/hydracomms/hydramessages/protobuff"
	"flag"
	"log"
	"strings"
	//"Hydra/hydracomms/hydramessages/thrift/gen-go/hydraThrift"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runClient(address string) {
	/*
		ship := &gotalkproto.Ship{
			Shipname:    "Hydra",
			CaptainName: "Jala",
			Crew: []*gotalkproto.Ship_CrewMember{
				&gotalkproto.Ship_CrewMember{1, "Kevin", 5, "Pilot"},
				&gotalkproto.Ship_CrewMember{2, "Jade", 4, "Tech"},
				&gotalkproto.Ship_CrewMember{3, "Wally", 3, "Enginneer"},
			},
		}


		ship := &gotalkThrift.Ship{
			Shipname:    "Hydra",
			CaptainName: "Jala",
			Crew: []*gotalkThrift.CrewMember{
				&gotalkThrift.CrewMember{1, "Kevin", 5, "Pilot"},
				&gotalkThrift.CrewMember{2, "Jade", 4, "Tech"},
				&hgotalkThrift.CrewMember{3, "Wally", 3, "Enginneer"},
			},
		}
	*/

	ship := &hydragob.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []hydragob.CrewMember{
			hydragob.CrewMember{1, "Kevin", 5, "Pilot"},
			hydragob.CrewMember{2, "Jade", 4, "Tech"},
			hydragob.CrewMember{3, "Wally", 3, "Enginneer"},
		},
	}

	if err := gotalkmessages.EncodeAndSend(gotalkmessages.GOB, ship, address); err != nil {
		log.Println(err)
	}
}

func runServer(address string) {
	for ship := range gotalkmessages.ListenAndDecode(gotalkmessages.GOB, address) {
		log.Println(ship)
	}
}
