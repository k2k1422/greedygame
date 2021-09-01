package response

var CodeMapping map[string]string

func Initialize() {
	CodeMapping = make(map[string]string)
	// Meta response codes

	CodeMapping["110"] = "Failed to decode"
	CodeMapping["111"] = "successful to register"

	// log.Println("[INFO] Initialized response mappings.")
}
