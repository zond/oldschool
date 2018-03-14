package oldschool

func init() {
	start := makeRoom(
		"Utanför grottan",
		func(s *state) []string {
			return []string{
				"Du står på en stig i en granskog.",
				"Solen tittar fram mellan träden ovanför dig.",
				"Bakom dig finns en klippa med en mörk grottöppning.",
				"Framför dig fortsätter stigen djupare in i skogen.",
				"Du ser en grind och höga torn bakom den.",
			}
		},
	)
	defaultRoom = start
	ghostCastle := makeRoom(
		"Spökslottet",
		func(s *state) []string {
			desc := []string{
				"Under de höga tornen ser du höga fönster och ett par stora dörrar som vaktas av riddarstatyer.",
				"När du närmar dig så ser du hur statyerna flyttar sig så att de står framför dörren.",
				"Mellan riddarstatyerna ser du en nyckel i ett hål.",
			}
			if s.s.Values["rockPos"] == "hole" {
				desc = append(desc, "I hålet ligger också en liten sten.")
			}
			return desc
		},
	)
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
		return []*room{cave, ghostCastle}
	}
	cave.exits = func(s *state) []*room {
		return []*room{start}
	}
	smallRock := makeThing("En liten sten")
	smallRock.actions = func(s *state) map[string]func(*state) {
		if s.s.Values["location"] == ghostCastle.title && s.held()[smallRock.name] {
			return map[string]func(*state){
				"Lägg stenen i hålet mellan riddarna": func(s *state) {
					s.s.Values["rockPos"] = "hole"
					delete(s.held(), smallRock.name)
				},
			}
		} else {
			return map[string]func(*state){}
		}
	}
	ghostCastle.actions = func(s *state) map[string]func(*state) {
		if !s.held()[smallRock.name] && s.s.Values["rockPos"] != "hole" {
			return map[string]func(*state){
				"Plocka upp en sten från marken": func(s *state) {
					s.held()[smallRock.name] = true
				},
			}
		} else {
			return map[string]func(*state){}
		}
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
