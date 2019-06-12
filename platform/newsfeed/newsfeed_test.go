package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{
		"A Title", "And a Post",
	})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{
		"A Title", "And a Post",
	})

	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
