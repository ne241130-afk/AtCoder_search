<script setup lang="ts">
import type { Problem } from "../types/problem"

import ProblemCard from "./ProblemCard.vue"

defineProps<{
  problems: Problem[]
}>()
</script>

<template>

<ProblemCard

v-for="problem in problems"

:key="problem.id"

:problem="problem"

/>

</template>