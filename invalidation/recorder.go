package invalidation
import "net/http"
type Recorder struct {
	writer http.ResponseWriter
	Status int
	Body []byte
	Headers map[string]string
}
func NewRecorder(w http.ResponseWriter) *Recorder {
	return &Recorder{
		writer: w,
		Status: 200,
		Headers: make(map[string]string),
	}
}
func (r *Recorder) Header() http.Header {
	return r.writer.Header()
}
func (r *Recorder) WriteHeader(status int) {
	r.Status = status
	r.writer.WriteHeader(status)
}
func (r *Recorder) Write(b []byte) (int, error) {
	r.Body = append(r.Body, b...)
	return r.writer.Write(b)
}