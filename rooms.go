package oldschool

import "log"

func init() {
	start := makeRoom(
		"Utanför grottan",
		func(s *state) []string {
			return []string{
				"Du står på en stig i en granskog.",
				"Solen tittar fram mellan träden ovanför dig.",
				"Bakom dig finns en klippa med en mörk grottöppning.",
				"Framför dig fortsätter stigen djupare in i skogen.",
			}
		},
	)
	defaultRoom = start
	cave := makeRoom(
		"Grottan",
		func(s *state) []string {
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
	start.exits = func(s *state) []*room {
		return []*room{cave}
	}
	cave.exits = func(s *state) []*room {
		return []*room{start}
	}

	lightOn := makeThing("Tänd ficklampa")
	lightOff := makeThing("Släckt ficklampa")
	defaultThings = map[string]bool{
		lightOff.name: true,
	}
	lightOn.actions = func(s *state) map[string]func(*state) {
		return map[string]func(*state){
			"Släck ficklampan": func(s *state) {
				s.swapHeld(lightOn.name, lightOff.name)
			},
		}
	}
	lightOff.actions = func(s *state) map[string]func(*state) {
		return map[string]func(*state){
			"Tänd ficklampan": func(s *state) {
				s.swapHeld(lightOff.name, lightOn.name)
			},
		}
	}
}
