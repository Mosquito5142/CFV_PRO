<template>
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-center">{{ card.name }}</h1>
      <div class="flex flex-col items-center">
        <img :src="card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md mb-4"/>
        <p class="text-lg"><strong>Clan:</strong> {{ card.clan }}</p>
        <p class="text-lg"><strong>Grade:</strong> {{ card.grade }}</p>
        <p class="text-lg"><strong>Power:</strong> {{ card.power }}</p>
        <p class="text-lg"><strong>Effect:</strong> {{ card.effect }}</p>
        <p class="text-lg"><strong>Flavor:</strong> {{ card.flavor }}</p>
        <p class="text-lg"><strong>Sets:</strong></p>
        <ul class="list-disc list-inside">
          <li v-for="set in card.sets" :key="set">{{ set }}</li>
        </ul>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  import { useRoute } from 'vue-router';
  import allCards from '/public/data/all_cards.json';
  
  const route = useRoute();
  const cardId = route.params.cardId;
  const card = ref({});
  
  onMounted(() => {
    card.value = allCards.find(c => c.id === parseInt(cardId));
  });
  </script>
  
  <style scoped>
  </style>