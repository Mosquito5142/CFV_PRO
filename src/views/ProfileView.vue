<script setup>
import { ref } from 'vue';

// Mock data - จะเชื่อมกับ backend จริงในอนาคต
const user = ref({
  name: 'Vanguard Fighter',
  avatar: null,
  level: 5,
  exp: 450,
  maxExp: 1000,
  joinDate: 'ธันวาคม 2024'
});

const stats = ref({
  decksCreated: 3,
  totalCards: 150,
  favoriteCards: 24
});

const recentDecks = ref([
  { id: 1, name: 'Royal Paladin', cardCount: 50, lastModified: '2 ชม. ที่แล้ว' },
  { id: 2, name: 'Shadow Paladin', cardCount: 48, lastModified: '1 วัน ที่แล้ว' },
  { id: 3, name: 'Kagero Rush', cardCount: 50, lastModified: '3 วัน ที่แล้ว' },
]);

const achievements = ref([
  { icon: '🏆', name: 'First Deck', description: 'สร้างเด็คแรก', unlocked: true },
  { icon: '⭐', name: 'Collector', description: 'เก็บการ์ด 100 ใบ', unlocked: true },
  { icon: '🔥', name: 'Clan Master', description: 'สร้างเด็คครบ 5 Clan', unlocked: false },
  { icon: '💎', name: 'Perfect Deck', description: 'สร้างเด็คที่มี 50 ใบพอดี', unlocked: true },
]);
</script>

<template>
  <div class="container mx-auto px-4 py-8">
    <!-- Profile Header -->
    <div class="card bg-base-200 shadow-xl mb-8 overflow-hidden">
      <!-- Cover Image -->
      <div class="h-32 bg-gradient-to-r from-primary via-secondary to-accent"></div>
      
      <div class="card-body -mt-16">
        <div class="flex flex-col md:flex-row items-center md:items-end gap-4">
          <!-- Avatar -->
          <div class="avatar placeholder">
            <div class="bg-neutral text-neutral-content rounded-full w-24 ring ring-primary ring-offset-base-100 ring-offset-2">
              <span class="text-3xl">🎴</span>
            </div>
          </div>
          
          <div class="text-center md:text-left flex-1">
            <h1 class="text-2xl font-bold">{{ user.name }}</h1>
            <p class="text-base-content/60">เข้าร่วมเมื่อ {{ user.joinDate }}</p>
            
            <!-- Level Progress -->
            <div class="flex items-center gap-2 mt-2">
              <div class="badge badge-primary">Lv. {{ user.level }}</div>
              <progress 
                class="progress progress-primary w-32" 
                :value="user.exp" 
                :max="user.maxExp"
              ></progress>
              <span class="text-sm text-base-content/60">{{ user.exp }}/{{ user.maxExp }} EXP</span>
            </div>
          </div>
          
          <button class="btn btn-outline btn-primary btn-sm">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
            แก้ไขโปรไฟล์
          </button>
        </div>
      </div>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
      <div class="stat bg-base-200 rounded-box shadow">
        <div class="stat-figure text-primary">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
        </div>
        <div class="stat-title">เด็คที่สร้าง</div>
        <div class="stat-value text-primary">{{ stats.decksCreated }}</div>
      </div>
      
      <div class="stat bg-base-200 rounded-box shadow">
        <div class="stat-figure text-secondary">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
          </svg>
        </div>
        <div class="stat-title">การ์ดทั้งหมด</div>
        <div class="stat-value text-secondary">{{ stats.totalCards }}</div>
      </div>
      
      <div class="stat bg-base-200 rounded-box shadow">
        <div class="stat-figure text-accent">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
          </svg>
        </div>
        <div class="stat-title">การ์ดโปรด</div>
        <div class="stat-value text-accent">{{ stats.favoriteCards }}</div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- Recent Decks -->
      <div class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-primary" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
            เด็คล่าสุด
          </h2>
          
          <div class="divide-y divide-base-300">
            <div 
              v-for="deck in recentDecks" 
              :key="deck.id"
              class="py-4 flex items-center justify-between hover:bg-base-300/50 px-2 rounded-lg transition-colors cursor-pointer"
            >
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-lg bg-gradient-to-br from-primary to-secondary flex items-center justify-center text-white font-bold">
                  {{ deck.name.charAt(0) }}
                </div>
                <div>
                  <p class="font-medium">{{ deck.name }}</p>
                  <p class="text-sm text-base-content/60">{{ deck.cardCount }} การ์ด</p>
                </div>
              </div>
              <div class="text-right">
                <span class="text-sm text-base-content/60">{{ deck.lastModified }}</span>
              </div>
            </div>
          </div>
          
          <div class="card-actions justify-center mt-4">
            <RouterLink to="/deck" class="btn btn-primary btn-sm btn-wide">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              สร้างเด็คใหม่
            </RouterLink>
          </div>
        </div>
      </div>

      <!-- Achievements -->
      <div class="card bg-base-200 shadow-xl">
        <div class="card-body">
          <h2 class="card-title">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-warning" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" />
            </svg>
            ความสำเร็จ
          </h2>
          
          <div class="grid grid-cols-2 gap-4">
            <div 
              v-for="achievement in achievements" 
              :key="achievement.name"
              :class="[
                'p-4 rounded-lg border-2 text-center transition-all',
                achievement.unlocked 
                  ? 'border-warning bg-warning/10' 
                  : 'border-base-300 opacity-50'
              ]"
            >
              <div class="text-3xl mb-2">{{ achievement.icon }}</div>
              <p class="font-medium text-sm">{{ achievement.name }}</p>
              <p class="text-xs text-base-content/60">{{ achievement.description }}</p>
              <div v-if="achievement.unlocked" class="badge badge-success badge-sm mt-2">Unlocked</div>
              <div v-else class="badge badge-ghost badge-sm mt-2">Locked</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>