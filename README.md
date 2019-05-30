# logrus-bolt-hook
Use this to log into [Bolt](https://github.com/boltdb/bolt).
This hook is created because i need to practice about go programming language. So it is still in initial development version. It only logs into bolt, with key initiated using [uuid](https://github.com/satori/go.uuid)
and the value is the entry itself.

# Usage
The hook could be configured using options provided:
```
type HookOptions struct {
	Dbpath      string
	FileMode    os.FileMode
	BoltOptions *bolt.Options
}
```
Example:

```
boltHook := boltlogrus.NewBoltHook(boltlogrus.Dbpath("customlog.db"))
log.AddHook(boltHook)
	log.WithFields(log.Fields{
		"animal": "Dog",
		"number": "1",
	}).Info("Kenny")
```
# Resources
This hook uses two dependencies of libary :
1. [logrus](https://github.com/sirupsen/logrus)
2. [bolt](https://github.com/boltdb/bolt)
3. [uuid](https://github.com/satori/go.uuid)

You could read more at those sources

# Contributions
Any contributions including Creating Issues & PR are appreciated. 
