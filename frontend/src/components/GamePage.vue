<template>
  <div>
    <h2>遊戲頁面</h2>
    <p>Player ID: {{ playerID }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      playerID: null,
    };
  },
  created() {
    this.playerID = this.$route.query.playerID;
  },
};
</script>
