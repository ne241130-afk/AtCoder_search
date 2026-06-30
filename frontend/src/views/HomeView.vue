<script setup lang="ts">

import {ref,onMounted,watch} from "vue"
import { useRoute, useRouter } from "vue-router"

import SearchBar from "../components/SearchBar.vue"
import ProblemList from "../components/ProblemList.vue"
import SortSelect from "../components/SortSelect.vue"

import {getProblems} from "../services/problem"

import type {Problem} from "../types/problem"
import type {SearchCondition} from "../types/searchCondition"

import TagFilter from "../components/TagFilter.vue"
import DifficultyFilter from "../components/DifficultyFilter.vue"
import { getTags } from "../services/problem"

const route = useRoute()
const router = useRouter()

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

function toSingleValue(value: unknown): string {
    if (Array.isArray(value) && value.length > 0) {
        return String(value[0])
    }

    if (typeof value === "string") {
        return value
    }

    return ""
}

function parseQuery(query: Record<string, unknown>) {
    condition.value.keyword = toSingleValue(query.keyword)

    const tagsParam = toSingleValue(query.tags)
    condition.value.tags = tagsParam
        ? tagsParam.split(",").map(tag => tag.trim()).filter(Boolean)
        : []

    condition.value.contestType = toSingleValue(query.contest)

    const minParam = toSingleValue(query.minDifficulty)
    condition.value.minDifficulty = minParam ? Number(minParam) : null

    const maxParam = toSingleValue(query.maxDifficulty)
    condition.value.maxDifficulty = maxParam ? Number(maxParam) : null

    sortOrder.value = toSingleValue(query.sort)
}

function buildQuery() {
    const query: Record<string, string> = {}

    if (condition.value.keyword) {
        query.keyword = condition.value.keyword
    }

    if (condition.value.tags.length > 0) {
        query.tags = condition.value.tags.join(",")
    }

    if (condition.value.contestType) {
        query.contest = condition.value.contestType
    }

    if (condition.value.minDifficulty !== null && condition.value.minDifficulty !== undefined) {
        query.minDifficulty = String(condition.value.minDifficulty)
    }

    if (condition.value.maxDifficulty !== null && condition.value.maxDifficulty !== undefined) {
        query.maxDifficulty = String(condition.value.maxDifficulty)
    }

    if (sortOrder.value) {
        query.sort = sortOrder.value
    }

    return query
}

function syncUrl() {
    router.replace({ query: buildQuery() })
}

async function search(){

    problems.value=await getProblems(condition.value)
    sortProblems()
    syncUrl()

}

onMounted(async () => {

    tags.value = await getTags()
    parseQuery(route.query)
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
    syncUrl()
}

watch(
    () => [condition.value.keyword, condition.value.tags, condition.value.contestType, condition.value.minDifficulty, condition.value.maxDifficulty],
    () => {
        syncUrl()
    },
    { deep: true }
)

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