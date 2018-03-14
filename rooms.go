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
	sword := makeThing("Svärd")
	libraryNook := makeRoom(
		"Vrå bredvid biblioteket",
		func(s *state) []string {
			return []string{
				"Du har krupit in i en liten vrå bredvid biblioteket.",
				"Den är precis stor nog för dig, här får inga drakar plats.",
				"Det luktar lite unket, och det enda ljuset kommer in under dörren ut till biblioteket.",
				"Du snubblar på tröskeln in i vrån och kommer åt en knapp på väggen, så att en hemlig dörr öppnas längst in i vrån.",
			}
		},
	)
	defaultRoom = start
	corridor := makeRoom(
		"Korridoren innanför vrån",
		func(s *state) []string {
			return []string{
				"Du går genom en dammig och mörk gammal korridor.",
				"Längs väggarna står murkna gamla möbler dammiga vita lakan liggande över.",
				"I hörnen är stora spindelnät med jättestora spindlar i.",
			}
		},
	)
	ghostBedroom := makeRoom(
		"Spöksovrummet",
		func(s *state) []string {
			rval := []string{
				"Du är i en sovsal med våningssängar längs väggarna.",
				"I sängarna sitter och ligger drösvis med spöken.",
				"Spökerna skriker av glädje när de ser dig.",
			}
			if s.s.Values["roomAction"] == "Prata med spökena" {
				rval = append(rval, "Det största och läskigaste spöket berättar att de inte vågat gå ut ur det här rummet på 200 år eftersom de är så rädda för draken.", "Nu är de jätteglada för att du kommit och kan skrämma bort draken från biblioteket.")
			}
			if s.s.Values["roomAction"] == "Berätta för spökena att draken smitit från biblioteket." {
				rval = append(rval, "Spökena blir ännu gladare, och börjar skrika av glädje igen.", "Men, berättar det största och läskigaste spöket, vi undrar om du kunde hjälpa oss lite till?", "Draken har flytt till sin skattkammare, och vi vill gärna ha hela slottet för oss själva.", "Kan du kanske skrämma bort draken från skattkammaren också?", "Spöket ler som en bilförsäljare.")
			}
			return rval
		},
	)
	castleLibrary := makeRoom(
		"Spökbiblioteket",
		func(s *state) []string {
			rval := []string{
				"Du står i ett jättestort bibliotek.",
				"Det är bokhyllor och böcker så långt du kan se, och flera våningar högt.",
				"Bredvid dig ser du en riddarrustning med ett svärd.",
			}
			if s.s.Values["roomAction"] == "Ta svärdet från rustningen" {
				rval = append(rval, "En av bokhyllorna glider undan med ett skrapande rasslande och du ser en liten dörr skymta fram bakom.")
			}
			if s.s.Values["roomAction"] == "Slå draken med svärdet" {
				rval = append(rval, "Draken väser till av skräck och flyger iväg över bokhyllorna.")
			} else {
				if s.s.Values["dragonState"] == "gone" {
					rval = append(rval, "Du ser spår i dammet efter något stort djur, eller monster?")
				} else {
					rval = append(
						rval,
						"Framför dig står en stor grön drake med gul mage, gula vingar och taggar längs hela kroppen.",
						"Draken kommer emot dig. Den vrålar och du känner dess stinkande andedräkt.",
					)
				}
			}
			return rval
		},
	)
	smallRock := makeThing("En liten sten")
	castleKey := makeThing("En slottsnyckel")
	ghostCastle := makeRoom(
		"Spökslottet",
		func(s *state) []string {
			desc := []string{
				"Under de höga tornen ser du höga fönster och ett par stora dörrar som vaktas av riddarstatyer.",
				"Mellan riddarstatyerna ser du en nyckel i ett hål.",
			}
			if s.s.Values["rockPos"] == "hole" {
				desc = append(desc, "I hålet ligger också en liten sten.")
			}
			if s.s.Values["roomAction"] == "Ta nyckeln från hålet" && s.s.Values["rockPos"] != "hole" {
				desc = append(desc, "Riddarstatyerna slår undan din hand när du försöker ta nyckeln.")
			}
			if !s.held()[castleKey.name] {
				desc = append(desc, "När du närmar dig så ser du hur statyerna flyttar sig så att de står framför dörren.")
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
	ghostCastle.exits = func(s *state) []*room {
		if s.held()[castleKey.name] {
			return []*room{start, castleLibrary}
		} else {
			return []*room{start}
		}
	}
	castleLibrary.exits = func(s *state) []*room {
		if s.held()[sword.name] {
			return []*room{ghostCastle, libraryNook}
		} else {
			return []*room{ghostCastle}
		}
	}
	castleLibrary.actions = func(s *state) map[string]func(*state) {
		if !s.held()[sword.name] {
			return map[string]func(*state){
				"Ta svärdet från rustningen": func(s *state) {
					s.held()[sword.name] = true
				},
			}
		} else if s.s.Values["dragonState"] == "gone" {
			return map[string]func(*state){}
		} else {
			return map[string]func(*state){
				"Slå draken med svärdet": func(s *state) {
					s.s.Values["dragonState"] = "gone"
				},
			}
		}
	}
	libraryNook.exits = func(s *state) []*room {
		return []*room{castleLibrary, corridor}
	}
	corridor.exits = func(s *state) []*room {
		return []*room{libraryNook, ghostBedroom}
	}
	ghostBedroom.exits = func(s *state) []*room {
		return []*room{corridor}
	}
	ghostBedroom.actions = func(s *state) map[string]func(*state) {
		rval := map[string]func(*state){
			"Prata med spökena": func(s *state) {
			},
		}
		if s.s.Values["dragonState"] == "gone" {
			rval["Berätta för spökena att draken smitit från biblioteket."] = func(s *state) {
			}
		}
		return rval
	}
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
		rval := map[string]func(*state){}
		if !s.held()[castleKey.name] {
			if s.s.Values["rockPos"] == "hole" {
				rval["Ta nyckeln från hålet"] = func(s *state) {
					s.held()[castleKey.name] = true
				}
			} else {
				rval["Ta nyckeln från hålet"] = func(s *state) {
				}
			}
		}
		if !s.held()[smallRock.name] && s.s.Values["rockPos"] != "hole" {
			rval["Plocka upp en sten från marken"] = func(s *state) {
				s.held()[smallRock.name] = true
			}
		}
		return rval
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
