package c2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type httplistener struct {
	Ip         string
	Port       string
	errorLog   *log.Logger
	requestLog *log.Logger
}

func (h *httplistener) taskHandler(w http.ResponseWriter, r *http.Request) {
	h.requestLog.Println("Received requests: ", r.Method, r.RemoteAddr, r.URL.Path, time.Now)
	task := map[string]any{
		"cmd":  "echo",
		"args": []string{"get shit on"},
	}
	jsonData, err := json.Marshal(task)
	if err != nil {
		h.errorLog.Println("Error Marshalling json: ")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func (h *httplistener) Listen() {
	err := os.MkdirAll("C2/logs", 0755)
	if err != nil {
		fmt.Printf("Could not create logs directory %v\n", err)
	}
	logFile, err := os.OpenFile("c2/logs/listener.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error Opening Log Files")
	}
	h.errorLog = log.New(logFile, "ERROR", log.LstdFlags)
	h.requestLog = log.New(logFile, "REQUEST", log.LstdFlags)

	// init http server
	mux := http.NewServeMux()
	// register handlers
	mux.HandleFunc("/task", h.taskHandler)

	server := &http.Server{
		Addr:     fmt.Sprintf(h.Ip + ":" + h.Port),
		ErrorLog: h.errorLog,
		Handler:  mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		h.errorLog.Fatalf("Server failed to start: %v", err)
	}

}
