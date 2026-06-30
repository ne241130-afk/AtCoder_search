<script setup lang="ts">
import { onMounted, ref } from "vue"

import SearchBar from "../components/SearchBar.vue"
import ProblemList from "../components/ProblemList.vue"

import { getProblems } from "../services/problem"

import type { Problem } from "../types/problem"

const problems = ref<Problem[]>([])

async function load(keyword = "") {
    problems.value = await getProblems(keyword)
}

onMounted(() => {
    load()
})
</script>

<template>

<div class="container">

<h1>

AtCoder Learning Hub

</h1>

<SearchBar

@search="load"

/>

<ProblemList

:problems="problems"

/>

</div>

</template>

<style scoped>

.container{

max-width:900px;

margin:auto;

padding:30px;

}

</style>