<template>
    <h2>hello profile require login</h2>
</template>

<script>
export default {
    layout: "navbar",
    middleware: 'auth'
}
</script>