/*
 * Test case for BRE
 *
 * Service
 *
 * API version: 1.0.0
 * Contact: gk@rekunov.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"github.com/xpoh/BRE-test/internal/service"
	"log"
	"net/http"
)

func AddContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	c := service.Content{}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(c.Id)
	err = service.App.SaveContent(c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func GetIdGet(w http.ResponseWriter, r *http.Request) {
	// TODO get content
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
