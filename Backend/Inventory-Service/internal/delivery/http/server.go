package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory-service/internal/usecase"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	mangaUC usecase.MangaUsecase
}

func (s *Server) Run(port string) any {
	panic("unimplemented")
}

func NewServer(mangaUC usecase.MangaUsecase) *Server {
	s := &Server{
		router:  mux.NewRouter(),
		mangaUC: mangaUC,
	}
	s.configureRouter()
	return s
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/manga", s.handleCreateManga).Methods("POST")
	s.router.HandleFunc("/manga/{id}", s.handleGetManga).Methods("GET")
	s.router.HandleFunc("/manga/{id}", s.handleUpdateManga).Methods("PUT")
	s.router.HandleFunc("/manga/{id}", s.handleDeleteManga).Methods("DELETE")
	s.router.HandleFunc("/manga", s.handleListManga).Methods("GET")
}

// Create Manga
func (s *Server) handleCreateManga(w http.ResponseWriter, r *http.Request) {
	var req createMangaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	manga, err := s.mangaUC.CreateManga(r.Context(), req.toDomain())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, manga)
}

// Get Manga by ID
func (s *Server) handleGetManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	manga, err := s.mangaUC.GetMangaByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Manga not found")
		return
	}

	respondWithJSON(w, http.StatusOK, manga)
}

// Update Manga
func (s *Server) handleUpdateManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req updateMangaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := s.mangaUC.UpdateManga(r.Context(), id, req.toDomain())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Manga updated successfully"})
}

// Delete Manga
func (s *Server) handleDeleteManga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := s.mangaUC.DeleteManga(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Manga deleted successfully"})
}

// List Manga with pagination
func (s *Server) handleListManga(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	manga, err := s.mangaUC.ListManga(r.Context(), page, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, manga)
}

// Helper functions
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Request/Response DTOs
type createMangaRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
}

func (r *createMangaRequest) toDomain() usecase.Manga {
	return usecase.Manga{
		Title:       r.Title,
		Author:      r.Author,
		Description: r.Description,
		Genres:      r.Genres,
	}
}

type updateMangaRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
}

func (r *updateMangaRequest) toDomain() usecase.Manga {
	return usecase.Manga{
		Title:       r.Title,
		Author:      r.Author,
		Description: r.Description,
		Genres:      r.Genres,
	}
}
