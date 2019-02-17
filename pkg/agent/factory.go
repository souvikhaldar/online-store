package agent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/souvikhaldar/online-store/pkg/db"
)

type Agent struct {
	AID          int
	Name         string
	AdharNum     string
	Availability bool
}

func AddAgent(w http.ResponseWriter, c *http.Request) {
	log.Print("--Running in AddAgent--")
	bod, er := ioutil.ReadAll(c.Body)
	if er != nil {
		log.Print("Error in reading from the request body: ", er)
		return
	}
	var agent Agent
	if e := json.Unmarshal(bod, &agent); e != nil {
		log.Print("Error in unmarshalling the agent: ", e)
		return
	}
	log.Printf("Agent recieved: %+v", agent)
	var aid int
	if e := db.DBdriver.QueryRow("insert into agt(name,adhar_num,availability) values ($1,$2,$3) returning aid", agent.Name, agent.AdharNum, true).Scan(&aid); e != nil {
		log.Print("Error in inserting to agt: ", e)
		return
	}
	fmt.Fprintf(w, "Successfully added new agent: %d", aid)
}
