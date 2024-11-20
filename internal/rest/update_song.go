package rest

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"online_song_library/internal/models"
	"online_song_library/internal/utils"
	"strconv"
)

// @Summary		Update an existing song by its ID.
// @Description	Updates an existing song with the provided details.
// @ID				update-song
// @Accept			json
// @Produce		json
// @Param			id		path		integer					true	"The ID of the song to be updated"
// @Param			Song	body		models.Song				true	"The updated song details"
// @Success		200		{object}	map[string]interface{}	"The song was updated successfully"
// @Failure		400		{object}	string					"Invalid request or bad request to API"
// @Failure		404		{object}	string					"Song not found"
// @Failure		500		{object}	string					"Internal server error"
// @Router			/api/v1/update/{id} [patch]
func (s *SongService) UpdateSong(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("got request: UpdateSong")
	idStr := r.URL.Path[len("/api/v1/update/"):]
	logrus.Debugf("got id: %v", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		logrus.WithField("id", id).Error("Invalid ID")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var Song models.Song
	err = json.NewDecoder(r.Body).Decode(&Song)
	if err != nil {
		logrus.WithField("id", id).Error("Failed to update Song")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Update(&Song)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, utils.Message(true, "Song updated"))
	logrus.WithField("id", id).Info("Song updated")
}
