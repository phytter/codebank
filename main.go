package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	_"github.com/lib/pg"
// )

func main() {
	fmt.Println("Teste")
	// db := stupDb()
	// defer db.Close()
	
	// cc := domain.NewCreditCard()
	// cc.Number = "1234"
	// cc.Name = "Alexandre"
	// cc.ExpirationYear = 2021
	// cc.ExpirationMonth = 7
	// cc.CVV = 123
	// cc.Limit = 1000
	// cc.Balance = 0

	// repo := repository.NewTransactionRepositoryDb(db)
	// repo.CreateCreditCard(*cc)
}

// func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
// 	transactionRepository := repository.NewTransactionRepositoryDb(db)
// 	useCase := usecase.NewUseCaseTransaction(transactionRepository)
// 	return useCase
// }

// func stupDb() *sql.DB {
// 	psqlInfo := fmt.Sprintf(format: "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
// 		a...:"db",
// 		"5432",
// 		"postgres",
// 		"root",
// 		"codebank",
// 	)
// 	db, err := sql.Open(driverName: "postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal(v...: "Error connect DB")
// 	}
// 	return db
// }