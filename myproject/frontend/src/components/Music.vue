<script setup>
// Код внутри компилируется каждый раз при создании экземпляра компонента
import {ref} from 'vue'
import NextIcon from '../assets/icons/NextIcon.vue'
import PlayIcon from '../assets/icons/PlayIcon.vue'
import PrevIcon from '../assets/icons/PrevIcon.vue'
import ChevronDownIcon from '../assets/icons/ChevronDownIcon.vue'

const songs = ref([])
let currentAudio = null // Храним текущий аудио-объект

function addSong(event) {
  const file = event.target.files[0]
  
  if (!file) return
  console.log('Добавлен файл:', file.name)
  
  const fileUrl = URL.createObjectURL(file)
  
  songs.value.push({
    name: file.name,
    url: fileUrl,
    file: file
  })
  
  event.target.value = ''
}

function playSong(song) {
  // Если уже есть играющий трек, останавливаем его
  if (currentAudio) {
    currentAudio.pause()
    currentAudio = null
  }
  
  currentAudio = new Audio(song.url)
  currentAudio.play()
  console.log('Playing:', song.name)
}

function pauseSong() {
  if (currentAudio) {
    currentAudio.pause()
    console.log('Paused:', currentAudio.src)
  } else {
    console.log('Нет активного трека для паузы')
  }
}
</script>

<template>
    <div class="music-player">
      <div class="music-container">

        <div class="playlist-info">
            <div class="playing-from">PLAYING FROM PLAYLIST:</div>
            <div class="name-chevron">
                <div class="playlist-name">Lofi Loft</div>
                <div class="chevron-down"><ChevronDownIcon/></div>
          </div>
        </div>
        <div class="album-art">
            <img src="../assets/images/album-image.png" alt="Album Cover">
        </div>
        <div class="song-info">
          <div class="song-title">grainy days</div>
          <div class="song-artist">moody</div>
        </div>
        <div class="progress-container">
          <div class="progress-bar">
            <div class="progress-line"></div>
            <div class="progress-indicator"></div>
            <div class="progress-dot"></div>
          </div>
          <div class="time-info">
            <span class="current-time">0:00</span>
            <span class="duration">2:43</span>
          </div>
        </div>
        <div class="player-controls">
          <button class="control-btn prev-btn"><PrevIcon/></button>
          <button @click="pauseSong" class="control-btn play-btn"><PlayIcon/></button>
          <button class="control-btn next-btn"><NextIcon/></button>
        </div>
      </div>

      <div class="music-container">
        <h1>🎵 Список песен</h1>
        <div class="song-list">
            <div
            class="song-item"
            v-for="song in songs"
            :key="song.id"
            @click="playSong(song)"
            >
            {{ song.name }}
            </div>
        </div>

        <!-- Кнопки управления -->
        <div class="controls">
            <button class="upload-button" @click="$refs.uploader.click()">
            Добавить песню
            <input
                ref="uploader"
                type="file"
                accept=".mp3,.mp4,.m4a,audio/*"
                @change="addSong"
                style="display: none"
            />
            </button>
        </div>
      </div>
    </div>
</template>

<style scoped>
.music-player {
  display: flex;
  background: #FFFFFF;  
  width: 100%;
  height: 100%;
} 

.music-container {
  position: relative;
  width: 50%;
  height: 100%;
  overflow: hidden;
}

.song-list {
  position: relative;
  width: 100%;
  max-height: 40%;
  background-color: #EE6983;
  border-radius: 8px;
  margin-bottom: 1rem;
  overflow-y: scroll;
}

.song-item {
  padding: 0.6rem;
  border-bottom: 1px solid #373737;
  cursor: pointer;
  transition: background 0.2s;
}

.playlist-info {
  position: relative;
  left: 24px;
  top: 20px;
}

.playing-from {
  position: absolute;
  font-family: 'Century Gothic';
  font-size: 10px;
  letter-spacing: 0.013em;
  color: #989898;
}

.playlist-name {
  position: absolute;
  font-family: 'Century Gothic';
  top: 14px;
  font-weight: 700;
  font-size: 12px;
  letter-spacing: 0.06em;
  color: #373737;
}

.chevron-down {
  position: absolute;
  top: 10px;
  left: 60px;
}


.album-art {
  position: absolute;
  width: 340px;
  height: 340px;
  left: 24px;
  top: 60px;
  border: 1px solid #989898;
}

.album-art img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.song-info {
  position: absolute;
  left: 24px;
  top: 420px;
  width: 262px;
  height: 20px;
}

.song-title {
  position: absolute;
  font-family: 'Century Gothic';
  font-style: normal;
  font-weight: 700;
  font-size: 24px;
  line-height: 29px;
  letter-spacing: 0.05em;
  color: #373737;
}

.song-artist {
  position: absolute;
  font-family: 'Century Gothic';
  left: 0;
  top: 37px;
  font-style: normal;
  font-weight: 700;
  font-size: 16px;
  line-height: 20px;
  letter-spacing: 0.05em;
  color: #989898;
}

.progress-container {
  position: absolute;
  width: 331.01px;
  left: 26px;
  top: 500px;
}

.progress-bar {
  position: relative;
  width: 100%;
  height: 100%;
}

.progress-line {
  position: absolute;
  width: 330px;
  border: 5px solid #8A9A9D;
}

.progress-indicator {
  position: absolute;
  width: 240px;
  height: 10px;
  background: #EE6983;
}

.progress-dot {
  position: absolute;
  width: 10px;
  height: 10px;
  left: 227px;
  top: 0;
  background: #EE6983;
  border-radius: 50%;
}

.time-info {
  position: absolute;
  width: 100%;
  top: 18px;
}

.current-time {
  position: absolute;
  left: 0;
  font-family: 'Century Gothic';
  font-style: normal;
  font-weight: 700;
  font-size: 12px;
  line-height: 15px;
  letter-spacing: 0.055em;
  color: #989898;
}

.duration {
  position: absolute;
  font-family: 'Century Gothic';
  right: 0;
  font-style: normal;
  font-weight: 700;
  font-size: 12px;
  line-height: 15px;
  letter-spacing: 0.055em;
  color: #989898;
}

.player-controls {
  position: relative;
  width: 100%;
  height: 56px;
  left: 0;
  top: 520px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 40px;
}

.control-btn {
  width: 40px;
  height: 40px;
  border: none;
  margin:auto;
  background: transparent;
}

.play-btn {
  width: 55px;
  height: 55px;
  background: linear-gradient(180deg, #EF90A2 -15.18%, #EE6983 84.82%);
  box-shadow: 0px 0px 2px #EE6983;
  border-radius: 50%;
}

</style>