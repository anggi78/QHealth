package nb

import (
	"log"
	"strings"
)

func PredictNaiveBayes(probabilities map[string]map[string]float64, diagnosis string) string {
    normalizedDiagnosis := strings.ToLower(strings.TrimSpace(diagnosis))
    maxPriority := "rendah" 
    maxProbability := 0.0

    if probs, exists := probabilities[normalizedDiagnosis]; exists {
        for priority, prob := range probs {
            if prob > maxProbability {
                maxProbability = prob
                maxPriority = priority
            }
        }
    }

    if maxProbability == 0 {
        log.Printf("Diagnosis '%s' not found, defaulting to 'rendah'", diagnosis)
    }

    return maxPriority
}