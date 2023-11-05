<script setup lang="ts">
import TextItem from './TextItem.vue'
</script>

<template>
    <TextItem>
        <template #heading> TEST </template>
        Dies ist ein Test!
    </TextItem>

    <TextItem>
        <template #heading> TEST </template>
        Dies ist ein Test!
    </TextItem>
    <TextItem>
        <template #heading> TEST </template>
        Dies ist ein Test!
    </TextItem>


</template>
