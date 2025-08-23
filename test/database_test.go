package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"idm/inner/common"
	"idm/inner/database"
)

func TestGetConfig_NoEnvFile(t *testing.T) {
	a := assert.New(t)
	cfg := common.GetConfig(".env_not_exists")

	a.Empty(cfg)
}

func TestGetConfig_HasEnv(t *testing.T) {
	a := assert.New(t)
	cfg := common.GetConfig(".env_test")

	a.NotEmpty(cfg)
}

func TestGetConfig_EnvVars(t *testing.T) {
	a := assert.New(t)
	cfg := common.GetConfig(".env_test")

	a.NotEmpty(cfg.DbDriverName)
	a.NotEmpty(cfg.Dsn)

	a.Equal("postgres", cfg.DbDriverName)
	a.Equal("host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", cfg.Dsn)
}

func TestDatabase_ConnectDbFailed(t *testing.T) {
	a := assert.New(t)
	var db = database.ConnectDb(".fake_env_test")

	if db == nil {
		a.Empty(db)
	}

	var clearDatabase = func() {
		db.MustExec("delete from employee")
	}
	defer func() {
		if r := recover(); r != nil {
			clearDatabase()
		}
	}()

}

func TestDatabase_ConnectDbSuccess(t *testing.T) {
	var db = database.ConnectDb(".env_test")
	var clearDatabase = func() {
		db.MustExec("delete from employee")
	}
	defer func() {
		if r := recover(); r != nil {
			clearDatabase()
		}
	}()

}
