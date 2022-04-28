package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"bitbucket.org/ajmal_azm/scraperP/model"
)

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.GetRequest
	vars := mux.Vars(r)
	req.PackageName = vars["package_name"]
	return req, nil
}

func makeGetPackageNameDetailsEndpoint(svc WebService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetRequest)
		d, err := svc.GetPackageNameDetails(req)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
}

func decodeGetLogRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.GetRequest
	vars := mux.Vars(r)
	req.PackageName = vars["package_name"]
	return req, nil
}

func makeGetLogEndpoint(svc WebService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.GetRequest)
		d, err := svc.GetChangeLogDetails(req)
		if err != nil {
			return nil, err
		}
		return d, nil
	}
}

func MakeHandler(svc WebService) http.Handler {
	r := mux.NewRouter()

	GetDetailsHandler := httptransport.NewServer(
		makeGetPackageNameDetailsEndpoint(svc),
		decodeGetRequest,
		encodeResponse,
	)

	GetLogHandler := httptransport.NewServer(
		makeGetLogEndpoint(svc),
		decodeGetLogRequest,
		encodeResponse,
	)

	r.Methods(http.MethodGet).Path("/scrap/{package_name}").Handler(GetDetailsHandler)
	r.Methods(http.MethodGet).Path("/scrap/log/{package_name}").Handler(GetLogHandler)
	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}
