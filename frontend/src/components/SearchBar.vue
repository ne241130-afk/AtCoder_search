<script setup lang="ts">

const keyword = defineModel<string>({
  required: true
})

const emit = defineEmits<{
  (e:"search"):void
}>()

</script>

<template>

<div class="search-container">

<input

v-model="keyword"

placeholder="問題名・コンテスト名"

@keyup.enter="emit('search')"

class="search-input"

/>

<button

class="search-button"

@click="emit('search')"

>

🔍

</button>

</div>

</template>

<style scoped>

.search-container{

display:flex;

gap:12px;

margin-bottom:30px;

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