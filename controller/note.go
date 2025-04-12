package controller

import (
	"go-notes-taker/common"
	"go-notes-taker/dto"
	"go-notes-taker/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NoteController interface {
	CreateNote(ctx *gin.Context)
	GetAllNotes(ctx *gin.Context)
	GetMyNotes(ctx *gin.Context)
	DeleteNote(ctx *gin.Context)
	UpdateNote(ctx *gin.Context)
}

type noteController struct {
	jwtService service.JWTService
	noteService service.NoteService
}

func NewNoteController(ns service.NoteService,  jwts service.JWTService) NoteController {
	return &noteController{
		noteService: ns,
		jwtService: jwts,
	}
}

func(nc *noteController) CreateNote(ctx *gin.Context) {
	var note dto.NoteCreateDTO

	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)

	userID, err := nc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err = ctx.ShouldBind(&note)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Gagal bind note: "+err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := nc.noteService.CreateNote(ctx.Request.Context(), note, userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Membuat Note", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Menambahkan Note", result)
	ctx.JSON(http.StatusOK, res)
}

func(nc *noteController) GetAllNotes(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := nc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := nc.noteService.GetAllNotes(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Daftar Note", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Daftar Note", result)
	ctx.JSON(http.StatusOK, res)
}

func(nc *noteController) GetMyNotes(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := nc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := nc.noteService.GetMyNotes(ctx.Request.Context(), userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan Daftar Note", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan Daftar Note", result)
	ctx.JSON(http.StatusOK, res)
}

func(nc *noteController) DeleteNote(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := nc.jwtService.GetUserIDByToken(token)
	ctx.Set("token", "")
	ctx.Set("userID", "")
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	noteID := ctx.Param("note_id")

	err = nc.noteService.DeleteNote(ctx.Request.Context(), userID, noteID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Menghapus Note", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Menghapus Note", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func(nc *noteController) UpdateNote(ctx *gin.Context) {
	var note dto.NoteUpdateDTO
	err := ctx.ShouldBind(&note)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Note", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	token := ctx.MustGet("token").(string)
	userID, err := nc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	note.ID, err = uuid.Parse(ctx.Param("note_id"))
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Gagal memproses Note ID", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	err = nc.noteService.UpdateNote(ctx.Request.Context(), note, userID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mengupdate Note", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := common.BuildResponse(true, "Berhasil Mengupdate Note", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}