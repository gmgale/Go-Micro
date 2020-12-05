package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gmgale/go_micro/4_RESTful_services/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.l.Println("GET hit")
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.l.Println("POST hit")
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect ID in the URI
		p.l.Println("PUT hit")
		u := r.URL.Path
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(u, -1)

		if len(g) != -1 {
			p.l.Println("Invalid URI more than one ID.")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one ID.")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			if len(g) != -1 {
				p.l.Println("Unale to convert to number")
				http.Error(rw, "Invalid URI", http.StatusBadRequest)
			}
			http.Error(rw, "invalid URI", http.StatusBadRequest)
		}

		p.l.Println("got id: ", id)

		p.updateProducts(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
