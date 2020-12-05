package transport

import (
	"context"

	cm "Tugas_6_Golang_Dikaandrajoni/Framework/common"
	"Tugas_6_Golang_Dikaandrajoni/Frameworkservices"

	log "github.com/Sirupsen/logrus"

	"github.com/go-kit/kit/endpoint"
)

func invalidRequest() cm.Message {
	return cm.Message{
		Result: &cm.Result{
			Code:   99,
			Remark: "Invalid Request",
		},
	}
}

func OrderEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(cm.Order); ok {
			return svc.OrderHandler(ctx, req), nil
		}

		log.WithField("Error", request).Info("Request in unknown format")
		return invalidRequest(), nil
	}
}

func CustomerEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(cm.Customer); ok {
			return svc.CustomerHandler(ctx, req), nil
		}

		log.WithField("Error", request).Info("Request in unknown format")
		return invalidRequest(), nil
	}
}

func ProductEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(cm.Product); ok {
			return svc.ProductHandler(ctx, req), nil
		}

		log.WithField("Error", request).Info("Request in unknown format")
		return invalidRequest(), nil
	}
}

func FaspayEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(cm.RequestFaspay); ok {
			return svc.FaspayHandler(ctx, req), nil
		}

		log.WithField("Error", request).Info("Request in unknown format")
		return invalidRequest(), nil
	}
}

func TripEndpoint(svc services.PaymentServices) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if req, ok := request.(cm.RequestTrip); ok {
			return svc.TripHandler(ctx, req), nil
		}

		log.WithField("Error", request).Info("Request in unknown format")
		return invalidRequest(), nil
	}
}
