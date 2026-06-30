<script setup lang="ts">

import {ref,onMounted} from "vue"

import SearchBar from "../components/SearchBar.vue"
import ProblemList from "../components/ProblemList.vue"
import SortSelect from "../components/SortSelect.vue"

import {getProblems} from "../services/problem"

import type {Problem} from "../types/problem"
import type {SearchCondition} from "../types/searchCondition"

import TagFilter from "../components/TagFilter.vue"
import DifficultyFilter from "../components/DifficultyFilter.vue"
import { getTags } from "../services/problem"

const tags = ref<string[]>([])

const problems=ref<Problem[]>([])

const sortOrder = ref<string>("")

const condition=ref<SearchCondition>({
    keyword:"",
    tags:[],
    contestType:"",
    minDifficulty:null,
    maxDifficulty:null
})

function sortProblems() {
    const nextProblems = [...problems.value]

    if (sortOrder.value === "asc") {
        nextProblems.sort((a, b) => a.difficulty - b.difficulty)
    } else if (sortOrder.value === "desc") {
        nextProblems.sort((a, b) => b.difficulty - a.difficulty)
    }

    problems.value = nextProblems
}

async function search(){

    problems.value=await getProblems(condition.value)
    sortProblems()

}

onMounted(async () => {

    tags.value = await getTags()

    await search()

})

function toggleTag(tag:string){

    if(condition.value.tags.includes(tag)){

        condition.value.tags =
            condition.value.tags.filter(t=>t!==tag)

    }else{

        condition.value.tags.push(tag)

    }

    search()

}

function handleSortChange(value: string) {
    sortOrder.value = value
    sortProblems()
}

onMounted(search)

</script>

<template>

<div class="container">

<h1>

AtCoder Learning Hub

</h1>

<SearchBar
  v-model="condition.keyword"
  @search="search"
/>

<TagFilter
  :tags="tags"
  :selected="condition.tags"
  @toggle="toggleTag"
/>

<label class="contest-select">
  <span>Contest</span>
  <select v-model="condition.contestType" @change="search">
    <option value="">すべて</option>
    <option value="ABC">[ABC]のみ</option>
    <option value="ARC">[ARC]のみ</option>
  </select>
</label>

<DifficultyFilter
  v-model:minDifficulty="condition.minDifficulty"
  v-model:maxDifficulty="condition.maxDifficulty"
  @change="search"
/>

<SortSelect
  v-model="sortOrder"
  @change="handleSortChange"
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

.contest-select {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
}

.contest-select select {
  padding: 8px 12px;
  border-radius: 8px;
  border: 1px solid #ccc;
}

</style>