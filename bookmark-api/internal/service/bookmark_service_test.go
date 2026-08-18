package service

import (
	"errors"
	"strings"
	"testing"

	"bookmark-api/internal/models"
)

// fakeBookmarkRepository is an in-memory stand-in for the real repository,
// used only in tests. It satisfies the BookmarkRepository interface, so
// BookmarkService can't tell the difference from the real thing.
type fakeBookmarkRepository struct {
	bookmarks []models.Bookmark
	nextID    uint
	failNext  bool
}

func newFakeRepo() *fakeBookmarkRepository {
	return &fakeBookmarkRepository{nextID: 1}
}

func (f *fakeBookmarkRepository) Create(bookmark *models.Bookmark) error {
	if f.failNext {
		return errors.New("simulated failure")
	}
	bookmark.ID = f.nextID
	f.nextID++
	f.bookmarks = append(f.bookmarks, *bookmark)
	return nil
}

func (f *fakeBookmarkRepository) FindAll(page, limit int, tag string) ([]models.Bookmark, int64, error) {
	var filtered []models.Bookmark
	for _, b := range f.bookmarks {
		if tag == "" || strings.Contains(b.Tags, tag) {
			filtered = append(filtered, b)
		}
	}

	total := int64(len(filtered))
	start := (page - 1) * limit
	if start > len(filtered) {
		return []models.Bookmark{}, total, nil
	}
	end := start + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], total, nil
}

func (f *fakeBookmarkRepository) Delete(id uint) error {
	for i, b := range f.bookmarks {
		if b.ID == id {
			f.bookmarks = append(f.bookmarks[:i], f.bookmarks[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func contains(haystack, needle string) bool {
	return len(needle) > 0 && (haystack == needle ||
		len(haystack) >= len(needle) &&
			(haystack[:len(needle)] == needle ||
				haystack[len(haystack)-len(needle):] == needle ||
				indexOf(haystack, needle) >= 0))
}

func indexOf(haystack, needle string) int {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func TestCreate_AssignsFieldsCorrectly(t *testing.T) {
	repo := newFakeRepo()
	svc := NewBookmarkService(repo)

	bookmark, err := svc.Create("Go Docs", "https://go.dev", "go,docs")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if bookmark.Title != "Go Docs" {
		t.Errorf("expected title %q, got %q", "Go Docs", bookmark.Title)
	}
	if bookmark.ID == 0 {
		t.Errorf("expected a non-zero ID after creation")
	}
}

func TestCreate_PropagatesRepositoryError(t *testing.T) {
	repo := newFakeRepo()
	repo.failNext = true
	svc := NewBookmarkService(repo)

	_, err := svc.Create("Broken", "https://example.com", "")
	if err == nil {
		t.Fatalf("expected an error when the repository fails, got nil")
	}
}

func TestList_FiltersByTag(t *testing.T) {
	repo := newFakeRepo()
	svc := NewBookmarkService(repo)

	svc.Create("Go Docs", "https://go.dev", "go,docs")
	svc.Create("Gin Framework", "https://gin-gonic.com", "go,web")
	svc.Create("Postgres Docs", "https://postgresql.org", "db")

	results, total, err := svc.List(1, 10, "go")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if total != 2 {
		t.Errorf("expected 2 results for tag 'go', got %d", total)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 items in results slice, got %d", len(results))
	}
}

func TestList_Pagination(t *testing.T) {
	repo := newFakeRepo()
	svc := NewBookmarkService(repo)

	for i := 0; i < 5; i++ {
		svc.Create("Bookmark", "https://example.com", "")
	}

	page1, total, err := svc.List(1, 2, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if total != 5 {
		t.Errorf("expected total 5, got %d", total)
	}
	if len(page1) != 2 {
		t.Errorf("expected 2 items on page 1 with limit 2, got %d", len(page1))
	}
}

func TestDelete_RemovesBookmark(t *testing.T) {
	repo := newFakeRepo()
	svc := NewBookmarkService(repo)

	bookmark, _ := svc.Create("To Delete", "https://example.com", "")

	if err := svc.Delete(bookmark.ID); err != nil {
		t.Fatalf("expected no error deleting, got: %v", err)
	}

	_, total, _ := svc.List(1, 10, "")
	if total != 0 {
		t.Errorf("expected 0 bookmarks after delete, got %d", total)
	}
}
