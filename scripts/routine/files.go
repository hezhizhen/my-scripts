package main

// files is a map from categories to their relevant notes (done and to do)
var files = map[string][]fileInfo{
	"learn": []fileInfo{
		{
			FileName: "15554915141859.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"todo": []fileInfo{
		{
			FileName: "15546245073427.md",
			Library:  MWeb3,
			Done:     false,
			Related: &fileInfo{
				FileName: "15447091793464.md",
				Library:  MWeb3,
				Done:     false,
			},
		},
	},
	"coursera": []fileInfo{
		{
			FileName: "15457126390831.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"book": []fileInfo{
		{
			FileName: "15489327025566.md",
			Library:  MWeb3,
			Done:     false,
		},
		{
			FileName: "15469501270291.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"clippings": []fileInfo{
		{
			FileName: "15561193881065.md",
			Library:  MWeb3,
			Done:     true,
		},
		{
			FileName: "15473875670445.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"morningroutine": []fileInfo{
		{
			FileName: "15468295067722.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"wishlist": []fileInfo{
		{
			FileName: "15482192105949.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"podcast": []fileInfo{
		{
			FileName: "15482471177025.md",
			Library:  MWeb3,
			Done:     false,
		},
		{
			FileName: "15489208657451.md",
			Library:  MWeb3,
			Done:     false,
		},
		{
			FileName: "15431467154381.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
	"readinglist": []fileInfo{
		{
			FileName: "15455748961215.md",
			Library:  MWeb3,
			Done:     false,
		},
	},
}
