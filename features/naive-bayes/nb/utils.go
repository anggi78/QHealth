package nb

import "strings"

var severeDiagnoses = map[string]bool{
	"heart failure":                 true,
	"atherosclerotic heart disease": true,
	"hypertensi":                    true,
	"ectopic pregnancy":             true,
	"hipertensi gestasional":        true,
	"prolonged pregnancy":           true,
	"post sc":                       true,
}

var moderateDiagnoses = map[string]bool{
	"lowbackpain":              true,
	"schizophrenia":            true,
	"dm tipe ii":               true,
	"dyspepsia":                true,
	"contraceptive management": true,
	"senile cataract":          true,
	"anxiety disorder":         true,
	"pulpitis":                 true,
	"asthma":                   true,
	"dermatitis":               true,
	"myalgia":                  true,
}

func IsSevereDiagnosis(diagnosis string) bool {
	normalizedDiagnosis := strings.ToLower(strings.TrimSpace(diagnosis))
	return severeDiagnoses[normalizedDiagnosis]
}

func IsModerateDiagnosis(diagnosis string) bool {
	normalizedDiagnosis := strings.ToLower(strings.TrimSpace(diagnosis))
	return moderateDiagnoses[normalizedDiagnosis]
}
