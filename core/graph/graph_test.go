package graph

var mockRailWays = map[string]map[string]int{
	"A": map[string]int{
		"B": 5,
		"D": 5,
		"E": 7,
	},
	"B": map[string]int{
		"C": 4,
	},
	"C": map[string]int{
		"D": 8,
		"E": 2,
	},
	"D": map[string]int{
		"C": 8,
		"E": 6,
	},
	"E": map[string]int{
		"B": 3,
	},
}
