// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"asb-queue-emulator/swagger/gen/restapi/operations"
	"asb-queue-emulator/swagger/handlers"
	"asb-queue-emulator/swagger/utils"
)

//go:generate swagger -q mixin ../../swagger/base.yaml ../../swagger/api.yaml ../../swagger/model.yaml -o ../../swagger/azure-servicebus-spec.yaml
//go:generate swagger generate server --target ../../gen --name AzureServiceBus --spec ../../azure-servicebus-spec.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.AzureServiceBusAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

var AzureServiceBusAPIContext utils.HandlerContext

type RawConsumer struct{}

func (c *RawConsumer) Consume(r io.Reader, target interface{}) error {
	if v, ok := target.(*string); ok {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(r)
		*v = buf.String()
		return err
	}
	return fmt.Errorf("Unsupported type")
}

func configureAPI(api *operations.AzureServiceBusAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()
	rawConsumer := &RawConsumer{}

	api.TxtConsumer = rawConsumer
	api.JSONConsumer = rawConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.PeekMessageHandler = handlers.NewPeekMessageHandler(AzureServiceBusAPIContext)

	api.DestructiveReadHandler = handlers.NewDestructiveReadHandler(AzureServiceBusAPIContext)

	api.SendMessageHandler = handlers.NewSendMessageHandler(AzureServiceBusAPIContext)

	api.DeleteMessageHandler = handlers.NewDeleteMessageHandler(AzureServiceBusAPIContext)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
