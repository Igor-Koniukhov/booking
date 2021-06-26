package handlers

import (
	"fmt"
	"github.com/Igor-Koniukhov/bookings/pkg/config"
	"github.com/Igor-Koniukhov/bookings/pkg/models"
	"github.com/Igor-Koniukhov/bookings/pkg/render"
	"net/http"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	//send the data to the template

	render.RenderTemplate(w, r,"about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"make-reservation.page.tmpl", &models.TemplateData{})
}

//Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"generals.page.tmpl", &models.TemplateData{})
}
//Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"majors.page.tmpl", &models.TemplateData{})
}
//Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"search-availability.page.tmpl", &models.TemplateData{})
}
//PostAvailability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request){
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("this is form value start: %v, and end: %v", start, end)))
}
//Contacts renders the contacts of organisation
func (m *Repository) Contacts(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"contacts.page.tmpl", &models.TemplateData{})
}
