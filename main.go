package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
)

var (
	logLevel int
	buf      bytes.Buffer // use a buffer to prevent the tests from printing
)

func init() {
	logLevel = 1
}

func main() {
	fmt.Println("vLogger happy path")
	buf.Reset()
	vLogger(0, "key", "value", "logger", "vLogger")
	fmt.Println(buf.String())

	fmt.Println("vLogger unbalanced key-pairs (last key doesn't get logged)")
	buf.Reset()
	vLogger(0, "key", "value", "logger", "vLogger", "unablanced")
	fmt.Println(buf.String())

	fmt.Println("vLogger int as a key")
	buf.Reset()
	vLogger(0, 1, 2, 3)
	fmt.Println(buf.String())

	fmt.Println("vLogger no key value pairs")
	buf.Reset()
	vLogger(0)
	fmt.Println(buf.String())

	fmt.Println("vLogger Chans")
	buf.Reset()
	vLogger(0, make(chan int), make(chan int))
	fmt.Println(buf.String())

	fmt.Println("vLogger invalid JSON")
	buf.Reset()
	vLogger(0, math.Inf(1), math.Inf(1))
	fmt.Println(buf.String())

	fmt.Println("mLogger happy path")
	buf.Reset()
	m := map[string]interface{}{"key": "value", "logger": "mLogger"}
	mLogger(0, m)
	fmt.Println(buf.String())

	fmt.Println("mLogger empty key and value")
	buf.Reset()
	m = map[string]interface{}{"": "", "logger": "mLogger"}
	mLogger(0, m)
	fmt.Println(buf.String())

	fmt.Println("mLogger chan as value")
	buf.Reset()
	m = map[string]interface{}{"key": make(chan int), "logger": "mLogger"}
	mLogger(0, m)
	fmt.Println(buf.String())

	fmt.Println("logger quoted key value pairs")
	buf.Reset()
	logger(0, `"key": "value", "logger": "logger"`)
	fmt.Println(buf.String())

	fmt.Println("logger plain key value pairs")
	buf.Reset()
	logger(0, `key: value, logger: logger`)
	fmt.Println(buf.String())

	fmt.Println("logger wrapped json")
	buf.Reset()
	logger(0, `{key: value, logger: logger}`)
	fmt.Println(buf.String())

	fmt.Println("logger quoted key value pairs json")
	buf.Reset()
	logger(0, `{ "key": "value", "logger": "logger" }`)
	fmt.Println(buf.String())

	fmt.Println("logMsg happy path")
	buf.Reset()
	logMsg(0, "key", "value")
	fmt.Println(buf.String())

	fmt.Println("logMetric happy path")
	buf.Reset()
	logMetric(0, "key", "value", "unit")
	fmt.Println(buf.String())
}

/*Variadic Logger
pros: easy syntax, clean code
cons: lots of logic and error checking needed in the function
*/
func vLogger(level int, values ...interface{}) {
	if level < logLevel {
		i := 1
		m := make(map[string]interface{})
		for i < (len(values)) {
			keyString := fmt.Sprintf("%v", values[i-1])
			m[keyString] = values[i]
			i = i + 2
		}

		m["logLevel"] = level
		// if the log message map fails marshell into JSON print the error and the raw map
		jsonLog, err := json.Marshal(m)
		if err != nil {
			fmt.Fprintf(&buf, "Err:%v, %s\n", err, m)
		} else {
			fmt.Fprintf(&buf, "%s\n", jsonLog)
		}

	}
}

/* Map Logger
pros: safer, less code
cons: need to create a new map everytime we want to log something, messy code
*/
func mLogger(level int, m map[string]interface{}) {
	if level < logLevel {
		m["logLevel"] = level
		// if the log message map fails marshell into JSON print the error and the raw map
		jsonLog, err := json.Marshal(m)
		if err != nil {
			fmt.Fprintf(&buf, "Err:%v, %s\n", err, m)
		} else {
			fmt.Fprintf(&buf, "%s\n", jsonLog)
		}
	}

}

/* Logger
pros: least code, safe, fastest
cons: manual formatting of kv pairs, messsy
*/
func logger(level int, msg string) {
	if level < logLevel {
		// if the log message map fails marshell into JSON print the error and the raw map
		jsonLog, err := json.Marshal(msg)
		if err != nil {
			fmt.Fprintf(&buf, "Err:%v, %s\n", err, msg)
		} else {
			fmt.Fprintf(&buf, "{ \"level\": %d, %s}\n", level, jsonLog)
		}
	}
}

/* are the above over complicating things? can we just log metrics and messages in a defined format, as per
https://hackernoon.com/tips-and-tricks-for-logging-and-monitoring-aws-lambda-functions-885af6da29a5
*/
func logMetric(level int, key string, value string, unit string) {
	if level < logLevel {
		fmt.Fprintf(&buf, "METRIC|%d|%s|%s|%s\n", level, key, value, unit)
	}
}

func logMsg(level int, key string, value string) {
	if level < logLevel {
		fmt.Fprintf(&buf, "MESSAGE|%d|%s|%s\n", level, key, value)
	}
}
