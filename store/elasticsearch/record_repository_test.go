package elasticsearch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/odpf/compass/asset"
	"github.com/odpf/compass/discovery"
	store "github.com/odpf/compass/store/elasticsearch"
	"github.com/stretchr/testify/assert"
)

func TestRecordRepository(t *testing.T) {
	ctx := context.Background()

	t.Run("CreateOrReplaceMany", func(t *testing.T) {
		var testCases = []struct {
			Title      string
			ShouldFail bool
			Setup      func(cli *elasticsearch.Client, assets []asset.Asset, typeName string) error
			PostCheck  func(cli *elasticsearch.Client, assets []asset.Asset, typeName string) error
			Type       string
			Assets     []asset.Asset
		}{
			{
				Title: "should succesfully write all the documents to the index for a valid type",
				Type:  "job",
				Assets: []asset.Asset{
					{
						URN: "dagger1",
						Data: map[string]interface{}{
							"foo": "bar",
						},
					},
					{
						URN: "dagger2",
						Data: map[string]interface{}{
							"foo": "bar",
						},
					},
					{
						URN: "dagger3",
						Data: map[string]interface{}{
							"foo": "bar",
						},
					},
				},
				PostCheck: func(cli *elasticsearch.Client, assets []asset.Asset, typeName string) error {
					searchReq := esapi.SearchRequest{
						Index: []string{typeName},
						Body:  strings.NewReader(`{"query":{"match_all":{}}}`),
					}
					res, err := searchReq.Do(context.Background(), cli)
					if err != nil {
						return fmt.Errorf("error querying elasticsearch: %w", err)
					}
					defer res.Body.Close()
					if res.IsError() {
						return fmt.Errorf("elasticsearch query returned error: %s", res.Status())
					}

					var response = struct {
						Hits struct {
							Hits []interface{} `json:"hits"`
						} `json:"hits"`
					}{}
					err = json.NewDecoder(res.Body).Decode(&response)
					if err != nil {
						return fmt.Errorf("error parsing elasticsearch response: %w", err)
					}
					if len(assets) != len(response.Hits.Hits) {
						return fmt.Errorf("expected elasticsearch index to contain %d assets, but had %d assets instead", len(assets), len(response.Hits.Hits))
					}

					return nil
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.Title, func(t *testing.T) {
				cli := esTestServer.NewClient()
				if testCase.Setup != nil {
					err := testCase.Setup(cli, testCase.Assets, testCase.Type)
					if err != nil {
						t.Errorf("error setting up testcase: %v", err)
					}
				}
				factory := store.NewRecordRepositoryFactory(cli)
				repo, err := factory.For(testCase.Type)
				if err != nil {
					t.Fatalf("error creating asset repository: %s", err)
				}

				err = repo.CreateOrReplaceMany(ctx, testCase.Assets)
				if testCase.ShouldFail {
					assert.Error(t, err)
				} else if err != nil {
					t.Errorf("repository returned unexpected error: %v", err)
					return
				}
				if testCase.PostCheck != nil {
					if err := testCase.PostCheck(cli, testCase.Assets, testCase.Type); err != nil {
						t.Error(err)
						return
					}
				}
			})
		}
	})

	cli := esTestServer.NewClient()
	rrf := store.NewRecordRepositoryFactory(cli)
	assetRepo, err := rrf.For("topic")
	if err != nil {
		t.Fatalf("failed to construct asset repository: %v", err)
		return
	}

	assets := insertRecord(ctx, t, assetRepo)

	t.Run("GetAll", func(t *testing.T) {
		type testCase struct {
			Description   string
			Filter        discovery.RecordFilter
			ResultsFile   string
			From          int
			Size          int
			ExpectedTotal int
		}

		var testCases = []testCase{
			{
				Description: "should handle nil filter and default sort by name",
				Filter:      nil,
				ResultsFile: "./testdata/assets-all.json",
			},
			{
				Description:   "should fetch certain offset and size if given",
				Filter:        nil,
				From:          2,
				Size:          3,
				ResultsFile:   "./testdata/assets-offset.json",
				ExpectedTotal: 10,
			},
			{
				Description: "should handle filter by service",
				Filter: map[string][]string{
					"service": {"rabbitmq"},
				},
				ResultsFile: "./testdata/assets-service.json",
			},
			{
				Description: "should support a single value filter",
				Filter: map[string][]string{
					"data.country": {"id"},
				},
				ResultsFile: "./testdata/assets-id.json",
			},
			{
				Description: "should support multi value filter",
				Filter: map[string][]string{
					"data.country": {"id", "vn"},
				},
				ResultsFile: "./testdata/assets-vn-id.json",
			},
			{
				Description: "should support multiple terms",
				Filter: map[string][]string{
					"data.country": {"th"},
					"data.title":   {"test_grant2"},
				},
				ResultsFile: "./testdata/assets-th-deployed.json",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.Description, func(t *testing.T) {
				expectedResults := []asset.Asset{}
				raw, err := ioutil.ReadFile(tc.ResultsFile)
				if err != nil {
					t.Fatalf("error reading results file: %v", err)
					return
				}
				err = json.Unmarshal(raw, &expectedResults)
				if err != nil {
					t.Fatalf("error parsing results file: %v", err)
					return
				}

				assetList, err := assetRepo.GetAll(ctx, discovery.GetConfig{
					Filters: tc.Filter,
					From:    tc.From,
					Size:    tc.Size,
				})
				if err != nil {
					t.Fatalf("error executing GetAll: %v", err)
					return
				}

				assert.Equal(t, len(expectedResults), assetList.Count)
				if reflect.DeepEqual(expectedResults, assetList.Data) == false {
					t.Error(incorrectResultsError(expectedResults, assetList.Data))
					return
				}

				if tc.ExpectedTotal > 0 {
					assert.Equal(t, tc.ExpectedTotal, assetList.Total)
				}
			})
		}
	})
	t.Run("GetByID", func(t *testing.T) {
		t.Run("data-based tests", func(t *testing.T) {
			for _, ast := range assets {
				assetFromRepo, err := assetRepo.GetByID(ctx, ast.URN)
				if err != nil {
					t.Errorf("unexpected error: GetByID(%q): %v", ast.URN, err)
					return
				}
				if reflect.DeepEqual(ast, assetFromRepo) == false {
					t.Error(incorrectResultsError(ast, assetFromRepo))
				}
			}
		})
		t.Run("should return an error if a non-existent asset is requested", func(t *testing.T) {
			var id = "this-doesnt-exists"
			_, err := assetRepo.GetByID(ctx, id)
			_, ok := err.(asset.NotFoundError)
			assert.True(t, ok)
		})
	})
	t.Run("Delete", func(t *testing.T) {
		t.Run("should delete asset from index", func(t *testing.T) {
			id := "delete-id-01"
			err := assetRepo.CreateOrReplaceMany(ctx, []asset.Asset{
				{
					URN:  id,
					Name: "To be deleted",
					Data: map[string]interface{}{
						"title": "To be deleted",
						"urn":   id,
					},
				},
			})
			if err != nil {
				t.Fatal(err)
			}

			err = assetRepo.Delete(ctx, id)
			assert.Nil(t, err)

			r, err := assetRepo.GetByID(ctx, id)
			assert.NotNil(t, err)
			assert.Equal(t, asset.Asset{}, r)
		})

		t.Run("should return custom error when asset could not be found", func(t *testing.T) {
			err := assetRepo.Delete(ctx, "not-found-id")
			assert.NotNil(t, err)
			assert.IsType(t, asset.NotFoundError{}, err)
		})
	})
}

func insertRecord(ctx context.Context, t *testing.T, repo discovery.RecordRepository) (assets []asset.Asset) {
	src, err := ioutil.ReadFile("./testdata/assets.json")
	if err != nil {
		t.Fatalf("error reading testdata: %v", err)
		return
	}

	err = json.Unmarshal(src, &assets)
	if err != nil {
		t.Fatalf("error unmarshalling testdata: %v", err)
		return
	}
	err = repo.CreateOrReplaceMany(ctx, assets)
	if err != nil {
		t.Fatalf("error writing testdata to elasticsearch: %v", err)
		return
	}

	return
}
