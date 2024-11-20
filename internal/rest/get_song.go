package rest

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"online_song_library/internal/utils"
	"strconv"
)

// @Summary		Get a list of songs
// @Description	Retrieves a list of songs with optional filtering and pagination.
// @ID				get-songs
// @Produce		json
// @Param			group	query		string					false	"Filter by group name"
// @Param			song	query		string					false	"Filter by song name"
// @Param			limit	query		integer					false	"Limit the number of results"	default(10)
// @Param			offset	query		integer					false	"Offset for pagination"			default(0)
// @Success		200		{object}	map[string]interface{}	"A map containing the list of songs"
// @Failure		500		{object}	string					"Internal server error"
// @Router			/api/v1/get [get]
func (s *SongService) GetSongs(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("got request: GetSongs")
	songs, err := s.GetPaginated()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Фильтрация по полям
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	logrus.Debugf("got group: %v, song: %v", group, song)
	if group != "" || song != "" {
		s.DB.Where("group LIKE ? OR song LIKE ?", "%"+group+"%", "%"+song+"%").Find(songs)
	}

	// Пагинация
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	if limit > 0 && offset >= 0 {
		s.DB.Limit(limit).Offset(offset).Find(&songs)
	} else {
		logrus.Warnf("limit or offset invalid: %v, %v", limit, offset)
		s.DB.Limit(0).Offset(0).Find(&songs)
	}

	w.WriteHeader(http.StatusOK)
	utils.Respond(w, map[string]interface{}{"data": songs})
	logrus.Info("Songs got")
}
