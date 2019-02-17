package agent

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/souvikhaldar/online-store/pkg/db"
)

func GetAvailAgent() (*Agent, error) {
	var agt Agent
	e := db.DBdriver.QueryRow("select * from agt where availability=true limit 1").Scan(&agt.AID, &agt.Name, &agt.AdharNum, &agt.Availability)
	if e != nil && e != sql.ErrNoRows {
		log.Print("Unable to get an availble agent: ", e)
		return nil, e
	} else if e != nil && e == sql.ErrNoRows {
		log.Print("No agent available: ", e)
		return nil, errors.New("No agent available")
	}
	return &agt, nil
}

func UpdateAgentAvail(agentID int) error {
	if _, e := db.DBdriver.Exec("update agt set availability=false where aid=$1", agentID); e != nil {
		er := fmt.Sprintf("Unable to update the availability of the agent to false: %s", e)
		return errors.New(er)
	}
	return nil
}
