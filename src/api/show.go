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

type ShowHandler struct {
	Shows  repository.IShowR
	Person repository.IPersonR
}

func NewShowHandler(r repository.IShowR, p repository.IPersonR) *ShowHandler {
	return &ShowHandler{
		Shows:  r,
		Person: p,
	}
}

func (h *ShowHandler) Register(r *gin.RouterGroup) {
	r.POST("/create", h.Create)
	r.POST("/update", h.Update)
	r.POST("/find", h.Find)
	r.POST("/get", h.Get)
	r.POST("/delete", h.Delete)
}

// todo add actors and others
func (h *ShowHandler) Create(c *gin.Context) {
	args := &models.ExtShow{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	show, err := h.Shows.Create(c, &args.Show)
	if err != nil {
		logrus.Error("error on creating", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}

	for _, pID := range args.Persons {
		err := h.Shows.AddPerson(c, pID.PersonID, show.ShowID)
		if err != nil {
			logrus.Error("error on adding person", err)
			apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, show)
}

func (h *ShowHandler) Update(c *gin.Context) {
	args := &models.ExtShow{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	show, err := h.Shows.Update(c, &args.Show)
	if err != nil {
		logrus.Error("error on updating", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.Shows.DeletePersonsFromShow(c, args.ShowID)
	if err != nil {
		logrus.Error("error on deleting persons", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	logrus.Debug(show)
	for _, pID := range args.Persons {
		err := h.Shows.AddPerson(c, pID.PersonID, show.ShowID)
		if err != nil {
			logrus.Error("error on adding person", err)
			apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, show)
}

type FindShowArgs struct {
	Query search.Show `json:"query"`
	Sort  search.Sort `json:"sort"`
}

func (a *FindShowArgs) Validate() {
	a.Sort.Validate()
}

func (h *ShowHandler) Find(c *gin.Context) {
	args := &FindShowArgs{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	args.Validate()
	shows, err := h.Shows.Find(c, args.Query, args.Sort)
	if err != nil {
		logrus.Error("error on finding", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, shows)
}

type GetShowArg struct {
	ID int `json:"id"`
}

type GetShowReply struct {
	models.ExtShow
	PersonList []models.Person `json:"personList"`
}

func (h *ShowHandler) Get(c *gin.Context) {
	args := &GetShowArg{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	extShow := models.ExtShow{}
	show, err := h.Shows.Get(c, args.ID)
	if err != nil {
		logrus.Error("error on getting", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	extShow.Show = *show
	persons, err := h.Person.GetByShowID(c, args.ID)
	if err != nil {
		logrus.Error("error on getting", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	extShow.Persons = persons
	c.JSON(http.StatusOK, extShow)
}

type DeleteShowArgs struct {
	ID int `json:"id"`
}

func (h *ShowHandler) Delete(c *gin.Context) {
	args := &GetShowArg{}
	if err := c.BindJSON(args); err != nil {
		logrus.Warn("error on binding json", err)
		apierr.Response(c, http.StatusBadRequest, apierr.ErrParseFailed)
		return
	}
	err := h.Shows.Delete(c, args.ID)
	if err != nil {
		logrus.Error("error on deleting", err)
		apierr.ResponseMsg(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
