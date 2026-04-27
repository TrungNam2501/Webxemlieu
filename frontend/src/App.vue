<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppHeader from '@/components/AppHeader.vue'
</script>

<template>
  <AppHeader />
  <RouterView />
  <footer class="mt-auto py-4 text-center text-xs text-slate-400">
    © {{ new Date().getFullYear() }} Kenda Rubber — Webxemlieu
  </footer>
</template>
