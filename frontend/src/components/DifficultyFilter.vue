<script setup lang="ts">
const minDifficulty = defineModel<number | null>('minDifficulty', { required: true })
const maxDifficulty = defineModel<number | null>('maxDifficulty', { required: true })

const emit = defineEmits<{
  (e: 'change'): void
}>()

const STEP = 400
const MIN = 0
const MAX = 4400

function normalizeValue(value: number | null) {
  if (value === null) {
    return null
  }

  const clamped = Math.min(Math.max(value, MIN), MAX)
  const rounded = Math.round(clamped / STEP) * STEP

  return Math.min(Math.max(rounded, MIN), MAX)
}

function handleInput(field: 'minDifficulty' | 'maxDifficulty', event: Event) {
  const target = event.target as HTMLInputElement
  const rawValue = target.value === '' ? null : Number(target.value)
  const value = normalizeValue(rawValue)

  if (field === 'minDifficulty') {
    minDifficulty.value = value
  } else {
    maxDifficulty.value = value
  }

  emit('change')
}
</script>

<template>
  <div class="difficulty-filter">
    <label class="filter-item">
      <span>最低Difficulty</span>
      <input
        type="number"
        :value="minDifficulty ?? ''"
        min="0"
        max="4400"
        step="400"
        @input="handleInput('minDifficulty', $event)"
      />
    </label>

    <label class="filter-item">
      <span>最大Difficulty</span>
      <input
        type="number"
        :value="maxDifficulty ?? ''"
        min="0"
        max="4400"
        step="400"
        @input="handleInput('maxDifficulty', $event)"
      />
    </label>
  </div>
</template>

<style scoped>
.difficulty-filter {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  margin-bottom: 25px;
}

.filter-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.filter-item span {
  font-size: 14px;
}

input {
  padding: 8px 12px;
  border-radius: 8px;
  border: 1px solid #ccc;
  min-width: 120px;
}
</style>
