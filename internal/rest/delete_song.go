package rest

import (
	"net/http"
	"online_song_library/internal/utils"
	"strconv"

	"github.com/sirupsen/logrus"
)

// @Summary		Delete a song by its ID.
// @Description	Deletes a song by its ID.
// @ID				delete-song
// @Produce		json
// @Param			id	path		integer					true	"The ID of the song to be deleted"
// @Success		200	{object}	map[string]interface{}	"The song was deleted successfully"
// @Failure		400	{object}	string					"Invalid ID"
// @Failure		404	{object}	string					"Song not found"
// @Failure		500	{object}	string					"Internal server error"
// @Router			/api/v1/delete/{id} [delete]
func (s *SongService) DeleteSong(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("got request: DeleteSonf")
	idStr := r.URL.Path[len("/api/v1/delete/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		logrus.WithField("id", id).Error("Invalid ID")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Delete(id)

	if err != nil {
		logrus.WithField("id", id).Error("Failed to delete Song")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.Respond(w, utils.Message(true, "Song deleted"))
	logrus.WithField("id", id).Info("Song deleted")
}
