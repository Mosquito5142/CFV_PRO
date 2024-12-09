import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProfileView from '../views/ProfileView.vue'
import ClanView from '../views/ClanView.vue'
import SetView from '../views/SetView.vue'
import CardDetailView from '../views/CardDetailView.vue'
import DeckView from '../views/DeckView.vue'
import testApi from '@/views/testApi.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/profile',
      name: 'profile',
      component: ProfileView,
    },
    {
      path: '/clan',
      name: 'clan',
      component: ClanView,
    },
    {
      path: '/set/:setName',
      name: 'set',
      component: SetView,
      props: true,
    },
    {
      path: '/card/:cardId',
      name: 'card',
      component: CardDetailView,
      props: true,
    },
    {
      path : '/deck',
      name: 'deck',
      component: DeckView,
    },
    {
      path : '/testApi',
      name: 'testApi',
      component: testApi,
    }
    
  ],
})

export default router