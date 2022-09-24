package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Date      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}