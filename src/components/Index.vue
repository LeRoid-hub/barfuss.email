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
        <template #heading> Subdomains </template>
        <a href="http://portainer.barfuss.email">Portainer </a> <br>
        <a href="http://nginx.barfuss.email">Nginx </a>
    </TextItem>


</template>
<style scoped>
    h1{
        margin-left: 1rem;
        padding-bottom: 1rem;
    }
</style> 
