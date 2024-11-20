package rest

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"online_song_library/internal/models"
)

func (s *SongService) GetInfo(group, song string) (*http.Response, error) {
	logrus.Debug("got request: GetInfo")
	baseUrl, err := url.Parse("http://0.0.0.0:8080/api/v1/info")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)
	baseUrl.RawQuery = params.Encode()

	return http.Get(baseUrl.String())
}

// @Summary		Get additional information about a song from an external API.
// @Description	Retrieves additional information about a song from an external API based on the group and song title.
// @ID				get-info
// @Produce		json
// @Param			group	query		string				true	"The name of the group."
// @Param			song	query		string				true	"The title of the song."
// @Success		200		{object}	models.SongDetail	"The detailed information about the song."
// @Failure		400		{object}	string				"Invalid request or bad request to API."
// @Failure		500		{object}	string				"Internal server error."
// @Router			/api/v1/get-info [get]
func (s *SongService) GetSongInfo(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	logrus.Debugf("got group: %v, song: %v", group, song)

	if group == "" || song == "" {
		http.Error(w, "Group and song are required", http.StatusBadRequest)
		return
	}

	songDetail := models.SongDetail{
		ReleaseDate: "16.07.2006",
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songDetail)
}
