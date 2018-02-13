package oldschool

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
		for thing := range s.held() {
			if thing == "Släckt ficklampa" {
				if s.action() == "Tänd ficklampan" {
					delete(s.held(), "Släckt ficklampa")
					s.held()["Tänd ficklampa"] = true
					return []string{"Släck ficklampan"}
				} else {
					return []string{"Tänd ficklampan"}
				}
			}
		}
		return nil
	}
}
