package jobs

import "net/http"

func createJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating job...."))
}

func listJobs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listing jobs...."))
}
func getJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting job...."))
}
func deleteJob(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting job...."))
}
