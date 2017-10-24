package main

import (
	"flag"
	"github.com/TheBookPeople/alterian/listmanager"
	"log"
)

var token = flag.String("token", "", "Access token")

func main() {
	flag.Parse()
	if *token == "" {
		log.Fatalln("Token not specified")
	}

	url := "https://uk56.em.sdlproducts.com/listmanager.asmx"

	auth := &listmanager.BasicAuth{}

	client := listmanager.NewSOAPClient(url, true, auth)
	req := listmanager.GetLists{
		Token: *token,
	}

	var resp listmanager.GetListsResponse
	if err := client.Call("DMWebServices/GetLists", req, &resp); err != nil {
		log.Fatal(err)
	}

	for _, list := range resp.GetListsResult.DMList {
		log.Println(list.Name)
	}
}
