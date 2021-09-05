package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"shows/src/api/apierr"
	"shows/src/models"
	"shows/src/models/search"
	"shows/src/repository"
)

type PersonHandler struct {
	Persons repository.IPersonR
	Shows   repository.IShowR
}

func NewPersonHandler(p repository.IPersonR, s repository.IShowR) *PersonHandler {
	return &PersonHandler{
		Persons: p,
		Shows:   s,
	}
}

func (h *PersonHandler) Register(r *gin.RouterGroup) {
	r.POST("/create", h.Create)
	r.POST("/update", h.Update)
	r.POST("/find", h.Find)
	r.POST("/get", h.Get)
	r.POST("/getAll", h.GetAll)
}

// todo add actors and others
func (h *PersonHandler) Create(c *gin.Context) {
	args := &models.ExtPerson{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	person, err := h.Persons.Create(c, &args.Person)
	if err != nil {
		logrus.Error("error on creating", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, person)
}

func (h *PersonHandler) Update(c *gin.Context) {
	args := &models.Person{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	show, err := h.Persons.Update(c, args)
	if err != nil {
		logrus.Error("error on updating", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, show)
}

type FindPersonArgs struct {
	Query search.Person `json:"query"`
	Sort  search.Sort   `json:"sort"`
}

func (a *FindPersonArgs) Validate() {
	a.Sort.Validate()
}

func (h *PersonHandler) Find(c *gin.Context) {
	args := &FindPersonArgs{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	args.Validate()
	persons, err := h.Persons.Find(c, "", args.Sort)
	if err != nil {
		logrus.Error("error on finding", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, persons)
}

type GetPersonArgs struct {
	ID int `json:"id"`
}

func (h *PersonHandler) Get(c *gin.Context) {
	args := &GetPersonArgs{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	extPerson := models.ExtPerson{}
	person, err := h.Persons.Get(c, args.ID)
	if err != nil {
		logrus.Error("error on getting", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	extPerson.Person = *person
	showRoles, err := h.Shows.ShowsByPersonID(c, args.ID)
	if err != nil {
		logrus.Error("error on getting", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	extPerson.Shows = showRoles
	c.JSON(http.StatusOK, extPerson)
}

func (h *PersonHandler) GetAll(c *gin.Context) {
	persons, err := h.Persons.Find(c, "", search.Sort{})
	if err != nil {
		logrus.Error("error on getting", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, persons)
}
