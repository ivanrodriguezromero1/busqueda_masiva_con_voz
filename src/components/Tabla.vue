<script setup>

defineProps({
  items: [],
})

</script> 
<template>
  <!-- <h1>{{ msg }}</h1>
  <div class="card">
    <button type="button" @click="count++">count is {{ count }}</button>
  </div> -->
  <div class="overflow-x-auto"> 
      <div class="items-center justify-center bg-gray-200 font-sans ">
          <div>
              <div class="bg-white shadow-md rounded ">
                  <table class="min-w-max w-full table-auto">
                      <thead>
                          <tr class="bg-gray-500 shadow-md rounded text-gray-100 uppercase text-sm leading-normal">
                              <th class="text-left  py-1 px-2">Subject</th>
                              <th class="text-left py-1 px-2">From</th>
                              <th class="text-left py-1 px-2">To</th>
                              <!-- <th class="text-center py-1 px-2">Date</th> -->
                              <th class="text-center py-1 px-2"></th>
                          </tr>
                      </thead>
                      <tbody class="text-gray-600 text-xs font-light">
                        <tr v-for="item in items" class="border-b border-gray-200 hover:bg-gray-100">
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
                              <!-- <td class="text-left py-1 px-2">
                                  <div class="flex items-center">
                                      <span class="font-medium">{{ item._source.Ruta }}</span>
                                  </div>
                              </td> -->
                              <td class="text-left py-1 px-2" @click="mostrar(item._source.Content)">
                                  <div class="flex item-center justify-center">
                                      <div class="w-4 mr-2 cursor-pointer transform hover:text-purple-500 hover:scale-110">
                                          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                                          </svg> 
                                      </div> 
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
  // data() {
  //   return {
  //     // items: [],
  //   };
  // },
  // async mounted() {
  //   const res = await this.axios.post("http://localhost:4080/api/enron_mail_20110402/_search", 
  //       {
  //         search_type: 'match',
  //         query:
  //         {
  //           term: 'Gas',
  //           start_time: '2021-06-02T14:28:31.894Z',
  //           end_time: '2023-12-02T15:28:31.894Z'
  //         },
  //         from: 0,
  //         max_results: 20,
  //         _source: []
  //       },
  //       {
  //         auth: {
  //           username: 'admin',
  //           password: 'Complexpass#123'
  //         },
  //         headers: {
  //         'content-type':'application/json'
  //         } 
  //       });
  //   this.items = res.data.hits.hits  
  //   // console.log(this.items[0]._source)
  //   // document.getElementsByClassName('vista')[0]
  //   //     .addEventListener('click', function (event) {
            
  //   //     });
  // },
  methods: {
    mostrar(Contenido) { 
      this.$emit('msg', Contenido)
    }
  }
  // async updated(){
  //    document.getElementsByClassName('vista')[0]
  //       .addEventListener('click', function (event) {
              
  //       }); 
  // },
};

</script>
<style scoped>
.read-the-docs {
  color: #888;
}
</style>
