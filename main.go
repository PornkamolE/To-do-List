package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"todolist/routes"
)

func LoadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return errors.New("no .env file found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			os.Setenv(key, value) // ตั้งค่าตัวแปรแวดล้อม
		}
	}
	return scanner.Err()
}

func main() {
	// โหลดค่าจาก .env
	if err := LoadEnv(); err != nil {
		log.Fatal(err)
	}

	// อ่านค่าพอร์ต ถ้าไม่มีให้ใช้ค่า default 8080
	port, exist := os.LookupEnv("PORT")
	if !exist {
		port = "8080"
		fmt.Println("PORT not set in .env, using default: 8080")
	}

	// เริ่ม HTTP Server
	err := http.ListenAndServe(":"+port, routes.Init())
	if err != nil {
		log.Fatal(err)
	}
}
