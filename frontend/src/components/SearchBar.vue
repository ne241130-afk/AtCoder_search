<script setup lang="ts">
import { ref } from "vue"

const keyword = ref("")

const emit = defineEmits<{
  (e: "search", keyword: string): void
}>()

function search() {
  emit("search", keyword.value)
}
</script>

<template>
  <div class="search-container">
    <input
      v-model="keyword"
      @keyup.enter="search"
      placeholder="問題名・コンテスト名で検索"
      class="search-input"
    />

    <button
      class="search-button"
      @click="search"
    >
      🔍
    </button>
  </div>
</template>

<style scoped>

.search-container{

    display:flex;

    gap:12px;

    margin:30px 0;

}

.search-input{

    flex:1;

    padding:12px;

    font-size:16px;

    border-radius:8px;

    border:1px solid #ccc;

}

.search-button{

    width:60px;

    border:none;

    border-radius:8px;

    cursor:pointer;

}

</style>