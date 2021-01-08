package normalizer

import (
	"log"
	"strings"
	"time"
)

func normDate(date string, layout string) (out string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var line string

	switch layout {
	case "01/02/2006":
		s := strings.Split(date, "/")
		if len(s[0]) == 1 {
			s[0] = "0" + s[0]
		}

		if len(s[1]) == 1 {
			s[1] = "0" + s[1]
		}

		line = s[0] + "/" + s[1] + "/" + s[2]
	case "20060102150405":
		s := strings.Split(date, ".")
		line = s[0]
	case "20060102":
		line = date

	}

	t, _ := time.Parse(layout, line)

	out = t.Format(time.RFC3339)

	return
}
