<template>
  <div class="p-4">
    <h2 class="text-2xl font-semibold mb-2">BẢNG: recipe_process_definition</h2>

    <table class="table-auto w-full border border-collapse border-gray-600 text-sm">
      <thead class="bg-gray-300">
        <tr>
          <th class="border border-gray-600 px-2 py-1">Oid</th>
          <th class="border border-gray-600 px-2 py-1">RecipeID</th>
          <th class="border border-gray-600 px-2 py-1">Name</th>
          <th class="border border-gray-600 px-2 py-1">Type</th>
          <th class="border border-gray-600 px-2 py-1">Configs (JSON)</th>
          <th class="border border-gray-600 px-2 py-1">ProductID</th>
          <th class="border border-gray-600 px-2 py-1">ProductType</th>
          <th class="border border-gray-600 px-2 py-1">LimitaryHour</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in items" :key="item.Oid" class="hover:bg-gray-100 align-top">
          <td class="border border-gray-600 px-2 py-1">{{ item.Oid }}</td>
          <td class="border border-gray-600 px-2 py-1">{{ item.RecipeID }}</td>
          <td class="border border-gray-600 px-2 py-1">{{ item.Name }}</td>
          <td class="border border-gray-600 px-2 py-1">{{ item.Type }}</td>
          <td class="border border-gray-600 px-2 py-1 whitespace-pre-wrap max-w-sm text-left">
            {{ JSON.stringify(item.Configs, null, 2) }}
          </td>
          <td class="border border-gray-600 px-2 py-1">{{ item.ProductID }}</td>
          <td class="border border-gray-600 px-2 py-1">{{ item.ProductType }}</td>
          <td class="border border-gray-600 px-2 py-1">{{ item.LimitaryHour }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      items: []
    }
  },
  mounted() {
    axios.get('/api/recipe-definitions')
      .then(res => {
        this.items = res.data
      })
      .catch(err => {
        console.error('Lỗi khi tải dữ liệu:', err)
      })
  }
}
</script>
