<script setup lang="ts">
import { onMounted, ref } from "vue"

import ProblemCard from "../components/ProblemCard.vue"

import { getProblems } from "../services/problem"

import type { Problem } from "../types/problem"

const problems = ref<Problem[]>([])
const keyword = ref("")

async function search() {
    problems.value = await getProblems(keyword.value)
}

onMounted(async () => {
  problems.value = await getProblems()
})
</script>

<template>
  <div class="container">

    <h1>
      AtCoder Learning Hub
    </h1>

    <input
        v-model="keyword"
        @input="search"
        placeholder="問題名を検索"
    />

    <ProblemCard
      v-for="problem in problems"
      :key="problem.id"
      :problem="problem"
    />

  </div>
</template>

<style scoped>

.container{

    width:900px;

    margin:auto;

}

</style>