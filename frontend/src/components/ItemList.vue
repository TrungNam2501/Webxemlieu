<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { api } from '@/api/client'
import type { RecipeProcessDefinition } from '@/types/recipe'

const items = ref<RecipeProcessDefinition[]>([])
const loading = ref(false)
const error = ref<string | null>(null)

function formatJson(value: unknown): string {
  if (value === null || value === undefined) return ''
  try {
    return JSON.stringify(value, null, 2)
  } catch {
    return String(value)
  }
}

async function load() {
  loading.value = true
  error.value = null
  try {
    const res = await api.get<RecipeProcessDefinition[]>('/recipe-definitions')
    items.value = res.data
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
  <section class="p-4 space-y-4">
    <header class="flex items-center justify-between gap-2 flex-wrap">
      <h2 class="text-2xl font-semibold">BẢNG: recipe_process_definition</h2>
      <button
        type="button"
        class="px-3 py-1.5 rounded border border-gray-400 bg-white hover:bg-gray-100 disabled:opacity-50"
        :disabled="loading"
        @click="load"
      >
        {{ loading ? 'Đang tải…' : 'Tải lại' }}
      </button>
    </header>

    <p v-if="error" class="rounded border border-red-400 bg-red-50 p-2 text-red-700">
      {{ error }}
    </p>

    <p v-if="loading && !items.length" class="text-gray-600">Đang tải dữ liệu…</p>

    <div v-else class="overflow-x-auto">
      <table class="table-auto w-full border border-collapse border-gray-600 text-sm">
        <thead class="bg-gray-300">
          <tr>
            <th class="border border-gray-600 px-2 py-1 text-left">Oid</th>
            <th class="border border-gray-600 px-2 py-1 text-left">RecipeID</th>
            <th class="border border-gray-600 px-2 py-1 text-left">Name</th>
            <th class="border border-gray-600 px-2 py-1 text-left">Type</th>
            <th class="border border-gray-600 px-2 py-1 text-left">Configs (JSON)</th>
            <th class="border border-gray-600 px-2 py-1 text-left">ProductID</th>
            <th class="border border-gray-600 px-2 py-1 text-left">ProductType</th>
            <th class="border border-gray-600 px-2 py-1 text-left">LimitaryHour</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in items"
            :key="item.Oid"
            class="hover:bg-gray-100 align-top"
          >
            <td class="border border-gray-600 px-2 py-1">{{ item.Oid }}</td>
            <td class="border border-gray-600 px-2 py-1">{{ item.RecipeID }}</td>
            <td class="border border-gray-600 px-2 py-1">{{ item.Name }}</td>
            <td class="border border-gray-600 px-2 py-1">{{ item.Type }}</td>
            <td
              class="border border-gray-600 px-2 py-1 whitespace-pre-wrap max-w-sm text-left font-mono"
            >
              {{ formatJson(item.Configs) }}
            </td>
            <td class="border border-gray-600 px-2 py-1">{{ item.ProductID }}</td>
            <td class="border border-gray-600 px-2 py-1">{{ item.ProductType }}</td>
            <td class="border border-gray-600 px-2 py-1 font-mono">
              {{ formatJson(item.LimitaryHour) }}
            </td>
          </tr>
          <tr v-if="!items.length && !loading">
            <td colspan="8" class="border border-gray-600 px-2 py-4 text-center text-gray-600">
              Không có dữ liệu.
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
