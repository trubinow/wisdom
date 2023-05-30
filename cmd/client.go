package main

import (
	"encoding/gob"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"wisdom/internal/models"
	"wisdom/pkg/hashcash"
)

func main() {
	log := logrus.WithFields(logrus.Fields{
		"service": "tcp-client",
	})

	// Convert port env variable value to an integer
	port, err := strconv.Atoi(os.Getenv("TCP_SERVER_PORT"))
	if err != nil {
		log.Errorf("error converting PORT variable to integer: %s", err.Error())
		return
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", os.Getenv("TCP_SERVER_HOST"), port))
	if err != nil {
		log.Errorf("error connecting to the server: %s", err.Error())
		return
	}
	defer conn.Close()

	//Send quotation-request to server
	var request = models.Request{
		Type: "quotation",
	}

	encoder := gob.NewEncoder(conn)
	if err = encoder.Encode(request); err != nil {
		log.Errorf("error encoding request: %s", err.Error())
		return
	}

	// Receive the number from the server
	var work models.Work
	decoder := gob.NewDecoder(conn)
	if err = decoder.Decode(&work); err != nil {
		log.Errorf("error decoding response: %s", err.Error())
		return
	}

	// Make work
	proof, err := hashcash.GenerateHashToken(work.Resource, work.Difficulty)
	if err != nil {
		log.Errorf("error generating hash token: %s", err.Error())
		return
	}

	// Send the proof to the server
	if err = encoder.Encode(proof); err != nil {
		log.Errorf("error sending proof: %s", err.Error())
		return
	}

	// Receive server response
	var quotation string
	if err = decoder.Decode(&quotation); err != nil {
		log.Errorf("error decoding response: %s", err.Error())
		return
	}

	if quotation == "" {
		log.Error("error getting quotation.")
	} else {
		log.Infof("quotation recieved: %s", quotation)
	}
}
