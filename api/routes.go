package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/odpf/columbus/api/handlers"
	"github.com/odpf/columbus/models"
	"github.com/odpf/columbus/tag"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger                  logrus.FieldLogger
	TypeRepository          models.TypeRepository
	RecordRepositoryFactory models.RecordRepositoryFactory
	RecordSearcher          models.RecordSearcher
	TagService              *tag.Service
	TagTemplateService      *tag.TemplateService
	LineageProvider         handlers.LineageProvider
}

func RegisterRoutes(router *mux.Router, config Config) {
	// By default mux will decode url and then match the decoded url against the route
	// we reverse the steps by telling mux to use encoded path to match the url
	// then we manually decode via custom middleware (decodeURLMiddleware).
	//
	// This is to allow urn that has "/" to be matched correctly to the route
	router.UseEncodedPath()
	router.Use(decodeURLMiddleware(config.Logger))

	typeHandler := handlers.NewTypeHandler(
		config.Logger.WithField("reporter", "type-handler"),
		config.TypeRepository,
		config.RecordRepositoryFactory,
	)
	searchHandler := handlers.NewSearchHandler(
		config.Logger.WithField("reporter", "search-handler"),
		config.RecordSearcher,
		config.TypeRepository,
	)
	lineageHandler := handlers.NewLineageHandler(
		config.Logger.WithField("reporter", "lineage-handler"),
		config.LineageProvider,
	)
	tagHandler := handlers.NewTagHandler(
		config.Logger.WithField("reporter", "tag-handler"),
		config.TagService,
	)
	tagTemplateHandler := handlers.NewTagTemplateHandler(
		config.Logger.WithField("reporter", "tag-template-handler"),
		config.TagTemplateService,
	)

	router.PathPrefix("/ping").Handler(handlers.NewHeartbeatHandler())
	setupV1TypeRoutes(router, "/v1/types", typeHandler)
	setupV1TagRoutes(router, "/v1/tags", tagHandler, tagTemplateHandler)

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
		HandlerFunc(typeHandler.DeleteRecord)

	router.Path(baseURL + "/{name}/records").
		Methods(http.MethodPut).
		HandlerFunc(typeHandler.IngestRecord)

	router.Path(baseURL+"/{name}/records/{id}").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(typeHandler.GetTypeRecord)

}

func setupV1TagRoutes(router *mux.Router, baseURL string, th *handlers.TagHandler, tth *handlers.TagTemplateHandler) {
	router.Methods(http.MethodPost).Path(baseURL).HandlerFunc(th.Create)

	url := baseURL + "/types/{type}/records/{record_urn}/templates/{template_urn}"
	router.Methods(http.MethodGet).Path(url).HandlerFunc(th.FindByRecordAndTemplate)
	router.Methods(http.MethodPut).Path(url).HandlerFunc(th.Update)
	router.Methods(http.MethodDelete).Path(url).HandlerFunc(th.Delete)

	router.Methods(http.MethodGet).Path(baseURL + "/types/{type}/records/{record_urn}").HandlerFunc(th.GetByRecord)

	templateURL := baseURL + "/templates"
	router.Methods(http.MethodGet).Path(templateURL).HandlerFunc(tth.Index)
	router.Methods(http.MethodPost).Path(templateURL).HandlerFunc(tth.Create)
	router.Methods(http.MethodGet).Path(templateURL + "/{template_urn}").HandlerFunc(tth.Find)
	router.Methods(http.MethodPut).Path(templateURL + "/{template_urn}").HandlerFunc(tth.Update)
	router.Methods(http.MethodDelete).Path(templateURL + "/{template_urn}").HandlerFunc(tth.Delete)
}
