<template>
    <div>
      <div class="absolute left-2 flex items-center pointer-events-none">
            <svg aria-hidden="true" class="my-6 w-6 h-6 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
        </div>
        <input type="text" class="absolute left-10 my-3 w-[calc(100%_-_54px)] bg-gray-50 border border-gray-300 text-gray-900 text-base rounded-lg focus:ring-blue-500 focus:border-blue-500 p-3 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Buscar" required>
        <button id="voz" class="absolute right-5 my-4 h-10 cursor-pointer bg-red-500 ml-2 rounded-full border-2 border-neutral-500 hover:border-neutral-700" type="button"  @click="hablando"></button>
    </div>
</template>
<script setup>
  defineProps({
    max_num:"",
  })
</script>
<script>
export default {
  name: "Buscador",
  data() {
    return {
      itemsTotales: []
    };
  },
  mounted() {
    const input = document.querySelector('input');
    input.addEventListener('input', 
    async (e)=>{this.envio(e.target.value)}
    );
  },
  methods: {
    async envio(texto){
      const res = await this.axios.post("http://localhost:4080/api/enron_mail_20110402/_search", 
        {
          search_type: 'match',
          query:
          {
            term: texto,
            start_time: '2021-06-02T14:28:31.894Z',
            end_time: '2023-12-02T15:28:31.894Z'
          },
          from: 0,
          max_results: parseInt(this.max_num),
          _source: ["Subject","From","To","Content"]
        },
        {
          auth: {
            username: 'admin',
            password: 'Complexpass#123'
          },
          headers: {
          'content-type':'application/json'
          } 
      });
      this.itemsTotales = res.data.hits.hits
      this.$emit('ItemsTotales', this.itemsTotales)
    },
    hablando(){
        const  recognition = new (window.SpeechRecognition || window.webkitSpeechRecognition || window.mozSpeechRecognition || window.msSpeechRecognition)();
        recognition.lang = "es-US";
        recognition.continuous = true;
        recognition.interimResults = true;
        recognition.onresult =  async (event)=> {
          var input = document.querySelector('input');          
          input.value = event.results[0][0].transcript;
          this.envio(input.value)
    
        };
        recognition.start();
      }       
    }
  }
</script>
