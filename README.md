# logrus-bolt-hook
Use this to log into [Bolt](https://github.com/boltdb/bolt).
This hook is created because i need to practice about go programming language. So it is still in initial development version. It only logs into bolt, with key initiated by default using [uuid](https://github.com/satori/go.uuid) and the value is the entry itself.

# Usage
The hook could be configured using options provided:
```
type HookOptions struct {
	IDGenerator UniqueID
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
To see the result of the logging, you could use [bolter](https://github.com/hasit/bolter)

# Custom IDGenerator
If you want to implement your own IDGenerator, you just need to implement the interface specified in `uniqueid.go`
Register your new IDGenerator by using `IDGenerator(youOwn)` inside the function of NewBoltHook. You could take a look in [hook_test.go](https://github.com/kennykarnama/logrus-bolt-hook/blob/master/hook_test.go)

# Silent the log
Sometimes you may need to not display the log message to your default terminal console. This may happen when you use multiple hook. Same thing will be logged twice inside your terminal console. hmm, you don't like, do you. So to prevent this, you could try this :

```
silentLog := logrus.New()
var b bytes.Buffer
silentLog.SetOutput(&b)
//Same goes here, add your newhook
```

# Resources
This hook uses two dependencies of libary :
1. [logrus](https://github.com/sirupsen/logrus)
2. [bolt](https://github.com/boltdb/bolt)
3. [uuid](https://github.com/satori/go.uuid)
4. [mockery](https://github.com/vektra/mockery) for generating mock of `uniqueid` interface
5. [testify](https://github.com/stretchr/testify) for assertion

You could read more at those sources

# Contributions
Any contributions including Creating Issues & PR are appreciated. 
