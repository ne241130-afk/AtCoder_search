<script setup lang="ts">
import type { Problem } from "../types/problem"

defineProps<{
  problem: Problem
}>()
</script>

<template>
  <div class="card">
    <div class="contest">
      {{ problem.contest }}
    </div>

    <h2>
      {{ problem.title }}
    </h2>

    <div class="difficulty">
      Difficulty : {{ problem.difficulty }}
    </div>

    <div class="tags">
      <span
        v-for="tag in problem.tags"
        :key="tag"
        class="tag"
      >
        {{ tag }}
      </span>
    </div>

    <button>
      詳細
    </button>
  </div>
</template>

<style scoped>

.card{
    background:white;
    border-radius:12px;
    padding:20px;
    margin:20px 0;
    box-shadow:0 2px 8px rgba(0,0,0,.15);
}

.tag{
    display:inline-block;
    margin-right:8px;
    padding:4px 10px;
    border-radius:20px;
    background:#2196f3;
    color:white;
    font-size:13px;
}

button{
    margin-top:15px;
}

</style>