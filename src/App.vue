<script setup>
import { RouterView } from 'vue-router'
import NavBar from './components/NavBar.vue'
</script>

<template>
  <div data-theme="cfvdark" class="min-h-screen bg-base-100">
    <NavBar />
    <main class="pb-8">
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
</style>