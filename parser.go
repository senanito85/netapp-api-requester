package main

import "encoding/json"

type (
	space struct {
		Size      float64 `json:"size"`
		Available float64 `json:"available"`
	}

	metric struct {
		Timestamp string `json:"timestamp"`
	}

	svm struct {
		Name string `json:"name"`
	}

	volume struct {
		Name   string `json:"name"`
		Svm    svm    `json:"svm"`
		Space  space  `json:"space"`
		Metric metric `json:"metric"`
	}

	response struct {
		Records []volume `json:"records"`
	}
)

func parseJSON(j []byte) (response, error) {
	r := &response{}
	err := json.Unmarshal(j, r)
	if err != nil {
		return response{}, err
	}

	return *r, nil
}
