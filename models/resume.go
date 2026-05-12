package models

type Meta struct {
	Name     string `yaml:"name"`
	Title    string `yaml:"title"`
	Email    string `yaml:"email"`
	Phone    string `yaml:"phone"`
	Location string `yaml:"location"`
	LinkedIn string `yaml:"linkedin"`
	GitHub   string `yaml:"github"`
	Website  string `yaml:"website"`
	Summary  string `yaml:"summary"`
}

type Experience struct {
	Company  string   `yaml:"company"`
	Title    string   `yaml:"title"`
	Location string   `yaml:"location"`
	Start    string   `yaml:"start"`
	End      string   `yaml:"end"`
	Current  bool     `yaml:"current"`
	Bullets  []string `yaml:"bullets"`
}

type Education struct {
	Institution string `yaml:"institution"`
	Degree      string `yaml:"degree"`
	Field       string `yaml:"field"`
	Start       int    `yaml:"start"`
	End         int    `yaml:"end"`
	GPA         string `yaml:"gpa"`
	Honors      string `yaml:"honors"`
}

type SkillGroup struct {
	Category string   `yaml:"category"`
	Items    []string `yaml:"items"`
}

type Project struct {
	Name        string   `yaml:"name"`
	URL         string   `yaml:"url"`
	Description string   `yaml:"description"`
	Tech        []string `yaml:"tech"`
}

type Certification struct {
	Name   string `yaml:"name"`
	Issuer string `yaml:"issuer"`
	Date   string `yaml:"date"`
	URL    string `yaml:"url"`
}

type Language struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type Resume struct {
	Meta           Meta            `yaml:"meta"`
	Experience     []Experience    `yaml:"experience"`
	Education      []Education     `yaml:"education"`
	Skills         []SkillGroup    `yaml:"skills"`
	Languages      []Language      `yaml:"languages"`
	Projects       []Project       `yaml:"projects"`
	Certifications []Certification `yaml:"certifications"`
}
