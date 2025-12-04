<template>
  <div class="container mx-auto p-4">
    <!-- Breadcrumb -->
    <div class="text-sm breadcrumbs mb-4">
      <ul>
        <li>
          <RouterLink to="/">Home</RouterLink>
        </li>
        <li>
          <RouterLink to="/clan">Clans</RouterLink>
        </li>
        <li class="text-primary">{{ card.name || 'Loading...' }}</li>
      </ul>
    </div>

    <div v-if="card.id" class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- Card Image -->
      <div class="flex justify-center">
        <div class="card bg-base-200 shadow-2xl overflow-hidden max-w-sm">
          <figure class="relative group">
            <img :src="card.imageurlen || card.imageurljp" :alt="card.name"
              class="w-full transition-transform duration-300 group-hover:scale-105" />
            <div
              class="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent opacity-0 group-hover:opacity-100 transition-opacity">
            </div>
          </figure>
        </div>
      </div>

      <!-- Card Details -->
      <div class="space-y-6">
        <!-- Title -->
        <div>
          <h1 class="text-3xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent mb-2">
            {{ card.name }}
          </h1>
          <div class="flex flex-wrap gap-2">
            <span class="badge badge-primary badge-lg">{{ card.clan }}</span>
            <span class="badge badge-secondary badge-lg">Grade {{ card.grade }}</span>
            <span v-if="card.triggereffect" class="badge badge-warning badge-lg">
              {{ card.triggereffect }}
            </span>
          </div>
        </div>

        <!-- Stats -->
        <div class="stats stats-vertical lg:stats-horizontal bg-base-200 shadow w-full">
          <div class="stat">
            <div class="stat-title">Power</div>
            <div class="stat-value text-primary text-2xl">{{ card.power || '-' }}</div>
          </div>
          <div class="stat">
            <div class="stat-title">Shield</div>
            <div class="stat-value text-secondary text-2xl">{{ card.shield || '-' }}</div>
          </div>
          <div class="stat">
            <div class="stat-title">Critical</div>
            <div class="stat-value text-accent text-2xl">{{ card.critical || '1' }}</div>
          </div>
        </div>

        <!-- Effect -->
        <div v-if="card.effect" class="card bg-base-200 shadow-lg">
          <div class="card-body">
            <h3 class="card-title text-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-primary" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              Effect
            </h3>
            <p class="text-base-content/80 whitespace-pre-wrap">{{ card.effect }}</p>
          </div>
        </div>

        <!-- Flavor -->
        <div v-if="card.flavor" class="card bg-base-200/50 shadow">
          <div class="card-body py-4">
            <p class="italic text-base-content/60">"{{ card.flavor }}"</p>
          </div>
        </div>

        <!-- Sets -->
        <div class="card bg-base-200 shadow-lg">
          <div class="card-body">
            <h3 class="card-title text-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-secondary" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
              </svg>
              Available in Sets
            </h3>
            <div class="flex flex-wrap gap-2">
              <RouterLink v-for="set in card.sets" :key="set" :to="`/set/${set}`"
                class="badge badge-outline hover:badge-primary transition-colors cursor-pointer">
                {{ set }}
              </RouterLink>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex gap-4">
          <button class="btn btn-primary flex-1">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            เพิ่มลงเด็ค
          </button>
          <button class="btn btn-outline btn-secondary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Loading State -->
    <div v-else class="flex justify-center items-center min-h-[50vh]">
      <span class="loading loading-spinner loading-lg text-primary"></span>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, RouterLink } from 'vue-router';
import allCards from '/public/data/all_cards.json';

const route = useRoute();
const cardId = route.params.cardId;
const card = ref({});

onMounted(() => {
  card.value = allCards.find(c => c.id === parseInt(cardId)) || {};
});
</script>

<style scoped></style>