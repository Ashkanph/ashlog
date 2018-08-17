A personalized logger for Golang using the log package.

## Features
* Can accept an array of io.Writer to log in (see [the example part](#example))
* Can choose one of 6 levels(Or easily add a new one):
    * TRACE: 0
	* DEBUG: 1
	* INFO: 2
	* WARNING: 3
	* ERROR: 4
	* CRITICAL: 5
* Can choose one of 4 Categories (Or easily add a new one):
    * LOGAPP
	* LOGDBM
	* LOGNET
	* LOGCONF
* Compressed readable log format:
```go
    @[20180816-205042] E/net: Websocket connection dropped
    @[20180816-205042] D/net: received: salam
```
* Log easily by import the package:
```go
    ashlog.Info(ashlog.LOGAPP, fmt.Sprintf("App Version: %s ", "1.2.3"))
    ashlog.Error(ashlog.LOGNET, "An error happened!")
```


## Usage
It has an initializer function (InitLogger) which receive two arguments:
* writers: an array of io.Writer which the ashlog must logs into it. 
* lvl: It define the level of logs. logs bigger than this wont be logged.

### Example
Below comes an example of initializing an ashlog which logs in both file and os.Stdout:

```go
    var fileAddr = "./ashlog.log"
	var writers []io.Writer

    // Opening the fileAddr (or creating it)
    if _, err := os.Stat(fileAddr); os.IsNotExist(err) {
        app.logFile, err = os.Create(fileAddr)
        writers = append(writers, app.logFile)
    } else {
        app.logFile, err = os.OpenFile(fileAddr, os.O_APPEND|os.O_WRONLY, 0664)
        if err != nil {
            log.Println("Can not open log file!")
        }else{
            writers = append(writers, app.logFile)
        }
    }

	// append os.Stdout
	writers = append(writers, os.Stdout)

	//Initialize ashlog
	if err := ashlog.InitLogger(writers, 5); err != nil {
		log.Println(err)
	}
```

