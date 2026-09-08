package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const summaryCacheTTL = 2 * time.Minute

// CachedSummaryRepository wraps SummaryRepository with a cache-aside layer:
// check Redis first, fall back to the real (expensive) query on a miss,
// and populate the cache before returning.
type CachedSummaryRepository struct {
	repo  *SummaryRepository
	cache *redis.Client
}

func NewCachedSummaryRepository(repo *SummaryRepository, cache *redis.Client) *CachedSummaryRepository {
	return &CachedSummaryRepository{repo: repo, cache: cache}
}

func summaryCacheKey(orgID uint) string {
	return fmt.Sprintf("org-summary:%d", orgID)
}

func (r *CachedSummaryRepository) GetSummary(ctx context.Context, orgID uint) (*OrgSummary, error) {
	key := summaryCacheKey(orgID)

	cached, err := r.cache.Get(ctx, key).Result()
	if err == nil {
		var summary OrgSummary
		if jsonErr := json.Unmarshal([]byte(cached), &summary); jsonErr == nil {
			return &summary, nil // cache hit
		}
	}

	summary, err := r.repo.GetSummary(orgID)
	if err != nil {
		return nil, err
	}

	if encoded, jsonErr := json.Marshal(summary); jsonErr == nil {
		r.cache.Set(ctx, key, encoded, summaryCacheTTL)
	}

	return summary, nil
}

// InvalidateSummary should be called whenever an invoice's status changes,
// since that's exactly what makes the cached summary stale.
func (r *CachedSummaryRepository) InvalidateSummary(ctx context.Context, orgID uint) {
	r.cache.Del(ctx, summaryCacheKey(orgID))
}
