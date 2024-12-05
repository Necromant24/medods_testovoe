package main

import (
	"medods/auth-service/config"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {

	// str := "Y3ajc0va8EIJISH2PgSReHsqdGduOcA9114Gw+7TykZSQ8sf180VNB6M6Vmt0ikvQm8gdbezSo8g8MGdykguCg=="

	// hash := sha256.Sum256([]byte(str))
	// preHashedPassword := hex.EncodeToString(hash[:])

	// // Генерация второго bcrypt хэша для той же строки
	// h2, err := bcrypt.GenerateFromPassword([]byte(preHashedPassword), bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println("Error generating hash: %v", err)
	// }

	// // Печатаем хэши для проверки
	// fmt.Printf("Hash 2: %s\n", h2)

	// hash2 := sha256.Sum256([]byte(str))
	// preHashedPassword2 := hex.EncodeToString(hash2[:])

	// // Сравнение хэша с паролем
	// err = bcrypt.CompareHashAndPassword(h2, []byte(preHashedPassword2)) // Сравниваем h2 с исходной строкой
	// if err != nil {
	// 	fmt.Println("Error comparing hashes:", err)
	// } else {
	// 	fmt.Println("Password matches the hash!")
	// }

	// return

	config.LoadConfiguration()

	port := config.GetConfig().HostPort

	usersController := InjectUsersController()
	tokensController := InjectTokensController()

	r := chi.NewRouter()
	r.HandleFunc("/users/{userId}", usersController.GetUserById)
	r.HandleFunc("/getpair/{userId}", tokensController.GetTokensPair)
	r.HandleFunc("/refresh", tokensController.RefreshTokensPair)

	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
