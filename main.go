package main


import ( "fmt" ; "net/http" ; "html/template")

//User is
type User struct {
	Name string
	Age uint16
	Money int16
	AvgGrades float64 
	Happiness float64
	Hobbies []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newHame string) {
	u.Name = newHame
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	//bob.setNewName("Alex")
	//fmt.Fprint(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ....
	//bob := User{name: "Bob", age: 25, money: -50, avgGrades: 4.2, happiness: 0.8}

	handleRequest()
}