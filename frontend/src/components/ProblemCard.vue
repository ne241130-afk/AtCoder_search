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

    <div class="actions">
      <a
        v-if="problem.atCoderUrl"
        :href="problem.atCoderUrl"
        target="_blank"
        rel="noopener noreferrer"
        class="action-link"
      >
        問題を開く
      </a>
      <a
        v-if="problem.problemsUrl"
        :href="problem.problemsUrl"
        target="_blank"
        rel="noopener noreferrer"
        class="action-link"
      >
        AtCoder Problems
      </a>
      <router-link :to="{ name: 'problem-detail', params: { id: problem.id } }" class="action-link detail-link">
        詳細を見る
      </router-link>
    </div>
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

.actions{
    display:flex;
    gap:10px;
    flex-wrap:wrap;
    margin-top:15px;
}

.action-link{
    display:inline-block;
    padding:8px 12px;
    border-radius:8px;
    background:#1976d2;
    color:white;
    text-decoration:none;
}

.detail-link{
    background:#2e7d32;
}

</style>