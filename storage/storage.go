package storage

import "github.com/boltdb/bolt"

// Storage is the struct to handle all storage operations on
type Storage struct {
	db *bolt.DB
}

// New returns a new Storage for a specific BoltDB instance
func New(db *bolt.DB) Storage {
	return Storage{
		db: db,
	}
}

// NewRun Adds a new run to the database
func (s *Storage) NewRun() Run {
	run := Run{
		ID: s.getNextIDIn("RUN"),
	}
	s.SaveRun(run)

	return run
}

// SaveRun saves a run to the database
func (s *Storage) SaveRun(run Run) error {
	runData, err := run.Encode()
	if err != nil {
		return err
	}
	s.saveData("RUN", run.ID, runData)
	return nil
}

// GetRun gets a run for a specific ID from the database
func (s *Storage) GetRun(id uint64) (Run, error) {
	return DecodeRun(s.getData("RUN", id))
}

// GetAllRuns gets all runs done by the system
func (s *Storage) GetAllRuns() ([]Run, error) {
	bytes := s.getAll("RUN")
	out := []Run{}

	for _, b := range bytes {
		run, _ := DecodeRun(b)
		out = append(out, run)
	}
	return out, nil
}
