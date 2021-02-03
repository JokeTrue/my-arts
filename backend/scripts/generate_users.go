package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"syreclabs.com/go/faker"
)

const QueryCreateUser = `INSERT INTO users (email, password, first_name, last_name, age, gender, location, biography) 
						 VALUES `

var genders = []string{"M", "F", "O"}

func main() {
	dsn := os.Args[1]
	if dsn == "" {
		log.Fatal("empty dsn")
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("failed to setup db connection")
	}
	defer db.Close()

	var usersCount int
	if err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&usersCount); err != nil {
		log.Fatal("failed to fetch users count")
	}

	emails := make([]string, 0, 10000)
	firstNames := make([]string, 0, 10000)
	lastNames := make([]string, 0, 10000)
	adresses := make([]string, 0, 10000)
	biographies := make([]string, 0, 10000)

	for j := 0; j < 30_000; j++ {
		emails = append(emails, faker.Internet().Email())
		firstNames = append(firstNames, faker.Name().FirstName())
		lastNames = append(lastNames, faker.Name().LastName())
		adresses = append(adresses, faker.Address().City()+", "+faker.Address().Country())
		biographies = append(biographies, faker.Hacker().SaySomethingSmart())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	needUsers := 1_000_000 - usersCount
	values := make([]string, 0, usersCount)
	for i := 0; i < needUsers; i++ {
		value := fmt.Sprintf(
			`("%s", "%s", "%s", "%s", %d, "%s", "%s", "%s")`,
			strconv.Itoa(rand.Intn(1_000_000-1000)+1000)+emails[rand.Intn(len(emails))],
			string(hashedPassword),
			firstNames[rand.Intn(len(firstNames))],
			lastNames[rand.Intn(len(lastNames))],
			rand.Intn(80-18)+18,
			genders[rand.Intn(len(genders))],
			adresses[rand.Intn(len(adresses))],
			biographies[rand.Intn(len(biographies))],
		)
		values = append(values, value)
	}

	batches := GetBatches(values, 5_000)
	for i, batch := range batches {
		query := QueryCreateUser + strings.Join(batch, ",\n")
		_, err = db.Exec(query)
		if err != nil {
			log.Printf("Failed to process %d of %d...\n", i+1, len(batches))
		} else {
			log.Printf("Processed %d of %d...\n", i+1, len(batches))
		}
	}
}

func GetBatches(slice []string, chunkSize int) [][]string {
	var batches [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		batches = append(batches, slice[i:end])
	}
	return batches
}
