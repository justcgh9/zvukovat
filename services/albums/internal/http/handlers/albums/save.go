package albums

import (
	"log/slog"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpError "github.com/justcgh9/zvukovat/services/albums/internal/http"
	"github.com/justcgh9/zvukovat/services/albums/internal/models"
)

type AlbumSaver interface {
    SaveAlbum(album models.Album) (models.Album, error)
}

type TrackAdder interface {
    AddTrackToAlbum(albumId, trackId string) error
}

type FileSaver interface {
    SaveFile(file multipart.File, header *multipart.FileHeader, uploadDir string, fileType string) (string, error)
}

func NewAlbumSaver(log *slog.Logger, albumSaver AlbumSaver, fileSaver FileSaver) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handlers.albums.save.NewAlbumSaver"

        log := log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )


        err := r.ParseMultipartForm(10 << 20)
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 400)
            render.JSON(w, r, httpError.NewHttpError(err.Error()))
            return
        }
        name := r.FormValue("name")
        artist := r.FormValue("artist")

        var picture string
        pictureFile, pictureHeader, err := r.FormFile("picture")
        if err != nil {
            log.Warn("the picture is empty")
            picture = ""
        } else {
            picture, err = fileSaver.SaveFile(pictureFile, pictureHeader, "files/", "picture")
            if err != nil {
                log.Error(err.Error())
                render.Status(r, 400)
                render.JSON(w, r, httpError.NewHttpError(err.Error()))
                return
            }
            log.Info("saved picture " + picture + " for album " + name + " and artist " + artist)
        }

        album := models.NewAlbum(name, artist, picture, make([]string, 0))
        log.Info("saved album", album )
        render.SetContentType(render.ContentTypeJSON)
        render.Status(r, 201)
        render.JSON(w, r, album)

    }
}

func NewTrackAdder(log *slog.Logger, trackAdder TrackAdder) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const op = "handlers.albums.save.NewAlbumSaver"

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
        track_id := r.URL.Query().Get("track_id")
        if track_id == "" {
            log.Error("the track_id is empty")
            render.Status(r, 400)
            render.JSON(w, r, httpError.NewHttpError("the track_id is empty"))
            return
        }

        err := trackAdder.AddTrackToAlbum(album_id, track_id)
        if err != nil {
            log.Error(err.Error())
            render.Status(r, 400)
            render.JSON(w, r, httpError.NewHttpError(err.Error()))
            return
        }

        log.Info("added track " + track_id + " to the album " + album_id)
        render.Status(r, 200)
    }
}
