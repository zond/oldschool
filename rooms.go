package oldschool

import "log"

var (
	start = makeRoom(
		"Utanför grottan",
		func(s state) []string {
			return []string{
				"Du står på en stig i en granskog.",
				"Solen tittar blygt fram mellan träden ovanför dig.",
				"Bakom dig finns en klippa med en mörk grottöppning.",
				"Framför dig fortsätter stigen djupare in i skogen.",
			}
		},
	)
	cave = makeRoom(
		"Grottan",
		func(s state) []string {
			if s.held()["Tänd ficklampa"] {
				return []string{
					"Du är inne i en grotta.",
					"Din ficklampa lyser upp i mörkret.",
					"Det droppar vatten från en stalaktit.",
					"Längst in i grottan ligger en liten hund och sover.",
				}
			} else {
				return []string{
					"Du är inne i en grotta.",
					"Det är kolsvart och du ser inget alls.",
					"Du hör något som droppar, och... något som andas?",
				}
			}
		},
	)
)

func init() {
	start.exits = func(s state) []*room {
		return []*room{cave}
	}
	cave.exits = func(s state) []*room {

		return []*room{start}
	}
	cave.actions = func(s state) []string {
		if s.held()["Släckt ficklampa"] {
			if s.action() == "Tänd ficklampan" {
				s.swapHeld("Släckt ficklampa", "Tänd ficklampa")
				log.Printf("swapped lamp, held are %+v", s.held())
			}
		} else if s.held()["Tänd ficklampa"] {
			if s.action() == "Släck ficklampan" {
				s.swapHeld("Tänd ficklampa", "Släckt ficklampa")
			}
		}
		if s.held()["Släckt ficklampa"] {
			return []string{"Tänd ficklampan"}
		} else if s.held()["Tänd ficklampa"] {
			return []string{"Släck ficklampan"}
		}
		return nil
	}
}
