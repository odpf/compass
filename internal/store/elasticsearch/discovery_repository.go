package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/odpf/compass/core/asset"
)

// DiscoveryRepository implements discovery.Repository
// with elasticsearch as the backing store.
type DiscoveryRepository struct {
	cli *Client
}

func NewDiscoveryRepository(cli *Client) *DiscoveryRepository {
	return &DiscoveryRepository{
		cli: cli,
	}
}

func (repo *DiscoveryRepository) Upsert(ctx context.Context, ast asset.Asset) error {
	if ast.ID == "" {
		return asset.ErrEmptyID
	}
	if !ast.Type.IsValid() {
		return asset.ErrUnknownType
	}

	idxExists, err := repo.cli.indexExists(ctx, ast.Service)
	if err != nil {
		return elasticSearchError(err)
	}

	if !idxExists {
		if err := repo.cli.CreateIdx(ctx, ast.Service); err != nil {
			return err
		}
	}

	body, err := repo.createUpsertBody(ast)
	if err != nil {
		return fmt.Errorf("error serialising payload: %w", err)
	}
	res, err := repo.cli.client.Bulk(
		body,
		repo.cli.client.Bulk.WithRefresh("true"),
		repo.cli.client.Bulk.WithContext(ctx),
	)
	if err != nil {
		return elasticSearchError(err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error response from elasticsearch: %s", errorReasonFromResponse(res))
	}
	return nil
}

func (repo *DiscoveryRepository) Delete(ctx context.Context, assetID string) error {
	if assetID == "" {
		return asset.ErrEmptyID
	}

	res, err := repo.cli.client.DeleteByQuery(
		[]string{"_all"},
		strings.NewReader(fmt.Sprintf(`{"query":{"terms":{"_id": ["%s"]}}}`, assetID)),
		repo.cli.client.DeleteByQuery.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("error deleting asset: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error response from elasticsearch: %s", errorReasonFromResponse(res))
	}

	return nil
}

func (repo *DiscoveryRepository) createUpsertBody(ast asset.Asset) (io.Reader, error) {
	payload := bytes.NewBuffer(nil)
	err := repo.writeInsertAction(payload, ast)
	if err != nil {
		return nil, fmt.Errorf("createBulkInsertPayload: %w", err)
	}

	err = json.NewEncoder(payload).Encode(ast)
	if err != nil {
		return nil, fmt.Errorf("error serialising asset: %w", err)
	}
	return payload, nil
}

func (repo *DiscoveryRepository) writeInsertAction(w io.Writer, ast asset.Asset) error {
	action := map[string]interface{}{
		"index": map[string]interface{}{
			"_index": ast.Service,
			"_id":    ast.ID,
		},
	}

	return json.NewEncoder(w).Encode(action)
}
