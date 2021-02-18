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

func Hodiny(dalsia int) (string, string, string) {
	hodiny := config.Hodiny
	minuty := config.Minuty
	t := time.Now()
	h := t.Hour()
	m := t.Minute()
	if t.Weekday() > 0 && t.Weekday() < 6{
		switch {
		case h <= hodiny[1] ||(h==hodiny[1] && m < minuty[1]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[0+dalsia]
			cas := sd.Casy[0+dalsia]
			link :=  sd.Linky[0+dalsia]
			return hod, link, cas
		case (h == hodiny[2] && m >= minuty[1]) || (h == hodiny[3] && m < minuty[3]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[1+dalsia]
			cas := sd.Casy[1+dalsia]
			link :=  sd.Linky[1+dalsia]
			return hod, link, cas
		case (h == hodiny[4] && m >= minuty[3]) || (h == hodiny[5] && m < minuty[5]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[2+dalsia]
			cas := sd.Casy[2+dalsia]
			link :=  sd.Linky[2+dalsia]
			return hod, link, cas
		case (h == hodiny[6] && m >= minuty[5]) || (h == hodiny[7] && m < minuty[7]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[3+dalsia]
			cas := sd.Casy[3+dalsia]
			link :=  sd.Linky[3+dalsia]
			return hod, link, cas
		case (h == hodiny[8] && m >= minuty[7]) || (h == hodiny[9] && m < minuty[9]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[4+dalsia]
			cas := sd.Casy[4+dalsia]
			link :=  sd.Linky[4+dalsia]
			return hod, link, cas
		case (h == hodiny[10] && m >= minuty[9]) || (h == hodiny[11] && m < minuty[11]):
			sd := GetSChoolday(0)
			hod := sd.Hodiny[5+dalsia]
			cas := sd.Casy[5+dalsia]
			link :=  sd.Linky[5+dalsia]
			return hod, link, cas
		case h == hodiny[12] && m >= minuty[11]:
			sd := GetSChoolday(0)
			hod := sd.Hodiny[6+dalsia]
			cas := sd.Casy[6+dalsia]
			link :=  sd.Linky[6+dalsia]
			return hod, link, cas
		default:
			link := "Momentalne nie je žiadna hodina"
			hod := ""
			cas := ""
			return hod, link, cas
		}
	} else {
		link := "Momentalne nie je žiadna hodina"
		hod := ""
		cas := ""
		return hod, link, cas
	}
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