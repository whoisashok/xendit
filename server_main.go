package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	serviceHTTPXendit "xendit/service"
	usecase "xendit/xendit_services/usecase"

	raven "github.com/getsentry/raven-go"
	logging "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	context "golang.org/x/net/context"

	cfg "xendit/config"
	endpointHTTPXendit "xendit/endpoint"
	transportHttpxendit "xendit/transport_http"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	if config.GetBool(`app.debug`) {
		fmt.Println("Payment Service RUN on DEBUG mode")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	raven.SetDSN(config.GetString(`sentry`))
}

func main() {

	httpServerAddr := config.GetString(`app.host`)

	urlVal := url.Values{}
	urlVal.Add("parseTime", "1")
	urlVal.Add("loc", "Asia/Jakarta")
	/*
		dbHost := config.GetString(`database.host`)
		dbPort := config.GetString(`database.port`)
		dbUser := config.GetString(`database.user`)
		dbPass := config.GetString(`database.pass`)
		dbName := config.GetString(`database.name`)
		db, err := gorm.Open("postgres", "host="+dbHost+" port="+dbPort+" user="+dbUser+" dbname="+dbName+" password="+dbPass+" sslmode=disable")
		defer db.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
	*/

	// raven.CaptureErrorAndWait(errors.New("custom error"), nil)
	// if err != nil && config.GetBool("app.debug") {
	// 	raven.CaptureErrorAndWait(err, nil)
	// 	//fmt.Println(err)
	// }

	logFile, err := os.OpenFile(config.GetString(`logfile`), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		panic(err)
	}
	defer logFile.Close()

	// Logging domain.
	var logger logging.Logger
	{
		w := logging.NewSyncWriter(logFile)
		logger = logging.NewLogfmtLogger(w)
		// AllowAll() AllowDebug() AllowInfo() AllowWarn() AllowError() AllowNone()
		// More info @ https://github.com/go-kit/kit/blob/master/log/level/level.go
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = logging.With(logger, "ts", logging.DefaultTimestampUTC)
		logger = logging.With(logger, "caller", logging.DefaultCaller)
	}

	ctx := context.Background()

	// Repository
	//repoMidTrans := repositoryPayment.NewMidTransRepositoryImpl(db, logger)

	// Usecase

	usecaseXendit := usecase.NewXenditUsecaseImpl(logger)

	// init service

	var xenditService serviceHTTPXendit.XenditService
	xenditService = serviceHTTPXendit.NewXenditServiceImpl(usecaseXendit, logger)
	errChan := make(chan error)

	// creating Endpoint struct
	endpointXendit := endpointHTTPXendit.XenditEndpoints{

		CreateDisbursementEndpoint:          endpointHTTPXendit.MakeCreateDisbursementEndpoint(xenditService),
		GetDisbursementByIdEndpoint:         endpointHTTPXendit.MakeGetDisbursementByIdEndpoint(xenditService),
		GetDisbursementByExternalIDEndpoint: endpointHTTPXendit.MakeGetDisbursementByExternalIDEndpoint(xenditService),
		DisbursementCallbackEndpoint:        endpointHTTPXendit.MakeDisbursementCallbackEndpoint(xenditService),
		CreateBulkDisbursementEndpoint:      endpointHTTPXendit.MakeCreateBulkDisbursementEndpoint(xenditService),
		BulkDisbursementCallbackEndpoint:    endpointHTTPXendit.MakeBulkDisbursementCallbackEndpoint(xenditService),
	}

	// HTTP handler
	httpHandlerXendit := transportHttpxendit.MakeXenditHttpHandler(ctx, endpointXendit, logger)

	// HTTP transport
	go func() {
		fmt.Println("Payment Service Http server at port: ", httpServerAddr)
		errChan <- http.ListenAndServe(httpServerAddr, httpHandlerXendit)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}
