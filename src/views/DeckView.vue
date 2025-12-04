<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-bold mb-6 text-center">
      <span class="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
        🎴 Deck Builder
      </span>
    </h1>

    <div class="flex flex-col lg:flex-row gap-4">
      <!-- Left side - Card List -->
      <div class="w-full lg:w-1/2 card bg-base-200 shadow-xl">
        <div class="card-body p-4">
          <h2 class="card-title text-lg mb-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-primary" fill="none" viewBox="0 0 24 24"
              stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
            Card Library
            <span class="badge badge-ghost badge-sm">{{ filteredCards.length }} cards</span>
          </h2>

          <!-- Filters -->
          <div class="space-y-2 mb-3">
            <!-- Search with clear button -->
            <div class="join w-full">
              <input type="text" v-model="searchQuery" placeholder="🔍 Search cards..."
                class="input input-bordered input-sm join-item w-full" />
              <button v-if="searchQuery" @click="searchQuery = ''" class="btn btn-sm btn-ghost join-item">✕</button>
            </div>

            <div class="flex gap-2">
              <select v-model="selectedClan" class="select select-bordered select-xs flex-1">
                <option value="">All Clans</option>
                <option v-for="clan in clans" :key="clan">{{ clan }}</option>
              </select>
              <select v-model="selectedGrade" class="select select-bordered select-xs w-20">
                <option value="">Grade</option>
                <option v-for="grade in grades" :key="grade">G{{ grade }}</option>
              </select>
              <select v-model="selectedTrigger" class="select select-bordered select-xs w-24">
                <option value="">Trigger</option>
                <option value="Heal">Heal</option>
                <option value="Critical">Crit</option>
                <option value="Draw">Draw</option>
                <option value="Stand">Stand</option>
                <option value="Front">Front</option>
                <option value="None">None</option>
              </select>
            </div>

            <!-- Quick filter buttons -->
            <div class="flex flex-wrap gap-1">
              <button v-for="g in ['0', '1', '2', '3']" :key="g" @click="selectedGrade = selectedGrade === g ? '' : g"
                :class="['btn btn-xs', selectedGrade === g ? 'btn-primary' : 'btn-ghost']">
                G{{ g }}
              </button>
              <button @click="resetFilters" class="btn btn-xs btn-ghost text-error">
                Reset
              </button>
            </div>
          </div>

          <!-- Cards Grid with count badge -->
          <div
            class="grid grid-cols-4 sm:grid-cols-5 md:grid-cols-6 lg:grid-cols-5 gap-1 max-h-[55vh] overflow-y-auto p-1">
            <div v-for="card in filteredCards" :key="card.id" class="relative cursor-pointer group"
              @click="addCardToDeck(card)">
              <img :src="card.imageurlen || card.imageurljp" :alt="card.name"
                class="rounded-lg shadow-md w-full hover:scale-105 transition-transform duration-200" loading="lazy" />
              <!-- Count badge -->
              <div v-if="getCardCountInDeck(card.name) > 0"
                class="absolute top-0 right-0 badge badge-primary badge-sm font-bold">
                {{ getCardCountInDeck(card.name) }}
              </div>
              <!-- Hover overlay -->
              <div
                class="absolute inset-0 bg-primary/30 opacity-0 group-hover:opacity-100 rounded-lg transition-opacity flex items-center justify-center">
                <span class="btn btn-circle btn-xs btn-primary shadow-lg">+</span>
              </div>
              <!-- Max indicator -->
              <div v-if="getCardCountInDeck(card.name) >= 4"
                class="absolute inset-0 bg-error/50 rounded-lg flex items-center justify-center">
                <span class="badge badge-error text-xs">MAX</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right side - Deck -->
      <div class="w-full lg:w-1/2 card bg-base-200 shadow-xl" id="Deck">
        <div class="card-body p-4">
          <!-- Deck Header with progress -->
          <div class="flex items-center justify-between mb-2">
            <h2 class="card-title text-lg">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-secondary" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
              My Deck
            </h2>
            <div class="flex items-center gap-2">
              <span
                :class="['badge badge-lg', nonGUnitCount === 50 ? 'badge-success' : nonGUnitCount > 50 ? 'badge-error' : 'badge-primary']">
                {{ nonGUnitCount }}/50
              </span>
              <button v-if="deck.length > 0" @click="clearDeck" class="btn btn-xs btn-ghost text-error"
                title="Clear deck">
                🗑️
              </button>
            </div>
          </div>

          <!-- Progress bar -->
          <progress class="progress progress-primary w-full h-2 mb-3" :value="nonGUnitCount" max="50"></progress>

          <!-- Compact Stats -->
          <div class="grid grid-cols-2 gap-2 mb-3">
            <!-- Grades -->
            <div class="bg-base-300 rounded-lg p-2 text-sm">
              <div class="font-semibold mb-1 text-xs opacity-70">GRADES</div>
              <div class="flex flex-wrap gap-x-3 gap-y-0.5">
                <span class="text-purple-400">G0: {{ countGrade0 }}</span>
                <span class="text-green-400">G1: {{ countGrade1 }}</span>
                <span class="text-orange-400">G2: {{ countGrade2 }}</span>
                <span class="text-amber-600">G3: {{ countGrade3 }}</span>
              </div>
            </div>

            <!-- Triggers -->
            <div class="bg-base-300 rounded-lg p-2 text-sm">
              <div class="font-semibold mb-1 text-xs opacity-70">
                TRIGGERS
                <span :class="['badge badge-xs ml-1', triggerUnitCount === 16 ? 'badge-success' : 'badge-warning']">
                  {{ triggerUnitCount }}/16
                </span>
              </div>
              <div class="flex flex-wrap gap-x-2 gap-y-0.5 text-xs">
                <span class="text-green-400">H:{{ healTriggerCount }}</span>
                <span class="text-yellow-400">C:{{ criticalTriggerCount }}</span>
                <span class="text-red-400">D:{{ drawTriggerCount }}</span>
                <span class="text-blue-400">S:{{ standTriggerCount }}</span>
                <span class="text-pink-400">F:{{ frontTriggerCount }}</span>
              </div>
            </div>
          </div>

          <!-- Deck validation alerts -->
          <div v-if="deckValidation.length > 0" class="mb-2">
            <div v-for="(msg, i) in deckValidation" :key="i"
              :class="['alert alert-sm py-1 px-2 text-xs', msg.type === 'error' ? 'alert-error' : 'alert-warning']">
              {{ msg.text }}
            </div>
          </div>

          <!-- Action buttons -->
          <div class="flex gap-2 mb-3">
            <button @click="sortDeck" class="btn btn-xs btn-outline btn-primary flex-1">
              ↕️ Sort by Grade
            </button>
            <button @click="sortDeckByName" class="btn btn-xs btn-outline flex-1">
              🔤 Sort by Name
            </button>
          </div>

          <!-- Grouped Deck Display -->
          <div class="max-h-[45vh] overflow-y-auto space-y-2">
            <!-- Group by unique cards -->
            <div v-for="group in groupedDeck" :key="group.card.id"
              class="flex items-center gap-2 bg-base-300 rounded-lg p-1 hover:bg-base-100 transition-colors">
              <!-- Card image -->
              <div class="relative w-12 flex-shrink-0">
                <img :src="group.card.imageurlen || group.card.imageurljp" :alt="group.card.name"
                  class="rounded shadow w-full cursor-pointer" @click="showCardPreview(group.card)" />
              </div>

              <!-- Card info -->
              <div class="flex-1 min-w-0">
                <p class="text-sm font-medium truncate">{{ group.card.name }}</p>
                <div class="flex gap-1">
                  <span class="badge badge-xs badge-outline">G{{ group.card.grade }}</span>
                  <span v-if="group.card.triggereffect" class="badge badge-xs badge-warning">
                    {{ group.card.triggereffect }}
                  </span>
                </div>
              </div>

              <!-- Quantity controls -->
              <div class="flex items-center gap-1">
                <button @click="removeOneCard(group.card)" class="btn btn-circle btn-xs btn-ghost text-error">
                  −
                </button>
                <span class="w-6 text-center font-bold">{{ group.count }}</span>
                <button @click="addCardToDeck(group.card)" :disabled="group.count >= 4 || !canAddCard(group.card)"
                  class="btn btn-circle btn-xs btn-ghost text-success">
                  +
                </button>
              </div>
            </div>

            <!-- Empty state -->
            <div v-if="deck.length === 0" class="text-center py-8 opacity-50">
              <div class="text-4xl mb-2">🃏</div>
              <p>Click cards from the library to add them</p>
            </div>
          </div>

          <!-- Bottom action -->
          <div class="mt-3 flex gap-2">
            <button @click="preparePreview" :disabled="deck.length === 0" class="btn btn-success flex-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
              Preview & Upload
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Floating scroll button for mobile -->
    <button @click="scrollToDeck" class="btn btn-circle btn-primary fixed bottom-4 right-4 shadow-lg lg:hidden z-50">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
      </svg>
    </button>

    <!-- Card Preview Modal -->
    <dialog id="card_preview_modal" class="modal">
      <div class="modal-box max-w-sm">
        <form method="dialog">
          <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
        </form>
        <div v-if="previewCard">
          <img :src="previewCard.imageurlen || previewCard.imageurljp" :alt="previewCard.name"
            class="w-full rounded-lg mb-4" />
          <h3 class="font-bold text-lg">{{ previewCard.name }}</h3>
          <div class="flex flex-wrap gap-2 my-2">
            <span class="badge badge-primary">{{ previewCard.clan }}</span>
            <span class="badge badge-secondary">Grade {{ previewCard.grade }}</span>
            <span v-if="previewCard.triggereffect" class="badge badge-warning">{{ previewCard.triggereffect }}</span>
          </div>
          <p v-if="previewCard.effect" class="text-sm opacity-70">{{ previewCard.effect }}</p>
          <div class="modal-action">
            <button @click="addCardToDeck(previewCard); closePreview()"
              :disabled="getCardCountInDeck(previewCard.name) >= 4 || !canAddCard(previewCard)"
              class="btn btn-primary btn-sm">
              Add to Deck
            </button>
            <form method="dialog">
              <button class="btn btn-ghost btn-sm">Close</button>
            </form>
          </div>
        </div>
      </div>
      <form method="dialog" class="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>

    <!-- Deck Preview/Upload Modal -->
    <dialog id="preview_modal" class="modal">
      <div class="modal-box max-w-4xl">
        <form method="dialog">
          <button class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
        </form>
        <h3 class="font-bold text-lg mb-4">
          <span class="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            🎴 Deck Preview
          </span>
        </h3>

        <div class="form-control mb-4">
          <label class="label py-1">
            <span class="label-text">Deck Name</span>
          </label>
          <input type="text" v-model="deckName" placeholder="Enter deck name..."
            class="input input-bordered input-primary" />
        </div>

        <!-- Visual deck preview -->
        <div id="deck-preview" class="grid grid-cols-8 sm:grid-cols-10 gap-1 mb-4 max-h-[50vh] overflow-y-auto">
          <div v-for="(card, index) in sortedDeckForPreview" :key="index">
            <img :src="card.imageurlen || card.imageurljp" :alt="card.name" class="rounded w-full" loading="lazy" />
          </div>
        </div>

        <div class="modal-action">
          <form method="dialog">
            <button class="btn btn-ghost">Cancel</button>
          </form>
          <button @click="uploadDeck" :disabled="isLoading || !deckName.trim()" class="btn btn-primary">
            <span v-if="isLoading" class="loading loading-spinner loading-sm"></span>
            <span v-else>
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 inline mr-1" fill="none" viewBox="0 0 24 24"
                stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" />
              </svg>
              Upload
            </span>
          </button>
        </div>
      </div>
      <form method="dialog" class="modal-backdrop">
        <button>close</button>
      </form>
    </dialog>

    <!-- Toast Notifications -->
    <div class="toast toast-end z-50">
      <div v-if="toastMessage"
        :class="['alert', toastType === 'error' ? 'alert-error' : toastType === 'success' ? 'alert-success' : 'alert-warning']">
        <span>{{ toastMessage }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue';
import axios from 'axios';
import allCards from '/public/data/all_cards.json';
import { useRouter } from 'vue-router';

const router = useRouter();
const selectedClan = ref('');
const selectedGrade = ref('');
const selectedTrigger = ref('');
const searchQuery = ref('');
const clans = ref([]);
const grades = ref([]);
const filteredCards = ref([]);
const deck = ref([]);
const deckName = ref('');
const isLoading = ref(false);
const toastMessage = ref('');
const toastType = ref('warning');
const previewCard = ref(null);

onMounted(() => {
  const uniqueClans = [...new Set(allCards.map(card => card.clan))].sort();
  clans.value = uniqueClans;
  const uniqueGrades = [...new Set(allCards.map(card => card.grade))].sort();
  grades.value = uniqueGrades;
  filterCards();
});

// Toast notifications
const showToast = (message, type = 'warning') => {
  toastMessage.value = message;
  toastType.value = type;
  setTimeout(() => {
    toastMessage.value = '';
  }, 2500);
};

// Filter cards
const filterCards = () => {
  let filtered = allCards;
  if (selectedClan.value) {
    filtered = filtered.filter(card => card.clan === selectedClan.value);
  }
  if (selectedGrade.value) {
    filtered = filtered.filter(card => card.grade === selectedGrade.value);
  }
  if (selectedTrigger.value) {
    if (selectedTrigger.value === 'None') {
      filtered = filtered.filter(card => !card.triggereffect);
    } else {
      filtered = filtered.filter(card => card.triggereffect && card.triggereffect.includes(selectedTrigger.value));
    }
  }
  if (searchQuery.value) {
    filtered = filtered.filter(card => card.name.toLowerCase().includes(searchQuery.value.toLowerCase()));
  }
  filteredCards.value = filtered;
};

const resetFilters = () => {
  selectedClan.value = '';
  selectedGrade.value = '';
  selectedTrigger.value = '';
  searchQuery.value = '';
};

// Card count helpers
const getCardCountInDeck = (cardName) => {
  return deck.value.filter(c => c.name === cardName).length;
};

const canAddCard = (card) => {
  const nonGUnitCount = deck.value.filter(c => c.cardtype !== 'G Unit').length;
  const cardCount = getCardCountInDeck(card.name);
  const triggerCount = deck.value.filter(c => c.cardtype === 'Trigger Unit').length;
  const healCount = deck.value.filter(c => c.triggereffect && c.triggereffect.includes('Heal')).length;

  if (cardCount >= 4) return false;
  if (card.cardtype !== 'G Unit' && nonGUnitCount >= 50) return false;
  if (card.cardtype === 'Trigger Unit' && triggerCount >= 16) return false;
  if (card.triggereffect && card.triggereffect.includes('Heal') && healCount >= 4) return false;

  return true;
};

// Add card to deck
const addCardToDeck = (card) => {
  if (!canAddCard(card)) {
    const cardCount = getCardCountInDeck(card.name);
    if (cardCount >= 4) {
      showToast(`${card.name} is at max (4)`, 'warning');
    } else if (deck.value.filter(c => c.cardtype === 'Trigger Unit').length >= 16) {
      showToast('Trigger limit reached (16)', 'error');
    } else if (deck.value.filter(c => c.triggereffect && c.triggereffect.includes('Heal')).length >= 4) {
      showToast('Heal trigger limit reached (4)', 'error');
    } else {
      showToast('Deck is full (50 cards)', 'error');
    }
    return;
  }

  deck.value.push(card);
  showToast(`Added ${card.name}`, 'success');
};

// Remove one card
const removeOneCard = (card) => {
  const index = deck.value.findIndex(c => c.id === card.id);
  if (index !== -1) {
    deck.value.splice(index, 1);
  }
};

// Clear deck
const clearDeck = () => {
  if (confirm('Clear entire deck?')) {
    deck.value = [];
    showToast('Deck cleared', 'warning');
  }
};

// Sort functions
const sortDeck = () => {
  deck.value.sort((a, b) => {
    if (a.grade === b.grade) {
      return a.name.localeCompare(b.name);
    }
    return a.grade - b.grade;
  });
  showToast('Sorted by grade', 'success');
};

const sortDeckByName = () => {
  deck.value.sort((a, b) => a.name.localeCompare(b.name));
  showToast('Sorted by name', 'success');
};

// Group deck for compact display
const groupedDeck = computed(() => {
  const groups = {};
  deck.value.forEach(card => {
    if (!groups[card.id]) {
      groups[card.id] = { card, count: 0 };
    }
    groups[card.id].count++;
  });
  return Object.values(groups).sort((a, b) => {
    if (a.card.grade === b.card.grade) {
      return a.card.name.localeCompare(b.card.name);
    }
    return a.card.grade - b.card.grade;
  });
});

// Sorted deck for preview
const sortedDeckForPreview = computed(() => {
  return [...deck.value].sort((a, b) => {
    if (a.grade === b.grade) return a.name.localeCompare(b.name);
    return a.grade - b.grade;
  });
});

// Deck validation
const deckValidation = computed(() => {
  const messages = [];
  if (nonGUnitCount.value > 50) {
    messages.push({ type: 'error', text: 'Deck exceeds 50 cards!' });
  }
  if (triggerUnitCount.value > 16) {
    messages.push({ type: 'error', text: 'Too many triggers!' });
  }
  if (nonGUnitCount.value > 0 && nonGUnitCount.value < 50) {
    messages.push({ type: 'warning', text: `Need ${50 - nonGUnitCount.value} more cards` });
  }
  if (triggerUnitCount.value < 16 && nonGUnitCount.value > 0) {
    messages.push({ type: 'warning', text: `Need ${16 - triggerUnitCount.value} more triggers` });
  }
  return messages;
});

// Card preview
const showCardPreview = (card) => {
  previewCard.value = card;
  document.getElementById('card_preview_modal').showModal();
};

const closePreview = () => {
  document.getElementById('card_preview_modal').close();
};

// Preview/Upload
const preparePreview = () => {
  document.getElementById('preview_modal').showModal();
};

const uploadDeck = () => {
  isLoading.value = true;
  const deckData = {
    name: deckName.value,
    cards: deck.value.map(card => card.id)
  };

  axios.post('/api/macros/s/AKfycbwxQSc_FdDlkU81payets_7BpBtqs61yDI_yLcIe_of7sArG3fhg-UZps8VIW3s486FvA/exec', deckData, {
    headers: { 'Content-Type': 'application/json' }
  })
    .then(response => {
      const text = response.data;
      if (text === "Deck data saved..." || typeof text === 'object') {
        document.getElementById('preview_modal').close();
        showToast('Deck uploaded successfully!', 'success');
        setTimeout(() => router.push('/'), 1000);
      } else {
        showToast('Deck uploaded!', 'success');
      }
    })
    .catch((error) => {
      showToast(`Upload failed: ${error.message}`, 'error');
    })
    .finally(() => {
      isLoading.value = false;
    });
};

const scrollToDeck = () => {
  document.getElementById('Deck')?.scrollIntoView({ behavior: 'smooth' });
};

// Computed stats
const nonGUnitCount = computed(() => deck.value.filter(c => c.cardtype !== 'G Unit').length);
const triggerUnitCount = computed(() => deck.value.filter(c => c.cardtype === 'Trigger Unit').length);
const healTriggerCount = computed(() => deck.value.filter(c => c.triggereffect?.includes('Heal')).length);
const criticalTriggerCount = computed(() => deck.value.filter(c => c.triggereffect?.includes('Critical')).length);
const drawTriggerCount = computed(() => deck.value.filter(c => c.triggereffect?.includes('Draw')).length);
const standTriggerCount = computed(() => deck.value.filter(c => c.triggereffect?.includes('Stand')).length);
const frontTriggerCount = computed(() => deck.value.filter(c => c.triggereffect?.includes('Front')).length);
const countGrade0 = computed(() => deck.value.filter(c => c.grade === '0').length);
const countGrade1 = computed(() => deck.value.filter(c => c.grade === '1').length);
const countGrade2 = computed(() => deck.value.filter(c => c.grade === '2').length);
const countGrade3 = computed(() => deck.value.filter(c => c.grade === '3').length);
const countGrade4 = computed(() => deck.value.filter(c => c.grade === '4').length);

watch([selectedClan, selectedGrade, selectedTrigger, searchQuery], filterCards);
</script>

<style scoped></style>