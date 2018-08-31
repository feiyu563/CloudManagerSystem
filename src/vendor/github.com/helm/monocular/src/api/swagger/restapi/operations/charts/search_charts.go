package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SearchChartsHandlerFunc turns a function with the right signature into a search charts handler
type SearchChartsHandlerFunc func(SearchChartsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchChartsHandlerFunc) Handle(params SearchChartsParams) middleware.Responder {
	return fn(params)
}

// SearchChartsHandler interface for that can handle valid search charts params
type SearchChartsHandler interface {
	Handle(SearchChartsParams) middleware.Responder
}

// NewSearchCharts creates a new http.Handler for the search charts operation
func NewSearchCharts(ctx *middleware.Context, handler SearchChartsHandler) *SearchCharts {
	return &SearchCharts{Context: ctx, Handler: handler}
}

/*SearchCharts swagger:route GET /v1/charts/search charts searchCharts

search charts

*/
type SearchCharts struct {
	Context *middleware.Context
	Handler SearchChartsHandler
}

func (o *SearchCharts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewSearchChartsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
