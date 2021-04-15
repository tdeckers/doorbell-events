package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type message struct {
	Message struct {
		Data string `json:"data"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// https://developers.google.com/nest/device-access/traits/device/doorbell-chime
type event struct {
	EventId        string `json:"eventId"`
	ResourceUpdate struct {
		Name   string                 `json:"name"`
		Events map[string]interface{} `json:"events"`
	} `json:"resourceUpdate"`
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	m := message{}

	// Focus on handling POST requests.
	if r.Method == "POST" {
		// Decode posted request into message struct
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			switch err {
			case io.EOF:
				http.Error(w, "EOF", http.StatusBadRequest)
				fmt.Println("EOF")
				return
			default:
				http.Error(w, fmt.Sprintf("json.NewDecoder: %v", err), http.StatusBadRequest)
				fmt.Printf("json.NewDecoder: %v\n", err)
				return
			}
		}
		// Decode the message data (which is the event from SDM)
		data, err := base64.StdEncoding.DecodeString(m.Message.Data)
		if err != nil {
			http.Error(w, fmt.Sprintf("base64.StdEncoding.DecodeString: %v", err), http.StatusBadRequest)
			fmt.Printf("base64.StdEncoding.DecodeString: %v\n", err)
			return
		}
		// Remove tabs and spaces from event json.  Openhab likes it clean.
		compactData := new(bytes.Buffer)
		err = json.Compact(compactData, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("json.Compact: %v", err), http.StatusBadRequest)
			fmt.Printf("json.Compact: %v\n", err)
			return
		}

		// Parse the event json into struct.  Validates format.
		evt := event{}
		err = json.Unmarshal(data, &evt)
		if err != nil {
			http.Error(w, fmt.Sprintf("json.Unmarshal: %v", err), http.StatusBadRequest)
			fmt.Printf("json.Unmarshal: %v\n", err)
			return
		}

		// Post event json straight to Openhab
		err = postData(compactData.Bytes())
		if err != nil {
			http.Error(w, err.Error(), 500)
			fmt.Println(err.Error())
			return
		}

		// TODO: if event has image, download image URL
		// For logging only: parse event and write to log.
		fmt.Println(evt.ResourceUpdate.Events)
		if _, ok := evt.ResourceUpdate.Events["sdm.devices.events.CameraMotion.Motion"]; ok {
			fmt.Println("Got motion")
		}
		if _, ok := evt.ResourceUpdate.Events["sdm.devices.events.DoorbellChime.Chime"]; ok {
			fmt.Println("Doorbell!	")
		}
	} else { // If not POST method
		fmt.Fprint(w, "ok")
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	}
}

func postData(data []byte) error {
	server := os.Getenv("OPENHAB_SERVER")
	user := os.Getenv("OPENHAB_USER")
	pwd := os.Getenv("OPENHAB_PWD")
	if server == "" {
		return fmt.Errorf("openhab_server can not be empty")
	}
	client := &http.Client{}
	r := bytes.NewReader(data)
	req, err := http.NewRequest("POST", server, r)
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "text/plain")
	req.SetBasicAuth(user, pwd)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status: %v", resp.Status)
	}
	return nil
}
