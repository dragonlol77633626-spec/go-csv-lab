package documents
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

// Задание 1: Чтение пользователей из users.csv
func task1() {
	fmt.Println("--- ЗАДАНИЕ 1: Чтение users.csv ---")
	const filename = "users.csv"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не удалось открыть файл %s", filename)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // Пропускаем заголовок

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Ошибка чтения csv файла")
	}

	for _, record := range records {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", record[0], record[1], record[2])
	}
	fmt.Println()
}

// Задание 2: Запись продуктов в products.csv
func task2() {
	fmt.Println("--- ЗАДАНИЕ 2: Запись продуктов ---")
	products := []Product{
		{ID: 1, Name: "Монитор", Price: 21990},
		{ID: 2, Name: "Видеокарта", Price: 69900},
		{ID: 3, Name: "Ноутбук", Price: 79500},
	}

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Не удалось создать csv файл")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "Price"}
	writer.Write(header)

	for _, product := range products {
		id := strconv.Itoa(product.ID)
		price := strconv.FormatFloat(product.Price, 'f', 2, 64)
		record := []string{id, product.Name, price}
		err := writer.Write(record)
		if err != nil {
			log.Fatalf("Ошибка записи строки: %v", err)
		}
	}
	fmt.Println("Файл products.csv успешно создан.")
	fmt.Println()
}

// Задание 3: Мост к JSON (Чтение из products.csv в структуры)
func task3() {
	fmt.Println("--- ЗАДАНИЕ 3: Парсинг в структуры ---")
	file, err := os.Open("products.csv")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Ошибка чтения данных из csv файла: %v", err)
	}

	loadedProducts := []Product{}
	records = records[1:] // Пропускаем заголовок

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("Ошибка чтения ID %s: %v", record[0], err)
		}
		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatalf("Ошибка чтения цены %s: %v", record[2], err)
		}
		product := Product{
			ID:    id,
			Name:  record[1],
			Price: price,
		}
		loadedProducts = append(loadedProducts, product)
	}

	for _, product := range loadedProducts {
		fmt.Printf("ID: %d, Name: %s, Price: %f\n", product.ID, product.Name, product.Price)
	}
	fmt.Println()
}

func main() {
	task1()
	task2()
	task3()
}
