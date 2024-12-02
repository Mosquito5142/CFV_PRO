<template>
    <nav class="bg-gray-800 p-4">
      <div class="container mx-auto flex justify-between items-center">
        <div class="text-white text-lg font-bold">My App</div>
        <div class="space-x-4">
          <RouterLink to="/" class="text-gray-300 hover:text-white">Home</RouterLink>
          <RouterLink to="/clan" class="text-gray-300 hover:text-white">Clan</RouterLink>
          <RouterLink to="/profile" class="text-gray-300 hover:text-white">Profile</RouterLink>
          <RouterLink to="/deck" class="text-gray-300 hover:text-white">deck</RouterLink>

        </div>
      </div>
    </nav>
  </template>
  
  <script setup>
  import { RouterLink } from 'vue-router'
  </script>
  
  <style scoped>
  </style>