package mod

type Site struct {
	Siteurl     string   `json:"siteurl"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Lang        string   `json:"lang"`
	Keywords    []string `json:"keywords"`
	Favicon     string   `json:"favicon"`
	Locale      string   `json:"locale"`
	Creator     string   `json:"creator"`
	Company     string   `json:"company"`
	Type        string   `json:"type"`
	Languages   []string `json:"languages"`
	Pages       []string `json:"pages"`

	Fb_app_id string `json:"fb_app_id"`
	Twitter   string `json:"twitter"`
	Contact   string `json:"contact"`
	Footer    string `json:"footer"`
}

type Home struct {
	Title       string `json:"title"`
	SubTitle    string `json:"subtitle"`
	Welcome     string `json:"welcome"`
	About       string `json:"about"`
	GetApp      string `json:"getapp"`
	Features    string `json:"features"`
	Feature1    string `json:"feature1"`
	Feature1txt string `json:"feature1txt"`
	Feature2    string `json:"feature2"`
	Feature2txt string `json:"feature2txt"`
	Feature3    string `json:"feature3"`
	Feature3txt string `json:"feature3txt"`
	Feature4    string `json:"feature4"`
	Feature4txt string `json:"feature4txt"`
	Gallery     string `json:"gallery"`
	Specs       string `json:"specs"`
	Spec1       string `json:"spec1"`
	Spec2       string `json:"spec2"`
	Spec3       string `json:"spec3"`
	Spec4       string `json:"spec4"`
	Moto1       string `json:"moto1"`
	Moto2       string `json:"moto2"`
}

type Story struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
}
type Software struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
}
type RoadMap struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
}

type Contact struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
}

type Resources struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Post struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Excerpt string `json:"excertpt"`
	Content string `json:"content"`
	Image   string `json:"img"`
}
