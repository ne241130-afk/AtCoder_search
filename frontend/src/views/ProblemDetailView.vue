<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getProblemById } from '../services/problem'
import type { Problem } from '../types/problem'

const route = useRoute()
const router = useRouter()

const problem = ref<Problem | null>(null)

onMounted(async () => {
  const id = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id
  if (id) {
    problem.value = await getProblemById(id)
  }
})
</script>

<template>
  <div class="container" v-if="problem">
    <button class="back-button" @click="router.back()">
      ← 一覧へ戻る
    </button>

    <h1>{{ problem.title }}</h1>

    <div class="info-block">
      <div><strong>Contest</strong><div>{{ problem.contest }}</div></div>
      <div><strong>Difficulty</strong><div>★★★★☆</div></div>
      <div><strong>タグ</strong><div class="tags">
        <span v-for="tag in problem.tags" :key="tag" class="tag">{{ tag }}</span>
      </div></div>
    </div>

    <hr />

    <h2>外部リンク</h2>
    <div class="actions">
      <a :href="problem.atCoderUrl" target="_blank" rel="noopener noreferrer" class="action-link">
        問題を開く
      </a>
      <a :href="problem.problemsUrl" target="_blank" rel="noopener noreferrer" class="action-link">
        AtCoder Problems
      </a>
    </div>

    <hr />

    <h2>学習情報</h2>
    <p><strong>Solved</strong>：未実装</p>
    <p><strong>Favorite</strong>：未実装</p>

    <hr />

    <h2>メモ</h2>
    <p>（今後実装予定）</p>
  </div>
</template>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 30px;
}

.back-button {
  border: none;
  background: transparent;
  color: #1976d2;
  cursor: pointer;
  padding: 0;
  margin-bottom: 20px;
}

.info-block {
  display: grid;
  gap: 16px;
  margin-top: 20px;
}

.tags {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  margin-top: 6px;
}

.tag {
  padding: 6px 10px;
  border-radius: 999px;
  background: #2196f3;
  color: white;
  font-size: 13px;
}

.actions {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.action-link {
  display: inline-block;
  padding: 8px 12px;
  border-radius: 8px;
  background: #1976d2;
  color: white;
  text-decoration: none;
}
</style>
