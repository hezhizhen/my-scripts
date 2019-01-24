package main

// files is a map from categories to their relevant notes (done and to do)
var files = map[string][]fileInfo{
	"coursera": []fileInfo{
		{
			FileName: "15457126390831.md",
			Title:    "《新教伦理与资本主义精神》导读",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"book": []fileInfo{
		{
			FileName: "15469501270291.md",
			Title:    "Go by Example",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"clippings": []fileInfo{
		{
			FileName: "15473875670445.md",
			Title:    "白夜行",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"morningroutine": []fileInfo{
		{
			FileName: "15468295067722.md",
			Title:    "My Morning Routine (Collection)",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"wishlist": []fileInfo{
		{
			FileName: "15482192105949.md",
			Title:    "WishList",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"podcast": []fileInfo{
		{
			FileName: "15482471177025.md",
			Title:    "忽左忽右",
			Library:  MWeb3,
			Done:     false,
		},
	},
}
