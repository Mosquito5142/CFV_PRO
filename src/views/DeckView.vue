<template>
  <div class="container mx-auto p-4">
    <h1 class="text-3xl font-bold mb-6 text-center">Deck View</h1>
    <div class="flex flex-col md:flex-row">
      <!-- Left side -->
      <div class="w-full md:w-1/2 p-4 border-b md:border-b-0 md:border-r border-gray-300" @dragover.prevent
        @drop="onDropFromDeck">
        <h2 class="text-2xl font-semibold mb-4">Card List</h2>
        <div class="mb-6">
          <label for="clan-select" class="block text-lg font-medium mb-2">Select Clan:</label>
          <select id="clan-select" v-model="selectedClan" class="w-full p-2 border border-gray-300 rounded-md mb-4">
            <option value="">Select clan</option>
            <option v-for="clan in clans" :key="clan">{{ clan }}</option>
          </select>
          <div class="flex space-x-4 mb-4">
            <div class="flex-1">
              <label for="grade-select" class="block text-lg font-medium mb-2">Select Grade:</label>
              <select id="grade-select" v-model="selectedGrade" class="w-full p-2 border border-gray-300 rounded-md">
                <option value="">Select grade</option>
                <option v-for="grade in grades" :key="grade">{{ grade }}</option>
              </select>
            </div>
            <div class="flex-1">
              <label for="trigger-select" class="block text-lg font-medium mb-2">Select Trigger:</label>
              <select id="trigger-select" v-model="selectedTrigger"
                class="w-full p-2 border border-gray-300 rounded-md">
                <option value="">Select trigger</option>
                <option value="Heal">Heal</option>
                <option value="Critical">Critical</option>
                <option value="Draw">Draw</option>
                <option value="Stand">Stand</option>
                <option value="Front">Front</option>
                <option value="None">non trigger</option>
              </select>
            </div>
          </div>
          <input type="text" v-model="searchQuery" placeholder="Search by card name"
            class="w-full p-2 border border-gray-300 rounded-md" />
        </div>
        <div class="grid grid-cols-3 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-2">
          <div v-for="card in filteredCards" :key="card.id"
            class="border border-gray-300 rounded-lg p-2 text-center bg-white shadow-md cursor-pointer" draggable="true"
            @dragstart="onDragStart(card, 'cardList')" @click="onRightClick(card, 'cardList')">
            <img :src="card.imageurlen || card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md"
              loading="lazy" />
            <p class="mt-2 text-sm font-bold">{{ card.name }}</p>
          </div>
        </div>
      </div>
      <!-- Right side -->
      <div :class="['w-full md:w-1/2 p-4', isDraggingOver ? 'bg-blue-100' : '']" @dragover.prevent="onDragOver"
        @dragleave="onDragLeave" @drop="onDropFromCardList" id="Deck">
        <h2 class="text-2xl font-semibold mb-4">Deck : {{ nonGUnitCount }}</h2>
        <div class="mb-4 flex flex-wrap gap-2 text-lg">
          <div class="w-2/2">
            <h3 class="text-xl font-semibold mb-2">Grades</h3>
            <p class="text-lg" style="color: purple;">Grade 0: {{ countGrade0 }}</p>
            <p class="text-lg" style="color: green;">Grade 1: {{ countGrade1 }}</p>
            <p class="text-lg" style="color: orange;">Grade 2: {{ countGrade2 }}</p>
            <p class="text-lg" style="color: brown;">Grade 3: {{ countGrade3 }}</p>
            <p class="text-lg" style="color: red;">Grade 4: {{ countGrade4 }}</p>
          </div>
          <div class="w-2/2">
            <p class="text-xl font-semibold mb-2">Trigger : {{ triggerUnitCount }}</p>
            <p class="text-lg" style="color: green;">Heal : {{ healTriggerCount }}</p>
            <p class="text-lg" style="color: gold;">Critical: {{ criticalTriggerCount }}</p>
            <p class="text-lg" style="color: red;">Draw : {{ drawTriggerCount }}</p>
            <p class="text-lg" style="color: blue;">Stand : {{ standTriggerCount }}</p>
            <p class="text-lg" style="color: pink;">Front : {{ frontTriggerCount }}</p>
          </div>
          <div>

          </div>
        </div>
        <button @click="sortDeck" class="mb-4 p-2 bg-blue-500 text-white rounded-md">Sort Deck</button>
        <div class="grid grid-cols-6 sm:grid-cols-6 md:grid-cols-6 lg:grid-cols-6 gap-2">
          <div v-for="(card, index) in deck" :key="index"
            class="border border-gray-300 rounded-lg text-center bg-white shadow-md" draggable="true"
            @dragstart="onDragStart(card, 'deck', index)" @click="onRightClick(card, 'deck', index)">
            <img :src="card.imageurlen || card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md"
              loading="lazy" />
          </div>
        </div>
        <div>
          <button @click="preparePreview"
            class="text-lg font-semibold mt-4 p-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition duration-300">Preview</button>
        </div>
      </div>
    </div>

    <!-- Floating Button -->
    <button @click="scrollToDeck"
      class="fixed bottom-4 right-4 bg-blue-500 text-white rounded-full p-3 shadow-lg hover:bg-blue-600 transition duration-300">
      ↓
    </button>

    <!-- Modal -->
    <div v-show="showPreview"
      class="fixed inset-0 bg-gray-800 bg-opacity-75 flex items-center justify-center overflow-auto">
      <div class="bg-white p-4 rounded-lg max-w-4xl w-full max-h-full overflow-auto">
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-2xl font-semibold">Deck Preview</h2>
          <button @click="showPreview = false" class="text-lg font-semibold">Close</button>
        </div>
        <div class="mb-4">
          <label for="deck-name" class="block text-lg font-medium mb-2">Deck Name:</label>
          <input type="text" id="deck-name" v-model="deckName" class="w-full border border-gray-300 rounded-md" />
        </div>
        <div id="deck-preview" class="grid grid-cols-8 sm:grid-cols-8 md:grid-cols-8 lg:grid-cols-10 gap-1">
          <div v-for="(card, index) in deck" :key="index"
            class="border border-gray-300 rounded-lg text-center bg-white shadow-md">
            <img :src="card.imageurlen || card.imageurljp" :alt="card.name" class="max-w-full h-auto rounded-md"
              loading="lazy" />
          </div>
        </div>
        <div class="mt-4 text-right">
          <button @click="uploadDeck" :disabled="isLoading"
            class="p-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition duration-300">
            <span v-if="isLoading">Uploading...</span>
            <span v-else>Upload</span>
          </button>
        </div>
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
const cards = ref([]);
const filteredCards = ref([]);
const deck = ref([]);
const deckName = ref('');
const isDraggingOver = ref(false);
const showPreview = ref(false);
const isLoading = ref(false); // เพิ่มตัวแปร isLoading

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

const onDragStart = (card, source, index) => {
  event.dataTransfer.setData('card', JSON.stringify(card));
  event.dataTransfer.setData('source', source);
  event.dataTransfer.setData('index', index);
};

const onDropFromCardList = (event) => {
  const card = JSON.parse(event.dataTransfer.getData('card'));
  const source = event.dataTransfer.getData('source');
  const nonGUnitCount = deck.value.filter(c => c.cardtype !== 'G Unit').length;

  if (source === 'cardList') {
    if (card.cardtype === 'G Unit' || nonGUnitCount < 50) {
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
  const nonGUnitCount = deck.value.filter(c => c.cardtype !== 'G Unit').length;

  if (source === 'cardList') {
    if (card.cardtype === 'G Unit' || nonGUnitCount < 50) {
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
    deck.value.splice(index, 1);
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

const preparePreview = () => {
  const deckPreview = document.getElementById('deck-preview');
  const images = deckPreview.querySelectorAll('img');

  // ตรวจสอบว่ารูปภาพทั้งหมดโหลดเสร็จหรือไม่
  const imagePromises = Array.from(images).map(img => {
    return new Promise((resolve) => {
      if (img.complete) {
        resolve();
      } else {
        img.onload = resolve;
        img.onerror = resolve; // กรณีโหลดภาพล้มเหลว
      }
    });
  });

  // รอให้รูปภาพทั้งหมดโหลดเสร็จ
  Promise.all(imagePromises).then(() => {
    showPreview.value = true;
  }).catch(error => {
    console.error('Error loading images:', error);
  });
};

const uploadDeck = () => {
  isLoading.value = true; // เริ่มการโหลด
  const deckData = {
    name: deckName.value,
    cards: deck.value.map(card => card.id)
  };

  console.log('Sending deck data:', deckData); // Log the data being sent

  axios.post('/api/macros/s/AKfycbwxQSc_FdDlkU81payets_7BpBtqs61yDI_yLcIe_of7sArG3fhg-UZps8VIW3s486FvA/exec', deckData, {
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(response => {
      const text = response.data;
      try {
        // Check if the response text is valid JSON
        const data = JSON.parse(text);
        console.log('Success:', data);
        alert('Deck uploaded successfully!');
        router.push('/'); // Redirect to the home path
      } catch (error) {
        // If parsing fails, handle the plain text response
        console.error('Error parsing JSON:', error);
        console.log('Response text:', text);
        if (text === "Deck data saved...") {
          alert('Deck uploaded successfully!');
          router.push('/'); // Redirect to the home path
        } else {
          alert('Deck uploaded, but received invalid JSON response.');
        }
      }
    })
    .catch((error) => {
      console.error('Error:', error);
      alert(`Failed to upload deck. Error: ${error.message}`);
    })
    .finally(() => {
      isLoading.value = false; // สิ้นสุดการโหลด
    });
};

const scrollToDeck = () => {
  const deckElement = document.getElementById('Deck');
  if (deckElement) {
    deckElement.scrollIntoView({ behavior: 'smooth' });
  }
};

const nonGUnitCount = computed(() => deck.value.filter(card => card.cardtype !== 'G Unit').length);
const triggerUnitCount = computed(() => deck.value.filter(card => card.cardtype === 'Trigger Unit').length);
const healTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Heal')).length);
const criticalTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Critical')).length);
const drawTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Draw')).length);
const standTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Stand')).length);
const frontTriggerCount = computed(() => deck.value.filter(card => card.triggereffect && card.triggereffect.includes('Front')).length);

const countGrade0 = computed(() => deck.value.filter(card => card.grade === '0').length);
const countGrade1 = computed(() => deck.value.filter(card => card.grade === '1').length);
const countGrade2 = computed(() => deck.value.filter(card => card.grade === '2').length);
const countGrade3 = computed(() => deck.value.filter(card => card.grade === '3').length);
const countGrade4 = computed(() => deck.value.filter(card => card.grade === '4').length);

watch([selectedClan, selectedGrade, selectedTrigger, searchQuery], filterCards);
</script>


<style scoped>
/* เพิ่มการตั้งค่า CSS เพื่อให้ modal สามารถเลื่อนและควบคุมได้ง่ายขึ้น */
.fixed.inset-0 {
  overflow: auto;
}

.max-h-full {
  max-height: 90vh;
}

/* Floating Button */
.fixed.bottom-4.right-4 {
  position: fixed;
  bottom: 1rem;
  right: 1rem;
  z-index: 1000;
}
</style>