package discord

import (
	"github.com/SAl0nKA/RozvrhBOT/config"
	"time"
)

func GetSChoolday(day time.Weekday) *config.SchoolDay {
	if day == 0 {
		day = time.Now().Weekday()
	}
	switch day {
	case 1:
		sd := config.SchoolDays[1]
		return sd
	case 2:
		sd := config.SchoolDays[2]
		return sd
	case 3:
		sd := config.SchoolDays[3]
		return sd
	case 4:
		sd := config.SchoolDays[4]
		return sd
	case 5:
		sd := config.SchoolDays[5]
		return sd
	default:
		sd := config.SchoolDays[0]
		return sd
	}
}

func Hodiny(dalsia int) (hod, link, cas string) {
	konce := config.KoniecHodin
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	sd := GetSChoolday(0)
	switch {
	case h <= konce[0].Hodina && m < konce[0].Minuta:
		hod = sd.Hodiny[0+dalsia]
		cas = sd.Casy[0+dalsia]
		link =  sd.Linky[0+dalsia]
	case (h == konce[0].Hodina && m > konce[0].Minuta) || (h == konce[1].Hodina && m < konce[1].Minuta):
		hod = sd.Hodiny[1+dalsia]
		cas = sd.Casy[1+dalsia]
		link = sd.Linky[1+dalsia]
	case (h == konce[1].Hodina && m > konce[1].Minuta) || (h == konce[2].Hodina && m < konce[2].Minuta):
		hod = sd.Hodiny[2+dalsia]
		cas = sd.Casy[2+dalsia]
		link =  sd.Linky[2+dalsia]
	case (h == konce[2].Hodina && m > konce[2].Minuta) || (h == konce[3].Hodina && m < konce[3].Minuta):
		hod = sd.Hodiny[3+dalsia]
		cas = sd.Casy[3+dalsia]
		link =  sd.Linky[3+dalsia]
	case (h == konce[3].Hodina && m > konce[3].Minuta) || (h == konce[4].Hodina && m < konce[4].Minuta):
		hod = sd.Hodiny[4+dalsia]
		cas = sd.Casy[4+dalsia]
		link =  sd.Linky[4+dalsia]
	case (h == konce[4].Hodina && m > konce[4].Minuta) || (h == konce[5].Hodina && m < konce[5].Minuta):
		hod = sd.Hodiny[5+dalsia]
		cas = sd.Casy[5+dalsia]
		link =  sd.Linky[5+dalsia]
	case (h == konce[5].Hodina && m > konce[5].Minuta) || (h == konce[6].Hodina && m < konce[6].Minuta):
		hod = sd.Hodiny[6+dalsia]
		cas = sd.Casy[6+dalsia]
		link =  sd.Linky[6+dalsia]
	case (h == konce[6].Hodina && m > konce[6].Minuta) || (h == konce[7].Hodina && m < konce[7].Minuta):
		hod = sd.Hodiny[7+dalsia]
		cas = sd.Casy[7+dalsia]
		link =  sd.Linky[7+dalsia]
	default:
		link = ""
		hod = ""
		cas = ""
	}
	return
}

func GetDayName(day time.Weekday)string {
	switch day {
	case 1:
		return "Pondelok"
	case 2:
		return "Utorok"
	case 3:
		return "Streda"
	case 4:
		return "Štvrtok"
	case 5:
		return "Piatok"
	default:
		return ""
	}
}