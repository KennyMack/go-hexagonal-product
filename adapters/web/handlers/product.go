package handlers

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/kennymack/go-hexagonal-product/adapters/dto"
	"github.com/kennymack/go-hexagonal-product/application"
	"net/http"
)

func getProduct (service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(res).Encode(product)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
	})
}

func createProduct (service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		var productDto dto.Product

		err := json.NewDecoder(req.Body).Decode(&productDto)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(res).Encode(product)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write(jsonError(err.Error()))
			return
		}

		res.WriteHeader(http.StatusOK)
	})
}

func postProductEnable (service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		product, err = service.Enable(product)

		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			_, _ = res.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(res).Encode(product)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
	})
}

func postProductDisable (service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}

		product, err = service.Disable(product)

		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			_, _ = res.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(res).Encode(product)

		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
	})
}

func MakeProductHandelrs(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(postProductEnable(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(postProductDisable(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}
