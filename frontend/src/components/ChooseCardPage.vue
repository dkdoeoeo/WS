<template>
  <div>
    <h2>選擇卡片系列</h2>
    <div>
      <button @click="selectCardSeries(1)">少女樂團系列</button>
      <button @click="selectCardSeries(2)">系列 2</button>
      <button @click="selectCardSeries(3)">系列 3</button>
    </div>

    <h2 v-if="selectedSeries !== null">選擇你的卡牌</h2>
    <div v-if="selectedSeries !== null" class="card-list">
      <div
        v-for="card in filteredCardList"
        :key="card.id"
        class="card-item"
        @click="incrementCardSelection(card)"
        @contextmenu.prevent="decrementCardSelection(card)"
      >
        <img :src="`http://localhost:8080/images/cards/${card.image}`" alt="Card image" class="card-image"/>
        <div class="card-count">{{ cardSelectionCount(card) }}</div>  <!-- 顯示已選擇數量 -->
        <p>{{ card.name }}</p>
      </div>
    </div>
    <button @click="submitDeckSelection" :disabled="selectedCardIds.length !== 1">
      提交選擇的卡牌
    </button>
    <p v-if="selectedCardIds.length === 1">已選擇 50 張卡牌，準備提交！</p>
  </div>
</template>

<script>
import { cardList } from "@/assets/cardList.js";

export default {
  data() {
    return {
      selectedSeries: null,  // 目前選擇的卡片系列
      cardList,
      selectedCardIds: [], // 存放選擇的卡牌 ID
    };
  },
  computed: {
    // 根據選擇的系列過濾卡片
    filteredCardList() {
      return this.cardList.filter(card => card.series === this.selectedSeries);
    }
  },
  methods: {
    logSelectedCardIds() {
    console.log(this.selectedCardIds);  // 輸出 selectedCardIds 的內容
    },
    // 選擇卡片系列
    selectCardSeries(series) {
      this.selectedSeries = series;
      this.selectedCardIds = [];
    },
    // 增加卡牌選擇數量（左鍵）
    incrementCardSelection(card) {
      const count = this.cardSelectionCount(card);
      const cardType3Count = this.cardSelectionCountByType(3);  // 計算已選擇 card_type = 3 的卡牌數量

      // 每種卡牌最多選擇 4 張，card_type = 3 的卡牌最多選擇 8 張
      if (count < 4 && (card.card_type !== 3 || cardType3Count < 8) && this.selectedCardIds.length<50) {
        this.selectedCardIds.push(card.id);
        this.logSelectedCardIds();
      }
    },
    // 減少卡牌選擇數量（右鍵）
    decrementCardSelection(card) {
      const count = this.cardSelectionCount(card);
      // 如果卡牌已選擇數量大於0，則可以減少
      if (count > 0) {
        const index = this.selectedCardIds.lastIndexOf(card.id);
        if (index !== -1) {
          this.selectedCardIds.splice(index, 1);
        }
      }
    },
    // 計算卡牌的選擇數量
    cardSelectionCount(card) {
      return this.selectedCardIds.filter(id => id === card.id).length;
    },
    // 計算某種 card_type 的選擇數量
    cardSelectionCountByType(cardType) {
      return this.selectedCardIds.filter(id => {
        const card = this.cardList.find(card => card.id === id);
        return card && card.card_type === cardType;
      }).length;
    },
    // 提交選擇的卡牌
    submitDeckSelection() {
      // 當卡牌選擇完畢，發送到後端
      fetch("http://localhost:8080/updateDeck", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ selectedCardIds: this.selectedCardIds }),
      })
      .then(response => response.json())
      .then(data => {
        console.log("選擇的卡牌已提交", data);
        // 可以在這裡使用 data 來進行後續操作
        this.$router.push("/matchmaking"); // 完成選擇後，跳轉到配對頁面
      })
      .catch(error => {
        console.error("錯誤:", error);
      });
    },
  },
};
</script>

<style scoped>
.card-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;  /* 圖片之間的間距 */
}

.card-item {
  position: relative;
  display: inline-block;
  margin: 10px;
  cursor: pointer;
  width: 150px;   /* 固定卡牌的寬度 */
  height: 200px;  /* 固定卡牌的高度 */
  text-align: center; /* 文字居中 */
  overflow: hidden;  /* 防止圖片溢出 */
}

.card-image {
  width: 100%;    /* 設定圖片寬度為 100%，自動填滿父容器 */
  height: 100%;   /* 設定圖片高度為 100%，自動填滿父容器 */
  object-fit: contain;  /* 讓圖片縮小填滿容器，保持比例，不會被裁剪 */
}

.card-count {
  position: absolute;
  top: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  padding: 5px;
  border-radius: 50%;
  font-size: 14px;
}

.card-item.selected {
  border: 2px solid green;
}
</style>
