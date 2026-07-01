<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import {
  attachProblemTag,
  createTag,
  detachProblemTag,
  getProblemById,
  getProblemStatus,
  getTags,
  updateProblemFavorite,
  updateProblemMemo,
  updateProblemSolved,
} from '../services/problem'
import type { Problem, Tag, UserProblemStatus } from '../types/problem'

const route = useRoute()
const router = useRouter()

const problem = ref<Problem | null>(null)
const status = ref<UserProblemStatus | null>(null)
const allTags = ref<Tag[]>([])
const selectedTagId = ref<number | ''>('')
const newTagName = ref('')
const memoDraft = ref('')
const saving = ref(false)
const tagSaving = ref(false)

const availableTags = computed(() => {
  if (!problem.value) return []
  const attachedIds = new Set((problem.value.tags ?? []).map(tag => tag.id))
  return allTags.value.filter(tag => !attachedIds.has(tag.id))
})

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

async function loadTags() {
  allTags.value = await getTags()
}

async function addExistingTag() {
  if (!problem.value || selectedTagId.value === '') return
  tagSaving.value = true
  try {
    await attachProblemTag(problem.value.id, selectedTagId.value)
    await loadProblem()
    selectedTagId.value = ''
  } finally {
    tagSaving.value = false
  }
}

async function addNewTag() {
  const name = newTagName.value.trim()
  if (!problem.value || !name) return
  tagSaving.value = true
  try {
    const tag = await createTag(name)
    await attachProblemTag(problem.value.id, tag.id)
    await loadTags()
    await loadProblem()
    newTagName.value = ''
  } finally {
    tagSaving.value = false
  }
}

async function removeTag(tagId: number) {
  if (!problem.value) return
  tagSaving.value = true
  try {
    await detachProblemTag(problem.value.id, tagId)
    await loadProblem()
  } finally {
    tagSaving.value = false
  }
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

onMounted(async () => {
  await Promise.all([loadProblem(), loadTags()])
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
      <div>
        <strong>タグ</strong>
        <div class="tags">
          <span v-for="tag in problem.tags" :key="tag.id" class="tag">
            {{ tag.name }}
            <button
              type="button"
              class="tag-remove"
              :disabled="tagSaving"
              @click="removeTag(tag.id)"
              aria-label="タグを削除"
            >
              ×
            </button>
          </span>
          <span v-if="problem.tags.length === 0" class="no-tags">タグなし</span>
        </div>
        <div class="tag-actions">
          <select v-model="selectedTagId" :disabled="tagSaving || availableTags.length === 0">
            <option value="">タグを選択</option>
            <option v-for="tag in availableTags" :key="tag.id" :value="tag.id">
              {{ tag.name }}
            </option>
          </select>
          <button
            type="button"
            class="tag-button"
            :disabled="tagSaving || selectedTagId === ''"
            @click="addExistingTag"
          >
            追加
          </button>
        </div>
        <div class="tag-actions">
          <input
            v-model="newTagName"
            type="text"
            placeholder="新しいタグ名"
            :disabled="tagSaving"
          />
          <button
            type="button"
            class="tag-button"
            :disabled="tagSaving || !newTagName.trim()"
            @click="addNewTag"
          >
            作成して追加
          </button>
        </div>
      </div>
    </div>

    <hr />

    <h2>外部リンク</h2>
    <div class="actions">
      <a :href="problem.atCoderUrl" target="_blank" rel="noopener noreferrer" class="action-link">
        問題を開く
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
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 10px;
  border-radius: 999px;
  background: #2196f3;
  color: white;
  font-size: 13px;
}

.tag-remove {
  border: none;
  background: rgba(255, 255, 255, 0.25);
  color: white;
  border-radius: 999px;
  width: 20px;
  height: 20px;
  cursor: pointer;
  line-height: 1;
  padding: 0;
}

.tag-remove:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.no-tags {
  color: #777;
  font-size: 14px;
}

.tag-actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  margin-top: 10px;
}

.tag-actions select,
.tag-actions input {
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 8px;
}

.tag-button {
  padding: 8px 12px;
  border: none;
  border-radius: 8px;
  background: #1976d2;
  color: white;
  cursor: pointer;
}

.tag-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
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
