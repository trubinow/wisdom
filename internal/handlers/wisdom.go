package handlers

import (
	"encoding/gob"
	"encoding/hex"
	"net"
	"wisdom/internal/helpers"
	"wisdom/internal/models"
	"wisdom/pkg/hashcash"
)

type Quotation interface {
	GetOne() (string, error)
}

type WisdomHandler struct {
	log              helpers.Logger
	quotationService Quotation
	difficulty       int
}

func NewWisdomHandler(logger helpers.Logger, quotationService Quotation, difficulty int) WisdomHandler {
	return WisdomHandler{
		log:              logger,
		quotationService: quotationService,
		difficulty:       difficulty,
	}
}

func (h WisdomHandler) Handle(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	var request models.Request
	if err := decoder.Decode(&request); err != nil {
		h.log.Errorf("error decoding data: %s", err.Error())
		return
	}

	if request.Type != "quotation" {
		h.log.Errorf("unknown request type: %s", request.Type)
		return
	}

	h.log.Info("quotation request received")

	resource, err := hashcash.GenerateNonce()
	if err != nil {
		h.log.Errorf("error generating resource: %s", err.Error())
		return
	}

	work := models.Work{
		Resource:   resource[:],
		Difficulty: h.difficulty,
	}

	// Send the work(challenge) to the client
	encoder := gob.NewEncoder(conn)
	if err = encoder.Encode(work); err != nil {
		h.log.Errorf("error encoding work(challenge): %s", err.Error())
		return
	}

	// Receive pow from the client
	var pow string
	if err = decoder.Decode(&pow); err != nil {
		h.log.Errorf("error decoding pow: %s", err.Error())
		return
	}

	// Check the pow
	if ok := hashcash.VerifyHashToken(pow, hex.EncodeToString(work.Resource), work.Difficulty); !ok {
		h.log.Errorf("hash token is not valid")
		return
	}

	// Get quotation
	quotation, err := h.quotationService.GetOne()
	if err != nil {
		h.log.Errorf("get quotation error: %s", err.Error())
		return
	}

	// Send quotation to the client
	if err = encoder.Encode(quotation); err != nil {
		h.log.Errorf("error sending quotation: %s", err.Error())
	}

}
