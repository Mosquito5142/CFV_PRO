<template>
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-center">Cards in {{ setName }}</h1>
      <div class="mb-6">
        <label for="clan-select" class="block text-lg font-medium mb-2">Select Clan:</label>
        <select id="clan-select" v-model="selectedClan" class="w-full p-2 border border-gray-300 rounded-md mb-4">
          <option value="">All Clans</option>
          <option v-for="clan in clansInSet" :key="clan">{{ clan }}</option>
        </select>
        <input type="text" v-model="searchQuery" placeholder="Search by card name" class="w-full p-2 border border-gray-300 rounded-md"/>
      </div>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <div v-for="card in filteredCards" :key="card.id" @click="goToCard(card.id)" class="border border-gray-300 rounded-lg p-4 text-center bg-white shadow-md cursor-pointer hover:bg-gray-100">
          <img :src="card.imageurlen || card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md"/>
          <p class="mt-2 text-lg font-bold">{{ card.name }}</p>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, computed } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import allCards from '/public/data/all_cards.json';
  
  const route = useRoute();
  const router = useRouter();
  const setName = route.params.setName;
  const cardsInSet = ref([]);
  const selectedClan = ref('');
  const clansInSet = ref([]);
  const searchQuery = ref('');
  
  onMounted(() => {
    cardsInSet.value = allCards.filter(card => card.sets.includes(setName));
    const uniqueClans = [...new Set(cardsInSet.value.map(card => card.clan))].sort();
    clansInSet.value = uniqueClans;
  });
  
  const filteredCards = computed(() => {
    let filtered = cardsInSet.value;
    if (selectedClan.value) {
      filtered = filtered.filter(card => card.clan === selectedClan.value);
    }
    if (searchQuery.value) {
      filtered = filtered.filter(card => card.name.toLowerCase().includes(searchQuery.value.toLowerCase()));
    }
    return filtered;
  });
  
  const goToCard = (cardId) => {
    router.push({ path: `/card/${cardId}` });
  };
  </script>
  
  <style scoped>
  </style>