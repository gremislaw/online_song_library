package rest

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"gorm.io/gorm"
	"net/http"
)

func handleRoutes(db *gorm.DB, mux *http.ServeMux) {
	SongSvc := NewSongService(db).(*SongService)

	mux.HandleFunc("/api/v1/create", SongSvc.CreateSong)
	mux.HandleFunc("/api/v1/delete/", SongSvc.DeleteSong)
	mux.HandleFunc("/api/v1/update/", SongSvc.UpdateSong)
	mux.HandleFunc("/api/v1/get", SongSvc.GetSongs)
	mux.HandleFunc("/api/v1/get/", SongSvc.GetSongText)
	mux.HandleFunc("/api/v1/info", SongSvc.GetSongInfo)
}

//go:embed docs/swagger.json
var spec []byte

func handleSwagger(mux *http.ServeMux) {
	mux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))
}

func CreateRestService(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	handleSwagger(mux)
	handleRoutes(db, mux)

	return mux
}
