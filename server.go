package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/tryy3/Recipe-Maker-Backend/graph"
	"github.com/tryy3/Recipe-Maker-Backend/graph/generated"
	"github.com/tryy3/go-cloudinary"

	firestore "cloud.google.com/go/firestore"

	"google.golang.org/api/option"
)

const defaultPort = "8090"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conf, err := readConfig()
	if err != nil {
		log.Fatal("Unable to read config: ", err)
	}

	opt := option.WithCredentialsFile("./configs/firebase_auth.json")
	db, err := firestore.NewClient(context.Background(), "bartender-c26d5", opt)
	if err != nil {
		log.Fatal("error connecting to firestore: ", err)
	}

	// TODO: Change this to not be hardcoded, either through env or config
	cloudinaryService, err := cloudinary.Dial(conf.Cloudinary)
	if err != nil {
		log.Fatal("error connecting to cloudinary: ", err)
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)


    srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Database: db, CloudinaryService: cloudinaryService}}))
    srv.AddTransport(&transport.Websocket{
        Upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				fmt.Println(r.Host)
                 return r.Host == "example.org"
            },
            ReadBufferSize:  1024,
            WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Cloudinary string `json:cloudinary`
}

// TODO: Add more configs and maybe redo this into a more proper config solution
func readConfig() (*Config, error) {
	f, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		return nil, err
	}

	conf := Config{}

	err = json.Unmarshal(f, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}