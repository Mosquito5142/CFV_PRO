<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useRouter } from 'vue-router';
import allCards from '/public/data/all_cards.json';

const clans = ref([]);
const selectedClan = ref('');
const searchQuery = ref('');
const sets = ref([]);
const router = useRouter();
const suggestions = ref([]);

onMounted(() => {
  const uniqueClans = [...new Set(allCards.map(card => card.clan))].sort();
  clans.value = uniqueClans;
});

const filteredSets = computed(() => {
  let filteredCards = allCards;
  if (selectedClan.value) {
    filteredCards = filteredCards.filter(card => card.clan === selectedClan.value);
  }
  if (searchQuery.value) {
    filteredCards = filteredCards.filter(card => card.name.toLowerCase().includes(searchQuery.value.toLowerCase()));
  }
  const allSets = filteredCards.flatMap(card => card.sets);
  return [...new Set(allSets)].sort();
});

const goToSet = (setName) => {
  router.push({ path: `/set/${setName}` });
};

watch(searchQuery, (newQuery) => {
  if (newQuery) {
    suggestions.value = allCards
      .filter(card => card.name.toLowerCase().includes(newQuery.toLowerCase()))
      .map(card => card.name)
      .slice(0, 5); // Limit to 5 suggestions
  } else {
    suggestions.value = [];
  }
});

const selectSuggestion = (suggestion) => {
  searchQuery.value = suggestion;
  suggestions.value = [];
};
</script>

<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-bold mb-6 text-center">Card Sets by Clan</h1>
    <div class="mb-6">
      <label for="clan-select" class="block text-lg font-medium mb-2">Select a Clan:</label>
      <select id="clan-select" v-model="selectedClan" class="w-full p-2 border border-gray-300 rounded-md mb-4">
        <option value="">Select a clan</option>
        <option v-for="clan in clans" :key="clan">{{ clan }}</option>
      </select>
      <input type="text" v-model="searchQuery" placeholder="Search by card name" class="w-full p-2 border border-gray-300 rounded-md"/>
      <ul v-if="suggestions.length" class="border border-gray-300 rounded-md mt-2 bg-white">
        <li v-for="suggestion in suggestions" :key="suggestion" @click="selectSuggestion(suggestion)" class="cursor-pointer p-2 hover:bg-gray-100">
          {{ suggestion }}
        </li>
      </ul>
    </div>
    <div v-if="filteredSets.length">
      <h2 class="text-2xl font-semibold mb-4">Sets:</h2>
      <ul class="space-y-2">
        <li v-for="set in filteredSets" :key="set" @click="goToSet(set)" class="cursor-pointer p-4 bg-gray-100 hover:bg-gray-200 rounded-md shadow-md">
          {{ set }}
        </li>
      </ul>
    </div>
    <div v-else>
      <p class="text-center text-gray-500">No sets found.</p>
    </div>
  </div>
</template>

<style scoped>
</style>