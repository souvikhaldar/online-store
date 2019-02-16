package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/souvikhaldar/online-store/pkg/db"
)

type pdt struct {
	PID   int
	Name  string
	Type  string
	Price float32
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	log.Print("--Running in AddProduct--")
	body, er := ioutil.ReadAll(r.Body)
	if er != nil {
		log.Print("Error in reading from request body", er)
		return
	}
	var p pdt
	e := json.Unmarshal(body, &p)
	if e != nil {
		log.Print("Error in Unmarshalling request body", e)
	}
	log.Printf("Product recieved: %+v", p)
	var id int
	if e := db.DBdriver.QueryRow("insert into pdt(name,type,price) values($1,$2,$3) returning pid", p.Name, p.Type, p.Price).Scan(&id); e != nil {
		log.Print("Error in insert new product: ", e)
		return
	}

	fmt.Fprintf(w, "Successfully inserted new product: %d", id)
}
