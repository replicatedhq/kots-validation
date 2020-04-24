package handlers

import (
	"net/http"
)

const redactResponse = `
apiVersion: troubleshoot.replicated.com/v1beta1
kind: Redactor
metadata:
  name: my-demo-web-redactor
spec:
  redactors:
  - name: replace the second literal string
    values:
    - redact-me-second
`

// Redact returns a hardcoded redaction spec
func Redact(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte(redactResponse))
}
