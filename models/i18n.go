package models

type I18nNav struct {
	Home          string `yaml:"home"`
	Summary       string `yaml:"summary"`
	Experience    string `yaml:"experience"`
	Skills        string `yaml:"skills"`
	Links         string `yaml:"links"`
	LangSwitch    string `yaml:"lang_switch"`
	LangSwitchURL string `yaml:"lang_switch_url"`
}

type I18nStatus struct {
	OpenToWork string `yaml:"open_to_work"`
}

type I18nActions struct {
	DownloadCV   string `yaml:"download_cv"`
	SendEmail    string `yaml:"send_email"`
	ViewGithub   string `yaml:"view_github"`
	ViewLinkedin string `yaml:"view_linkedin"`
	ViewProject  string `yaml:"view_project"`
}

type I18nHero struct {
	Tagline      string `yaml:"tagline"`
	CTAPrimary   string `yaml:"cta_primary"`
	CTASecondary string `yaml:"cta_secondary"`
}

type QuickFact struct {
	Emoji string `yaml:"emoji"`
	Label string `yaml:"label"`
}

type I18nSummary struct {
	Title      string      `yaml:"title"`
	Greeting   string      `yaml:"greeting"`
	Paragraphs []string    `yaml:"paragraphs"`
	QuickFacts []QuickFact `yaml:"quick_facts"`
}

type I18nExperience struct {
	Title   string     `yaml:"title"`
	Present string     `yaml:"present"`
	Bullets [][]string `yaml:"bullets"`
}

type I18nSkills struct {
	Title          string `yaml:"title"`
	TechTitle      string `yaml:"tech_title"`
	LanguagesTitle string `yaml:"languages_title"`
	EducationTitle string `yaml:"education_title"`
	CertsTitle     string `yaml:"certs_title"`
}

type I18nLinks struct {
	Title         string `yaml:"title"`
	ProjectsTitle string `yaml:"projects_title"`
	ContactTitle  string `yaml:"contact_title"`
	ContactBody   string `yaml:"contact_body"`
}

type I18nFooter struct {
	Built string `yaml:"built"`
}

type I18nContactModal struct {
	Title              string `yaml:"title"`
	Subtitle           string `yaml:"subtitle"`
	NameLabel          string `yaml:"name_label"`
	NamePlaceholder    string `yaml:"name_placeholder"`
	SubjectLabel       string `yaml:"subject_label"`
	SubjectPlaceholder string `yaml:"subject_placeholder"`
	MessageLabel       string `yaml:"message_label"`
	MessagePlaceholder string `yaml:"message_placeholder"`
	Send               string `yaml:"send"`
	Cancel             string `yaml:"cancel"`
}

type I18n struct {
	Nav          I18nNav          `yaml:"nav"`
	Status       I18nStatus       `yaml:"status"`
	Actions      I18nActions      `yaml:"actions"`
	Hero         I18nHero         `yaml:"hero"`
	Summary      I18nSummary      `yaml:"summary"`
	Experience   I18nExperience   `yaml:"experience"`
	Skills       I18nSkills       `yaml:"skills"`
	Links        I18nLinks        `yaml:"links"`
	Footer       I18nFooter       `yaml:"footer"`
	ContactModal I18nContactModal `yaml:"contact_modal"`
}
