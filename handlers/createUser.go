package handlers

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// // Парсинг JSON из тела запроса
	// var user User
	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	// 	return
	// }

	// // Открываем соединение с базой данных
	// db, err := sql.Open("pgx", "user=postgres password=2048 host=localhost port=5432 dbname=goblogdb sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // SQL-запрос для вставки пользователя
	// _, err = db.Exec(`INSERT INTO users (user_name, user_email, password_hash) VALUES ($1, $2, $3) RETURNING user_id`,
	// 	user.UserName, user.UserEmail, user.PasswordHash)
	// if err != nil {
	// 	http.Error(w, "Failed to create user", http.StatusInternalServerError)
	// 	return
	// }

	// // Возвращаем успешный ответ с созданным пользователем
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(user)

}
