<script setup>

defineProps({
  items: [],
})

</script> 
<template>
  <div class="overflow-x-auto overflow-y-scroll"> 
      <div class="items-center justify-center bg-zinc-50 font-sans ">
          <div>
              <div class="h-[410px] bg-zinc-50 shadow-md rounded ">
                  <table class="bg-zinc-50 min-w-max w-full table-auto">
                      <thead>
                          <tr class="bg-gray-500 shadow-md rounded text-gray-100 uppercase text-sm leading-normal">
                              <th class="text-left  py-1 px-2">Subject</th>
                              <th class="text-left py-1 px-2">From</th>
                              <th class="text-left py-1 px-2">To</th>
                              <th class="text-center py-1 px-2"></th>
                          </tr>
                      </thead>
                      <tbody class="cursor-pointer text-gray-600 text-xs font-light">
                        <tr v-for="item in items" @click="mostrar(item._source.Content)" 
                            class="border-b border-gray-200 hover:bg-gray-100">
                              <td class="text-left py-1 px-2">
                                  <div class="flex items-center">
                                      <span class="font-medium">{{ item._source.Subject }}</span>
                                  </div>
                              </td>
                              <td class="text-left py-1 px-2">
                                  <div class="flex items-center">
                                      <span class="font-medium">{{ item._source.From }}</span>
                                  </div>
                              </td>
                              <td class="text-left py-1 px-2">
                                  <div class="flex items-center">
                                      <span class="font-medium">{{ item._source.To }}</span>
                                  </div>
                              </td>
                          </tr>
                      </tbody>
                  </table>
              </div>
          </div>
      </div>
    </div>
</template>
<script>
export default {
  name: "Tabla",
  methods: {
    mostrar(Contenido) { 
      this.$emit('msg', Contenido)
    }
  }
};

</script>
<style scoped>
</style>
