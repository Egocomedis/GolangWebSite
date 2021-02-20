package main


import ( "fmt" ; "net/http" ; "html/template")

//User is
type User struct {
	name string
	age uint16
	money int16
	avgGrades float64 
	happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newHame string) {
	u.name = newHame
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
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