package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/boltdb/bolt"
)

func TestRun(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer db.Close()
	defer os.Remove("my.db")

	storage := New(db)
	run := storage.NewRun()

	run.Repo = "test"
	err = storage.SaveRun(run)
	assert.NoError(t, err)

	run2, err := storage.GetRun(run.ID)
	assert.NoError(t, err)

	assert.Equal(t, run, run2)
}

func TestGetAllRuns(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer db.Close()
	defer os.Remove("my.db")

	storage := New(db)
	run := storage.NewRun()

	run.Repo = "test"
	err = storage.SaveRun(run)
	assert.NoError(t, err)

	runs, err := storage.GetAllRuns()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(runs))
}

func TestGetAllFirstToLast(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer db.Close()
	defer os.Remove("my.db")

	storage := New(db)
	run := storage.NewRun()

	run.Repo = "test"
	err = storage.SaveRun(run)
	assert.NoError(t, err)

	bytes := storage.getAll("RUN")
	runs := []Run{}

	for _, b := range bytes {
		runs, _ := DecodeRun(b)
		out = append(out, run)
	}

	assert.Equal(t, 1, len(runs))
}
