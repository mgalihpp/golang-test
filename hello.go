// // Copyright 2011 The Go Authors. All rights reserved.
// // Use of this source code is governed by a BSD-style
// // license that can be found in the LICENSE file.

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// const (
// 	win            = 100 // The winning score in a game of Pig
// 	gamesPerSeries = 10  // The number of games per series to simulate
// )

// // A score includes scores accumulated in previous turns for each player,
// // as well as the points scored by the current player in this turn.
// type score struct {
// 	player, opponent, thisTurn int
// }

// // An action transitions stochastically to a resulting score.
// type action func(current score) (result score, turnIsOver bool)

// // roll returns the (result, turnIsOver) outcome of simulating a die roll.
// // If the roll value is 1, then thisTurn score is abandoned, and the players'
// // roles swap.  Otherwise, the roll value is added to thisTurn.
// func roll(s score) (score, bool) {
// 	outcome := rand.Intn(6) + 1 // A random int in [1, 6]
// 	if outcome == 1 {
// 		return score{s.opponent, s.player, 0}, true
// 	}
// 	return score{s.player, s.opponent, outcome + s.thisTurn}, false
// }

// // stay returns the (result, turnIsOver) outcome of staying.
// // thisTurn score is added to the player's score, and the players' roles swap.
// func stay(s score) (score, bool) {
// 	return score{s.opponent, s.player + s.thisTurn, 0}, true
// }

// // A strategy chooses an action for any given score.
// type strategy func(score) action

// // stayAtK returns a strategy that rolls until thisTurn is at least k, then stays.
// func stayAtK(k int) strategy {
// 	return func(s score) action {
// 		if s.thisTurn >= k {
// 			return stay
// 		}
// 		return roll
// 	}
// }

// // play simulates a Pig game and returns the winner (0 or 1).
// func play(strategy0, strategy1 strategy) int {
// 	strategies := []strategy{strategy0, strategy1}
// 	var s score
// 	var turnIsOver bool
// 	currentPlayer := rand.Intn(2) // Randomly decide who plays first
// 	for s.player+s.thisTurn < win {
// 		action := strategies[currentPlayer](s)
// 		s, turnIsOver = action(s)
// 		if turnIsOver {
// 			currentPlayer = (currentPlayer + 1) % 2
// 		}
// 	}
// 	return currentPlayer
// }

// // roundRobin simulates a series of games between every pair of strategies.
// func roundRobin(strategies []strategy) ([]int, int) {
// 	wins := make([]int, len(strategies))
// 	for i := 0; i < len(strategies); i++ {
// 		for j := i + 1; j < len(strategies); j++ {
// 			for k := 0; k < gamesPerSeries; k++ {
// 				winner := play(strategies[i], strategies[j])
// 				if winner == 0 {
// 					wins[i]++
// 				} else {
// 					wins[j]++
// 				}
// 			}
// 		}
// 	}
// 	gamesPerStrategy := gamesPerSeries * (len(strategies) - 1) // no self play
// 	return wins, gamesPerStrategy
// }

// // ratioString takes a list of integer values and returns a string that lists
// // each value and its percentage of the sum of all values.
// // e.g., ratios(1, 2, 3) = "1/6 (16.7%), 2/6 (33.3%), 3/6 (50.0%)"
// func ratioString(vals ...int) string {
// 	total := 0
// 	for _, val := range vals {
// 		total += val
// 	}
// 	s := ""
// 	for _, val := range vals {
// 		if s != "" {
// 			s += ", "
// 		}
// 		pct := 100 * float64(val) / float64(total)
// 		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
// 	}
// 	return s
// }

// func main() {
// 	strategies := make([]strategy, win)
// 	for k := range strategies {
// 		strategies[k] = stayAtK(k + 1)
// 	}
// 	wins, games := roundRobin(strategies)

// 	for k := range strategies {
// 		fmt.Printf("Wins, losses staying at k =% 4d: %s\n",
// 			k+1, ratioString(wins[k], games-wins[k]))
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	// Membuat seed acak berdasarkan waktu
// 	rand.Seed(time.Now().UnixNano())

// 	// Menghasilkan angka acak antara 1 dan 100
// 	secretNumber := rand.Intn(100) + 1
// 	fmt.Println("Halo! Ayo main tebak angka.")
// 	fmt.Println("Saya telah memilih sebuah angka antara 1 dan 100.")

// 	// Menebak angka
// 	var guess int
// 	for {
// 		fmt.Print("Tebak angka: ")
// 		fmt.Scanln(&guess)

// 		// Memeriksa tebakan
// 		if guess < secretNumber {
// 			fmt.Println("Terlalu kecil! Coba lagi.")
// 		} else if guess > secretNumber {
// 			fmt.Println("Terlalu besar! Coba lagi.")
// 		} else {
// 			fmt.Println("Selamat! Anda menebak dengan benar!")
// 			break
// 		}
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// // Struct untuk menyimpan data video
// type Video struct {
// 	ID    string `json:"id"`
// 	Title string `json:"title"`
// 	Description string `json:"desc"`
// }

// func main() {
// 	// Mendaftarkan handler untuk route "/videos"
// 	http.HandleFunc("/videos", getVideos)

// 	// Mulai server pada port 8080
// 	log.Println("Server is running on port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // Handler untuk route "/videos"
// func getVideos(w http.ResponseWriter, r *http.Request) {
// 	// Membuat data dummy video
// 	videos := []Video{
// 		{ID: "1", Title: "Video 1"},
// 		{ID: "2", Title: "Video 2"},
// 		{ID: "3", Title: "Video 3", Description: "apa aaja"},
// 	}

// 	// Set header Content-Type ke application/json
// 	w.Header().Set("Content-Type", "application/json")

// 	// Mengirim response berupa JSON yang berisi data video
// 	json.NewEncoder(w).Encode(videos)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// URL endpoint yang akan di-fetch
// 	url := "https://yourtube-six.vercel.app/api/trpc/video.getRandomVideo?batch=1&input=%7B%220%22%3A%7B%22json%22%3A40%7D%7D"

// 	// Lakukan HTTP GET request ke URL
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalf("Failed to fetch data : %v", err)
// 	}
// 	defer response.Body.Close()

// 	// Dekode response JSON ke dalam slice yang sesuai
// 	var jsonData []interface{}
// 	if err := json.NewDecoder(response.Body).Decode(&jsonData); err != nil {
// 		log.Fatalf("Failed to decode JSON: %v", err)
// 	}

// 	// Cetak data yang diambil dari response JSON
// 	for index, item := range jsonData {
// 		fmt.Printf("Item %d: %+v\n", index, item)
// 	}
// }

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Struct untuk menyimpan informasi pengguna
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

func main() {
	// Koneksi ke database MySQL
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golangtest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat tabel pengguna jika belum ada
	createTable()

	// Menjalankan server HTTP
	http.HandleFunc("/login", handleLogin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Membuat tabel pengguna jika belum ada
func createTable() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler untuk proses login
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query ke database untuk mencari user dengan username dan password yang sesuai
	row := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", user.Username, user.Password)
	var userID int
	err = row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Jika user ditemukan, kirim response sukses
	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}
