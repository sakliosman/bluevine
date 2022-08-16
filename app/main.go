package main

import (
	"fmt"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var logger *zap.Logger

func init() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world from osman\n")
	logger.Info("request applied",
		zap.String("requestHost", req.Host),
		zap.String("method", req.Method),
		zap.String("remoteAddr", req.RemoteAddr),
	)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
