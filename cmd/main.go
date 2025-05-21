package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	// Server settings
	listenAddr = ":8080"

	// Header configuration
	userHeader   = "X-Forwarded-User"
	authHeader   = "Authorization"
	bearerPrefix = "Bearer "
)

func main() {

	klog.InitFlags(nil)

	cmd := &cobra.Command{
		Use:   "showmethatoken",
		Short: "Simple HTTP server that logs and displays configurable auth headers",
		Run:   runServer,
	}

	// Include standard Go flag set to support flags like -v for klog.
	cmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// Define CLI flags for server configuration.
	cmd.Flags().StringVar(&listenAddr, "listen", listenAddr, "HTTP server listen address")
	cmd.Flags().StringVar(&userHeader, "user-header", userHeader, "Header name containing the user identity")
	cmd.Flags().StringVar(&authHeader, "auth-header", authHeader, "Header name containing the bearer token")
	cmd.Flags().StringVar(&bearerPrefix, "bearer-prefix", bearerPrefix, "Prefix to trim from the bearer token")

	// Execute the root Cobra command.
	if err := cmd.Execute(); err != nil {
		klog.Fatalf("Error: %v", err)
	}
}

// runServer configures and launches the HTTP server.
func runServer(cmd *cobra.Command, args []string) {
	http.HandleFunc("/", handleRequest)

	klog.Infof("Listening on %s", listenAddr)
	klog.Infof("Using headers: user=%q, auth=%q, prefix=%q", userHeader, authHeader, bearerPrefix)

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		klog.Fatalf("Server failed: %v", err)
	}
}

// handleRequest handles incoming HTTP requests and logs the user and token.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get(userHeader)
	rawToken := r.Header.Get(authHeader)
	trimmedToken := trimBearerPrefix(rawToken)

	klog.Infof("Request from user=%q, token=%q", user, redactToken(trimmedToken))
	fmt.Fprintf(w, "Bearer: %s\n\nLogin with this token:\n\nkubectl config set-credentials %s --token=%s", trimmedToken, user, trimmedToken)
}

// redactToken masks the middle 80% of the token string to avoid exposing secrets.
// Keeps the first and last 10% of the token visible. If too short, replaces entirely.
func redactToken(token string) string {
	n := len(token)
	if n < 10 {
		return strings.Repeat("*", n)
	}

	edge := max(1, n/10)
	return token[:edge] + strings.Repeat("*", n-2*edge) + token[n-edge:]
}

// trimBearerPrefix removes the configured prefix from the token if it exists.
func trimBearerPrefix(token string) string {
	if strings.HasPrefix(token, bearerPrefix) {
		return strings.TrimPrefix(token, bearerPrefix)
	}
	return token
}
