package common

type Keyword string

const (
	KeywordUnknown  Keyword = ""
	KeywordFeature  Keyword = "feat"
	KeywordBugFix   Keyword = "fix"
	KeywordDocs     Keyword = "docs"
	KeywordStyle    Keyword = "style"
	KeywordRefactor Keyword = "refactor"
	KeywordPref     Keyword = "pref"
	KeywordTest     Keyword = "test"
	KeywordChore    Keyword = "chore"
	KeywordRevert   Keyword = "revert"

	TimeFormat          = "2006-01-02 15:04:05"
	GroupKeywordFeature = "Feature"
	GroupKeywordOther   = "Other"
)

var KeywordList = []Keyword{
	KeywordFeature,
	KeywordBugFix,
	KeywordDocs,
	KeywordStyle,
	KeywordRefactor,
	KeywordPref,
	KeywordTest,
	KeywordChore,
	KeywordRevert,
}
