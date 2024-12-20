package Read

import (
	"fmt"
	"net/http"
)

func ReadTasks(w http.ResponseWriter, req *http.Request) {
	db, err := connectReadDB()

	if err != nil {
		fmt.Printf("ERR OCCURED AT READTASKS FUNCTION 1st ERROR\n %s", err)
	}

	// if no error then fetch all the Tasks.
	data, err := db.Query("SELECT * FROM Tasks;")

	if err != nil {
		fmt.Printf("ERR OCCURED AT READTASKS FUNCTION 2nd ERROR\n %s", err)
	}

	fmt.Println(data)

}
