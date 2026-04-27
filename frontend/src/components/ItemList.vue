<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { api } from '@/api/client'
import StatCard from '@/components/StatCard.vue'
import type { RecipeProcessDefinition } from '@/types/recipe'

const items = ref<RecipeProcessDefinition[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const lastFetched = ref<Date | null>(null)
const keyword = ref('')

function formatJson(value: unknown): string {
  if (value === null || value === undefined || value === '') return ''
  try {
    return JSON.stringify(value, null, 2)
  } catch {
    return String(value)
  }
}

function formatTime(d: Date | null): string {
  if (!d) return '—'
  return d.toLocaleTimeString('vi-VN', { hour12: false })
}

const TYPE_STYLES: Record<string, string> = {
  mixing: 'bg-amber-50 text-amber-700 ring-amber-200',
  cure: 'bg-rose-50 text-rose-700 ring-rose-200',
  curing: 'bg-rose-50 text-rose-700 ring-rose-200',
  build: 'bg-sky-50 text-sky-700 ring-sky-200',
  building: 'bg-sky-50 text-sky-700 ring-sky-200',
  inspection: 'bg-violet-50 text-violet-700 ring-violet-200',
}

function typeBadgeClass(type: string | null | undefined): string {
  const key = (type ?? '').toLowerCase()
  return TYPE_STYLES[key] ?? 'bg-slate-100 text-slate-700 ring-slate-200'
}

const filteredItems = computed(() => {
  const q = keyword.value.trim().toLowerCase()
  if (!q) return items.value
  return items.value.filter((item) =>
    [item.Name, item.RecipeID, item.Type, item.ProductID, item.ProductType, item.Oid]
      .map((v) => (v ?? '').toString().toLowerCase())
      .some((v) => v.includes(q)),
  )
})

async function load() {
  loading.value = true
  error.value = null
  try {
    const res = await api.get<RecipeProcessDefinition[]>('/recipe-definitions')
    items.value = Array.isArray(res.data) ? res.data : []
    lastFetched.value = new Date()
  } catch (err) {
    console.error('Lỗi khi tải dữ liệu:', err)
    error.value = err instanceof Error ? err.message : 'Lỗi không xác định khi tải dữ liệu.'
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>

<template>
  <section class="max-w-7xl mx-auto w-full px-4 sm:px-6 lg:px-8 py-6 space-y-6">
    <!-- Title + actions -->
    <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
      <div>
        <p class="text-xs uppercase tracking-wider text-emerald-700 font-semibold">
          KVMES · MES dataset
        </p>
        <h1 class="text-2xl sm:text-3xl font-bold text-slate-900 mt-1">
          recipe_process_definition
        </h1>
        <p class="text-sm text-slate-500 mt-1">
          Bảng định nghĩa quy trình sản xuất — tối đa 100 dòng đầu.
        </p>
      </div>

      <button
        type="button"
        class="inline-flex items-center gap-2 rounded-lg bg-emerald-600 px-3.5 py-2 text-sm font-medium text-white shadow-sm ring-1 ring-emerald-700/40 hover:bg-emerald-700 focus:outline-none focus:ring-2 focus:ring-emerald-400 disabled:opacity-50 disabled:cursor-not-allowed transition"
        :disabled="loading"
        @click="load"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="h-4 w-4"
          :class="{ 'animate-spin': loading }"
          aria-hidden="true"
        >
          <path
            fill-rule="evenodd"
            d="M15.312 11.424a5.5 5.5 0 01-9.201 2.466l-.312-.311h2.433a.75.75 0 000-1.5H3.989a.75.75 0 00-.75.75v4.242a.75.75 0 001.5 0v-2.43l.31.31a7 7 0 0011.712-3.138.75.75 0 00-1.449-.39zm1.23-3.723a.75.75 0 00.219-.53V2.929a.75.75 0 00-1.5 0V5.36l-.31-.31A7 7 0 003.239 8.188a.75.75 0 101.448.389A5.5 5.5 0 0113.89 6.11l.311.31h-2.432a.75.75 0 000 1.5h4.243a.75.75 0 00.53-.219z"
            clip-rule="evenodd"
          />
        </svg>
        {{ loading ? 'Đang tải…' : 'Tải lại' }}
      </button>
    </div>

    <!-- Stats -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-3">
      <StatCard :label="'Tổng số dòng'" :value="items.length.toString()" accent="emerald" />
      <StatCard
        :label="'Đang hiển thị'"
        :value="filteredItems.length.toString()"
        accent="sky"
      />
      <StatCard :label="'Lần tải gần nhất'" :value="formatTime(lastFetched)" accent="violet" />
    </div>

    <!-- Error -->
    <div
      v-if="error"
      class="rounded-xl border border-rose-200 bg-rose-50 p-4 flex items-start gap-3"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 20 20"
        fill="currentColor"
        class="h-5 w-5 text-rose-500 mt-0.5 shrink-0"
        aria-hidden="true"
      >
        <path
          fill-rule="evenodd"
          d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z"
          clip-rule="evenodd"
        />
      </svg>
      <div class="flex-1 min-w-0">
        <p class="text-sm font-semibold text-rose-800">Không tải được dữ liệu</p>
        <p class="text-sm text-rose-700 mt-0.5 break-words">{{ error }}</p>
      </div>
      <button
        type="button"
        class="text-sm font-medium text-rose-700 hover:text-rose-900"
        @click="load"
      >
        Thử lại
      </button>
    </div>

    <!-- Card chứa bảng -->
    <div class="bg-white rounded-2xl shadow-sm ring-1 ring-slate-200 overflow-hidden">
      <!-- Search bar -->
      <div class="px-4 sm:px-5 py-3 border-b border-slate-200 flex items-center gap-3">
        <div class="relative flex-1 max-w-md">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            class="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-slate-400"
            aria-hidden="true"
          >
            <path
              fill-rule="evenodd"
              d="M9 3.5a5.5 5.5 0 100 11 5.5 5.5 0 000-11zM2 9a7 7 0 1112.452 4.391l3.328 3.329a.75.75 0 11-1.06 1.06l-3.329-3.328A7 7 0 012 9z"
              clip-rule="evenodd"
            />
          </svg>
          <input
            v-model="keyword"
            type="search"
            placeholder="Tìm theo Name / RecipeID / Type / Product…"
            class="w-full pl-9 pr-3 py-2 text-sm bg-slate-50 border border-slate-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-300 focus:border-emerald-400"
          />
        </div>
        <span class="text-xs text-slate-400 hidden sm:inline">
          {{ filteredItems.length }} / {{ items.length }} dòng
        </span>
      </div>

      <!-- Loading skeleton -->
      <div v-if="loading && !items.length" class="divide-y divide-slate-100">
        <div v-for="n in 6" :key="n" class="px-4 sm:px-5 py-3 flex gap-4 animate-pulse">
          <div class="h-3 w-32 bg-slate-200 rounded" />
          <div class="h-3 w-24 bg-slate-200 rounded" />
          <div class="h-3 w-40 bg-slate-200 rounded" />
          <div class="h-3 w-20 bg-slate-200 rounded" />
          <div class="h-3 flex-1 bg-slate-200 rounded" />
        </div>
      </div>

      <!-- Empty state -->
      <div
        v-else-if="!filteredItems.length"
        class="px-6 py-16 text-center text-slate-500 space-y-2"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="1.5"
          class="mx-auto h-10 w-10 text-slate-300"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M3.75 12h16.5m-16.5 3.75h16.5M3.75 19.5h16.5M5.625 4.5h12.75a1.875 1.875 0 010 3.75H5.625a1.875 1.875 0 010-3.75z"
          />
        </svg>
        <p class="text-sm font-medium text-slate-700">Không có dữ liệu</p>
        <p class="text-xs text-slate-400">Thử bỏ bớt từ khoá tìm kiếm hoặc bấm "Tải lại".</p>
      </div>

      <!-- Table -->
      <div v-else class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead
            class="sticky top-0 bg-slate-50/80 backdrop-blur text-slate-600 text-xs uppercase tracking-wider"
          >
            <tr>
              <th class="px-4 py-2.5 text-left font-semibold">Oid</th>
              <th class="px-4 py-2.5 text-left font-semibold">RecipeID</th>
              <th class="px-4 py-2.5 text-left font-semibold">Name</th>
              <th class="px-4 py-2.5 text-left font-semibold">Type</th>
              <th class="px-4 py-2.5 text-left font-semibold">Configs</th>
              <th class="px-4 py-2.5 text-left font-semibold">ProductID</th>
              <th class="px-4 py-2.5 text-left font-semibold">ProductType</th>
              <th class="px-4 py-2.5 text-left font-semibold">LimitaryHour</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-slate-100">
            <tr
              v-for="item in filteredItems"
              :key="item.Oid"
              class="align-top hover:bg-emerald-50/40 transition"
            >
              <td class="px-4 py-2.5 font-mono text-[11px] text-slate-500 whitespace-nowrap">
                {{ item.Oid }}
              </td>
              <td class="px-4 py-2.5 font-mono text-xs text-slate-700 whitespace-nowrap">
                {{ item.RecipeID }}
              </td>
              <td class="px-4 py-2.5 font-medium text-slate-900">
                {{ item.Name }}
              </td>
              <td class="px-4 py-2.5">
                <span
                  v-if="item.Type"
                  class="inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium ring-1 ring-inset"
                  :class="typeBadgeClass(item.Type)"
                >
                  {{ item.Type }}
                </span>
                <span v-else class="text-slate-300">—</span>
              </td>
              <td class="px-4 py-2.5">
                <pre
                  v-if="formatJson(item.Configs)"
                  class="font-mono text-[11px] leading-snug bg-slate-50 text-slate-700 rounded-md px-2 py-1.5 max-w-xs whitespace-pre-wrap break-words"
                  >{{ formatJson(item.Configs) }}</pre
                >
                <span v-else class="text-slate-300">—</span>
              </td>
              <td class="px-4 py-2.5 font-mono text-xs text-slate-700 whitespace-nowrap">
                {{ item.ProductID || '—' }}
              </td>
              <td class="px-4 py-2.5 text-slate-700 whitespace-nowrap">
                {{ item.ProductType || '—' }}
              </td>
              <td class="px-4 py-2.5">
                <pre
                  v-if="formatJson(item.LimitaryHour)"
                  class="font-mono text-[11px] leading-snug bg-slate-50 text-slate-700 rounded-md px-2 py-1.5 max-w-[16rem] whitespace-pre-wrap break-words"
                  >{{ formatJson(item.LimitaryHour) }}</pre
                >
                <span v-else class="text-slate-300">—</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>
</template>
