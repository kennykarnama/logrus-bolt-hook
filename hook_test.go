package boltlogrus_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"

	boltlogrus "github.com/kennykarnama/logrus-bolt-hook"
	log "github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	t.Run("LogSuccess", func(t *testing.T) {
		mockGenerator := boltlogrus.NewMockUniqueID()
		boltHook := boltlogrus.NewBoltHook(boltlogrus.Dbpath("customlog.db"), boltlogrus.IDGenerator(mockGenerator))
		log.AddHook(boltHook)
		var b bytes.Buffer
		log.SetOutput(&b)
		log.WithFields(log.Fields{
			"animal": "Dog",
			"number": "1",
		}).Info("zap")
		db, err := bolt.Open("customlog.db", 0600, nil)
		defer db.Close()
		assert.Nil(t, err, "Opendb Error must be nil")
		err = db.View(func(tx *bolt.Tx) error {
			buck := tx.Bucket([]byte("logs"))
			v := buck.Get([]byte("UNIK"))
			s := string(v)
			assert.Equal(t, b.String(), s, "Log saved in db & in output buffer should be same")
			return nil
		})
	})
	//test cleanup
	clearDBFiles()
}

func clearDBFiles() {
	files, err := filepath.Glob("*.db")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}
