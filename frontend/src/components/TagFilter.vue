<script setup lang="ts">

defineProps<{

    tags:string[]

    selected:string[]

}>()

const emit=defineEmits<{

(e:"toggle",tag:string):void

}>()

</script>

<template>

<div class="tag-list">

<button

v-for="tag in tags"

:key="tag"

@click="emit('toggle',tag)"

:class="{active:selected.includes(tag)}"

>

{{tag}}

</button>

</div>

</template>

<style scoped>

.tag-list{

display:flex;

gap:10px;

flex-wrap:wrap;

margin-bottom:25px;

}

button{

padding:8px 16px;

border:none;

border-radius:20px;

cursor:pointer;

background:#eeeeee;

}

.active{

background:#1976d2;

color:white;

}

</style>