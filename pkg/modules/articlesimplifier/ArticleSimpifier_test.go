package articlesimplifier

import (
	"testing"
)

func TestArticleSimplifier(t *testing.T) {
	simplifier := ArticleSimplifier{}
	_, err := simplifier.SimplifyArticle("http://example.com/article")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
