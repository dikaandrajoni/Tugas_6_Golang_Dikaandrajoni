package middleware

import (
	"time"

	"context"

	cm "Tugas_6_Golang_Dikaandrajoni/Framework/common"
	"Tugas_6_Golang_Dikaandrajoni/Framework/services"

	log "github.com/Sirupsen/logrus"
)

func BasicMiddleware() services.ServiceMiddleware {
	return func(next services.PaymentServices) services.PaymentServices {
		return BasicMiddlewareStruct{next}
	}
}

type BasicMiddlewareStruct struct {
	services.PaymentServices
}

func (mw BasicMiddlewareStruct) OrderHandler(ctx context.Context, request cm.Order) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("OrderHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("OrderHandler begins")

	return mw.PaymentServices.OrderHandler(ctx, request)

}

func (mw BasicMiddlewareStruct) CustomerHandler(ctx context.Context, request cm.Customer) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("CustomerHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("CustomerHandler begins")

	return mw.PaymentServices.CustomerHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) ProductHandler(ctx context.Context, request cm.Product) cm.Message {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("ProductHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("ProductHandler begins")

	return mw.PaymentServices.ProductHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) FaspayHandler(ctx context.Context, request cm.RequestFaspay) cm.ResponseFaspay {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("FaspayHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("FaspayHandler begins")

	return mw.PaymentServices.FaspayHandler(ctx, request)
}

func (mw BasicMiddlewareStruct) TripHandler(ctx context.Context, request cm.RequestTrip) cm.ResponseTrip {
	defer func(begin time.Time) {
		log.WithField("execTime", float64(time.Since(begin).Nanoseconds())/float64(1e6)).Info("TripHandler ends")
	}(time.Now())

	log.WithField("request", request).Info("TripHandler begins")

	return mw.PaymentServices.TripHandler(ctx, request)
}
