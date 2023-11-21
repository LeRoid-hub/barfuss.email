<template>
    <div>
        <button @click="downloadFile">Download</button>
    </div>
</template>

<script>
    export default {
        name: 'DownloadButton',
        props: {
            file: {
                type: String,
                required: true,
            },
            name: {
                type: String,
                required: true,
            },
        },
        methods: {
            downloadFile() {
                const link = document.createElement('a')
                link.href = this.file
                link.download = this.name
                link.click()
            },
        },
    }
</script>

