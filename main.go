package main

import (
	"fmt"
	"lib/internal/SQL"
	"lib/internal/files"
	"log"
	"time"
)

func setupSettings(confFile string) error {
	if confFile == "" {
		confFile = files.JoinRunPath("settings.json")
	}

	var err error
	settings, err = LoadSettings(confFile)
	if err != nil {
		return fmt.Errorf("unable to load settings: %w", err)
	}

	return nil
}

func main() {

	for i := 1; i < 100; i++ {
		run()
	}
}

func run() {

	start := time.Now()

	fmt.Println(start)

	if err := setupSettings("settings.json"); err != nil {
		log.Fatalf("Error loading settings: %v\n", err)
	}
	//fmt.Println(settings.InputDir)
	//fmt.Println(settings.filename)
	//fmt.Println(settings.OutputDir)
	SQL.Exect()

}

//func mainsite() {
//	//создаем и запускаем в работу роутер для обслуживания запросов
//	r := httprouter.New()
//	routes(r)
//	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
//	//вторым параметром передается роутер, который будет работать с запросами
//	err := http.ListenAndServe("localhost:4444", r)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//func routes(r *httprouter.Router) {
//	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
//	r.ServeFiles("/public/*filepath", http.Dir("public"))
//	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
//	r.GET("/", controller.StartPage)
//	r.GET("/users", controller.GetUsers)
//}
