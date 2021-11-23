package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ApogeeNetworking/apgnet"
	"github.com/ApogeeNetworking/apgnet/apgsn"
	"github.com/subosito/gotenv"
)

var host string

func init() {
	gotenv.Load()
	host = os.Getenv("HOST")
}

func main() {
	apg := apgnet.NewService(host, false)

	// sites, _ := apg.Platform.GetSites()
	// for _, site := range sites {
	// 	fmt.Println(site)
	// }
	// site, _ := apg.Platform.GetSite("60a4044ab9174de3ee175d23")
	// fmt.Println(site)
	inc := &apgsn.CreateIncidentReq{
		Description:       "Lockhaven - 825 - AP Down - ap01.lockhaven.outdr.2.pa",
		ShortDescription:  "Lockhaven - 825 - AP Down",
		Category:          "Break/Fix NOC Networking",
		SubCategory:       "Infrastructure Apogee",
		CompanyID:         "582bcf26db9118d0a07dd77648961956",
		BusinessServiceID: "58b1c1d81b768cd09830433fbd4bcb63",
		AssignmentGroupID: "efb0cfebdbfa6f002b38ed384b961977",
		CreatedBy:         "csmith@apogee.us",
		Priority:          "6",
		Impact:            "5",
		Severity:          "3",
	}
	ninc, err := apg.Snow.CreateIncidentFlow(inc)
	if err != nil {
		log.Fatal(err)
	}
	d, _ := json.Marshal(ninc)
	fmt.Println(string(d))
	// user, _ := apg.Snow.GetUser("swomack@apogee.us")
	// fmt.Println(user)
	// wo, err := apg.Snow.GetWorkOrderByID("WO0042396")
	// if err != nil {
	// 	log.Printf("%v", err)
	// 	return
	// }
	// fmt.Println(wo)
}
