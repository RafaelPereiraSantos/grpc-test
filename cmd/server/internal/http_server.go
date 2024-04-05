package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartHTTPServer(cdpImplementation *CDPImplementation) {
	fmt.Println("Starting server")

	http.HandleFunc("/v1/identify", func(w http.ResponseWriter, r *http.Request) {

		userRequest := IdentifyUser{}

		err := json.NewDecoder(r.Body).Decode(&userRequest)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		err = cdpImplementation.Identify(context.Background(), userRequest)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil))
}
