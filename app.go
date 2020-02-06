package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/gorilla/sessions"
	"google.golang.org/appengine"
)

var (
	rooms         = map[string]*room{}
	things        = map[string]*thing{}
	store         = sessions.NewCookieStore([]byte("session-key"))
	defaultRoom   *room
	defaultThings map[string]bool
)

type state struct {
	w http.ResponseWriter
	r *http.Request
	s *sessions.Session
}

func (s *state) String() string {
	return fmt.Sprintf("%+v", s.s)
}

func (s *state) held() map[string]bool {
	return s.s.Values["held"].(map[string]bool)
}

func (s *state) action() string {
	return s.r.Form.Get("action")
}

func (s *state) swapHeld(a, b string) {
	delete(s.held(), a)
	s.held()[b] = true
}

func (s *state) save() error {
	return s.s.Save(s.r, s.w)
}

type room struct {
	title   string
	desc    func(*state) []string
	exits   func(*state) []*room
	actions func(*state) map[string]func(*state)
}

func makeRoom(
	title string,
	desc func(*state) []string,
) *room {
	r := &room{
		title: title,
		desc:  desc,
		exits: func(s *state) []*room {
			return nil
		},
		actions: func(s *state) map[string]func(*state) {
			return map[string]func(*state){}
		},
	}
	rooms[r.title] = r
	return r
}

type thing struct {
	name    string
	actions func(*state) map[string]func(*state)
}

func makeThing(
	name string,
) *thing {
	t := &thing{
		name: name,
		actions: func(s *state) map[string]func(*state) {
			return map[string]func(*state){}
		},
	}
	things[t.name] = t
	return t
}

func (ro *room) render(s *state) {
	if action, ok := s.s.Values["roomAction"].(string); ok && action != "" {
		if f, found := ro.actions(s)[action]; found {
			f(s)
		}
	}

	if action, ok := s.s.Values["thingAction"].(string); ok && action != "" {
		if th, found := things[s.s.Values["thing"].(string)]; found {
			if f, found := th.actions(s)[action]; found {
				f(s)
			}
		}
	}

	actionsUL := []string{}
	for action := range ro.actions(s) {
		actionsUL = append(
			actionsUL,
			fmt.Sprintf(
				"<li><form method='post' action='/'><input type='hidden' name='location' value='%s'><input type='hidden' name='roomAction' value='%s'><input type='submit' value='%s'></form></li>",
				ro.title,
				action,
				action,
			),
		)
	}

	for th := range s.held() {
		for action := range things[th].actions(s) {
			actionsUL = append(
				actionsUL,
				fmt.Sprintf(
					"<li><form method='post' action='/'><input type='hidden' name='location' value='%s'><input type='hidden' name='thingAction' value='%s'><input type='submit' value='%s'><input type='hidden' name='thing' value='%s'></form></li>",
					ro.title,
					action,
					action,
					th,
				),
			)
		}
	}

	exitsUL := sort.StringSlice{}
	for _, exit := range ro.exits(s) {
		exitsUL = append(
			exitsUL,
			fmt.Sprintf(
				"<li><form method='post' action='/'><input type='hidden' name='location' value='%s'><input type='submit' value='%s'></form></li>",
				exit.title,
				exit.title,
			),
		)
	}
	sort.Sort(exitsUL)

	descDIVs := []string{}
	for _, line := range ro.desc(s) {
		descDIVs = append(descDIVs, fmt.Sprintf("<div>%s</div>", line))
	}

	thingsUL := sort.StringSlice{}
	for thing := range s.held() {
		thingsUL = append(thingsUL, fmt.Sprintf("<li>%s</li>", thing))
	}
	sort.Sort(thingsUL)

	delete(s.s.Values, "roomAction")
	delete(s.s.Values, "thingAction")
	s.save()

	s.w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(s.w, `
<html>
<head>
<title>%s</title>
<style>
body {
	background-color: black;
	color: white;
	font-size: x-large;
	margin: auto;
	width: 50%%;
}
</style>
</head>
<body>
<h1>%s</h1>
%s
`,
		ro.title,
		ro.title,
		strings.Join(descDIVs, ""),
	)
	if len(thingsUL) > 0 {
		fmt.Fprintf(s.w, `
<h3>Du har</h3>
<ul>
%s
</ul>
`,
			strings.Join(thingsUL, ""))
	}
	if len(actionsUL) > 0 {
		fmt.Fprintf(s.w, `
<h3>Du kan</h3>
<ul>
%s
</ul>
`,
			strings.Join(actionsUL, ""))
	}
	if len(exitsUL) > 0 {
		fmt.Fprintf(s.w, `
<h3>Du kan gå till</h3>
<ul>
%s
</ul>
`,
			strings.Join(exitsUL, ""))
	}
	fmt.Fprintf(s.w, `
</body>
</html>
`)

}

func root(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "oldschool")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("Got session %+v", session)

	saveOnce := &sync.Once{}
	saveSession := func() {
		saveOnce.Do(func() {
			log.Printf("Saving session %+v", session)
			if err := session.Save(r, w); err != nil {
				http.Error(w, err.Error(), 500)
				panic(err)
			}
		})
	}

	defer saveSession()

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	s := &state{
		w: w,
		r: r,
		s: session,
	}
	if location := r.Form.Get("location"); location != "" {
		s.s.Values["location"] = location
	}
	if action := r.Form.Get("roomAction"); action != "" {
		s.s.Values["roomAction"] = action
	}
	if action := r.Form.Get("thingAction"); action != "" {
		s.s.Values["thingAction"] = action
		s.s.Values["thing"] = r.Form.Get("thing")
	}

	if r.Method == "POST" {
		saveSession()
		log.Println("Redirect due to POST")
		http.Redirect(w, r, "/", 303)
		return
	}

	if action, found := s.s.Values["action"]; found {
		s.r.Form.Set("action", action.(string))
		delete(s.s.Values, "action")
	}

	_, ok := s.s.Values["held"].(map[string]bool)
	if !ok {
		s.s.Values["held"] = defaultThings
		saveSession()
	}
	location, ok := s.s.Values["location"].(string)
	if !ok {
		s.s.Values["location"] = defaultRoom.title
		saveSession()
		location = defaultRoom.title
	}

	if source, found := rooms[location]; found {
		source.render(s)
	} else {
		http.Error(w, "Gulp", 500)
	}
}

func main() {
	gob.Register(map[string]bool{})
	http.Handle("/", http.HandlerFunc(root))
	appengine.Main()
}
