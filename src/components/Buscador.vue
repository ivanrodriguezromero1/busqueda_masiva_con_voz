<template>
    <div>
        <input type="text" class=" my-4 w-2/4 bg-gray-50 border border-gray-300 text-gray-900 text-base rounded-lg focus:ring-blue-500 focus:border-blue-500 p-3 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Buscar" required>
        <button id="voz" class="cursor-pointer bg-red-500 ml-2  h-10 rounded-full border-2 border-neutral-500 hover:border-neutral-700" type="button"  @click="hablando"></button>
    </div>
</template>
<script>

export default {
  name: "Buscador",
  data() {
    return {
      items: []
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
          max_results: 20,
          _source: []
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
    this.items = res.data.hits.hits
    this.$emit('Items', this.items)

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
    }, 
  }

</script>
