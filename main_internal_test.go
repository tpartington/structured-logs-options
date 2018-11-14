package main

import (
	"math"
	"testing"
)

func TestVLogger(t *testing.T) {
	vLogger(0, "key", "value", "logger", "vLogger")
}

func TestVLoggerUnablanced(t *testing.T) {
	vLogger(0, "key", "value", "logger", "vLogger", "unbalanced")
}

func TestVLoggerNonStringKey(t *testing.T) {
	vLogger(0, 2, 3)
}

func TestVLoggerNoPairs(t *testing.T) {
	vLogger(0)
}

func TestVLoggerChans(t *testing.T) {
	vLogger(0, make(chan int), make(chan int))
}

func TestVLoggerInvalidJSON(t *testing.T) {
	vLogger(0, math.Inf(1), math.Inf(1))
}

func TestMLogger(t *testing.T) {
	m := map[string]interface{}{"key": "value", "logger": "mLogger"}
	mLogger(0, m)
}

func TestMLoggerEmptyKeyValue(t *testing.T) {
	m := map[string]interface{}{"": "", "logger": "mLogger"}
	mLogger(0, m)

}

func TestMLoggerChanValue(t *testing.T) {
	m := map[string]interface{}{"key": make(chan int), "logger": "mLogger"}
	mLogger(0, m)

}

func TestLogger(t *testing.T) {
	logger(0, "\"error\": \"some-type-of-error\", \"msg\": \"logger\"")
}

func BenchmarkVlogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vLogger(0, "key", "value", "logger", "vLogger")
	}
}

func BenchmarkMlogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := map[string]interface{}{"key": "value", "logger": "mLogger"}
		mLogger(0, m)
	}
}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger(0, "\"key\": \"value\", \"logger\": \"logger\"")
	}
}

func BenchmarkLogMetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logMetric(0, "key", "value", "unit")
	}
}

func BenchmarkVloggerNOOP(b *testing.B) {
	logLevel = 0
	for i := 0; i < b.N; i++ {
		vLogger(0, "key", "value", "logger", "vLogger")
	}
}

func BenchmarkMloggerNOOP(b *testing.B) {
	logLevel = 0
	for i := 0; i < b.N; i++ {
		m := map[string]interface{}{"key": "value", "logger": "mLogger"}
		mLogger(0, m)
	}
}

func BenchmarkLoggerNOOP(b *testing.B) {
	logLevel = 0
	for i := 0; i < b.N; i++ {
		logger(0, "\"key\": \"value\", \"logger\": \"logger\"")
	}
}

func BenchmarkLogMetricNOOP(b *testing.B) {
	logLevel = 0
	for i := 0; i < b.N; i++ {
		logMetric(0, "key", "value", "unit")
	}
}
