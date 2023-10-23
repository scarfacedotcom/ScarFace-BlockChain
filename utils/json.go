// package utils

// import "encoding/json"

//	func JsonStatus(message string) []byte {
//		m, _ := json.Marshal(struct {
//			Message string `json: "message"`
//		}{
//			Message: message,
//		})
//		return m
//	}
package utils

import "encoding/json"

func JsonResponse(status string, message string) []byte {
	response, _ := json.Marshal(struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: message,
	})
	return response
}
