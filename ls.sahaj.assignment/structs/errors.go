package structs

import "fmt"

type SentenceNotFound struct {
	Text string
}

func (s *SentenceNotFound) Error() string {
	return fmt.Sprintf("%s sentence not found", s.Text)
}

type MatchNotFound struct {
	Text string
}

func (s *MatchNotFound) Error() string {
	return fmt.Sprintf("%s match not found", s.Text)
}
