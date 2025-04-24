import {createApp} from 'vue' // createApp - создает экземпляр приложения
import App from './App.vue' // App - корневой компонент
import './style.css'; // тут стили главного окна

createApp(App).mount('#app') // mount - монтирует экземпляр приложения Vue в в HTML-элемент, который соответствует CSS-селектору #app. 
// Vue найдёт этот элемент с id="app" и заменит его содержимое на наше Vue-приложение. - она в index.html
