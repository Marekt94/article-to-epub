package misc

import "testing"

func TestAdaptUrlToFileName(t *testing.T) {
	const cUrl = "https://claude.com/blog/how-claude-code-works-in-large-codebases-best-practices-and-where-to-start?utm_source=unknownews#COMMENT"
	const cExpected = "how-claude-code-works-in-large-c"

	res, _ := AdaptUrlToFileName(cUrl)

	if res != cExpected {
		t.Errorf("Expected %s but have %s", cExpected, res)
	}
}

func TestAdaptUrlToFileNameShort(t *testing.T) {
	const cUrl = "https://claude.com/blog/how-claude-code-works?utm_source=unknownews#COMMENT"
	const cExpected = "how-claude-code-works"

	res, _ := AdaptUrlToFileName(cUrl)

	if res != cExpected {
		t.Errorf("Expected %s but have %s", cExpected, res)
	}
}

func TestAdaptUrlToFileNameWhenURLWrong(t *testing.T) {
	const cUrl = "ps:/claude.com/blog/how-claude-code-works?utm_source=unknownews#COMMENT"
	const cExpected = "how-claude-code-works"

	res, _ := AdaptUrlToFileName(cUrl)

	if res != cExpected {
		t.Errorf("Expected %s but have %s", cExpected, res)
	}
}

func TestAdaptUrlToFileNameWhenURLEmpty(t *testing.T) {
	const cUrl = ""
	const cExpected = DEFAULT_FILE_NAME

	res, _ := AdaptUrlToFileName(cUrl)

	if res != cExpected {
		t.Errorf("Expected %s but have %s", cExpected, res)
	}
}

func TestAdaptUrlToFileNameWhenPathToFile(t *testing.T) {
	const cUrl = "C://marek//doc//file.html"
	const cExpected = "file.html"

	res, _ := AdaptUrlToFileName(cUrl)

	if res != cExpected {
		t.Errorf("Expected %s but have %s", cExpected, res)
	}
}
