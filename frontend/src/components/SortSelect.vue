<script setup lang="ts">
const model = defineModel<string>({ required: true })

const emit = defineEmits<{
  (e: 'change', value: string): void
}>()

function handleChange(event: Event) {
  const target = event.target as HTMLSelectElement
  const value = target.value
  model.value = value
  emit('change', value)
}
</script>

<template>
  <label class="sort-select">
    <span>Sort</span>
    <select :value="model" @change="handleChange">
      <option value="">選択してください</option>
      <option value="asc">Difficulty 昇順</option>
      <option value="desc">Difficulty 降順</option>
    </select>
  </label>
</template>

<style scoped>
.sort-select {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
}

select {
  padding: 8px 12px;
  border-radius: 8px;
  border: 1px solid #ccc;
}
</style>
