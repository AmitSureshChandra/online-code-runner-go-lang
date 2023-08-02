package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"online-compiler/internal/docker"
	"online-compiler/internal/runner"
)

type RunRequest struct {
	Input    string `json:"input"`
	Code     string `json:"code"`
	Compiler string `json:"compiler"`
}

type RunRes struct {
	Time   float64 `json:"time"`
	Output string  `json:"output"`
}

func HandleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Parse the JSON request body into the RunRequest struct
	var req RunRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	output, start, stop := runner.Run(req.Code, req.Input, docker.GetContainerName(req.Compiler))

	response, err := json.Marshal(RunRes{
		Time:   stop.Sub(start).Seconds(),
		Output: string(output),
	})

	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
