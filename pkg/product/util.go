package product

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/souvikhaldar/online-store/pkg/agent"
	"github.com/souvikhaldar/online-store/pkg/db"
)

func GetReqProduct(pid int) (*Product, error) {
	var pdt Product
	e := db.DBdriver.QueryRow("select * from pdt where pid=$1", pid).Scan(&pdt.PID, &pdt.Name, &pdt.Type, &pdt.Price)
	if e != nil && e != sql.ErrNoRows {
		return nil, fmt.Errorf("Error in fetching the required product: %s", e)
	} else if e != nil && e == sql.ErrNoRows {
		return nil, fmt.Errorf("Required product not found")
	}
	return &pdt, nil
}

func Purchase(w http.ResponseWriter, r *http.Request) {
	log.Print("--Running in Purchase--")
	queryVal := r.URL.Query()
	pid := queryVal.Get("product_id")
	if len(pid) < 0 {
		http.Error(w, "Unable to fetch the product_id, eg.- /product_id=33", 400)
		return
	}
	pidInt, er := strconv.Atoi(pid)
	if er != nil {
		http.Error(w, fmt.Sprintf("Unable to convert pid string to int: %s", er), 500)
		return
	}
	pdt, e := GetReqProduct(pidInt)
	if e != nil {
		http.Error(w, fmt.Sprintf("Unable to get the required product: %s", e), 500)
		return
	}
	agt, e := agent.GetAvailAgent()
	if e != nil {
		http.Error(w, fmt.Sprintf("Unable to assign agent: %s", e), 404)
		return
	}
	var purID int
	if e := db.DBdriver.QueryRow("insert into purchase(pid,aid) values($1,$2) returning pur_id", pdt.PID, agt.AID).Scan(&purID); e != nil {
		http.Error(w, fmt.Sprintf("Unable to insert to purchase table: %s", e), 500)
		return
	}
	if e := agent.UpdateAgentAvail(agt.AID); e != nil {
		http.Error(w, fmt.Sprintf("Unable to update the agent's availability: %s", e), 500)
		return
	}
	fmt.Fprintf(w, "Successfully purchased product: %d with purchase ID: %d; Agent assigned: %d ", pdt.PID, purID, agt.AID)
}
