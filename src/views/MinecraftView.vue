
<script setup lang="ts">
import DownloadFile from '../components/DownloadFile.vue'
import DownloadButton from '../components/DownloadButton.vue'
import ToolingIcon from '../components/icons/IconTooling.vue'
</script>
<template>
  <div class="about">
    <h1>Minecraft</h1>
    <DownloadFile>
    <template #icon> <ToolingIcon /> </template>
        <template #heading> Test </template>
        <template #desc> Ich hoffe das funktioniert </template>
        <template #file><DownloadButton file="./files/Download.png" name="Hallo.png" /></template>
    </DownloadFile>
      </div>
</template>

