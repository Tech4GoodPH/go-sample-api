// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/Tech4GoodPH/go-sample-api.git/restapi/operations"
	"github.com/Tech4GoodPH/go-sample-api.git/restapi/operations/query_posts"
	"github.com/Tech4GoodPH/go-sample-api.git/restapi/operations/query_rates"
	"github.com/Tech4GoodPH/go-sample-api.git/restapi/operations/report"
)

//go:generate swagger generate server --target ..\..\go-sample-api.git --name GoSample --spec ..\swagger.yml

func configureFlags(api *operations.GoSampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GoSampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ReportDeleteReportHandler == nil {
		api.ReportDeleteReportHandler = report.DeleteReportHandlerFunc(func(params report.DeleteReportParams) middleware.Responder {
			return middleware.NotImplemented("operation report.DeleteReport has not yet been implemented")
		})
	}
	if api.QueryRatesGetAvgratesHandler == nil {
		api.QueryRatesGetAvgratesHandler = query_rates.GetAvgratesHandlerFunc(func(params query_rates.GetAvgratesParams) middleware.Responder {
			return middleware.NotImplemented("operation query_rates.GetAvgrates has not yet been implemented")
		})
	}
	if api.QueryPostsGetPostsHandler == nil {
		api.QueryPostsGetPostsHandler = query_posts.GetPostsHandlerFunc(func(params query_posts.GetPostsParams) middleware.Responder {
			return middleware.NotImplemented("operation query_posts.GetPosts has not yet been implemented")
		})
	}
	if api.QueryPostsGetPostsMessiestHandler == nil {
		api.QueryPostsGetPostsMessiestHandler = query_posts.GetPostsMessiestHandlerFunc(func(params query_posts.GetPostsMessiestParams) middleware.Responder {
			return middleware.NotImplemented("operation query_posts.GetPostsMessiest has not yet been implemented")
		})
	}
	if api.QueryPostsGetPostsMostactiveHandler == nil {
		api.QueryPostsGetPostsMostactiveHandler = query_posts.GetPostsMostactiveHandlerFunc(func(params query_posts.GetPostsMostactiveParams) middleware.Responder {
			return middleware.NotImplemented("operation query_posts.GetPostsMostactive has not yet been implemented")
		})
	}
	if api.QueryPostsGetPostsNewestHandler == nil {
		api.QueryPostsGetPostsNewestHandler = query_posts.GetPostsNewestHandlerFunc(func(params query_posts.GetPostsNewestParams) middleware.Responder {
			return middleware.NotImplemented("operation query_posts.GetPostsNewest has not yet been implemented")
		})
	}
	if api.QueryPostsGetPostsTidiestHandler == nil {
		api.QueryPostsGetPostsTidiestHandler = query_posts.GetPostsTidiestHandlerFunc(func(params query_posts.GetPostsTidiestParams) middleware.Responder {
			return middleware.NotImplemented("operation query_posts.GetPostsTidiest has not yet been implemented")
		})
	}
	if api.QueryRatesGetRatesIDHandler == nil {
		api.QueryRatesGetRatesIDHandler = query_rates.GetRatesIDHandlerFunc(func(params query_rates.GetRatesIDParams) middleware.Responder {
			return middleware.NotImplemented("operation query_rates.GetRatesID has not yet been implemented")
		})
	}
	if api.ReportGetReportHandler == nil {
		api.ReportGetReportHandler = report.GetReportHandlerFunc(func(params report.GetReportParams) middleware.Responder {
			return middleware.NotImplemented("operation report.GetReport has not yet been implemented")
		})
	}
	if api.ReportPatchReportHandler == nil {
		api.ReportPatchReportHandler = report.PatchReportHandlerFunc(func(params report.PatchReportParams) middleware.Responder {
			return middleware.NotImplemented("operation report.PatchReport has not yet been implemented")
		})
	}
	if api.ReportPostReportHandler == nil {
		api.ReportPostReportHandler = report.PostReportHandlerFunc(func(params report.PostReportParams) middleware.Responder {
			return middleware.NotImplemented("operation report.PostReport has not yet been implemented")
		})
	}

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
