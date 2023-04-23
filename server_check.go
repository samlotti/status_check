package main

import (
	"bytes"
	"fmt"
	"github.com/samlotti/status_check/checker"
	"log"
	"net/http"
	"os"
	"strings"
)

var Version = "v0.1.3"

func processRequest(w http.ResponseWriter, r *http.Request) {
	allPassed := true
	b := bytes.Buffer{}
	for _, check := range checker.GConfig.Rules {
		rr := check.CheckRule()
		if check.CheckRule().Passed {
			b.Write([]byte(" (   ) "))
		} else {
			allPassed = false
			b.Write([]byte(" ( ! ) "))
		}
		b.Write([]byte(" ->  "))
		b.Write([]byte(fmt.Sprintf("%-24s", rr.Name)))
		if len(rr.Extra) > 0 {
			b.Write([]byte("   note: "))
			b.Write([]byte(rr.Extra))
		}

		b.Write([]byte("\n"))
	}

	if !allPassed {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}

	w.Write([]byte(fmt.Sprintf("Server status check, %s\n\n", Version)))
	w.Write([]byte(strings.Repeat("-", 80)))
	w.Write([]byte("\n"))

	w.Write(b.Bytes())
	w.Write([]byte(strings.Repeat("-", 80)))
	w.Write([]byte("\n"))

	if !allPassed {
		w.Write([]byte("Status: ERRORS\n\n"))
	} else {
		w.Write([]byte("Status: OK\n\n"))
	}

}

func main() {
	configName := "server_check.conf"
	if len(os.Args) > 1 {
		configName = os.Args[1]
	}
	checker.LoadConfig(configName)

	http.HandleFunc("/", processRequest)

	sport := fmt.Sprintf(":%s", checker.GConfig.Port)
	fmt.Printf("Listening on %s\n", sport)
	log.Fatal(http.ListenAndServe(sport, nil))

}
