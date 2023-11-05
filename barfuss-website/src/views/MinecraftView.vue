
<script setup lang="ts">
import DownloadFile from '../components/DownloadFile.vue'

import ToolingIcon from '../components/icons/IconTooling.vue'
</script>
<template>
  <div class="about">
    <h1>Minecraft</h1>
    <DownloadFile>
    <template #icon> <ToolingIcon /> </template>
        <template #heading> Test </template>
        <template #desc> Ich hoffe das funktioniert </template>
        <template #file><a href="./files/Download.png">Download</a></template>
    </DownloadFile>
  </div>
</template>

