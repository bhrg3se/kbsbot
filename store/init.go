package store

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"kbsbot/models"
	"os"
	"path/filepath"
	"time"
)

type Store interface {
	GetConfig() models.Config
	GetJWTPrivateKey() *rsa.PrivateKey
	GetJWTPublicKey() *rsa.PublicKey
}

var State Store

// NewRealStore creates a new store with all dependencies like database
func NewRealStore(config models.Config) RealStore {

	privateKey := initJWTKeys()

	//create database connection
	db := createDBPool(config)
	return RealStore{
		db:     db,
		config: config,
		jwtKey: struct {
			public  *rsa.PublicKey
			private *rsa.PrivateKey
		}{public: &privateKey.PublicKey, private: privateKey},
	}
}

// initJWTKeys opens private key file for jwt. If it does not exists, it creates one.
func initJWTKeys() *rsa.PrivateKey {

	f, err := os.Open("/etc/kbsbot/jwt.key")
	if err != nil {
		if os.IsNotExist(err) {
			key, errGen := rsa.GenerateKey(rand.Reader, 2048)
			if errGen != nil {
				logrus.Fatalf("could not generate private key file: %v", errGen)
			}
			keyBytes := x509.MarshalPKCS1PrivateKey(key)

			err = ioutil.WriteFile("/etc/kbsbot/jwt.key", keyBytes, os.ModeType)
			if err != nil {
				logrus.Fatalf("could not write private key file: %v", err)
			}
			return key
		}
		logrus.Fatalf("could not open private key file: %v", err)
	}
	defer f.Close()
	keyBytes, err := ioutil.ReadAll(f)
	if err != nil {
		logrus.Fatalf("could not read private key file: %v", err)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err != nil {
		logrus.Fatalf("could not parse private key file: %v", err)
	}

	return privateKey
}

// createDBPool creates the connection to postgres database
func createDBPool(config models.Config) *sql.DB {
	var str string

	if config.Database.SSL {

		caCert, _ := filepath.Abs(config.Database.CaCertPath)
		userCert, _ := filepath.Abs(config.Database.UserCertPath)
		userKey, _ := filepath.Abs(config.Database.UserKeyPath)

		str = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=verify-full&sslrootcert=%s&sslcert=%s&sslkey=%s",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
			caCert,
			userCert,
			userKey,
		)

	} else {
		str = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)
	}

	db, err := sql.Open("postgres", str)
	if err != nil {
		panic(err.Error())
	}

	//Check if the connection is successful by establishing a connection.
	//Retry upto 10 times if connection is not successful
	for retryCount := 0; retryCount < 10; retryCount++ {
		err = db.Ping()
		if err == nil {
			logrus.Info("database connection successful")
			return db
		}

		logrus.Error(err)
		logrus.Info("could not connect to database: retrying...")
		time.Sleep(time.Second)
	}

	panic("could not connect to database")

}
