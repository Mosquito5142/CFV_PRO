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
      .slice(0, 5);
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
    <!-- Header -->
    <div class="text-center mb-8">
      <h1 class="text-4xl font-bold mb-2">
        <span class="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          🎴 Card Sets by Clan
        </span>
      </h1>
      <p class="text-base-content/70">เลือก Clan และค้นหา Set ที่ต้องการ</p>
    </div>

    <!-- Filters Card -->
    <div class="card bg-base-200 shadow-xl mb-8">
      <div class="card-body">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <!-- Clan Select -->
          <div class="form-control">
            <label class="label" for="clan-select">
              <span class="label-text text-lg font-medium">Select a Clan</span>
            </label>
            <select 
              id="clan-select" 
              v-model="selectedClan" 
              class="select select-bordered select-primary w-full"
            >
              <option value="">ทุก Clan</option>
              <option v-for="clan in clans" :key="clan">{{ clan }}</option>
            </select>
          </div>

          <!-- Search Input -->
          <div class="form-control">
            <label class="label">
              <span class="label-text text-lg font-medium">ค้นหาการ์ด</span>
            </label>
            <div class="relative">
              <input 
                type="text" 
                v-model="searchQuery" 
                placeholder="Search by card name..." 
                class="input input-bordered input-primary w-full pr-10"
              />
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 absolute right-3 top-1/2 -translate-y-1/2 text-base-content/50" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
            
            <!-- Suggestions Dropdown -->
            <ul v-if="suggestions.length" class="menu bg-base-300 rounded-box mt-2 shadow-lg absolute z-50 w-full">
              <li v-for="suggestion in suggestions" :key="suggestion">
                <a @click="selectSuggestion(suggestion)" class="hover:bg-primary hover:text-white">
                  {{ suggestion }}
                </a>
              </li>
            </ul>
          </div>
        </div>

        <!-- Active Filters -->
        <div v-if="selectedClan || searchQuery" class="flex flex-wrap gap-2 mt-4">
          <div v-if="selectedClan" class="badge badge-primary badge-lg gap-2">
            {{ selectedClan }}
            <button @click="selectedClan = ''" class="btn btn-ghost btn-xs">✕</button>
          </div>
          <div v-if="searchQuery" class="badge badge-secondary badge-lg gap-2">
            "{{ searchQuery }}"
            <button @click="searchQuery = ''" class="btn btn-ghost btn-xs">✕</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Results -->
    <div v-if="filteredSets.length">
      <div class="flex items-center justify-between mb-4">
        <h2 class="text-2xl font-semibold">
          Sets
          <span class="badge badge-primary ml-2">{{ filteredSets.length }}</span>
        </h2>
      </div>
      
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <div 
          v-for="set in filteredSets" 
          :key="set" 
          @click="goToSet(set)" 
          class="card bg-base-200 shadow-lg hover:shadow-xl hover:scale-105 transition-all duration-300 cursor-pointer group"
        >
          <div class="card-body p-4">
            <div class="flex items-center gap-3">
              <div class="w-12 h-12 rounded-lg bg-gradient-to-br from-primary to-secondary flex items-center justify-center text-white font-bold text-xl group-hover:scale-110 transition-transform">
                {{ set.charAt(0) }}
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="font-semibold truncate">{{ set }}</h3>
                <p class="text-sm text-base-content/60">Card Set</p>
              </div>
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-primary opacity-0 group-hover:opacity-100 transition-opacity" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-16">
      <div class="text-6xl mb-4">🔍</div>
      <h3 class="text-xl font-semibold mb-2">ไม่พบ Set</h3>
      <p class="text-base-content/60">ลองเปลี่ยน Clan หรือคำค้นหาใหม่</p>
    </div>
  </div>
</template>

<style scoped>
</style>