package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/odpf/columbus/api/handlers"
	"github.com/odpf/columbus/models"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger                  logrus.FieldLogger
	TypeRepository          models.TypeRepository
	RecordRepositoryFactory models.RecordRepositoryFactory
	RecordV1Searcher        models.RecordV1Searcher
	LineageProvider         handlers.LineageProvider
}

func RegisterRoutes(router *mux.Router, config Config) {
	typeHandler := handlers.NewTypeHandler(
		config.Logger.WithField("reporter", "type-handler"),
		config.TypeRepository,
		config.RecordRepositoryFactory,
	)
	searchHandler := handlers.NewSearchHandler(
		config.Logger.WithField("reporter", "search-handler"),
		config.RecordV1Searcher,
		config.TypeRepository,
	)

	lineageHandler := handlers.NewLineageHandler(
		config.Logger.WithField("reporter", "lineage-handler"),
		config.LineageProvider,
	)

	router.PathPrefix("/ping").Handler(handlers.NewHeartbeatHandler())
	setupV1TypeRoutes(router, "/v1/types", typeHandler)
	setupV2TypeRoutes(router, "/v2/types", typeHandler)

	router.Path("/v1/search").
		Methods(http.MethodGet).
		HandlerFunc(searchHandler.Search)

	router.PathPrefix("/v1/lineage/{type}/{id}").
		Methods(http.MethodGet).
		HandlerFunc(lineageHandler.GetLineage)

	router.PathPrefix("/v1/lineage").
		Methods(http.MethodGet).
		HandlerFunc(lineageHandler.ListLineage)
}

func setupV1TypeRoutes(router *mux.Router, baseURL string, typeHandler *handlers.TypeHandler) {
	router.Path(baseURL).
		Methods(http.MethodGet).
		HandlerFunc(typeHandler.GetAll)

	router.Path(baseURL+"/{name}").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(typeHandler.GetType)

	router.Path(baseURL+"/{name}/records").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(typeHandler.ListTypeRecords)

	router.Path(baseURL).
		Methods(http.MethodPut).
		HandlerFunc(typeHandler.CreateOrReplaceType)

	router.Path(baseURL + "/{name}").
		Methods(http.MethodDelete).
		HandlerFunc(typeHandler.DeleteType)

	router.Path(baseURL + "/{name}/records/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(typeHandler.DeleteRecordV1)

	router.Path(baseURL + "/{name}/records").
		Methods(http.MethodPut).
		HandlerFunc(typeHandler.IngestRecordV1)

	router.Path(baseURL+"/{name}/records/{id}").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(typeHandler.GetTypeRecord)
}

func setupV2TypeRoutes(router *mux.Router, baseURL string, typeHandler *handlers.TypeHandler) {
	router.Path(baseURL + "/{name}/records").
		Methods(http.MethodPut).
		HandlerFunc(typeHandler.IngestRecordV2)
}
