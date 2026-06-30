<script setup lang="ts">

import {ref,onMounted} from "vue"

import SearchBar from "../components/SearchBar.vue"
import ProblemList from "../components/ProblemList.vue"

import {getProblems} from "../services/problem"

import type {Problem} from "../types/problem"
import type {SearchCondition} from "../types/searchCondition"

import TagFilter from "../components/TagFilter.vue"
import { getTags } from "../services/problem"

const tags = ref<string[]>([])

const problems=ref<Problem[]>([])

const condition=ref<SearchCondition>({
    keyword:"",
    tags:[]
})

async function search(){

    problems.value=await getProblems(condition.value)

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