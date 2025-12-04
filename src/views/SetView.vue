<template>
  <div class="container mx-auto p-4">
    <!-- Header -->
    <div class="mb-8">
      <div class="flex items-center gap-2 text-sm text-base-content/60 mb-2">
        <RouterLink to="/clan" class="hover:text-primary">Clans</RouterLink>
        <span>/</span>
        <span class="text-primary">{{ setName }}</span>
      </div>
      <h1 class="text-3xl font-bold">
        <span class="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          📦 {{ setName }}
        </span>
      </h1>
    </div>

    <!-- Filters Card -->
    <div class="card bg-base-200 shadow-xl mb-8">
      <div class="card-body p-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <!-- Clan Select -->
          <div class="form-control">
            <label class="label py-1" for="clan-select">
              <span class="label-text font-medium">Filter by Clan</span>
            </label>
            <select id="clan-select" v-model="selectedClan" class="select select-bordered select-sm w-full">
              <option value="">All Clans</option>
              <option v-for="clan in clansInSet" :key="clan">{{ clan }}</option>
            </select>
          </div>

          <!-- Search Input -->
          <div class="form-control">
            <label class="label py-1">
              <span class="label-text font-medium">Search</span>
            </label>
            <div class="join w-full">
              <input type="text" v-model="searchQuery" placeholder="Card name..."
                class="input input-bordered input-sm join-item w-full" />
              <button class="btn btn-primary btn-sm join-item">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </button>
            </div>
          </div>
        </div>

        <!-- Results count -->
        <div class="text-sm text-base-content/60 mt-2">
          พบ {{ filteredCards.length }} การ์ด
        </div>
      </div>
    </div>

    <!-- Cards Grid -->
    <div v-if="filteredCards.length"
      class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-4">
      <div v-for="card in filteredCards" :key="card.id" @click="goToCard(card.id)"
        class="card bg-base-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-105 cursor-pointer group">
        <figure class="pt-2 px-2">
          <img :src="card.imageurlen || card.imageurljp" :alt="card.name" class="rounded-lg w-full" loading="lazy" />
        </figure>
        <div class="card-body p-3">
          <h3 class="text-sm font-semibold line-clamp-2 group-hover:text-primary transition-colors">
            {{ card.name }}
          </h3>
          <div class="flex flex-wrap gap-1 mt-1">
            <span class="badge badge-outline badge-xs">G{{ card.grade }}</span>
            <span v-if="card.triggereffect" class="badge badge-warning badge-xs">
              {{ card.triggereffect }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-16">
      <div class="text-6xl mb-4">🃏</div>
      <h3 class="text-xl font-semibold mb-2">ไม่พบการ์ด</h3>
      <p class="text-base-content/60">ลองเปลี่ยนตัวกรองใหม่</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter, RouterLink } from 'vue-router';
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
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>