package rest

import (
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"online_song_library/internal/utils"
	"strconv"
	"strings"
)

// @Summary		Get the text of a song by its ID.
// @Description	Retrieves the text of a song by its ID with optional pagination.
// @ID				get-song-text
// @Produce		json
// @Param			id		path		integer					true	"The ID of the song"
// @Param			limit	query		integer					false	"Limit the number of lines"	default(10)
// @Param			offset	query		integer					false	"Offset for pagination"		default(0)
// @Success		200		{object}	map[string]interface{}	"The text of the song"
// @Failure		400		{object}	string					"Invalid ID"
// @Failure		404		{object}	string					"Song not found"
// @Failure		500		{object}	string					"Internal server error"
// @Router			/api/v1/get/{id} [get]
func (s *SongService) GetSongText(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("got request: GetSongText")
	idStr := r.URL.Path[len("/api/v1/get/"):]
	logrus.Debugf("got id: %v", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		logrus.WithField("id", id).Error("Invalid ID")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	song, err := s.GetTextPaginated(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	text := song.Text
	lines := strings.Split(text, "\n")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if limit > 0 && offset >= 0 && offset+limit < len(lines) {
		lines = lines[offset : offset+limit]
	} else {
		logrus.Warnf("limit or offset invalid: %v, %v", limit, offset)
		lines = lines[0:0]
	}

	w.WriteHeader(http.StatusOK)
	utils.Respond(w, map[string]interface{}{"text": strings.Join(lines, "\n")})
	logrus.WithField("id", id).Info("Song text got")
}
