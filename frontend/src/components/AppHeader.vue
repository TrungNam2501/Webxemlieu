<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header
    class="sticky top-0 z-20 backdrop-blur bg-white/80 border-b border-slate-200 shadow-sm"
  >
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-14 flex items-center gap-3">
      <RouterLink
        :to="{ name: 'recipes' }"
        class="flex items-center gap-2 text-slate-900 hover:text-emerald-700"
      >
        <span
          class="inline-flex h-9 w-9 items-center justify-center rounded-xl bg-gradient-to-br from-emerald-500 to-emerald-700 text-white font-bold shadow-md"
        >
          K
        </span>
        <span class="flex flex-col leading-tight">
          <span class="text-base font-semibold tracking-tight">Webxemlieu</span>
          <span class="text-[11px] text-slate-500">Cao su Kenda</span>
        </span>
      </RouterLink>

      <nav class="ml-6 hidden sm:flex items-center gap-1 text-sm">
        <RouterLink
          :to="{ name: 'recipes' }"
          class="px-3 py-1.5 rounded-lg text-slate-600 hover:bg-slate-100 hover:text-slate-900 transition"
          active-class="bg-emerald-50 text-emerald-700 font-medium"
        >
          Recipes
        </RouterLink>
      </nav>

      <div class="ml-auto flex items-center gap-2 text-xs text-slate-500">
        <span
          class="inline-flex items-center gap-1 rounded-full bg-emerald-50 px-2 py-1 text-emerald-700 ring-1 ring-emerald-200"
        >
          <span class="h-1.5 w-1.5 rounded-full bg-emerald-500"></span>
          Online
        </span>
      </div>
    </div>
  </header>
</template>
