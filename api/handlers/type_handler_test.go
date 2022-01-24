package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/odpf/salt/log"

	"github.com/odpf/columbus/api/handlers"
	"github.com/odpf/columbus/asset"
	"github.com/odpf/columbus/lib/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeHandler(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		type testCase struct {
			Description  string
			ExpectStatus int
			Setup        func(tc *testCase, er *mocks.TypeRepository)
			PostCheck    func(t *testing.T, tc *testCase, resp *http.Response) error
		}

		var testCases = []testCase{
			{
				Description:  "should return 500 status code if failing to fetch types",
				ExpectStatus: http.StatusInternalServerError,
				Setup: func(tc *testCase, er *mocks.TypeRepository) {
					er.On("GetAll", context.Background()).Return(map[asset.Type]int{}, errors.New("failed to fetch type"))
				},
			},
			{
				Description:  "should return 500 status code if failing to fetch counts",
				ExpectStatus: http.StatusInternalServerError,
				Setup: func(tc *testCase, er *mocks.TypeRepository) {
					er.On("GetAll", context.Background()).Return(map[asset.Type]int{}, errors.New("failed to fetch assets count"))
				},
			},
			{
				Description:  "should return all valid types with its record count",
				ExpectStatus: http.StatusOK,
				Setup: func(tc *testCase, er *mocks.TypeRepository) {
					er.On("GetAll", context.Background()).Return(map[asset.Type]int{
						asset.Type("table"): 10,
						asset.Type("topic"): 30,
						asset.Type("job"):   15,
					}, nil)
				},
				PostCheck: func(t *testing.T, tc *testCase, resp *http.Response) error {
					actual, err := ioutil.ReadAll(resp.Body)
					require.NoError(t, err)

					expected, err := json.Marshal([]map[string]interface{}{
						{"name": "table", "count": 10},
						{"name": "job", "count": 15},
						{"name": "dashboard", "count": 0},
						{"name": "topic", "count": 30},
					})
					require.NoError(t, err)

					assert.JSONEq(t, string(expected), string(actual))

					return nil
				},
			},
		}
		for _, tc := range testCases {
			t.Run(tc.Description, func(t *testing.T) {
				er := new(mocks.TypeRepository)
				logger := log.NewNoop()
				defer er.AssertExpectations(t)
				tc.Setup(&tc, er)

				handler := handlers.NewTypeHandler(logger, er)
				rr := httptest.NewRequest("GET", "/", nil)
				rw := httptest.NewRecorder()

				handler.Get(rw, rr)
				if rw.Code != tc.ExpectStatus {
					t.Errorf("expected handler to return %d status, was %d instead", tc.ExpectStatus, rw.Code)
					return
				}

				if tc.PostCheck != nil {
					if err := tc.PostCheck(t, &tc, rw.Result()); err != nil {
						t.Error(err)
					}
				}
			})
		}
	})
}
