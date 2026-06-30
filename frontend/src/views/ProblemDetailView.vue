<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getProblemById, getProblemStatus, updateProblemFavorite, updateProblemMemo, updateProblemSolved } from '../services/problem'
import type { Problem, UserProblemStatus } from '../types/problem'

const route = useRoute()
const router = useRouter()

const problem = ref<Problem | null>(null)
const status = ref<UserProblemStatus | null>(null)
const memoDraft = ref('')
const saving = ref(false)

async function loadProblem() {
  const id = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id
  if (!id) {
    return
  }
  problem.value = await getProblemById(id)
  const nextStatus = await getProblemStatus(id)
  status.value = nextStatus
  memoDraft.value = nextStatus.memo
}

async function toggleSolved() {
  if (!problem.value) return
  const nextValue = !status.value?.solved
  await updateProblemSolved(problem.value.id, nextValue)
  if (status.value) {
    status.value.solved = nextValue
  }
}

async function toggleFavorite() {
  if (!problem.value) return
  const nextValue = !status.value?.favorite
  await updateProblemFavorite(problem.value.id, nextValue)
  if (status.value) {
    status.value.favorite = nextValue
  }
}

async function saveMemo() {
  if (!problem.value) return
  saving.value = true
  try {
    await updateProblemMemo(problem.value.id, memoDraft.value)
    if (status.value) {
      status.value.memo = memoDraft.value
    }
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  loadProblem()
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
    <div class="status-actions">
      <button class="toggle-button" :class="{ active: status?.solved }" @click="toggleSolved">
        {{ status?.solved ? '✓ Solved' : '○ Solved' }}
      </button>
      <button class="toggle-button" :class="{ active: status?.favorite }" @click="toggleFavorite">
        {{ status?.favorite ? '★ Favorite' : '☆ Favorite' }}
      </button>
    </div>

    <hr />

    <h2>メモ</h2>
    <textarea v-model="memoDraft" rows="6" class="memo-input" placeholder="学習メモを入力"></textarea>
    <button class="save-button" @click="saveMemo" :disabled="saving">
      {{ saving ? '保存中...' : '保存' }}
    </button>
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
