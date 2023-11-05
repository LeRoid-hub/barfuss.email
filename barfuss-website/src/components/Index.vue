<script setup lang="ts">
import TextItem from './TextItem.vue'
</script>

<template>
    <h1>Tools</h1>
    <TextItem>
        <template #heading> CouchDB - Obsidian </template>
        <a href="http://barfuss.email:5984/_utils">Utils </a>
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
<style scoped>
    h1{
        margin-left: 1rem;
        padding-bottom: 1rem;
    }
</style> 
