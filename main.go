package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {

	http.HandleFunc("/people", peopleHandler)      // HandleFunc - метод из пакета http, который позволяет настроить пути
	http.HandleFunc("/health", healthCheckHandler) // тоже самое почти

	log.Fatal(http.ListenAndServe(":8081", nil)) // логируем ошибку, прослушиваем локальный порт 8080, nil - используем DefaultServeMux - мультиплексирование трафика между различными маршрутами мультиплексирование - возможность выполнения множеста запросов в одной TCP сессии, то есть программа готова анализировать команды и выполнять их по этому порту
}

func peopleHandler(w http.ResponseWriter, r *http.Request) { // ResponseWriter - будущий ответ на запрос, принимает любые параметры, мы можем работать с самим запросом, изменять его
	switch r.Method { // определяем методы запроса
	case http.MethodGet: // запрос на получение
		getPeople(w) // вызываем функцию и прокидываем переменную для записи ответа на запрос
	case http.MethodPost: // запрос на установления данных
		postPerson(w, r) //вызываем функцию и прокидываем переменные для записи ответа на запрос и ещё чего-то, не понял чего
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed) // если иной запрос, то ошибка, не знаю как можно получить иной запрос
	}
}

func getPeople(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(people) // NewEncoder - возвращает новый кодировщик, который записывает данные в w, то есть ответ на запрос, Encode - кодирует в json, преобразуем слайс структур people в json объект для отправки на сервер и получения ответа
	fmt.Fprintf(w, "get people: '%v'", people)
	w.WriteHeader(http.StatusOK)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person // создаём экземпляр структуры

	err := json.NewDecoder(r.Body).Decode(&person) // NewDecoder(r.Body) - декодирует тело запроса, Decode(&person) - декодируем  указатель на структуру ??
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // обработка ошибок
		return
	}

	people = append(people, person)
	fmt.Fprintf(w, "post new person: '%v'", person)
	w.WriteHeader(http.StatusOK)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http web-server works correctly")
}
