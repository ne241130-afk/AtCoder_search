<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  page: number
  total: number
  limit: number
}>()

const emit = defineEmits<{
  (e: 'change', page: number): void
}>()

const totalPages = computed(() => Math.max(1, Math.ceil(props.total / props.limit)))

function changePage(nextPage: number) {
  if (nextPage < 1 || nextPage > totalPages.value) {
    return
  }

  emit('change', nextPage)
}

function pageNumbers() {
  const pages: number[] = []
  const maxVisible = 5
  let start = Math.max(1, props.page - 2)
  const end = Math.min(totalPages.value, start + maxVisible - 1)
  start = Math.max(1, end - maxVisible + 1)

  for (let i = start; i <= end; i += 1) {
    pages.push(i)
  }

  return pages
}
</script>

<template>
  <div class="pagination" v-if="total > 0">
    <button :disabled="page <= 1" @click="changePage(page - 1)">
      &lt; 前へ
    </button>

    <button
      v-for="pageNumber in pageNumbers()"
      :key="pageNumber"
      :class="{ active: pageNumber === page }"
      @click="changePage(pageNumber)"
    >
      {{ pageNumber }}
    </button>

    <button :disabled="page >= totalPages" @click="changePage(page + 1)">
      次へ &gt;
    </button>
  </div>
</template>

<style scoped>
.pagination {
  display: flex;
  gap: 8px;
  justify-content: center;
  margin-top: 20px;
  flex-wrap: wrap;
}

button {
  padding: 8px 12px;
  border: 1px solid #1976d2;
  background: white;
  color: #1976d2;
  border-radius: 8px;
  cursor: pointer;
}

button.active {
  background: #1976d2;
  color: white;
}

button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
