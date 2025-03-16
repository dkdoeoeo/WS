<template>
    <div>
      <h2>等待配對對手...</h2>
      <p>目前正在等待對手加入，請稍候。</p>
      <button @click="cancelMatchmaking">取消配對</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      cancelMatchmaking() {
        this.$router.push("/"); // 返回選擇卡牌頁面
      },
    },
  };
  </script>
  