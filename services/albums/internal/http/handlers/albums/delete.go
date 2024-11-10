package albums

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	httpError "github.com/justcgh9/zvukovat/services/albums/internal/http"
	"github.com/go-chi/render"
)

type AlbumRemover interface {
    RemoveAlbum(album_id string) error
}

func NewAlbumDeleter(log *slog.Logger, albumRemover AlbumRemover) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handlers.albums.delete.NewAlbumDeleter"

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
        err := albumRemover.RemoveAlbum(album_id)
        if err != nil {
            log.Error(err.Error())
            //TODO: Add custom errors and switch on 'em
            render.Status(r, http.StatusNotFound)
            render.JSON(w, r, httpError.NewHttpError(err.Error()))
            return
        }

        render.Status(r, http.StatusOK)
    }
}
