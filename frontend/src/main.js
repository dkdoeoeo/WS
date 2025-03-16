import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router'; // 引入 Vue Router

// 引入你的頁面組件
import ChooseCardPage from './components/ChooseCardPage.vue';
import MatchmakingPage from './components/MatchmakingPage.vue';
import GamePage from './components/GamePage.vue';

// 設定路由
const routes = [
  { path: '/', component: ChooseCardPage },           // 首頁：選擇卡牌
  { path: '/matchmaking', component: MatchmakingPage }, // 配對頁面
  { path: '/game', component: GamePage },              // 遊戲頁面
];

// 創建 Vue Router 實例
const router = createRouter({
  history: createWebHistory(), // 使用 HTML5 History 模式
  routes, // 設定路由
});

// 創建 Vue 實例並掛載
createApp(App).use(router).mount('#app');
