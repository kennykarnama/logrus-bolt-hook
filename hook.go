package boltlogrus

import (
	"os"

	"github.com/boltdb/bolt"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//HookOptions is struct to
//give the ability for user
//in order to configure bolt hook
type HookOptions struct {
	Dbpath      string
	FileMode    os.FileMode
	BoltOptions *bolt.Options
}

//HookOption will handle the options given
type HookOption func(*HookOptions)

//Dbpath is a function to set the option
func Dbpath(dbpath string) HookOption {
	return func(args *HookOptions) {
		args.Dbpath = dbpath
	}
}

//Filemode sets the filemode used in bolt
func Filemode(filemode os.FileMode) HookOption {
	return func(args *HookOptions) {
		args.FileMode = filemode
	}
}

//BoltOptions set the options used by bolt db
func BoltOptions(opts *bolt.Options) HookOption {
	return func(args *HookOptions) {
		args.BoltOptions = opts
	}
}

type boltHook struct {
	db *bolt.DB
}

//NewBoltHook will return new hook for logrus
func NewBoltHook(options ...HookOption) log.Hook {

	defaultOptions := &HookOptions{
		Dbpath:      "log.db",
		FileMode:    0600,
		BoltOptions: nil,
	}

	for _, option := range options {
		option(defaultOptions)
	}

	boltDb, err := bolt.Open(defaultOptions.Dbpath, defaultOptions.FileMode, defaultOptions.BoltOptions)

	if err != nil {
		log.Fatal(err)
	}

	return &boltHook{boltDb}
}

func (bh *boltHook) Fire(entry *log.Entry) error {
	err := bh.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("logs"))
		if err != nil {
			return err
		}
		uid := uuid.NewV4()
		str, err := entry.String()
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(uid.String()), []byte(str))
		return err
	})
	err = bh.Flush()
	return err
}

func (bh *boltHook) Levels() []log.Level {
	return log.AllLevels
}

func (bh *boltHook) Flush() error {
	return bh.db.Close()
}
