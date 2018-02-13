package oldschool

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

var (
	rooms = map[string]*room{}
	store = sessions.NewCookieStore([]byte("session-key"))
)

type state struct {
	w http.ResponseWriter
	r *http.Request
	s *sessions.Session
}

func (s state) held() map[string]bool {
	if _, found := s.s.Values["held"]; !found {
		s.s.Values["held"] = map[string]bool{
			"Släckt ficklampa": true,
		}
	}
	return s.s.Values["held"].(map[string]bool)
}

func (s state) action() string {
	return s.r.Form.Get("action")
}

type room struct {
	title   string
	desc    func(state) []string
	exits   func(state) []*room
	actions func(state) []string
}

func makeRoom(
	title string,
	desc func(state) []string,
) *room {
	r := &room{
		title: title,
		desc:  desc,
		exits: func(s state) []*room {
			return nil
		},
		actions: func(s state) []string {
			return nil
		},
	}
	rooms[r.title] = r
	return r
}

func (ro *room) render(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "oldschool")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer session.Save(r, w)

	s := state{
		w: w,
		r: r,
		s: session,
	}

	thingsUL := []string{}
	exitsUL := []string{}
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

	actionsUL := []string{}
	for _, action := range ro.actions(s) {
		actionsUL = append(
			actionsUL,
			fmt.Sprintf(
				"<li><form method='post' action='/'><input type='hidden' name='location' value='%s'><input type='hidden' name='action' value='%s'><input type='submit' value='%s'></form></li>",
				ro.title,
				action,
				action,
			),
		)
	}

	descDIVs := []string{}
	for _, line := range ro.desc(s) {
		descDIVs = append(descDIVs, fmt.Sprintf("<div>%s</div>", line))
	}

	for thing := range s.held() {
		thingsUL = append(thingsUL, fmt.Sprintf("<li>%s</li>", thing))
	}

	fmt.Fprintf(w, `
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
	if len(actionsUL) > 0 {
		fmt.Fprintf(w, `
<h3>Du har</h3>
<ul>
%s
</ul>
`,
			strings.Join(thingsUL, ""))
	}
	if len(actionsUL) > 0 {
		fmt.Fprintf(w, `
<h3>Du kan</h3>
<ul>
%s
</ul>
`,
			strings.Join(actionsUL, ""))
	}
	if len(exitsUL) > 0 {
		fmt.Fprintf(w, `
<h3>Du kan gå till</h3>
<ul>
%s
</ul>
`,
			strings.Join(exitsUL, ""))
	}
	fmt.Fprintf(w, `
</body>
</html>
`)

}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=UTF-8")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	location := r.Form.Get("location")
	if source, found := rooms[location]; found {
		source.render(w, r)
	} else {
		start.render(w, r)
	}
}

func init() {
	http.Handle("/", http.HandlerFunc(root))
}
