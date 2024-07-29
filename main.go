package main

import (
	"fmt"
	"log"
	"net/http"
)

var mas = "https://get.activated.win"

func main() {

	http.HandleFunc("/{method...}", func(w http.ResponseWriter, r *http.Request) {

		m := r.PathValue("method")
		switch m {
		case "hwid":
			fmt.Fprintf(w, "& ([ScriptBlock]::Create((irm %s))) /HWID /S", mas)
			return

		case "ohook":
			fmt.Fprintf(w, "& ([ScriptBlock]::Create((irm %s))) /Ohook /S", mas)
			return
		case "kms38":
			fmt.Fprintf(w, "& ([ScriptBlock]::Create((irm %s))) /KMS38 /S", mas)
			return

		case "kms":
			fmt.Fprintf(w, "& ([ScriptBlock]::Create((irm %s))) /KMS-ActAndRenewalTask /S", mas)
			return
		}

		fmt.Fprintf(w, "irm %s | iex", mas)
	})

	fmt.Println("Starting Service")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
