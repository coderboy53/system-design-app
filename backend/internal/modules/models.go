package modules

import "github.com/lib/pq"

type Module struct {
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	ModuleOrder int16  `json:"module_order" db:"module_order"`
	TopicCount  int16  `json:"topic_count" db:"topic_count"`
	Overview    string `json:"overview" db:"overview"`
}

type Topic struct {
	Id            string         `json:"id"`
	Title         string         `json:"title"`
	Module        string         `json:"module"`
	Order         int16          `json:"order"`
	Minutes       int16          `json:"est_minutes"`
	Timelines     pq.StringArray `json:"timelines"`
	Tags          pq.StringArray `json:"tags"`
	Prerequisites pq.StringArray `json:"prerequisites"`
	Related       pq.StringArray `json:"related"`
	Source        string         `json:"source"`
	WordCount     int            `json:"word_count"`
	Sections      pq.StringArray `json:"sections"`
	HasSelfCheck  bool           `json:"has_self_check"`
	Body          string         `json:"body,omitempty"`
}
