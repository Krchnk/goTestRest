package main

import (
	"fmt"
	"net/http"

	"test/internal/storage/psql"
	"test/pkg/domain"
	"test/pkg/storage"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("start")

	// Инициализация базы данных
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Автоматическая миграция структуры Msg
	db.AutoMigrate(&storage.Msg{})

	// Создание экземпляра структуры PSQL
	p := psql.NewPSQL(db)

	fmt.Println("start server")
	mux := http.NewServeMux()

	mux.HandleFunc("/msg", func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()
		msg := queryParams.Get("msg") // Получаем значение параметра "msg"

		// Создание нового сообщения
		newMsg := &domain.Msg{
			TimeStamp: uint64(time.Time.Unix(time.Now())),
			Text:      msg,
		}

		createdMsg, err := p.CreateMsg(newMsg)
		if err != nil {
			fmt.Println("Ошибка при создании сообщения:", err)
		} else {
			fmt.Printf("Сообщение создано: ID = %d, Text = %s\n", createdMsg.ID, createdMsg.Text)
		}

		fmt.Fprintf(w, "response")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}

}
