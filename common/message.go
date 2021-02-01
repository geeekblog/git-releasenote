package common

var (
	messageBugFix                  string
	messagePerformanceOptimization string
	messageNewestTag               string
)

func init() {
	if messageBugFix == "" {
		messageBugFix = "bug fix"
	}

	if messagePerformanceOptimization == "" {
		messagePerformanceOptimization = "performance optimization"
	}

	if messageNewestTag == "" {
		messageNewestTag = "NEWEST"
	}
}
