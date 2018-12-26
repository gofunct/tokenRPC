package cmd

import (
	"encoding/json"
	kitlog "github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net"
	"net/http"
	"net/textproto"
	"os"
	"os/signal"
	"strings"
	"time"
)

func logFormatter() handlers.LogFormatter {
	// Setup logrus
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	level, err := logrus.ParseLevel(ethConfig.LogLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}

	return func(writer io.Writer, params handlers.LogFormatterParams) {

		host, _, err := net.SplitHostPort(params.Request.RemoteAddr)
		if err != nil {
			host = params.Request.RemoteAddr
		}

		uri := params.Request.RequestURI

		// Requests using the CONNECT method over HTTP/2.0 must use
		// the authority field (aka r.Host) to identify the target.
		// Refer: https://httpwg.github.io/specs/rfc7540.html#CONNECT
		if params.Request.ProtoMajor == 2 && params.Request.Method == "CONNECT" {
			uri = params.Request.Host
		}
		if uri == "" {
			uri = params.URL.RequestURI()
		}

		duration := int64(time.Now().Sub(params.TimeStamp) / time.Millisecond)

		fields := logrus.Fields{
			"host":       host,
			"url":        uri,
			"duration":   duration,
			"status":     params.StatusCode,
			"method":     params.Request.Method,
			"request":    params.Request.RequestURI,
			"remote":     params.Request.RemoteAddr,
			"size":       params.Size,
			"referer":    params.Request.Referer(),
			"user_agent": params.Request.UserAgent(),
			"request_id": params.Request.Header.Get("x-request-id"),
		}

		// Only append headers if explicitly enabled
		if gatewayConfig.LogHeaders {
			if headers, err := json.Marshal(params.Request.Header); err == nil {
				fields["headers"] = string(headers)
			} else {
				fields["header_error"] = err.Error()
			}
		}

		logrus.WithFields(fields).WithTime(params.TimeStamp).Infof("%s %s %d", params.Request.Method, uri, params.StatusCode)
	}
}

func allowCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		corsAllowOrigin := gatewayConfig.CorsAllowOrigin
		if corsAllowOrigin == "*" {
			if origin := req.Header.Get("Origin"); origin != "" {
				corsAllowOrigin = origin
			}
		}
		w.Header().Set("Access-Control-Allow-Origin", corsAllowOrigin)
		w.Header().Set("Access-Control-Allow-Credentials", gatewayConfig.CorsAllowCredentials)
		w.Header().Set("Access-Control-Allow-Methods", gatewayConfig.CorsAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", gatewayConfig.CorsAllowHeaders)
		if req.Method == "OPTIONS" && req.Header.Get("Access-Control-Request-Method") != "" {
			return
		}
		handler.ServeHTTP(w, req)
	})
}

// sanitizeApiPrefix forces prefix to be non-empty and end with a slash.
func sanitizeApiPrefix(prefix string) string {
	if len(prefix) == 0 || prefix[len(prefix)-1:] != "/" {
		return prefix + "/"
	}
	return prefix
}

// isPermanentHTTPHeader checks whether hdr belongs to the list of
// permenant request headers maintained by IANA.
// http://www.iana.org/assignments/message-headers/message-headers.xml
// From https://github.com/grpc-ecosystem/grpc-gateway/blob/7a2a43655ccd9a488d423ea41a3fc723af103eda/runtime/context.go#L157
func isPermanentHTTPHeader(hdr string) bool {
	switch hdr {
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":
		return true
	}
	return false
}

// isReserved returns whether the key is reserved by gRPC.
func isReserved(key string) bool {
	return strings.HasPrefix(key, "Grpc-")
}

// incomingHeaderMatcher converts an HTTP header name on http.Request to
// grpc metadata. Permanent headers (i.e. User-Agent) are prepended with
// "grpc-gateway". Headers that start with start with "Grpc-" (reserved
// by grpc) are prepended with "X-". Other headers are forwarded as is.
func incomingHeaderMatcher(key string) (string, bool) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	if isPermanentHTTPHeader(key) {
		return runtime.MetadataPrefix + key, true
	}
	if isReserved(key) {
		return "X-" + key, true
	}

	// The Istio service mesh dislikes when you pass the Content-Length header
	if key == "Content-Length" {
		return "", false
	}

	return key, true
}

// outgoingHeaderMatcher transforms outgoing metadata into HTTP headers.
// We return any response metadata as is.
func outgoingHeaderMatcher(metadata string) (string, bool) {
	return metadata, true
}

// SignalRunner runs a runner function until an interrupt signal is received, at which point it
// will call stopper.
func SignalRunner(runner, stopper func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	go func() {
		runner()
	}()

	logrus.Info("hit Ctrl-C to shutdown")
	select {
	case <-signals:
		stopper()
	}
}

func GetDefaultConfig() *viper.Viper {
	return defaultViper
}

func GetKitLogger() kitlog.Logger {
	return kitLog
}
