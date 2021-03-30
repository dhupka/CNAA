package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	//"strconv"
)

func main() {
	db := database{"1": "task1", "2": "task2"}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/desc", db.desc)
	http.HandleFunc("/resolve", db.resolve)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var mutex = &sync.RWMutex{}
type database map[string]string 

//list
func (db database) list(w http.ResponseWriter, req *http.Request) {
	mutex.RLock()                 //locks for reading
	for id, task := range db { //print items in the database
		fmt.Fprintf(w, "%s: %s\n", id, task)
	}
	mutex.RUnlock()
}

//desc?id=2
func (db database) desc(w http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	id := req.URL.Query().Get("id")
	if task, ok := db[id]; ok {
		fmt.Fprintf(w, "%s\n", task)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404: task not in database
		fmt.Fprintf(w, "no such item: %q\n", task)
	}
	mutex.RUnlock()
}

//create?id=4&task=task4
func (db database) create(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	task := req.URL.Query().Get("task")
	//mutex.Lock() 
	db[id] = task
	//mutex.Unlock()
	fmt.Fprint(w, "ID: ", id, ". Task: ", task, "\n")
}

///update?id=3&task=newtask3 
func (db database) update(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	task := req.URL.Query().Get("task")


	if _, found := db[id]; found {
		mutex.Lock()
		db[id] = task
		mutex.Unlock()
		fmt.Fprint(w, "Updated task: ", id, ". Item task:  ", task, " \n") 
	} else {
		w.WriteHeader(http.StatusNotFound) // 404: id not in database
		fmt.Fprintf(w, "no such id: %q\n", id)
	}
}

//delete?id=1
func (db database) resolve(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if _, found := db[id]; found {
		mutex.Lock()
		delete(db, id)
		mutex.Unlock()
		fmt.Fprintf(w,"Deleted id %s\n", id)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404: id not in database
		fmt.Fprintf(w, "no such id: %q\n", id)
	}
}