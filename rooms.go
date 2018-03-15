package oldschool

func init() {
	fork := makeThing("Potatisgaffel i nysilver")
	lightOn := makeThing("Tänd ficklampa")
	lightOff := makeThing("Släckt ficklampa")
	start := makeRoom(
		"Utanför grottan",
		func(s *state) []string {
			rval := []string{
				"Du står på en stig i en granskog.",
				"Solen tittar fram mellan träden ovanför dig.",
				"Bredvid dig står ett ihåligt träd.",
				"Bakom dig finns en klippa med en mörk grottöppning.",
				"Inifrån grottan hörs ett dånande muller.",
				"Framför dig ser du en grind, och höga torn bakom den.",
			}
			if s.s.Values["roomAction"] == "Tryck in knappen i det ihåliga trädet." {
				rval = append(rval, "Du sträcker in armen så långt du kan, men kan inte riktigt nå ner. Och knappen är liksom halvt gömd bakom en knöl i stammen. Du behöver något smalt och lite krokigt för att komma åt den!")
			}
			if s.s.Values["roomAction"] == "Leta i det ihåliga trädet." {
				rval = append(rval, "Inne i det ihåliga trädet ser du, längst ner bland de murkna löven och nötterna ekorrarna glömde förra året, en liten knapp!")
			}
			if s.s.Values["roomAction"] == "Använd gaffeln för att peta in knappen i trädet." {
				rval = append(
					rval,
					"Knappen trycks in med ett svagt klick.",
					"Inifrån grottan blir det plötsligt tyst.",
				)
			}
			return rval
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
				"Längs väggarna står murkna gamla möbler med dammiga vita lakan liggande över.",
				"I hörnen är stora spindelnät med jättestora spindlar i.",
			}
		},
	)
	treasureRoom := makeRoom(
		"Skattkammaren",
		func(s *state) []string {
			rval := []string{
				"Du är inne i slottets skattkammare.",
				"Den är upplyst av facklor och små oljelampor.",
				"Du står bland högar av guld, ädelstenar, chokladbollar och sura kryptoniter.",
				"Längs väggarna står rustningar i guld och silver, och i taket hänger kristallkronor med julgransdekorationer och små kottar dinglande i presentsnören.",
				"Taket är så högt upp att du inte ens kan se det.",
			}
			if s.s.Values["roomAction"] == "Slå draken med svärdet." {
				rval = append(
					rval,
					"Draken väser av skräck, mumlar något som låter som 'neeeej, inte igeen!!', och flyger upp i mörkret och försvinner.",
				)
			}
			if s.s.Values["treasureState"] == "dragonGone" {
				if s.s.Values["roomAction"] != "Slå draken med svärdet." {
					rval = append(
						rval,
						"Du har befriat slottet från draken! Hurra!! Nu blir spökena glada.",
						"Spöken behöver ju inte guld... du kanske kan ta med dig lite?",
					)
				} else {
					rval = append(
						rval,
						"Bland guldet ser du spår av draken som tidigare bodde här.",
						"Spöken behöver ju inte guld... du kanske kan ta med dig lite?",
					)
				}
			} else {
				rval = append(
					rval,
					"Mitt i allt guldet vilar draken. Den ser väldigt arg ut!",
					"Den väser, och ringlar mot dig!!!",
				)
			}
			if s.s.Values["roomAction"] == "Ta en näve guld och ädelstenar." {
				rval = append(
					rval,
					"Du rullar runt bland guld och skatter och fyller fickorna med så mycket du får plats med.",
					"Nu blir det skatteåterbäring!",
				)
			}
			return rval
		},
	)
	ghostBedroom := makeRoom(
		"Spöksovrummet",
		func(s *state) []string {
			rval := []string{}
			if s.s.Values["ghostInstructions"] == "yes" {
				rval = []string{
					"Du är i en sovsal med våningssängar längs väggarna.",
					"En del spöken, bleka och urvattnade, svävar runt mellan sängarna.",
				}
			} else {
				rval = []string{
					"Du är i en sovsal med våningssängar längs väggarna.",
					"I sängarna sitter och ligger drösvis med spöken.",
					"Spökena skriker av glädje när de ser dig.",
				}
			}
			if s.s.Values["roomAction"] == "Prata med spökena." {
				rval = append(rval, "Det största och läskigaste spöket berättar att de inte vågat gå ut ur det här rummet på 200 år eftersom de är så rädda för draken.", "Nu är de jätteglada för att du kommit och kan skrämma bort draken från biblioteket.")
			}
			if s.s.Values["roomAction"] == "Berätta för spökena att draken smitit från biblioteket." {
				rval = append(
					rval,
					"Spökena blir ännu gladare, och börjar skrika av glädje igen.",
					"Men, berättar det största och läskigaste spöket, vi undrar om du kunde hjälpa oss lite till?",
					"Draken har flytt till sin skattkammare, och vi vill gärna ha hela slottet för oss själva.",
					"Kan du kanske skrämma bort draken från skattkammaren också?",
					"Spöket ler som en bilförsäljare och börjar, tillsammans med alla de andra spökena, blekna bort.")
			}
			if s.s.Values["roomAction"] == "Fråga spökena hur man hittar till skattkammaren." {
				rval = append(
					rval,
					"Spökena viskar nästan ohörbart: Följ den gröööönaaa sooooleeeen...",
				)
			}
			return rval
		},
	)
	darkStairs := makeRoom(
		"Den mörka trappan",
		func(s *state) []string {
			rval := []string{
				"Du står i en mörk och brant trappa.",
			}
			if s.held()[lightOn.name] {
				rval = append(rval, "Ficklampan lyser upp trappan och den ser väldigt brant och hal ut.")
			} else {
				rval = append(rval, "Det är kolmörkt, och du ser inget alls.", "Se upp så att du inte snubblar!")
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
			}
			if s.held()[sword.name] {
				rval = append(rval, "Bredvid dig ser du en riddarrustning.")
			} else {
				rval = append(rval, "Bredvid dig ser du en riddarrustning med ett svärd.")
			}
			if s.s.Values["roomAction"] == "Ta svärdet från rustningen." {
				rval = append(rval, "En av bokhyllorna glider undan med ett skrapande rasslande och du ser en liten dörr skymta fram bakom.")
			}
			if s.s.Values["roomAction"] == "Slå draken med svärdet." {
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
			if s.s.Values["roomAction"] == "Leta efter böcker om zombieläkare." {
				rval = append(
					rval,
					"Efter att du letat en stund hittar du en bok om läkare som blivit zombies.",
					"I boken står det att alla läkare som blir zombies blir otroligt engagerade och seriösa i sitt yrke, och aldrig skulle lämna en patient i sticket.",
				)
			}
			if s.s.Values["zombieState"] == "library" {
				rval = append(
					rval,
					"Bland bokhyllorna vacklar en zombie med läkarrock och stetoskop omkring och stönar: -Vaaaarrrrr äääärrr patieeeenteeeen....",
				)
			}
			return rval
		},
	)
	smallRock := makeThing("Liten sten")
	castleKey := makeThing("Slottsnyckel")
	ghostCastle := makeRoom(
		"Spökslottet",
		func(s *state) []string {
			desc := []string{
				"Under de höga tornen ser du höga fönster och ett par stora dörrar som vaktas av riddarstatyer.",
			}
			if !s.held()[castleKey.name] {
				desc = append(desc, "Mellan riddarstatyerna ser du en nyckel i ett hål.")
			}
			if s.s.Values["rockPos"] == "hole" {
				desc = append(desc, "I hålet ligger också en liten sten.")
			}
			if s.s.Values["roomAction"] == "Ta nyckeln från hålet." && s.s.Values["rockPos"] != "hole" {
				desc = append(desc, "Riddarstatyerna slår undan din hand när du försöker ta nyckeln.")
			}
			if !s.held()[castleKey.name] {
				desc = append(desc, "När du närmar dig så ser du hur statyerna flyttar sig så att de står framför dörren.")
			}
			if s.s.Values["greenSun"] == "yes" {
				desc = append(
					desc,
					"Det är ett väldigt konstigt grönt ljus som skiner uppe i himlen, nästan som en grön sol.",
					"Ljuset får det högsta tornet på slottet att kasta en skugga mot en stenvägg.",
					"Skuggan pekar på en väldigt liten diskret knapp på väggen.",
				)

			}
			if s.s.Values["roomAction"] == "Tryck in den lilla knappen." {
				desc = append(
					desc,
					"Med ett mäktigt muller så glider en stor sten i muren undan och avslöjar en trappa ner i en mörk och kuslig gång.",
				)
			}
			return desc
		},
	)
	buffetRoom := makeRoom(
		"Potatisbuffén",
		func(s *state) []string {
			s.s.Values["zombieSeen"] = "yes"
			rval := []string{
				"Du är i en stor matsal.",
				"På långa bord är en enorm dignande potatisbuffé uppdukad.",
				"Det finns stekt, kokt, råriven och grillad potatis. Bakad potatis, potatismos och potatisbullar.",
			}
			if s.s.Values["roomAction"] == "Säg till zombien att det hänt en olycka i biblioteket och att någon gjort illa sig där." {
				rval = append(
					rval,
					"Zombien spärrar upp ögonen och stönar: -Måååste taaaa haaaaand ooooom patieeeeentttttt!",
					"Sedan ramlar och raglar den uppför trappen och försvinner ur synhåll.",
				)
			} else {
				if s.s.Values["zombieState"] != "library" {
					rval = append(
						rval,
						"Vid borden går en zombie i vit rock och stetoskop runt och plockar bland potatisarna.",
						"När den ser dig ger den ifrån sig ett lågt stönande morr och börjar stappla mot dig med utsträckta armar.",
					)
				}
			}
			if s.s.Values["roomAction"] == "Smaka på potatisbuffén." {
				rval = append(
					rval,
					"Du går runt och plockar bland potatisarna, men upptäcker att de bakade potatisarna är varma som lava.",
					"Som tur är finns det en liten gaffel du kan använda för att slippa bränna fingarnar av dig.",
					"Potatisen är delikat med lite smör och gräslök!",
				)
			}
			return rval
		},
	)
	caveTunnel := makeRoom(
		"Tunneln i underjorden",
		func(s *state) []string {
			rval := []string{
				"Du befinner dig djupt ner i underjorden, i en smutsig och mörk gång.",
			}
			if s.held()[lightOn.name] {
				rval = append(rval, "Gången fortsätter långt framåt och neråt.")
			} else {
				rval = append(rval, "Det är kolmörkt och du ser ingenting.")
			}
			return rval
		},
	)
	cave := makeRoom(
		"Grottan",
		func(s *state) []string {
			rval := []string{}
			if s.held()["Tänd ficklampa"] {
				rval = append(
					rval,
					"Du är inne i en grotta.",
					"Din ficklampa lyser upp i mörkret.",
					"Längst in i grottan ligger en liten hund och sover.",
				)
				if s.s.Values["treeButtonPushed"] == "yes" {
					rval = append(
						rval,
						"I ett hörn av grottan, där det verkar ha forsat ett vattenfall, är en mörk öppning till en gång ner i underjorden.",
					)
				} else {
					rval = append(
						rval,
						"I ett hörn av grottan forsar ett litet vattenfall.",
					)
				}
			} else {
				rval = append(
					rval,
					"Du är inne i en grotta.",
					"Det är kolsvart och du ser inget alls.",
				)
				if s.s.Values["treeButtonPushed"] != "yes" {
					rval = append(
						rval,
						"Du hör ett muller av vatten som forsar, och... något som andas?",
					)
				}
			}
			return rval
		},
	)
	start.exits = func(s *state) []*room {
		return []*room{cave, ghostCastle}
	}
	start.actions = func(s *state) map[string]func(*state) {
		rval := map[string]func(*state){}
		if s.s.Values["treeButtonFound"] == "yes" {
			if s.held()[fork.name] {
				rval["Använd gaffeln för att peta in knappen i trädet."] = func(s *state) {
					s.s.Values["treeButtonPushed"] = "yes"
				}
			} else {
				rval["Tryck in knappen i det ihåliga trädet."] = func(s *state) {
				}
			}
		} else {
			rval["Leta i det ihåliga trädet."] = func(s *state) {
				s.s.Values["treeButtonFound"] = "yes"
			}
		}
		return rval
	}
	caveTunnel.exits = func(s *state) []*room {
		if s.held()[lightOn.name] {
			return []*room{cave, treasureRoom}
		} else {
			return []*room{}
		}
	}
	cave.exits = func(s *state) []*room {
		rval := []*room{start}
		if s.s.Values["treeButtonPushed"] == "yes" {
			rval = append(rval, caveTunnel)
		}
		return rval
	}
	darkStairs.exits = func(s *state) []*room {
		exits := []*room{ghostCastle}
		if s.held()[lightOn.name] {
			exits = append(exits, buffetRoom)
		}
		return exits
	}
	buffetRoom.exits = func(s *state) []*room {
		return []*room{darkStairs}
	}
	treasure := makeThing("Guld och ädelstenar i massor")
	treasureRoom.actions = func(s *state) map[string]func(*state) {
		rval := map[string]func(*state){}
		if s.s.Values["treasureState"] == "dragonGone" {
			rval["Ta en näve guld och ädelstenar."] = func(s *state) {
				s.held()[treasure.name] = true
			}
		} else {
			rval["Slå draken med svärdet."] = func(s *state) {
				s.s.Values["treasureState"] = "dragonGone"
			}
		}
		return rval
	}
	buffetRoom.actions = func(s *state) map[string]func(*state) {
		rval := map[string]func(*state){}
		if s.s.Values["bookFound"] == "yes" && s.s.Values["zombieState"] != "library" {
			rval["Säg till zombien att det hänt en olycka i biblioteket och att någon gjort illa sig där."] = func(s *state) {
				s.s.Values["zombieState"] = "library"
			}
		}
		if s.s.Values["zombieState"] == "library" {
			rval["Smaka på potatisbuffén."] = func(s *state) {
				s.held()[fork.name] = true
			}
		}
		return rval
	}
	ghostCastle.exits = func(s *state) []*room {
		exits := []*room{start}
		if s.held()[castleKey.name] {
			exits = append(exits, castleLibrary)
		}
		if s.s.Values["buttonPressed"] == "yes" {
			exits = append(exits, darkStairs)
		}
		return exits
	}
	castleLibrary.exits = func(s *state) []*room {
		if s.held()[sword.name] {
			return []*room{ghostCastle, libraryNook}
		} else {
			return []*room{ghostCastle}
		}
	}
	castleLibrary.actions = func(s *state) map[string]func(*state) {
		rval := map[string]func(*state){}
		if s.s.Values["zombieState"] != "library" {
			if !s.held()[sword.name] {
				rval["Ta svärdet från rustningen."] = func(s *state) {
					s.held()[sword.name] = true
				}
			} else if s.s.Values["dragonState"] != "gone" {
				rval["Slå draken med svärdet."] = func(s *state) {
					s.s.Values["dragonState"] = "gone"
				}
			}
			if s.s.Values["zombieSeen"] == "yes" {
				rval["Leta efter böcker om zombieläkare."] = func(s *state) {
					s.s.Values["bookFound"] = "yes"
				}
			}
		}
		return rval
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
			"Prata med spökena.": func(s *state) {
			},
		}
		if s.s.Values["dragonState"] == "gone" && s.s.Values["ghostInstructions"] != "yes" {
			rval["Berätta för spökena att draken smitit från biblioteket."] = func(s *state) {
				s.s.Values["ghostInstructions"] = "yes"
			}
		}
		if s.s.Values["ghostInstructions"] == "yes" {
			rval["Fråga spökena hur man hittar till skattkammaren."] = func(s *state) {
				s.s.Values["greenSun"] = "yes"
			}
		}
		return rval
	}
	smallRock.actions = func(s *state) map[string]func(*state) {
		if s.s.Values["location"] == ghostCastle.title && s.held()[smallRock.name] {
			return map[string]func(*state){
				"Lägg stenen i hålet mellan riddarna.": func(s *state) {
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
				rval["Ta nyckeln från hålet."] = func(s *state) {
					s.held()[castleKey.name] = true
				}
			} else {
				rval["Ta nyckeln från hålet."] = func(s *state) {
				}
			}
		}
		if !s.held()[smallRock.name] && s.s.Values["rockPos"] != "hole" {
			rval["Plocka upp en sten från marken."] = func(s *state) {
				s.held()[smallRock.name] = true
			}
		}
		if s.s.Values["greenSun"] == "yes" {
			rval["Tryck in den lilla knappen."] = func(s *state) {
				s.s.Values["buttonPressed"] = "yes"
			}
		}
		return rval
	}

	defaultThings = map[string]bool{
		lightOff.name: true,
	}
	lightOn.actions = func(s *state) map[string]func(*state) {
		return map[string]func(*state){
			"Släck ficklampan.": func(s *state) {
				s.swapHeld(lightOn.name, lightOff.name)
			},
		}
	}
	lightOff.actions = func(s *state) map[string]func(*state) {
		return map[string]func(*state){
			"Tänd ficklampan.": func(s *state) {
				s.swapHeld(lightOff.name, lightOn.name)
			},
		}
	}
}
