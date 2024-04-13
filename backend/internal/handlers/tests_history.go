package handlers

// func (h *Handler) TestsHistory(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		ErrorHandler(w, http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != http.MethodGet {
// 		ErrorHandler(w, http.StatusMethodNotAllowed)
// 		return
// 	}

// 	tmpl, err := template.ParseFiles("templates/main.html")
// 	if err != nil {
// 		ErrorHandler(w, http.StatusInternalServerError)
// 		return
// 	}

// 	userID := r.FormValue("id")
// 	if _, err := strconv.Atoi(userID); err != nil {
// 		ErrorHandler(w, http.StatusBadRequest)
// 		return
// 	}

// 	tests, err := GetTestsByUserID(userID, "../../storage/storage.db")
// 	if err != nil {
// 		ErrorHandler(w, http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, tests)
// 	if err != nil {
// 		ErrorHandler(w, http.StatusInternalServerError)
// 		return
// 	}
// }

// func GetTestsByUserID(userID, dbName string) ([]models.Test, error) {
// 	db, err := sql.Open("sqlite3", dbName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Ensure the database connection is valid
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Query tests with the given user ID
// 	entries, err := db.Query("SELECT id, user_id, test_type_id, score, feedback FROM tests WHERE user_id = ?", userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer entries.Close()

// 	// Iterate over the rows and scan into Test structs
// 	var tests []models.Test
// 	for entries.Next() {
// 		var test models.Test
// 		err := entries.Scan(&test.ID, &test.UserID, &test.TestTypeID, &test.Score, &test.Feedback)
// 		if err != nil {
// 			return nil, err
// 		}
// 		tests = append(tests, test)
// 	}

// 	return tests, nil
// }
