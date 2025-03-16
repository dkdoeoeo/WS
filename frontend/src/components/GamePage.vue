<template>
    <div>
      <h2>遊戲開始！</h2>
      <p>這裡顯示遊戲進行中的內容。</p>
      <button @click="endGame">結束遊戲</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      endGame() {
        this.$router.push("/"); // 返回選擇卡牌頁面
      },
    },
  };
  </script>
  