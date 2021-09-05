package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"shows/src/api/apierr"
	"shows/src/repository"
)

type GenreHandler struct {
	Genres repository.IGenreR
	Shows  repository.IShowR
}

func NewGenreHandler(p repository.IGenreR, s repository.IShowR) *GenreHandler {
	return &GenreHandler{
		Genres: p,
		Shows:  s,
	}
}

func (h *GenreHandler) Register(r *gin.RouterGroup) {
	r.POST("/create", h.Create)
	r.POST("/get", h.Get)
	r.POST("/getAll", h.GetAll)
}

type CreateGenreArgs struct {
	Name string `json:"name"`
}

func (h *GenreHandler) Create(c *gin.Context) {
	args := &CreateGenreArgs{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	genre, err := h.Genres.Create(c, args.Name)
	if err != nil {
		logrus.Error("error on creating genre", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, genre)
}

func (h *GenreHandler) GetAll(c *gin.Context) {
	genres, err := h.Genres.GetAll(c)
	if err != nil {
		logrus.Error("error on getting all genres", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, genres)
}

type GetGenreArgs struct {
	ID int `json:"id"`
}

func (h *GenreHandler) Get(c *gin.Context) {
	args := &GetGenreArgs{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	shows, err := h.Shows.GetByGenre(c, args.ID)
	if err != nil {
		logrus.Error("error on creating genre", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, shows)
}
