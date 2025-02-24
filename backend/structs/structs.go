package structs

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type ArtistsResponse struct {
	ArtistsList []ArtistResponse `json:"artists_list"`
}

type ArtistResponse struct {
	ArtistID      int    `json:"artist_id"`
	ArtistName    string `json:"artist_name"`
	Country       string `json:"country"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
	Popularity    int    `json:"popularity"`
	ArtistPicture string `json:"artist_picture"`
}

type RegisterRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	FullName     string `json:"full_name"`
	BirthDate    string `json:"birth_date"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string
	Password string
	Role     string
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type GetArtistRequest struct {
	ArtistName string `json:"artist_name"`
}

type GetArtistResponse struct {
	Artist ArtistResponse    `json:"artist"`
	Songs  []SongsForArtists `json:"songs"`
}

type SongsForArtists struct {
	SongID         int    `json:"song_id"`
	SongTitle      string `json:"song_title"`
	SongPopularity int    `json:"song_popularity"`
	Duration       int    `json:"duration"`
	Picture        string `json:"picture"`
	SongURL        string `json:"song_url"`
}

type ListenSongRequest struct {
	SongID int `json:"song_id"`
}

type TrackPlay struct {
	SongID string `json:"songId"`
}

type User struct {
	UserID            int    `json:"user_id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	FullName          string `json:"full_name"`
	BirthDate         string `json:"birth_date"`
	ProfilePictureURL string `json:"profile_pic_url"`
	Role              string `json:"role"`
}

type UpdateProfileRequest struct {
	FullName       string `json:"full_name"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_pic_url"`
}

type AddArtistRequest struct {
	ArtistName    string `json:"artist_name"`
	Country       string `json:"country"`
	Genre         string `json:"genre"`
	Description   string `json:"description"`
	ArtistPicture string `json:"artist_picture"`
}

type AddSongRequest struct {
	ArtistID     int    `json:"artist_id"`
	SongName     string `json:"song_name"`
	SongGenre    string `json:"genre"`
	SongDuration int    `json:"duration"`
	SongDate     string `json:"date"`
	SongLanguage string `json:"language"`
	SongFile     string `json:"file"`
}

type ArtistForReport struct {
	ArtistName string
	Genre      string
	Popularity int
}

type Base64Excel struct {
	Base64 string `json:"base64"`
}
