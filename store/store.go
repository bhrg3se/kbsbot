package store

import (
	"crypto/rsa"
	"database/sql"
	"kbsbot/models"
)

type RealStore struct {
	db     *sql.DB
	config models.Config
	jwtKey struct {
		public  *rsa.PublicKey
		private *rsa.PrivateKey
	}
}

// GetConfig returns config
func (rs RealStore) GetConfig() models.Config {
	return rs.config
}

// GetJWTPrivateKey gets the private key used for generating JWT tokens
func (rs RealStore) GetJWTPrivateKey() *rsa.PrivateKey {
	return rs.jwtKey.private
}

// GetJWTPublicKey gets the private key used to verify JWT tokens
func (rs RealStore) GetJWTPublicKey() *rsa.PublicKey {
	return rs.jwtKey.public
}
