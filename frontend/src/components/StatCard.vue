<script setup lang="ts">
type Accent = 'emerald' | 'sky' | 'violet' | 'rose'

const props = withDefaults(
  defineProps<{
    label: string
    value: string
    accent?: Accent
  }>(),
  { accent: 'emerald' },
)

const ACCENT_BAR: Record<Accent, string> = {
  emerald: 'bg-gradient-to-b from-emerald-500 to-emerald-700',
  sky: 'bg-gradient-to-b from-sky-500 to-sky-700',
  violet: 'bg-gradient-to-b from-violet-500 to-violet-700',
  rose: 'bg-gradient-to-b from-rose-500 to-rose-700',
}

const ACCENT_TEXT: Record<Accent, string> = {
  emerald: 'text-emerald-700',
  sky: 'text-sky-700',
  violet: 'text-violet-700',
  rose: 'text-rose-700',
}
</script>

<template>
  <div
    class="relative overflow-hidden rounded-2xl bg-white p-4 shadow-sm ring-1 ring-slate-200"
  >
    <div class="absolute inset-y-0 left-0 w-1" :class="ACCENT_BAR[props.accent]" />
    <p class="text-xs font-medium uppercase tracking-wider text-slate-500">
      {{ props.label }}
    </p>
    <p class="mt-1 text-2xl font-bold tabular-nums" :class="ACCENT_TEXT[props.accent]">
      {{ props.value }}
    </p>
  </div>
</template>
