package go_logging

import (
	"reflect"
	"testing"

	logrus "github.com/sirupsen/logrus"
)

func Test_convertStringToLogLevel(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name string
		args args
		want logrus.Level
	}{
		{"Debug", args{"debug"}, logrus.DebugLevel},
		{"Info", args{"info"}, logrus.InfoLevel},
		{"Warn", args{"warn"}, logrus.WarnLevel},
		{"Default", args{""}, logrus.InfoLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStringToLogLevel(tt.args.logLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertStringToLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangeLogLevel(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Info", args{"Info"}},
		{"Debug", args{"Debug"}},
		{"Warn", args{"Warn"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChangeLogLevel(tt.args.logLevel)
		})
	}
}

func TestInitLog(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Info", args{"Info"}},
		{"Debug", args{"Debug"}},
		{"Warn", args{"Warn"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLog(tt.args.logLevel)
		})
	}
}
