<template>
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-center">Deck View</h1>
      <div class="flex">
        <!-- Left side -->
        <div
          class="w-1/2 p-4 border-r border-gray-300"
          @dragover.prevent
          @drop="onDropFromDeck"
        >
          <h2 class="text-2xl font-semibold mb-4">Card List</h2>
          <div class="mb-6">
            <label for="clan-select" class="block text-lg font-medium mb-2">Select Clan:</label>
            <select id="clan-select" v-model="selectedClan" class="w-full p-2 border border-gray-300 rounded-md mb-4">
              <option value="">กรุณาเลือก</option>
              <option v-for="clan in clans" :key="clan">{{ clan }}</option>
            </select>
            <label for="grade-select" class="block text-lg font-medium mb-2">Select Grade:</label>
            <select id="grade-select" v-model="selectedGrade" class="w-full p-2 border border-gray-300 rounded-md mb-4">
              <option value="">กรุณาเลือก</option>
              <option v-for="grade in grades" :key="grade">{{ grade }}</option>
            </select>
            <input type="text" v-model="searchQuery" placeholder="Search by card name" class="w-full p-2 border border-gray-300 rounded-md"/>
          </div>
          <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            <div
              v-for="card in filteredCards"
              :key="card.id"
              class="border border-gray-300 rounded-lg p-4 text-center bg-white shadow-md cursor-pointer"
              draggable="true"
              @dragstart="onDragStart(card, 'cardList')"
              @contextmenu.prevent="onRightClick(card, 'cardList')"
            >
              <img :src="card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md" loading="lazy"/>
              <p class="mt-2 text-lg font-bold">{{ card.name }}</p>
            </div>
          </div>
        </div>
        <!-- Right side -->
        <div
          :class="['w-1/2 p-4', isDraggingOver ? 'bg-blue-100' : '']"
          @dragover.prevent="onDragOver"
          @dragleave="onDragLeave"
          @drop="onDropFromCardList"
        >
          <h2 class="text-2xl font-semibold mb-4">Deck</h2>
          <div class="mb-4">
            <p class="text-lg">Total Cards: {{ nonGUnitCount }}</p>
            <p class="text-lg">Trigger Units: {{ triggerUnitCount }}</p>
            <p class="text-lg">Heal Triggers: {{ healTriggerCount }}</p>
            <p class="text-lg">Critical Triggers: {{ criticalTriggerCount }}</p>
            <p class="text-lg">Draw Triggers: {{ drawTriggerCount }}</p>
            <p class="text-lg">Stand Triggers: {{ standTriggerCount }}</p>
            <p class="text-lg">Front Triggers: {{ frontTriggerCount }}</p>
          </div>
          <button @click="sortDeck" class="mb-4 p-2 bg-blue-500 text-white rounded-md">Sort Deck</button>
          <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
            <div
              v-for="(card, index) in deck"
              :key="index"
              class="border border-gray-300 rounded-lg p-4 text-center bg-white shadow-md"
              draggable="true"
              @dragstart="onDragStart(card, 'deck', index)"
              @contextmenu.prevent="onRightClick(card, 'deck', index)"
            >
              <img :src="card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md" loading="lazy"/>
              <p class="mt-2 text-lg font-bold">{{ card.name }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, watch, computed } from 'vue';
  import allCards from '/public/data/all_cards.json';
  
  const selectedClan = ref('');
  const selectedGrade = ref('');
  const searchQuery = ref('');
  const clans = ref([]);
  const grades = ref([]);
  const cards = ref([]);
  const filteredCards = ref([]);
  const deck = ref([]);
  const isDraggingOver = ref(false);
  
  onMounted(() => {
    const uniqueClans = [...new Set(allCards.map(card => card.clan))].sort();
    clans.value = uniqueClans;
    const uniqueGrades = [...new Set(allCards.map(card => card.grade))].sort();
    grades.value = uniqueGrades;
  });
  
  const filterCards = () => {
    let filtered = allCards;
    if (selectedClan.value) {
      filtered = filtered.filter(card => card.clan === selectedClan.value);
    }
    if (selectedGrade.value) {
      filtered = filtered.filter(card => card.grade === selectedGrade.value);
    }
    if (searchQuery.value) {
      filtered = filtered.filter(card => card.name.toLowerCase().includes(searchQuery.value.toLowerCase()));
    }
    filteredCards.value = filtered;
  };
  
  const onDragStart = (card, source, index) => {
    event.dataTransfer.setData('card', JSON.stringify(card));
    event.dataTransfer.setData('source', source);
    event.dataTransfer.setData('index', index);
  };
  
  const onDropFromCardList = (event) => {
    const card = JSON.parse(event.dataTransfer.getData('card'));
    const source = event.dataTransfer.getData('source');
    if (source === 'cardList') {
      const nonGUnitCount = deck.value.filter(c => c.cardtype !== 'G Unit').length;
      if (nonGUnitCount < 50) {
        const cardCount = deck.value.filter(c => c.name === card.name).length;
        const triggerUnitCount = deck.value.filter(c => c.cardtype === 'Trigger Unit').length;
        const healTriggerCount = deck.value.filter(c => c.triggereffect && c.triggereffect.includes('Heal')).length;
        if (cardCount < 4) {
          if (card.cardtype === 'Trigger Unit' && triggerUnitCount >= 16) {
            alert('การ์ดประเภท Trigger Unit สามารถใส่ได้สูงสุด 16 ใบ');
          } else if (card.triggereffect && card.triggereffect.includes('Heal') && healTriggerCount >= 4) {
            alert('การ์ดที่มี Heal สามารถใส่ได้สูงสุด 4 ใบ');
          } else {
            deck.value.push(card);
          }
        } else {
          alert('การ์ดชื่อเดียวกันสามารถใส่ได้สูงสุด 4 ใบ');
        }
      } else {
        alert('เด็คสามารถใส่การ์ดได้สูงสุด 50 ใบ');
      }
    } else if (source === 'deck') {
      const index = event.dataTransfer.getData('index');
      deck.value.splice(index, 1);
    }
    isDraggingOver.value = false;
  };
  
  const onDropFromDeck = (event) => {
    const card = JSON.parse(event.dataTransfer.getData('card'));
    const index = event.dataTransfer.getData('index');
    deck.value.splice(index, 1);
  };
  
  const onRightClick = (card, source, index) => {
    if (source === 'cardList') {
      onDropFromCardList({ dataTransfer: { getData: () => JSON.stringify(card) } });
    } else if (source === 'deck') {
      onDropFromDeck({ dataTransfer: { getData: () => JSON.stringify(card), getData: () => index } });
    }
  };
  
  const onDragOver = () => {
    isDraggingOver.value = true;
  };
  
  const onDragLeave = () => {
    isDraggingOver.value = false;
  };
  
  const sortDeck = () => {
    deck.value.sort((a, b) => {
      if (a.grade === b.grade) {
        return a.name.localeCompare(b.name);
      }
      return a.grade - b.grade;
    });
  };
  
  const nonGUnitCount = computed(() => deck.value.filter(card => card.cardtype !== 'G Unit').length);
  const triggerUnitCount = computed(() => deck.value.filter(card => card.cardtype === 'Trigger Unit').length);
  const healTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Heal')).length);
  const criticalTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Critical')).length);
  const drawTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Draw')).length);
  const standTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Stand')).length);
  const frontTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Front')).length);
  
  watch([selectedClan, selectedGrade, searchQuery], filterCards);
  </script>
  
  <style scoped>
  </style>