package main

import (
	"log"
	api "music/api"
	"net/http"
)

func main() {
	// Устанавливаем маршрут для пути

	http.HandleFunc("/artist/", api.CORS(api.GetArtist))
	http.HandleFunc("/register", api.CORS(api.Register))
	http.HandleFunc("/auth", api.CORS(api.Auth))
	http.HandleFunc("/artists", api.CORS(api.GetArtists))
	http.HandleFunc("/albums", api.GetAlbums)
	http.HandleFunc("/song", api.ListenSong)
	http.HandleFunc("/api", api.HelloFromServer)
	http.HandleFunc("/track_play", api.CORS(api.TrackPlay))
	http.HandleFunc("/user", api.CORS(api.GetUser))
	http.HandleFunc("/update_profile", api.CORS(api.UpdateProfile))
	http.HandleFunc("/subscribes", api.CORS(api.GetSubscribedArtists))
	http.HandleFunc("/like", api.CORS(api.LikeSong))
	http.HandleFunc("/get_liked", api.CORS(api.LikedSongs))
	http.HandleFunc("/unlike", api.CORS(api.Unlike))
	http.HandleFunc("/check_role", api.CORS(api.CheckRole))
	http.HandleFunc("/add_artist", api.CORS(api.AddArtist))
	http.HandleFunc("/delete_artist/", api.CORS(api.DeleteArtist))
	http.HandleFunc("/delete_song/", api.CORS(api.DeleteSong))
	http.HandleFunc("/edit_artist/", api.CORS(api.EditArtist))
	http.HandleFunc("/add_song/", api.CORS(api.AddSong))
	http.HandleFunc("/report", api.CORS(api.GetReport))
	http.Handle("/images/band/", http.StripPrefix("/images/band/", http.FileServer(http.Dir("D:/res/artist_profile_photo"))))
	http.Handle("/images/song/", http.StripPrefix("/images/song/", http.FileServer(http.Dir("D:/res/song_photo"))))
	http.Handle("/images/user/", http.StripPrefix("/images/user/", http.FileServer(http.Dir("D:/res/user_profile_photo"))))
	http.Handle("/music/", http.StripPrefix("/music/", http.FileServer(http.Dir("D:/res/music"))))
	http.Handle("/rep", http.StripPrefix("/rep", http.FileServer(http.Dir("D:/res"))))
	// Запускаем сервер на порту 8080
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
