package rest

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"online_song_library/internal/models"
	"online_song_library/internal/utils"
)

// @Summary		Create a new song.
// @Description	Creates a new song with the provided details.
// @ID				create-song
// @Accept			json
// @Produce		json
// @Param			Song	body		models.Song				true	"The song details to be created"
// @Success		200		{object}	map[string]interface{}	"The song was created successfully"
// @Failure		400		{object}	string					"Invalid request or bad request to API"
// @Failure		500		{object}	string					"Internal server error"
// @Router			/api/v1/create [post]
func (s *SongService) CreateSong(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("got request: CreateSong")
	var Song models.Song
	err := json.NewDecoder(r.Body).Decode(&Song)
	if err != nil {
		logrus.Error("Failed to create Song")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Запрос в АПИ для получения дополнительной информации
	group := Song.Group
	song := Song.Title
	resp, err := s.GetInfo(group, song)
	if err != nil {
		logrus.Error("Bad request to API")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	var songDetail models.SongDetail
	err = json.NewDecoder(resp.Body).Decode(&songDetail)
	if err != nil {
		logrus.Error("Bad response from API")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Song.Text = songDetail.Text
	Song.Link = songDetail.Link
	Song.ReleaseDate = songDetail.ReleaseDate

	err = s.Create(&Song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.Respond(w, utils.Message(true, "Song created"))
	logrus.Info("Song created")
}
