package articlesimplifier

import (
	"testing"
)

func TestArticleSimplifier(t *testing.T) {
	simplifier := &ArticleSimplifier{}
	_, _, _, err := simplifier.SimplifyArticleInt([]byte("https://fs.blog/mental-models/?utm_source=unknownews"), 30)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
