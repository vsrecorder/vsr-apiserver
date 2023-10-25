package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"github.com/vsrecorder/vsr-apiserver/pkg/controllers"
	"github.com/vsrecorder/vsr-apiserver/pkg/infrastructures"
	"github.com/vsrecorder/vsr-apiserver/pkg/repositories"
	"github.com/vsrecorder/vsr-apiserver/pkg/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	userName := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	firebaseProjectId := os.Getenv("FIREBASE_PROJECT_ID")
	firebaseCredentialsFilePath := os.Getenv("FIREBASE_CREDENTIAL_FILE_PATH")

	r := gin.Default()
	m := ginmetrics.GetMonitor()

	m.SetMetricPath("/api/v1alpha/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.25, 0.5, 0.75, 1.0, 2.5, 5.0, 7.5, 10.0})
	m.Use(r)

	{
		opt := option.WithCredentialsFile(firebaseCredentialsFilePath)
		config := &firebase.Config{ProjectID: firebaseProjectId}

		app, err := firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}

		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}

		db, err := infrastructures.NewMySQL(userName, password, dbHostname, dbPort, dbName)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		controllers.NewUserController(
			r,
			services.NewUserService(
				repositories.NewUserRepository(auth),
				repositories.NewRecordRepository(db),
				repositories.NewGameRepository(db),
				repositories.NewDeckRepository(db),
			),
		).RegisterRoutes("/api/v1alpha")
	}

	{
		db, err := infrastructures.NewMySQL(userName, password, dbHostname, dbPort, dbName)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		controllers.NewOfficialEventController(
			r,
			services.NewOfficialEventService(
				repositories.NewOfficialEventRepository(db),
				repositories.NewRecordRepository(db),
			),
		).RegisterRoutes("/api/v1alpha")
	}

	{
		db, err := infrastructures.NewMySQL(userName, password, dbHostname, dbPort, dbName)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		controllers.NewRecordController(
			r,
			services.NewRecordService(
				repositories.NewRecordRepository(db),
				repositories.NewGameRepository(db),
				repositories.NewOfficialEventRepository(db),
			),
		).RegisterRoutes("/api/v1alpha")
	}

	{
		db, err := infrastructures.NewMySQL(userName, password, dbHostname, dbPort, dbName)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		controllers.NewGameController(
			r,
			services.NewGameService(
				repositories.NewGameRepository(db),
				repositories.NewRecordRepository(db),
			),
		).RegisterRoutes("/api/v1alpha")
	}

	{
		db, err := infrastructures.NewMySQL(userName, password, dbHostname, dbPort, dbName)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		controllers.NewDeckController(
			r,
			services.NewDeckService(
				repositories.NewDeckRepository(db),
				repositories.NewRecordRepository(db),
			),
		).RegisterRoutes("/api/v1alpha")
	}

	if err := r.Run(":8913"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
