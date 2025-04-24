<script setup>
import Playlist from './components/Playlist.vue'
</script>

<template>
  <!-- Сюда потом добавлю менюшку верхнюю  -->
  <!-- Шаблон для картинки меню: <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>  -->
  <Playlist/>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
