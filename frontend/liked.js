let likedSongs = [];
let audioPlayer;
let isPlaying = false;
let currentSongUrl = '';
let currentProgressBar = null;
let currentButton = null;

// Функция для загрузки лайкнутых песен
async function loadLikedSongs() {
    try {
        const response = await fetch('http://localhost:8080/get_liked', {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('jwt')}`
            }
        });
        likedSongs = await response.json();
        displayLikedSongs();
    } catch (error) {
        console.error('Error loading liked songs:', error);
    }
}
async function trackPlayOnServer(songId) {
    let token = localStorage.getItem('jwt');
    try {
        const response = await fetch('http://localhost:8080/track_play', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify({ songId })
        });

        if (!response.ok) {
            switch (response.status) {
                case 400:
                    throw new Error('Bad Request: Проверьте отправленные данные.');
                case 401:
                    throw new Error('Unauthorized: Неверный или отсутствующий токен.');
                case 403:
                    throw new Error('Forbidden: У вас нет доступа.');
                case 404:
                    throw new Error('Not Found: Ресурс не найден.');
                case 500:
                    throw new Error('Internal Server Error: Проблема на сервере.');
                default:
                    throw new Error(`Unexpected error: ${response.statusText}`);
            }
        }

        return true; // Успешное отслеживание
    } catch (error) {
        console.error('Error tracking play:', error);
        return false; // Ошибка отслеживания
    }
}

// Отображение лайкнутых песен
function displayLikedSongs() {
    const songsContainer = document.getElementById('songs-container');
    songsContainer.innerHTML = ''; // Очистка контейнера

    likedSongs.forEach((song, index) => {
        const songItem = document.createElement('div');
        songItem.classList.add('song-item');
        songItem.innerHTML = `
            <img src="${song.picture}" alt="${song.song_title}">
            <div class="song-details">
                <h3>${song.song_title}</h3>
                <p>Duration: ${song.duration}s</p>
                <p>Popularity: ${song.song_popularity}</p>
                <div class="progress-container">
                    <div class="progress-bar" data-index="${index}"></div>
                </div>
            </div>
            <div class="song-actions">
                <button class="play-button" data-id="${song.song_id}" data-audio="${song.song_url}">▶ Play</button>
                <button class="unlike-button" data-id="${song.song_id}">Unlike</button>
            </div>
        `;
        songsContainer.appendChild(songItem);
    });

    // Добавляем обработчики событий
    const playButtons = document.querySelectorAll('.play-button');
    playButtons.forEach(button => {
        button.addEventListener('click', event => {
            const audioUrl = event.target.dataset.audio;
            const progressBar = event.target.closest('.song-item').querySelector('.progress-bar');
            playAudio(audioUrl, progressBar, event.target);
        });
    });

    const unlikeButtons = document.querySelectorAll('.unlike-button');
    unlikeButtons.forEach(button => {
        button.addEventListener('click', event => {
            const songId = event.target.dataset.id;
            unlikeSong(songId);
        });
    });
}

// Удаление песни из лайкнутых
async function unlikeSong(songId) {
    data = {songId: songId}
    console.log(data)
    try {
        const response = await fetch(`http://localhost:8080/unlike`, {
            method: 'DELETE',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('jwt')}`
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            likedSongs = likedSongs.filter(song => song.song_id !== parseInt(songId));
            displayLikedSongs();
        } else {
            console.error('Failed to unlike song');
        }
    } catch (error) {
        console.error('Error unliking song:', error);
    }
}

// Воспроизведение песни
async function playAudio(url, progressBar, button) {
   
    // Если это та же песня, переключаем воспроизведение
    if (audioPlayer && audioPlayer.src === url) {
        togglePlayPause();
        return;
    }

    // Останавливаем текущую песню
    if (audioPlayer && isPlaying) {
        audioPlayer.pause();
        if (currentButton) {
            currentButton.textContent = '▶ Play';
        }
    }

    // Создаём аудиоплеер, если его ещё нет
    if (!audioPlayer) {
        audioPlayer = new Audio();
        audioPlayer.addEventListener('timeupdate', updateProgress);
        audioPlayer.addEventListener('ended', onAudioEnd);
    }

    // Воспроизведение новой песни
    var songId = button.dataset.id;
    let isTracked = await trackPlayOnServer(songId);
    if (!isTracked) {
        alert("КУПИ ПОДПИСКУ!!!");
        return; // Завершаем выполнение функции, если возникла ошибка
    }
    audioPlayer.src = url;
    audioPlayer.play();
    isPlaying = true;
    currentSongUrl = url;
    currentProgressBar = progressBar;
    currentButton = button;
    button.textContent = '⏸ Pause';
}

function togglePlayPause() {
    if (isPlaying) {
        audioPlayer.pause();
    } else {
        audioPlayer.play();
    }
    isPlaying = !isPlaying;
    if (currentButton) {
        currentButton.textContent = isPlaying ? '⏸ Pause' : '▶ Play';
    }
}

function updateProgress() {
    if (currentProgressBar && audioPlayer.duration > 0) {
        const progress = (audioPlayer.currentTime / audioPlayer.duration) * 100;
        currentProgressBar.style.width = `${progress}%`;
    }
}

function onAudioEnd() {
    if (currentProgressBar) {
        currentProgressBar.style.width = '0%';
    }
    isPlaying = false;
    if (currentButton) {
        currentButton.textContent = '▶ Play';
    }
}

// Загружаем лайкнутые песни при загрузке страницы
window.onload = loadLikedSongs;
