<script setup>
import { RouterLink } from 'vue-router';
import { ref, onMounted } from 'vue';
import allCards from '/public/data/all_cards.json';

const totalCards = ref(0);
const totalClans = ref(0);
const recentCards = ref([]);

onMounted(() => {
  totalCards.value = allCards.length;
  totalClans.value = [...new Set(allCards.map(card => card.clan))].length;
  // Get 6 random cards for showcase
  const shuffled = [...allCards].sort(() => 0.5 - Math.random());
  recentCards.value = shuffled.slice(0, 6);
});
</script>

<template>
  <div class="min-h-screen">
    <!-- Hero Section -->
    <div class="hero min-h-[60vh] relative overflow-hidden">
      <!-- Animated Background -->
      <div class="absolute inset-0 bg-gradient-to-br from-primary/20 via-base-100 to-secondary/20"></div>
      <div class="absolute inset-0 opacity-20">
        <div class="absolute top-10 left-10 w-72 h-72 bg-primary rounded-full filter blur-3xl animate-pulse"></div>
        <div class="absolute bottom-10 right-10 w-96 h-96 bg-secondary rounded-full filter blur-3xl animate-pulse delay-1000"></div>
      </div>
      
      <div class="hero-content text-center relative z-10">
        <div class="max-w-2xl">
          <h1 class="text-5xl md:text-7xl font-bold mb-6">
            <span class="bg-gradient-to-r from-primary via-accent to-secondary bg-clip-text text-transparent">
              CFV Pro
            </span>
          </h1>
          <p class="text-xl md:text-2xl mb-4 text-base-content/80">
            🎴 Cardfight!! Vanguard Deck Builder
          </p>
          <p class="text-lg text-base-content/60 mb-8">
            สร้างเด็ค ค้นหาการ์ด และจัดการคอลเลคชันของคุณได้ง่ายๆ
          </p>
          <div class="flex flex-wrap justify-center gap-4">
            <RouterLink to="/deck" class="btn btn-primary btn-lg gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              สร้างเด็คใหม่
            </RouterLink>
            <RouterLink to="/clan" class="btn btn-outline btn-secondary btn-lg gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
              ค้นหาการ์ด
            </RouterLink>
          </div>
        </div>
      </div>
    </div>

    <!-- Stats Section -->
    <div class="container mx-auto px-4 -mt-16 relative z-20">
      <div class="stats stats-vertical lg:stats-horizontal shadow-xl bg-base-200 w-full">
        <div class="stat">
          <div class="stat-figure text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
          </div>
          <div class="stat-title">การ์ดทั้งหมด</div>
          <div class="stat-value text-primary">{{ totalCards.toLocaleString() }}</div>
          <div class="stat-desc">ในฐานข้อมูล</div>
        </div>
        
        <div class="stat">
          <div class="stat-figure text-secondary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
            </svg>
          </div>
          <div class="stat-title">จำนวน Clan</div>
          <div class="stat-value text-secondary">{{ totalClans }}</div>
          <div class="stat-desc">พร้อมให้เลือกใช้</div>
        </div>
        
        <div class="stat">
          <div class="stat-figure text-accent">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
          <div class="stat-title">Deck Builder</div>
          <div class="stat-value text-accent">Free</div>
          <div class="stat-desc">ใช้งานได้ฟรี!</div>
        </div>
      </div>
    </div>

    <!-- Features Section -->
    <div class="container mx-auto px-4 py-16">
      <h2 class="text-3xl font-bold text-center mb-12">
        <span class="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          ฟีเจอร์หลัก
        </span>
      </h2>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- Feature 1 -->
        <div class="card bg-base-200 shadow-xl hover:shadow-2xl transition-all duration-300 hover:-translate-y-2">
          <figure class="px-10 pt-10">
            <div class="w-16 h-16 rounded-full bg-primary/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-primary" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </figure>
          <div class="card-body items-center text-center">
            <h3 class="card-title">ค้นหาการ์ด</h3>
            <p class="text-base-content/70">ค้นหาการ์ดจาก Clan, Grade, หรือชื่อได้ง่ายๆ</p>
            <div class="card-actions">
              <RouterLink to="/clan" class="btn btn-primary btn-sm">ค้นหาเลย</RouterLink>
            </div>
          </div>
        </div>

        <!-- Feature 2 -->
        <div class="card bg-base-200 shadow-xl hover:shadow-2xl transition-all duration-300 hover:-translate-y-2">
          <figure class="px-10 pt-10">
            <div class="w-16 h-16 rounded-full bg-secondary/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-secondary" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
            </div>
          </figure>
          <div class="card-body items-center text-center">
            <h3 class="card-title">สร้างเด็ค</h3>
            <p class="text-base-content/70">ลากวางการ์ดเพื่อสร้างเด็คของคุณ พร้อมระบบตรวจสอบกฎ</p>
            <div class="card-actions">
              <RouterLink to="/deck" class="btn btn-secondary btn-sm">สร้างเด็ค</RouterLink>
            </div>
          </div>
        </div>

        <!-- Feature 3 -->
        <div class="card bg-base-200 shadow-xl hover:shadow-2xl transition-all duration-300 hover:-translate-y-2">
          <figure class="px-10 pt-10">
            <div class="w-16 h-16 rounded-full bg-accent/20 flex items-center justify-center">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-accent" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12" />
              </svg>
            </div>
          </figure>
          <div class="card-body items-center text-center">
            <h3 class="card-title">แชร์เด็ค</h3>
            <p class="text-base-content/70">อัพโหลดและแชร์เด็คของคุณกับเพื่อนๆ</p>
            <div class="card-actions">
              <RouterLink to="/deck" class="btn btn-accent btn-sm">เริ่มต้น</RouterLink>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Card Showcase -->
    <div class="container mx-auto px-4 pb-16">
      <h2 class="text-3xl font-bold text-center mb-8">
        <span class="bg-gradient-to-r from-accent to-primary bg-clip-text text-transparent">
          🌟 ตัวอย่างการ์ด
        </span>
      </h2>
      
      <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-4">
        <div 
          v-for="card in recentCards" 
          :key="card.id" 
          class="card bg-base-200 shadow-lg hover:shadow-xl transition-all duration-300 hover:scale-105"
        >
          <figure class="pt-2 px-2">
            <img 
              :src="card.imageurlen || card.imageurljp" 
              :alt="card.name" 
              class="rounded-lg w-full"
              loading="lazy"
            />
          </figure>
          <div class="card-body p-2">
            <p class="text-xs text-center font-medium truncate">{{ card.name }}</p>
          </div>
        </div>
      </div>
      
      <div class="text-center mt-8">
        <RouterLink to="/clan" class="btn btn-outline btn-primary btn-wide">
          ดูการ์ดทั้งหมด
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.delay-1000 {
  animation-delay: 1s;
}
</style>
