<template>
  <div>
    <h2>等待配對對手...</h2>
    <p>目前正在等待對手加入，請稍候。</p>
    <button @click="cancelMatchmaking">取消配對</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      socket: null,  // WebSocket 連線對象
      playerGameID: null, // 玩家 ID
      matchStatus: "" // 配對狀態
    };
  },
  mounted() {

    const playerId = localStorage.getItem("player_id");
    if (!playerId) {
      console.error("未找到玩家 ID，請重新選擇卡組！");
      this.$router.push("/"); // 導回選擇卡牌頁面
    } else {
      console.log("玩家 ID:", playerId);
    }
    // 當頁面加載時建立 WebSocket 連線
    this.socket = new WebSocket("ws://localhost:8080/ws");

    // 當 WebSocket 連線成功時
    this.socket.onopen = () => {
      console.log("WebSocket 連線成功！");
      // 發送配對請求
      this.socket.send(JSON.stringify({ 
        type: "player_id", // 自定義的訊息類型
        player_id: playerId, // 傳送數字型玩家 ID
      }));
    };

    // 當收到來自後端的訊息時
    this.socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.playerGameID !== undefined) {
        console.log("收到玩家遊戲 ID:", data.playerGameID);
        this.playerGameID = data.playerGameID;
        // 配對成功，跳轉到遊戲頁面
        this.$router.push({ path: '/game', query: { playerGameID: this.playerGameID } });
      }
      if (data.message) {
        this.matchStatus = data.message;
      }
    };

    // 當 WebSocket 連線關閉時
    this.socket.onclose = () => {
      console.log("WebSocket 連線已關閉");
    };

    // 當 WebSocket 發生錯誤時
    this.socket.onerror = (error) => {
      console.log("WebSocket 發生錯誤:", error);
    };
  },
  methods: {
    // 取消配對並返回到卡牌選擇頁面
    cancelMatchmaking() {
      if (this.socket) {
        this.socket.close(); // 關閉 WebSocket 連線
      }
      this.$router.push("/"); // 返回到卡牌選擇頁面
    }
  },
  beforeUnmount() {
    // 當頁面銷毀時關閉 WebSocket 連線
    if (this.socket) {
      this.socket.close();
    }
  }
};
</script>
