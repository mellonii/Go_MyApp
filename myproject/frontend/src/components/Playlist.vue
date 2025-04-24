<script setup>
// Код внутри компилируется каждый раз при создании экземпляра компонента
import {ref} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'

/*
const data = reactive({
  name: "",
  resultText: "Введите путь до песни:",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}
*/

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
  <main>
    <div class="container">
      <h1>🎵 Мой плеер</h1>

      <!-- Список песен -->
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
        <button @click="pauseSong">Пауза</button>
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
  </main>
</template>


<style scoped>

.container {
  max-width: 400px;
  margin: 2rem auto;
  font-family: sans-serif;
  color: #fff;
  background-color: #222;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
}

h1 {
  text-align: center;
  margin-bottom: 1rem;
}

.song-list {
  max-height: 200px;
  overflow-y: auto;
  background-color: #333;
  border-radius: 8px;
  padding: 0.5rem;
  margin-bottom: 1rem;
}

.song-item {
  padding: 0.6rem;
  border-bottom: 1px solid #444;
  cursor: pointer;
  transition: background 0.2s;
}

.song-item:hover {
  background-color: #444;
}

.controls {
  display: flex;
  justify-content: space-between;
  gap: 0.5rem;
}

button {
  flex: 1;
  padding: 0.6rem;
  border: none;
  background-color: #42b883;
  color: #fff;
  border-radius: 8px;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.2s;
}

button:hover {
  background-color: #369f6b;
}

</style>