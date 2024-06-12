package server

import (
	
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"nhooyr.io/websocket"
)
type user struct{
	Password string `json:"password"`
	Email string `json:"email"`
}
func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.AllowAll().Handler)
	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Get("/websocket", s.websocketHandler)

	r.Post("/create",s.CreateUserHandler)
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())

	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns:[]string{"*"},
	})

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))

		if err != nil {
			break
		}
		time.Sleep(time.Second * 1)
	}
}
func (s *Server) CreateUserHandler(w http.ResponseWriter,r *http.Request){
	var user user
	json.NewDecoder(r.Body).Decode(&user)
	
	
	if  user.Email=="" || user.Password=="" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	b,err:=json.Marshal(user)
	if err!=nil{
		log.Print(err)
	}
	w.Write(b)
}