package main

import (
	"database/sql"
	"fmt"
	"log"
	"numberniceic/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	// 🚀 [1. แก้ไข] เราจะ Import 'lib/pq'
	// และลบ Import ของ 'pgx' ทั้งหมด
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	engine := html.New("./views", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupRoutes(app, db)

	log.Println("Starting server on port :8000...")
	err = app.Listen(":8000")
	if err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}

func InitDB() (*sql.DB, error) {
	// 1. โหลด .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// 2. อ่านค่า (เหมือนเดิม)
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 3. 🚀 [2. แก้ไข] สร้าง DSN สำหรับ 'lib/pq'
	//    (Driver นี้เข้าใจ client_encoding ในรูปแบบ Key-Value ได้ดี)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=utf8",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName,
	)

	// 4. 🚀 [3. แก้ไข] เปลี่ยนชื่อ Driver เป็น "postgres"
	//    (นี่คือชื่อ Driver ที่ 'lib/pq' ลงทะเบียนไว้)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// 5. Ping (เหมือนเดิม)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully!")
	return db, nil
}
