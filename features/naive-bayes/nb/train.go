package nb

import (
	"strings"
)

func TrainNaiveBayes(data []Dataset) map[string]map[string]float64 {
    classCounts := make(map[string]int)
    diagnosisCounts := make(map[string]map[string]int)

    for _, record := range data {
        normalizedDiagnosis := strings.ToLower(strings.TrimSpace(record.Diagnosis))
        if record.Priority != "" {
            classCounts[record.Priority]++
        } else {
            classCounts["unknown"]++
        }

        if _, exists := diagnosisCounts[normalizedDiagnosis]; !exists {
            diagnosisCounts[normalizedDiagnosis] = make(map[string]int)
        }
        diagnosisCounts[normalizedDiagnosis][record.Priority]++
    }

    probabilities := make(map[string]map[string]float64)
    for diagnosis, priorities := range diagnosisCounts {
        probabilities[diagnosis] = make(map[string]float64)
        for priority, count := range priorities {
            if classCounts[priority] > 0 {
                probabilities[diagnosis][priority] = float64(count) / float64(classCounts[priority])
            }
        }
    }

    return probabilities
}
