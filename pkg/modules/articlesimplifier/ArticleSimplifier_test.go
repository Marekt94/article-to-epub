package articlesimplifier

import (
	"os"
	"testing"
)

func TestArticleSimplifierFromURL(t *testing.T) {
	simplifier := &ArticleSimplifierFromURL{}
	_, _, _, err := simplifier.SimplifyArticleInt([]byte("https://fs.blog/mental-models/?utm_source=unknownews"), 30)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestArticleSimplifierForHTML(t *testing.T) {
	simplifier := &ArticleSimplifierFormHTML{}

	strm, err := os.ReadFile("testdata\\test.html")
	if err != nil {
		t.Error(err.Error())
	}
	_, _, _, err = simplifier.SimplifyArticle(strm)
	if err != nil {
		t.Error(err.Error())
	}
}
