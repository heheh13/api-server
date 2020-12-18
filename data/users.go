package data

//Profile defines the structure of the user profile
type Profile struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Skills Skills `json:"skills"`
}

//Skills store the skill of a user
type Skills struct {
	Language []string `json:"language"`
	Tools    []string `json:"tools"`
	Endorsed int      `json:"endorsed"`
}

//User profile
type User []Profile

// Users store data
var Users = User{
	Profile{
		ID:   "1",
		Name: "Mehedi Hasan",
		Skills: Skills{
			Language: []string{"c++,go"},
			Tools:    []string{"git", "linux"},
		},
	},
	Profile{
		ID:   "2",
		Name: "Sahadat",
		Skills: Skills{
			Language: []string{"c++,go"},
			Tools:    []string{"git", "linux"},
			Endorsed: 5,
		},
	},
	Profile{
		ID:   "3",
		Name: "Pulak",
		Skills: Skills{
			Language: []string{"c++,go"},
			Tools:    []string{"git", "linux"},
			Endorsed: 5,
		},
	},
	Profile{
		ID:   "4",
		Name: "Sakib",
		Skills: Skills{
			Language: []string{"c++,go"},
			Tools:    []string{"git", "linux"},
			Endorsed: 5,
		},
	},
	Profile{
		ID:   "5",
		Name: "Prangan",
		Skills: Skills{
			Language: []string{"c++,go"},
			Tools:    []string{"git", "linux"},
			Endorsed: 5,
		},
	},
}
