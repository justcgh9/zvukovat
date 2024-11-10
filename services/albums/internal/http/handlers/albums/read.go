package albums

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpError "github.com/justcgh9/zvukovat/services/albums/internal/http"
	"github.com/justcgh9/zvukovat/services/albums/internal/models"
)

type AlbumReader interface {
    ReadAlbum(album_id string) (models.Album, error)
}

func NewAlbumReader(log *slog.Logger, albumReader AlbumReader) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handlers.albums.read.NewAlbumReader"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        album_id := r.URL.Query().Get("album_id")
        if album_id == "" {
            log.Error("the album_id is empty")
            render.Status(r, 400)
            render.JSON(w, r, httpError.NewHttpError("the album_id is empty"))
            return
        }

        album, err := albumReader.ReadAlbum(album_id)
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 400)
            render.JSON(w, r, httpError.NewHttpError(err.Error()))
            return
        }

        render.Status(r, 200)
        render.JSON(w, r, album)

    }
}
