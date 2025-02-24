async function likeSong(songId) {
    let token = localStorage.getItem('jwt');
    try {
        const response = await fetch('http://localhost:8080/like', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify({ songId })
        });

        if (response.ok) {
            alert("Песня лайкнута!");
        } else {

            switch(response.status){
                case 401: 
                console.error('Ошибка при лайковании песни:', response.status);
                alert("Авторизируйтесь!");
                break;
                case 403: 
                console.error('Ошибка при лайковании песни:', response.status);
                alert("Купи подписку!");
                break;
                case 500: 
                console.error('Ошибка при лайковании песни:', response.status);
                alert("Вы уже добавили эту песню в свои любимые!!");
                break;
    }
            
        }
    } catch (error) {
        console.error('Error liking song:', error);
        alert("Ошибка при лайковании песни. Попробуйте позже.");
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

async function playAudio(url, progressBar, button, songId) {
    // Если выбрана та же песня, переключаем состояние воспроизведения
    if (audioPlayer && audioPlayer.src === url) {
        togglePlayPause();
        return;
    }

    let isTracked = await trackPlayOnServer(songId);
    if (!isTracked) {
        alert("КУПИ ПОДПИСКУ!!!");
        return; // Завершаем выполнение функции, если возникла ошибка
    }

    // Останавливаем текущую песню и сбрасываем кнопку плей
    if (audioPlayer && isPlaying) {
        audioPlayer.pause();
        if (currentButton) {
            currentButton.textContent = '▶ Play';
        }
    }

    // Создаем или переиспользуем элемент аудио
    if (!audioPlayer) {
        audioPlayer = document.createElement('audio');
        audioPlayer.id = 'audio-player';
        audioPlayer.addEventListener('timeupdate', updateProgress);
        audioPlayer.addEventListener('ended', onAudioEnd);
        document.body.appendChild(audioPlayer);
    }

    // Начинаем воспроизведение новой песни
    audioPlayer.src = url;
    currentSongUrl = url;
    currentProgressBar = progressBar;
    audioPlayer.play();
    isPlaying = true;
    currentButton = button;  // Сохраняем текущую кнопку
    button.textContent = '⏸ Pause';  // Меняем текст кнопки на Pause
    updatePlayPauseButton();
}

function togglePlayPause() {
    if (isPlaying) {
        audioPlayer.pause();
    } else {
        audioPlayer.play();
    }
    isPlaying = !isPlaying;
    updatePlayPauseButton();
}

function updatePlayPauseButton() {
    const playPauseButtons = document.querySelectorAll('.play-button');
    playPauseButtons.forEach(button => {
        if (audioPlayer.src && button.dataset.audio === audioPlayer.src) {
            button.textContent = isPlaying ? '⏸ Pause' : '▶ Play';
        }
    });
}

function updateProgress() {
    if (currentProgressBar && audioPlayer.duration > 0) {
        const progress = (audioPlayer.currentTime / audioPlayer.duration) * 100;
        currentProgressBar.style.width = `${progress}%`;
    }
}

function onAudioEnd() {
    const progressBar = document.getElementById('progress-bar');
    progressBar.style.width = '0%';
    isPlaying = false;
    updatePlayPauseButton();
    if (currentButton) {
        currentButton.textContent = '▶ Play';
    }
}