package main

import (
	"context"
	"log"
	"time"

	"github.com/jaswdr/faker"
	"gitlab.com/voxe-analytics/internal/config"
	"gitlab.com/voxe-analytics/internal/core/repository"
	"gitlab.com/voxe-analytics/internal/core/repository/psql/sqlc"
	"gitlab.com/voxe-analytics/internal/pkg/logger"
	"gitlab.com/voxe-analytics/pkg/logger/factory"
	"gopkg.in/guregu/null.v4/zero"
)

const (
	defNumbers      = 100
	defParentNumber = 10
)

var (
	fake = faker.New()
	db   repository.Store

	organizationsId []string
)

func insertLicense() {
	for i := 0; i < defNumbers; i++ {
		db.LicenseInsertOne(context.Background(), sqlc.LicenseInsertOneParams{
			LicenseType:      1,
			DocumentNumber:   zero.StringFrom(fake.RandomStringWithLength(10)),
			GrantedDate:      zero.TimeFrom(time.Now()),
			Lifetime:         zero.IntFrom(1),
			OrganizationName: fake.Company().Name(),
			StirNumber:       zero.StringFrom(fake.RandomStringWithLength(10)),
			ReestrNumber:     zero.StringFrom(fake.RandomStringWithLength(10)),
			WorkCategory:     zero.StringFrom(fake.RandomStringWithLength(10)),
			DocFile:          zero.StringFrom(fake.RandomStringWithLength(10)),
		})
	}
}

func insertOrganizations() {
	for i := 0; i < defParentNumber; i++ {
		id, err := db.OrganizationInsertOne(context.Background(), sqlc.OrganizationInsertOneParams{
			Name:        fake.Company().Name(),
			FullName:    fake.Company().Name(),
			PhoneNumber: fake.RandomStringWithLength(10),
			Location:    zero.StringFrom(fake.Address().Address()),
		})

		if err != nil {
			log.Fatal(err)
		}

		organizationsId = append(organizationsId, id)
	}
}

func insertPayments() {
	for _, id := range organizationsId {
		for i := 0; i < defNumbers; i++ {
			db.PaymentInsertOne(context.Background(), sqlc.PaymentInsertOneParams{
				OrganizationID: id,
				Amount:         fake.Int64Between(100000, 10000000),
				Requisites:     zero.StringFrom(fake.RandomStringWithLength(10)),
				Status:         fake.Int32Between(1, 3),
				Type:           fake.Int32Between(1, 3),
			})
		}
	}
}

func insertDeclrations() {
	for _, id := range organizationsId {
		for i := 0; i < defNumbers; i++ {
			db.DeclarationInsertOne(context.Background(), sqlc.DeclarationInsertOneParams{
				OrganizationID:           id,
				DangerRate:               fake.Int32Between(3, 10),
				ReasonsOfDanger:          zero.StringFrom(fake.RandomStringWithLength(10)),
				ConverageOfTheDangerArea: zero.StringFrom(fake.RandomStringWithLength(10)),
				Proof:                    zero.StringFrom(fake.RandomStringWithLength(10)).String,
				LocationInfo:             zero.StringFrom(fake.RandomStringWithLength(10)),
				ResidentsInfo:            zero.StringFrom(fake.RandomStringWithLength(10)),
				LifeInsurance:            zero.StringFrom(fake.RandomStringWithLength(10)),
				TechDocument:             zero.StringFrom(fake.RandomStringWithLength(10)),
			})
		}

	}
}

func main() {
	cfg := config.Load()

	customLog, err := factory.Build(&cfg.Logging)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLogger(customLog)

	db = repository.New(context.Background(), cfg)
	logger.Log.Info("Seeding license table...")
	insertLicense()
	logger.Log.Info("Seeding organization table...")
	insertOrganizations()
	logger.Log.Info("Seeding declaration table...")
	insertDeclrations()
	logger.Log.Info("Seeding payment table...")
	insertPayments()
}
