package services

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type QuotationService struct {
	quotations []string
}

func NewQuotationService() *QuotationService {
	return &QuotationService{}
}

func (s *QuotationService) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	s.quotations = lines

	return nil
}

func (s *QuotationService) GetOne() (string, error) {
	if len(s.quotations) == 0 {
		return "", fmt.Errorf("no quotations loaded")
	}

	index := rand.Intn(len(s.quotations))

	return s.quotations[index], nil
}
