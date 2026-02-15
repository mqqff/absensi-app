package main

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/mqqff/absensi-app/domain/entity"
	bcrypt "github.com/mqqff/absensi-app/pkg/bcrypt"
	uuid "github.com/mqqff/absensi-app/pkg/uuid"
	"github.com/xuri/excelize/v2"

	userRepo "github.com/mqqff/absensi-app/internal/app/user/repository"

	"github.com/mqqff/absensi-app/internal/infra/database"
	"github.com/mqqff/absensi-app/internal/infra/env"
	"github.com/mqqff/absensi-app/pkg/log"
)

const SeedersFilePath = "data/seeders/"
const SeedersDevPath = SeedersFilePath + "dev/"
const SeedersProdPath = SeedersFilePath + "prod/"

func main() {
	psqlDB := database.NewPgsqlConn()
	defer psqlDB.Close()

	var path string
	if env.AppEnv.AppEnv == "production" {
		path = SeedersProdPath
	} else {
		path = SeedersDevPath
	}

	seedUsers(path, psqlDB)
}

func seedUsers(path string, db *sqlx.DB) {
	bcrypt := bcrypt.Bcrypt
	uuid := uuid.UUID

	path += "users.xlsx"

	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Error(log.LogInfo{
			"error": err,
		}, "[seed][seedUsers] Error reading file")
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Error(log.LogInfo{
			"error": err,
		}, "[seed][seedUsers] Error getting rows")
	}

	userRepo := userRepo.NewUserRepository(db)

	for i, row := range rows {
		if i == 0 { // skip header
			continue
		}
		log.Info(log.LogInfo{
			"row": row,
		}, "[seed][seedUsers] Inserting user")

		hashedPassword, _ := bcrypt.Hash(strings.TrimSpace(row[2]))

		id, _ := uuid.NewV7()

		user := entity.User{
			ID:       id,
			Name:     strings.TrimSpace(row[0]),
			Email:    strings.TrimSpace(row[1]),
			Password: hashedPassword,
		}

		err := userRepo.CreateUser(context.Background(), user)

		if err != nil {
			log.Error(log.LogInfo{
				"error": err.Error(),
			}, "[seed][seedUsers] Error inserting user")
			return
		}
	}
}
