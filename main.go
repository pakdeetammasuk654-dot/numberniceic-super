package main

import (
	"database/sql"
	"fmt"
	"log"
	"numberniceic/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	app := fiber.New()

	routes.SetupRoutes(app, db)

	log.Println("Starting server on port :8000...")
	err = app.Listen(":8000")
	if err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}

}

func InitDB() (*sql.DB, error) {
	// 1. โหลดไฟล์ .env
	err := godotenv.Load() // โหลด .env จาก path ปัจจุบัน
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		// เราไม่จำเป็นต้อง Fatal ที่นี่ เพราะ ENV อาจถูกตั้งค่าไว้ที่ Server โดยตรง
	}

	// 2. อ่านค่าจาก Environment Variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 3. สร้าง Connection String (DSN)
	// format: "postgres://username:password@host:port/dbname?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	// 4. เปิดการเชื่อมต่อ
	// เราใช้ "pgx" เป็นชื่อไดรเวอร์
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// 5. ตรวจสอบว่าเชื่อมต่อสำเร็จจริง
	// sql.Open() ไม่ได้เชื่อมต่อทันที แต่จะรอจนกว่าจะใช้งานจริง
	// เราจึงใช้ Ping() เพื่อยืนยันว่าทุกอย่างถูกต้อง
	err = db.Ping()
	if err != nil {
		db.Close() // ปิดการเชื่อมต่อถ้า Ping ไม่ผ่าน
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully!")
	return db, nil
}
